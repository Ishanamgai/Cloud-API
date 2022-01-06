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

package cx

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	cxpb "google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3beta1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newEntityTypesClientHook clientHook

// EntityTypesCallOptions contains the retry settings for each method of EntityTypesClient.
type EntityTypesCallOptions struct {
	ListEntityTypes  []gax.CallOption
	GetEntityType    []gax.CallOption
	CreateEntityType []gax.CallOption
	UpdateEntityType []gax.CallOption
	DeleteEntityType []gax.CallOption
}

func defaultEntityTypesGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("dialogflow.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("dialogflow.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://dialogflow.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultEntityTypesCallOptions() *EntityTypesCallOptions {
	return &EntityTypesCallOptions{
		ListEntityTypes: []gax.CallOption{
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
		GetEntityType: []gax.CallOption{
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
		CreateEntityType: []gax.CallOption{
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
		UpdateEntityType: []gax.CallOption{
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
		DeleteEntityType: []gax.CallOption{
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
	}
}

// internalEntityTypesClient is an interface that defines the methods availaible from Dialogflow API.
type internalEntityTypesClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListEntityTypes(context.Context, *cxpb.ListEntityTypesRequest, ...gax.CallOption) *EntityTypeIterator
	GetEntityType(context.Context, *cxpb.GetEntityTypeRequest, ...gax.CallOption) (*cxpb.EntityType, error)
	CreateEntityType(context.Context, *cxpb.CreateEntityTypeRequest, ...gax.CallOption) (*cxpb.EntityType, error)
	UpdateEntityType(context.Context, *cxpb.UpdateEntityTypeRequest, ...gax.CallOption) (*cxpb.EntityType, error)
	DeleteEntityType(context.Context, *cxpb.DeleteEntityTypeRequest, ...gax.CallOption) error
}

// EntityTypesClient is a client for interacting with Dialogflow API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service for managing EntityTypes.
type EntityTypesClient struct {
	// The internal transport-dependent client.
	internalClient internalEntityTypesClient

	// The call options for this service.
	CallOptions *EntityTypesCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *EntityTypesClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *EntityTypesClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *EntityTypesClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListEntityTypes returns the list of all entity types in the specified agent.
func (c *EntityTypesClient) ListEntityTypes(ctx context.Context, req *cxpb.ListEntityTypesRequest, opts ...gax.CallOption) *EntityTypeIterator {
	return c.internalClient.ListEntityTypes(ctx, req, opts...)
}

// GetEntityType retrieves the specified entity type.
func (c *EntityTypesClient) GetEntityType(ctx context.Context, req *cxpb.GetEntityTypeRequest, opts ...gax.CallOption) (*cxpb.EntityType, error) {
	return c.internalClient.GetEntityType(ctx, req, opts...)
}

// CreateEntityType creates an entity type in the specified agent.
func (c *EntityTypesClient) CreateEntityType(ctx context.Context, req *cxpb.CreateEntityTypeRequest, opts ...gax.CallOption) (*cxpb.EntityType, error) {
	return c.internalClient.CreateEntityType(ctx, req, opts...)
}

// UpdateEntityType updates the specified entity type.
//
// Note: You should always train a flow prior to sending it queries. See the
// training
// documentation (at https://cloud.google.com/dialogflow/cx/docs/concept/training).
func (c *EntityTypesClient) UpdateEntityType(ctx context.Context, req *cxpb.UpdateEntityTypeRequest, opts ...gax.CallOption) (*cxpb.EntityType, error) {
	return c.internalClient.UpdateEntityType(ctx, req, opts...)
}

// DeleteEntityType deletes the specified entity type.
//
// Note: You should always train a flow prior to sending it queries. See the
// training
// documentation (at https://cloud.google.com/dialogflow/cx/docs/concept/training).
func (c *EntityTypesClient) DeleteEntityType(ctx context.Context, req *cxpb.DeleteEntityTypeRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteEntityType(ctx, req, opts...)
}

// entityTypesGRPCClient is a client for interacting with Dialogflow API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type entityTypesGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing EntityTypesClient
	CallOptions **EntityTypesCallOptions

	// The gRPC API client.
	entityTypesClient cxpb.EntityTypesClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewEntityTypesClient creates a new entity types client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service for managing EntityTypes.
func NewEntityTypesClient(ctx context.Context, opts ...option.ClientOption) (*EntityTypesClient, error) {
	clientOpts := defaultEntityTypesGRPCClientOptions()
	if newEntityTypesClientHook != nil {
		hookOpts, err := newEntityTypesClientHook(ctx, clientHookParams{})
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
	client := EntityTypesClient{CallOptions: defaultEntityTypesCallOptions()}

	c := &entityTypesGRPCClient{
		connPool:          connPool,
		disableDeadlines:  disableDeadlines,
		entityTypesClient: cxpb.NewEntityTypesClient(connPool),
		CallOptions:       &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *entityTypesGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *entityTypesGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *entityTypesGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *entityTypesGRPCClient) ListEntityTypes(ctx context.Context, req *cxpb.ListEntityTypesRequest, opts ...gax.CallOption) *EntityTypeIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListEntityTypes[0:len((*c.CallOptions).ListEntityTypes):len((*c.CallOptions).ListEntityTypes)], opts...)
	it := &EntityTypeIterator{}
	req = proto.Clone(req).(*cxpb.ListEntityTypesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*cxpb.EntityType, string, error) {
		resp := &cxpb.ListEntityTypesResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.entityTypesClient.ListEntityTypes(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetEntityTypes(), resp.GetNextPageToken(), nil
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

func (c *entityTypesGRPCClient) GetEntityType(ctx context.Context, req *cxpb.GetEntityTypeRequest, opts ...gax.CallOption) (*cxpb.EntityType, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetEntityType[0:len((*c.CallOptions).GetEntityType):len((*c.CallOptions).GetEntityType)], opts...)
	var resp *cxpb.EntityType
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.entityTypesClient.GetEntityType(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *entityTypesGRPCClient) CreateEntityType(ctx context.Context, req *cxpb.CreateEntityTypeRequest, opts ...gax.CallOption) (*cxpb.EntityType, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateEntityType[0:len((*c.CallOptions).CreateEntityType):len((*c.CallOptions).CreateEntityType)], opts...)
	var resp *cxpb.EntityType
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.entityTypesClient.CreateEntityType(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *entityTypesGRPCClient) UpdateEntityType(ctx context.Context, req *cxpb.UpdateEntityTypeRequest, opts ...gax.CallOption) (*cxpb.EntityType, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "entity_type.name", url.QueryEscape(req.GetEntityType().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateEntityType[0:len((*c.CallOptions).UpdateEntityType):len((*c.CallOptions).UpdateEntityType)], opts...)
	var resp *cxpb.EntityType
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.entityTypesClient.UpdateEntityType(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *entityTypesGRPCClient) DeleteEntityType(ctx context.Context, req *cxpb.DeleteEntityTypeRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteEntityType[0:len((*c.CallOptions).DeleteEntityType):len((*c.CallOptions).DeleteEntityType)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.entityTypesClient.DeleteEntityType(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// EntityTypeIterator manages a stream of *cxpb.EntityType.
type EntityTypeIterator struct {
	items    []*cxpb.EntityType
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
	InternalFetch func(pageSize int, pageToken string) (results []*cxpb.EntityType, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *EntityTypeIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *EntityTypeIterator) Next() (*cxpb.EntityType, error) {
	var item *cxpb.EntityType
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *EntityTypeIterator) bufLen() int {
	return len(it.items)
}

func (it *EntityTypeIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
