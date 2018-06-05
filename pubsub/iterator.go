// Copyright 2016 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pubsub

import (
	"sync"
	"time"

	vkit "cloud.google.com/go/pubsub/apiv1"
	"cloud.google.com/go/pubsub/internal/distribution"
	"golang.org/x/net/context"
	pb "google.golang.org/genproto/googleapis/pubsub/v1"
)

// newMessageIterator starts a new streamingMessageIterator.  Stop must be called on the messageIterator
// when it is no longer needed.
// subName is the full name of the subscription to pull messages from.
// ctx is the context to use for acking messages and extending message deadlines.
func newMessageIterator(ctx context.Context, subc *vkit.SubscriberClient, subName string, po *pullOptions) *streamingMessageIterator {
	ps := newPullStream(ctx, subc.StreamingPull, subName, int32(po.ackDeadline.Seconds()))
	return newStreamingMessageIterator(ctx, ps, po, subc, subName)
}

type streamingMessageIterator struct {
	ctx        context.Context
	po         *pullOptions
	ps         *pullStream
	subc       *vkit.SubscriberClient
	subName    string
	kaTicker   *time.Ticker  // keep-alive (deadline extensions)
	ackTicker  *time.Ticker  // message acks
	nackTicker *time.Ticker  // message nacks (more frequent than acks)
	pingTicker *time.Ticker  //  sends to the stream to keep it open
	failed     chan struct{} // closed on stream error
	stopped    chan struct{} // closed when Stop is called
	drained    chan struct{} // closed when stopped && no more pending messages
	wg         sync.WaitGroup

	mu                 sync.Mutex
	ackTimeDist        *distribution.D
	keepAliveDeadlines map[string]time.Time
	pendingAcks        map[string]bool
	pendingNacks       map[string]bool
	pendingModAcks     map[string]bool // ack IDs whose ack deadline is to be modified
	err                error           // error from stream failure
}

func newStreamingMessageIterator(ctx context.Context, ps *pullStream, po *pullOptions, subc *vkit.SubscriberClient, subName string) *streamingMessageIterator {
	// TODO: make kaTicker frequency more configurable. (ackDeadline - 5s) is a
	// reasonable default for now, because the minimum ack period is 10s. This
	// gives us 5s grace.
	keepAlivePeriod := po.ackDeadline - 5*time.Second
	kaTicker := time.NewTicker(keepAlivePeriod)

	// Ack promptly so users don't lose work if client crashes.
	ackTicker := time.NewTicker(100 * time.Millisecond)
	nackTicker := time.NewTicker(100 * time.Millisecond)
	pingTicker := time.NewTicker(30 * time.Second)
	it := &streamingMessageIterator{
		ctx:                ctx,
		ps:                 ps,
		po:                 po,
		subc:               subc,
		subName:            subName,
		kaTicker:           kaTicker,
		ackTicker:          ackTicker,
		nackTicker:         nackTicker,
		pingTicker:         pingTicker,
		failed:             make(chan struct{}),
		stopped:            make(chan struct{}),
		drained:            make(chan struct{}),
		ackTimeDist:        distribution.New(int(maxAckDeadline/time.Second) + 1),
		keepAliveDeadlines: map[string]time.Time{},
		pendingAcks:        map[string]bool{},
		pendingNacks:       map[string]bool{},
		pendingModAcks:     map[string]bool{},
	}
	it.wg.Add(1)
	go it.sender()
	return it
}

// Subscription.receive will call stop on its messageIterator when finished with it.
// Stop will block until Done has been called on all Messages that have been
// returned by Next, or until the context with which the messageIterator was created
// is cancelled or exceeds its deadline.
func (it *streamingMessageIterator) stop() {
	it.mu.Lock()
	select {
	case <-it.stopped:
	default:
		close(it.stopped)
	}
	it.checkDrained()
	it.mu.Unlock()
	it.wg.Wait()
}

// checkDrained closes the drained channel if the iterator has been stopped and all
// pending messages have either been n/acked or expired.
//
// Called with the lock held.
func (it *streamingMessageIterator) checkDrained() {
	select {
	case <-it.drained:
		return
	default:
	}
	select {
	case <-it.stopped:
		if len(it.keepAliveDeadlines) == 0 {
			close(it.drained)
		}
	default:
	}
}

// Called when a message is acked/nacked.
func (it *streamingMessageIterator) done(ackID string, ack bool, receiveTime time.Time) {
	it.ackTimeDist.Record(int(time.Since(receiveTime) / time.Second))
	it.mu.Lock()
	defer it.mu.Unlock()
	delete(it.keepAliveDeadlines, ackID)
	if ack {
		it.pendingAcks[ackID] = true
	} else {
		it.pendingNacks[ackID] = true
	}
	it.checkDrained()
}

// fail is called when a stream method returns a permanent error.
func (it *streamingMessageIterator) fail(err error) {
	it.mu.Lock()
	if it.err == nil {
		it.err = err
		close(it.failed)
	}
	it.mu.Unlock()
}

// receive makes a call to the stream's Recv method and returns
// its messages.
func (it *streamingMessageIterator) receive() ([]*Message, error) {
	// Stop retrieving messages if the context is done, the stream
	// failed, or the iterator's Stop method was called.
	select {
	case <-it.ctx.Done():
		return nil, it.ctx.Err()
	default:
	}
	it.mu.Lock()
	err := it.err
	it.mu.Unlock()
	if err != nil {
		return nil, err
	}
	// Receive messages from stream. This may block indefinitely.
	res, err := it.ps.Recv()
	// The pullStream handles retries, so any error here is fatal.
	if err != nil {
		it.fail(err)
		return nil, err
	}
	msgs, err := convertMessages(res.ReceivedMessages)
	if err != nil {
		it.fail(err)
		return nil, err
	}
	// We received some messages. Remember them so we can keep them alive. Also,
	// do a receipt mod-ack.
	maxExt := time.Now().Add(it.po.maxExtension)
	ackIDs := map[string]bool{}
	it.mu.Lock()
	now := time.Now()
	for _, m := range msgs {
		m.receiveTime = now
		addRecv(m.ID, m.ackID, now)
		m.doneFunc = it.done
		it.keepAliveDeadlines[m.ackID] = maxExt
		// The receipt mod-ack uses the subscription's configured ack deadline. Don't
		// change the mod-ack if the message is going to be nacked. This is possible
		// if there are retries.
		if !it.pendingNacks[m.ackID] {
			ackIDs[m.ackID] = true
		}
	}
	it.mu.Unlock()
	if !it.sendModAck(ackIDs, trunc32(int64(it.po.ackDeadline.Seconds()))) {
		return nil, it.err
	}
	return msgs, nil
}

// sender runs in a goroutine and handles all sends to the stream.
func (it *streamingMessageIterator) sender() {
	defer it.wg.Done()
	defer it.kaTicker.Stop()
	defer it.ackTicker.Stop()
	defer it.nackTicker.Stop()
	defer it.pingTicker.Stop()
	defer it.ps.CloseSend()

	done := false
	for !done {
		sendAcks := false
		sendNacks := false
		sendModAcks := false
		sendPing := false
		select {
		case <-it.ctx.Done():
			// Context canceled or timed out: stop immediately, without
			// another RPC.
			return

		case <-it.failed:
			// Stream failed: nothing to do, so stop immediately.
			return

		case <-it.drained:
			// All outstanding messages have been marked done:
			// nothing left to do except make the final calls.
			it.mu.Lock()
			sendAcks = (len(it.pendingAcks) > 0)
			sendNacks = (len(it.pendingNacks) > 0)
			// No point in sending modacks.
			done = true

		case <-it.kaTicker.C:
			it.mu.Lock()
			it.handleKeepAlives()
			sendModAcks = (len(it.pendingModAcks) > 0)

		case <-it.nackTicker.C:
			it.mu.Lock()
			sendNacks = (len(it.pendingNacks) > 0)

		case <-it.ackTicker.C:
			it.mu.Lock()
			sendAcks = (len(it.pendingAcks) > 0)

		case <-it.pingTicker.C:
			it.mu.Lock()
			// Ping only if we are processing messages.
			sendPing = (len(it.keepAliveDeadlines) > 0)
		}
		// Lock is held here.
		var acks, nacks, modAcks map[string]bool
		if sendAcks {
			acks = it.pendingAcks
			it.pendingAcks = map[string]bool{}
		}
		if sendNacks {
			nacks = it.pendingNacks
			it.pendingNacks = map[string]bool{}
		}
		if sendModAcks {
			modAcks = it.pendingModAcks
			it.pendingModAcks = map[string]bool{}
		}
		it.mu.Unlock()
		// Make Ack and ModAck RPCs.
		if sendAcks {
			if !it.sendAck(acks) {
				return
			}
		}
		if sendNacks {
			// Nack indicated by modifying the deadline to zero.
			if !it.sendModAck(nacks, 0) {
				return
			}
		}
		if sendModAcks {
			if !it.sendModAck(modAcks, trunc32(int64(it.po.ackDeadline.Seconds()))) {
				return
			}
		}
		if sendPing {
			it.pingStream()
		}
	}
}

// handleKeepAlives modifies the pending request to include deadline extensions
// for live messages. It also purges expired messages.
//
// Called with the lock held.
func (it *streamingMessageIterator) handleKeepAlives() {
	now := time.Now()
	for id, expiry := range it.keepAliveDeadlines {
		if expiry.Before(now) {
			// This delete will not result in skipping any map items, as implied by
			// the spec at https://golang.org/ref/spec#For_statements, "For
			// statements with range clause", note 3, and stated explicitly at
			// https://groups.google.com/forum/#!msg/golang-nuts/UciASUb03Js/pzSq5iVFAQAJ.
			delete(it.keepAliveDeadlines, id)
		} else {
			// This will not conflict with a nack, because nacking removes the ID from keepAliveDeadlines.
			it.pendingModAcks[id] = true
		}
	}
	it.checkDrained()
}

func (it *streamingMessageIterator) sendAck(m map[string]bool) bool {
	return it.sendAckIDRPC(m, func(ids []string) error {
		addAcks(ids)
		return it.subc.Acknowledge(it.ctx, &pb.AcknowledgeRequest{
			Subscription: it.subName,
			AckIds:       ids,
		})
	})
}

func (it *streamingMessageIterator) sendModAck(m map[string]bool, deadlineSecs int32) bool {
	return it.sendAckIDRPC(m, func(ids []string) error {
		addModAcks(ids, deadlineSecs)
		return it.subc.ModifyAckDeadline(it.ctx, &pb.ModifyAckDeadlineRequest{
			Subscription:       it.subName,
			AckDeadlineSeconds: deadlineSecs,
			AckIds:             ids,
		})
	})
}

func (it *streamingMessageIterator) sendAckIDRPC(ackIDSet map[string]bool, call func([]string) error) bool {
	ackIDs := make([]string, 0, len(ackIDSet))
	for k := range ackIDSet {
		ackIDs = append(ackIDs, k)
	}
	var toSend []string
	for len(ackIDs) > 0 {
		toSend, ackIDs = splitRequestIDs(ackIDs, maxPayload)
		if err := call(toSend); err != nil {
			// The underlying client handles retries, so any error is fatal to the
			// iterator.
			it.fail(err)
			return false
		}
	}
	return true
}

// Send a message to the stream to keep it open. The stream will close if there's no
// traffic on it for a while. By keeping it open, we delay the start of the
// expiration timer on messages that are buffered by gRPC or elsewhere in the
// network. This matters if it takes a long time to process messages relative to the
// default ack deadline, and if the messages are small enough so that many can fit
// into the buffer.
func (it *streamingMessageIterator) pingStream() {
	// Ignore error; if the stream is broken, this doesn't matter anyway.
	_ = it.ps.Send(&pb.StreamingPullRequest{})
}

func splitRequestIDs(ids []string, maxSize int) (prefix, remainder []string) {
	size := reqFixedOverhead
	i := 0
	for size < maxSize && i < len(ids) {
		size += overheadPerID + len(ids[i])
		i++
	}
	if size > maxSize {
		i--
	}
	return ids[:i], ids[i:]
}
