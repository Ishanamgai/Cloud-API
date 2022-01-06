// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package storage

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	storagepb "google.golang.org/genproto/googleapis/cloud/bigquery/storage/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newBigQueryReadClientHook clientHook

// BigQueryReadCallOptions contains the retry settings for each method of BigQueryReadClient.
type BigQueryReadCallOptions struct {
	CreateReadSession []gax.CallOption
	ReadRows          []gax.CallOption
	SplitReadStream   []gax.CallOption
}

func defaultBigQueryReadGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("bigquerystorage.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("bigquerystorage.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://bigquerystorage.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultBigQueryReadCallOptions() *BigQueryReadCallOptions {
	return &BigQueryReadCallOptions{
		CreateReadSession: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ReadRows: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		SplitReadStream: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// internalBigQueryReadClient is an interface that defines the methods availaible from BigQuery Storage API.
type internalBigQueryReadClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateReadSession(context.Context, *storagepb.CreateReadSessionRequest, ...gax.CallOption) (*storagepb.ReadSession, error)
	ReadRows(context.Context, *storagepb.ReadRowsRequest, ...gax.CallOption) (storagepb.BigQueryRead_ReadRowsClient, error)
	SplitReadStream(context.Context, *storagepb.SplitReadStreamRequest, ...gax.CallOption) (*storagepb.SplitReadStreamResponse, error)
}

// BigQueryReadClient is a client for interacting with BigQuery Storage API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// BigQuery Read API.
//
// The Read API can be used to read data from BigQuery.
type BigQueryReadClient struct {
	// The internal transport-dependent client.
	internalClient internalBigQueryReadClient

	// The call options for this service.
	CallOptions *BigQueryReadCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *BigQueryReadClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *BigQueryReadClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *BigQueryReadClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateReadSession creates a new read session. A read session divides the contents of a
// BigQuery table into one or more streams, which can then be used to read
// data from the table. The read session also specifies properties of the
// data to be read, such as a list of columns or a push-down filter describing
// the rows to be returned.
//
// A particular row can be read by at most one stream. When the caller has
// reached the end of each stream in the session, then all the data in the
// table has been read.
//
// Data is assigned to each stream such that roughly the same number of
// rows can be read from each stream. Because the server-side unit for
// assigning data is collections of rows, the API does not guarantee that
// each stream will return the same number or rows. Additionally, the
// limits are enforced based on the number of pre-filtered rows, so some
// filters can lead to lopsided assignments.
//
// Read sessions automatically expire 6 hours after they are created and do
// not require manual clean-up by the caller.
func (c *BigQueryReadClient) CreateReadSession(ctx context.Context, req *storagepb.CreateReadSessionRequest, opts ...gax.CallOption) (*storagepb.ReadSession, error) {
	return c.internalClient.CreateReadSession(ctx, req, opts...)
}

// ReadRows reads rows from the stream in the format prescribed by the ReadSession.
// Each response contains one or more table rows, up to a maximum of 100 MiB
// per response; read requests which attempt to read individual rows larger
// than 100 MiB will fail.
//
// Each request also returns a set of stream statistics reflecting the current
// state of the stream.
func (c *BigQueryReadClient) ReadRows(ctx context.Context, req *storagepb.ReadRowsRequest, opts ...gax.CallOption) (storagepb.BigQueryRead_ReadRowsClient, error) {
	return c.internalClient.ReadRows(ctx, req, opts...)
}

// SplitReadStream splits a given ReadStream into two ReadStream objects. These
// ReadStream objects are referred to as the primary and the residual
// streams of the split. The original ReadStream can still be read from in
// the same manner as before. Both of the returned ReadStream objects can
// also be read from, and the rows returned by both child streams will be
// the same as the rows read from the original stream.
//
// Moreover, the two child streams will be allocated back-to-back in the
// original ReadStream. Concretely, it is guaranteed that for streams
// original, primary, and residual, that original[0-j] = primary[0-j] and
// original[j-n] = residual[0-m] once the streams have been read to
// completion.
func (c *BigQueryReadClient) SplitReadStream(ctx context.Context, req *storagepb.SplitReadStreamRequest, opts ...gax.CallOption) (*storagepb.SplitReadStreamResponse, error) {
	return c.internalClient.SplitReadStream(ctx, req, opts...)
}

// bigQueryReadGRPCClient is a client for interacting with BigQuery Storage API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type bigQueryReadGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing BigQueryReadClient
	CallOptions **BigQueryReadCallOptions

	// The gRPC API client.
	bigQueryReadClient storagepb.BigQueryReadClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewBigQueryReadClient creates a new big query read client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// BigQuery Read API.
//
// The Read API can be used to read data from BigQuery.
func NewBigQueryReadClient(ctx context.Context, opts ...option.ClientOption) (*BigQueryReadClient, error) {
	clientOpts := defaultBigQueryReadGRPCClientOptions()
	if newBigQueryReadClientHook != nil {
		hookOpts, err := newBigQueryReadClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := BigQueryReadClient{CallOptions: defaultBigQueryReadCallOptions()}

	c := &bigQueryReadGRPCClient{
		connPool:           connPool,
		disableDeadlines:   disableDeadlines,
		bigQueryReadClient: storagepb.NewBigQueryReadClient(connPool),
		CallOptions:        &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *bigQueryReadGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *bigQueryReadGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *bigQueryReadGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *bigQueryReadGRPCClient) CreateReadSession(ctx context.Context, req *storagepb.CreateReadSessionRequest, opts ...gax.CallOption) (*storagepb.ReadSession, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "read_session.table", url.QueryEscape(req.GetReadSession().GetTable())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateReadSession[0:len((*c.CallOptions).CreateReadSession):len((*c.CallOptions).CreateReadSession)], opts...)
	var resp *storagepb.ReadSession
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.bigQueryReadClient.CreateReadSession(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *bigQueryReadGRPCClient) ReadRows(ctx context.Context, req *storagepb.ReadRowsRequest, opts ...gax.CallOption) (storagepb.BigQueryRead_ReadRowsClient, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "read_stream", url.QueryEscape(req.GetReadStream())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	var resp storagepb.BigQueryRead_ReadRowsClient
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.bigQueryReadClient.ReadRows(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *bigQueryReadGRPCClient) SplitReadStream(ctx context.Context, req *storagepb.SplitReadStreamRequest, opts ...gax.CallOption) (*storagepb.SplitReadStreamResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).SplitReadStream[0:len((*c.CallOptions).SplitReadStream):len((*c.CallOptions).SplitReadStream)], opts...)
	var resp *storagepb.SplitReadStreamResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.bigQueryReadClient.SplitReadStream(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
