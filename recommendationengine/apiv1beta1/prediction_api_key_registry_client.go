// Copyright 2021 Google LLC
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

package recommendationengine

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	"github.com/golang/protobuf/proto"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	recommendationenginepb "google.golang.org/genproto/googleapis/cloud/recommendationengine/v1beta1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newPredictionApiKeyRegistryClientHook clientHook

// PredictionApiKeyRegistryCallOptions contains the retry settings for each method of PredictionApiKeyRegistryClient.
type PredictionApiKeyRegistryCallOptions struct {
	CreatePredictionApiKeyRegistration []gax.CallOption
	ListPredictionApiKeyRegistrations  []gax.CallOption
	DeletePredictionApiKeyRegistration []gax.CallOption
}

func defaultPredictionApiKeyRegistryGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("recommendationengine.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("recommendationengine.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://recommendationengine.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultPredictionApiKeyRegistryCallOptions() *PredictionApiKeyRegistryCallOptions {
	return &PredictionApiKeyRegistryCallOptions{
		CreatePredictionApiKeyRegistration: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ListPredictionApiKeyRegistrations: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		DeletePredictionApiKeyRegistration: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// internalPredictionApiKeyRegistryClient is an interface that defines the methods availaible from Recommendations AI.
type internalPredictionApiKeyRegistryClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreatePredictionApiKeyRegistration(context.Context, *recommendationenginepb.CreatePredictionApiKeyRegistrationRequest, ...gax.CallOption) (*recommendationenginepb.PredictionApiKeyRegistration, error)
	ListPredictionApiKeyRegistrations(context.Context, *recommendationenginepb.ListPredictionApiKeyRegistrationsRequest, ...gax.CallOption) *PredictionApiKeyRegistrationIterator
	DeletePredictionApiKeyRegistration(context.Context, *recommendationenginepb.DeletePredictionApiKeyRegistrationRequest, ...gax.CallOption) error
}

// PredictionApiKeyRegistryClient is a client for interacting with Recommendations AI.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service for registering API keys for use with the predict method. If you
// use an API key to request predictions, you must first register the API key.
// Otherwise, your prediction request is rejected. If you use OAuth to
// authenticate your predict method call, you do not need to register an API
// key. You can register up to 20 API keys per project.
type PredictionApiKeyRegistryClient struct {
	// The internal transport-dependent client.
	internalClient internalPredictionApiKeyRegistryClient

	// The call options for this service.
	CallOptions *PredictionApiKeyRegistryCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *PredictionApiKeyRegistryClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *PredictionApiKeyRegistryClient) setGoogleClientInfo(...string) {
	c.internalClient.setGoogleClientInfo()
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *PredictionApiKeyRegistryClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreatePredictionApiKeyRegistration register an API key for use with predict method.
func (c *PredictionApiKeyRegistryClient) CreatePredictionApiKeyRegistration(ctx context.Context, req *recommendationenginepb.CreatePredictionApiKeyRegistrationRequest, opts ...gax.CallOption) (*recommendationenginepb.PredictionApiKeyRegistration, error) {
	return c.internalClient.CreatePredictionApiKeyRegistration(ctx, req, opts...)
}

// ListPredictionApiKeyRegistrations list the registered apiKeys for use with predict method.
func (c *PredictionApiKeyRegistryClient) ListPredictionApiKeyRegistrations(ctx context.Context, req *recommendationenginepb.ListPredictionApiKeyRegistrationsRequest, opts ...gax.CallOption) *PredictionApiKeyRegistrationIterator {
	return c.internalClient.ListPredictionApiKeyRegistrations(ctx, req, opts...)
}

// DeletePredictionApiKeyRegistration unregister an apiKey from using for predict method.
func (c *PredictionApiKeyRegistryClient) DeletePredictionApiKeyRegistration(ctx context.Context, req *recommendationenginepb.DeletePredictionApiKeyRegistrationRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeletePredictionApiKeyRegistration(ctx, req, opts...)
}

// predictionApiKeyRegistryGRPCClient is a client for interacting with Recommendations AI over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type predictionApiKeyRegistryGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing PredictionApiKeyRegistryClient
	CallOptions **PredictionApiKeyRegistryCallOptions

	// The gRPC API client.
	predictionApiKeyRegistryClient recommendationenginepb.PredictionApiKeyRegistryClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewPredictionApiKeyRegistryClient creates a new prediction api key registry client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service for registering API keys for use with the predict method. If you
// use an API key to request predictions, you must first register the API key.
// Otherwise, your prediction request is rejected. If you use OAuth to
// authenticate your predict method call, you do not need to register an API
// key. You can register up to 20 API keys per project.
func NewPredictionApiKeyRegistryClient(ctx context.Context, opts ...option.ClientOption) (*PredictionApiKeyRegistryClient, error) {
	clientOpts := defaultPredictionApiKeyRegistryGRPCClientOptions()
	if newPredictionApiKeyRegistryClientHook != nil {
		hookOpts, err := newPredictionApiKeyRegistryClientHook(ctx, clientHookParams{})
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
	client := PredictionApiKeyRegistryClient{CallOptions: defaultPredictionApiKeyRegistryCallOptions()}

	c := &predictionApiKeyRegistryGRPCClient{
		connPool:                       connPool,
		disableDeadlines:               disableDeadlines,
		predictionApiKeyRegistryClient: recommendationenginepb.NewPredictionApiKeyRegistryClient(connPool),
		CallOptions:                    &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *predictionApiKeyRegistryGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *predictionApiKeyRegistryGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *predictionApiKeyRegistryGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *predictionApiKeyRegistryGRPCClient) CreatePredictionApiKeyRegistration(ctx context.Context, req *recommendationenginepb.CreatePredictionApiKeyRegistrationRequest, opts ...gax.CallOption) (*recommendationenginepb.PredictionApiKeyRegistration, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreatePredictionApiKeyRegistration[0:len((*c.CallOptions).CreatePredictionApiKeyRegistration):len((*c.CallOptions).CreatePredictionApiKeyRegistration)], opts...)
	var resp *recommendationenginepb.PredictionApiKeyRegistration
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.predictionApiKeyRegistryClient.CreatePredictionApiKeyRegistration(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *predictionApiKeyRegistryGRPCClient) ListPredictionApiKeyRegistrations(ctx context.Context, req *recommendationenginepb.ListPredictionApiKeyRegistrationsRequest, opts ...gax.CallOption) *PredictionApiKeyRegistrationIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListPredictionApiKeyRegistrations[0:len((*c.CallOptions).ListPredictionApiKeyRegistrations):len((*c.CallOptions).ListPredictionApiKeyRegistrations)], opts...)
	it := &PredictionApiKeyRegistrationIterator{}
	req = proto.Clone(req).(*recommendationenginepb.ListPredictionApiKeyRegistrationsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*recommendationenginepb.PredictionApiKeyRegistration, string, error) {
		var resp *recommendationenginepb.ListPredictionApiKeyRegistrationsResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.predictionApiKeyRegistryClient.ListPredictionApiKeyRegistrations(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetPredictionApiKeyRegistrations(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}
	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()
	return it
}

func (c *predictionApiKeyRegistryGRPCClient) DeletePredictionApiKeyRegistration(ctx context.Context, req *recommendationenginepb.DeletePredictionApiKeyRegistrationRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeletePredictionApiKeyRegistration[0:len((*c.CallOptions).DeletePredictionApiKeyRegistration):len((*c.CallOptions).DeletePredictionApiKeyRegistration)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.predictionApiKeyRegistryClient.DeletePredictionApiKeyRegistration(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// PredictionApiKeyRegistrationIterator manages a stream of *recommendationenginepb.PredictionApiKeyRegistration.
type PredictionApiKeyRegistrationIterator struct {
	items    []*recommendationenginepb.PredictionApiKeyRegistration
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*recommendationenginepb.PredictionApiKeyRegistration, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *PredictionApiKeyRegistrationIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *PredictionApiKeyRegistrationIterator) Next() (*recommendationenginepb.PredictionApiKeyRegistration, error) {
	var item *recommendationenginepb.PredictionApiKeyRegistration
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *PredictionApiKeyRegistrationIterator) bufLen() int {
	return len(it.items)
}

func (it *PredictionApiKeyRegistrationIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
