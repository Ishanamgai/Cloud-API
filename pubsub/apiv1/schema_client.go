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

package pubsub

import (
	"context"
	"fmt"
	"math"
	"net/url"

	"github.com/golang/protobuf/proto"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	pubsubpb "google.golang.org/genproto/googleapis/pubsub/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var newSchemaClientHook clientHook

// SchemaCallOptions contains the retry settings for each method of SchemaClient.
type SchemaCallOptions struct {
	CreateSchema    []gax.CallOption
	GetSchema       []gax.CallOption
	ListSchemas     []gax.CallOption
	DeleteSchema    []gax.CallOption
	ValidateSchema  []gax.CallOption
	ValidateMessage []gax.CallOption
}

func defaultSchemaGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("pubsub.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("pubsub.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://pubsub.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultSchemaCallOptions() *SchemaCallOptions {
	return &SchemaCallOptions{
		CreateSchema:    []gax.CallOption{},
		GetSchema:       []gax.CallOption{},
		ListSchemas:     []gax.CallOption{},
		DeleteSchema:    []gax.CallOption{},
		ValidateSchema:  []gax.CallOption{},
		ValidateMessage: []gax.CallOption{},
	}
}

// internalSchemaClient is an interface that defines the methods availaible from Cloud Pub/Sub API.
type internalSchemaClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateSchema(context.Context, *pubsubpb.CreateSchemaRequest, ...gax.CallOption) (*pubsubpb.Schema, error)
	GetSchema(context.Context, *pubsubpb.GetSchemaRequest, ...gax.CallOption) (*pubsubpb.Schema, error)
	ListSchemas(context.Context, *pubsubpb.ListSchemasRequest, ...gax.CallOption) *SchemaIterator
	DeleteSchema(context.Context, *pubsubpb.DeleteSchemaRequest, ...gax.CallOption) error
	ValidateSchema(context.Context, *pubsubpb.ValidateSchemaRequest, ...gax.CallOption) (*pubsubpb.ValidateSchemaResponse, error)
	ValidateMessage(context.Context, *pubsubpb.ValidateMessageRequest, ...gax.CallOption) (*pubsubpb.ValidateMessageResponse, error)
}

// SchemaClient is a client for interacting with Cloud Pub/Sub API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service for doing schema-related operations.
type SchemaClient struct {
	// The internal transport-dependent client.
	internalClient internalSchemaClient

	// The call options for this service.
	CallOptions *SchemaCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *SchemaClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *SchemaClient) setGoogleClientInfo(...string) {
	c.internalClient.setGoogleClientInfo()
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *SchemaClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateSchema creates a schema.
func (c *SchemaClient) CreateSchema(ctx context.Context, req *pubsubpb.CreateSchemaRequest, opts ...gax.CallOption) (*pubsubpb.Schema, error) {
	return c.internalClient.CreateSchema(ctx, req, opts...)
}

// GetSchema gets a schema.
func (c *SchemaClient) GetSchema(ctx context.Context, req *pubsubpb.GetSchemaRequest, opts ...gax.CallOption) (*pubsubpb.Schema, error) {
	return c.internalClient.GetSchema(ctx, req, opts...)
}

// ListSchemas lists schemas in a project.
func (c *SchemaClient) ListSchemas(ctx context.Context, req *pubsubpb.ListSchemasRequest, opts ...gax.CallOption) *SchemaIterator {
	return c.internalClient.ListSchemas(ctx, req, opts...)
}

// DeleteSchema deletes a schema.
func (c *SchemaClient) DeleteSchema(ctx context.Context, req *pubsubpb.DeleteSchemaRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteSchema(ctx, req, opts...)
}

// ValidateSchema validates a schema.
func (c *SchemaClient) ValidateSchema(ctx context.Context, req *pubsubpb.ValidateSchemaRequest, opts ...gax.CallOption) (*pubsubpb.ValidateSchemaResponse, error) {
	return c.internalClient.ValidateSchema(ctx, req, opts...)
}

// ValidateMessage validates a message against a schema.
func (c *SchemaClient) ValidateMessage(ctx context.Context, req *pubsubpb.ValidateMessageRequest, opts ...gax.CallOption) (*pubsubpb.ValidateMessageResponse, error) {
	return c.internalClient.ValidateMessage(ctx, req, opts...)
}

// schemaGRPCClient is a client for interacting with Cloud Pub/Sub API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type schemaGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing SchemaClient
	CallOptions **SchemaCallOptions

	// The gRPC API client.
	schemaClient pubsubpb.SchemaServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewSchemaClient creates a new schema service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service for doing schema-related operations.
func NewSchemaClient(ctx context.Context, opts ...option.ClientOption) (*SchemaClient, error) {
	clientOpts := defaultSchemaGRPCClientOptions()
	if newSchemaClientHook != nil {
		hookOpts, err := newSchemaClientHook(ctx, clientHookParams{})
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
	client := SchemaClient{CallOptions: defaultSchemaCallOptions()}

	c := &schemaGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		schemaClient:     pubsubpb.NewSchemaServiceClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *schemaGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *schemaGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *schemaGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *schemaGRPCClient) CreateSchema(ctx context.Context, req *pubsubpb.CreateSchemaRequest, opts ...gax.CallOption) (*pubsubpb.Schema, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateSchema[0:len((*c.CallOptions).CreateSchema):len((*c.CallOptions).CreateSchema)], opts...)
	var resp *pubsubpb.Schema
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.schemaClient.CreateSchema(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *schemaGRPCClient) GetSchema(ctx context.Context, req *pubsubpb.GetSchemaRequest, opts ...gax.CallOption) (*pubsubpb.Schema, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetSchema[0:len((*c.CallOptions).GetSchema):len((*c.CallOptions).GetSchema)], opts...)
	var resp *pubsubpb.Schema
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.schemaClient.GetSchema(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *schemaGRPCClient) ListSchemas(ctx context.Context, req *pubsubpb.ListSchemasRequest, opts ...gax.CallOption) *SchemaIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListSchemas[0:len((*c.CallOptions).ListSchemas):len((*c.CallOptions).ListSchemas)], opts...)
	it := &SchemaIterator{}
	req = proto.Clone(req).(*pubsubpb.ListSchemasRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*pubsubpb.Schema, string, error) {
		var resp *pubsubpb.ListSchemasResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.schemaClient.ListSchemas(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetSchemas(), resp.GetNextPageToken(), nil
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

func (c *schemaGRPCClient) DeleteSchema(ctx context.Context, req *pubsubpb.DeleteSchemaRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteSchema[0:len((*c.CallOptions).DeleteSchema):len((*c.CallOptions).DeleteSchema)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.schemaClient.DeleteSchema(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *schemaGRPCClient) ValidateSchema(ctx context.Context, req *pubsubpb.ValidateSchemaRequest, opts ...gax.CallOption) (*pubsubpb.ValidateSchemaResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ValidateSchema[0:len((*c.CallOptions).ValidateSchema):len((*c.CallOptions).ValidateSchema)], opts...)
	var resp *pubsubpb.ValidateSchemaResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.schemaClient.ValidateSchema(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *schemaGRPCClient) ValidateMessage(ctx context.Context, req *pubsubpb.ValidateMessageRequest, opts ...gax.CallOption) (*pubsubpb.ValidateMessageResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ValidateMessage[0:len((*c.CallOptions).ValidateMessage):len((*c.CallOptions).ValidateMessage)], opts...)
	var resp *pubsubpb.ValidateMessageResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.schemaClient.ValidateMessage(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SchemaIterator manages a stream of *pubsubpb.Schema.
type SchemaIterator struct {
	items    []*pubsubpb.Schema
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
	InternalFetch func(pageSize int, pageToken string) (results []*pubsubpb.Schema, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *SchemaIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *SchemaIterator) Next() (*pubsubpb.Schema, error) {
	var item *pubsubpb.Schema
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *SchemaIterator) bufLen() int {
	return len(it.items)
}

func (it *SchemaIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
