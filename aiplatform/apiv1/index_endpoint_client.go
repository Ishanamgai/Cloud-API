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

package aiplatform

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	aiplatformpb "google.golang.org/genproto/googleapis/cloud/aiplatform/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newIndexEndpointClientHook clientHook

// IndexEndpointCallOptions contains the retry settings for each method of IndexEndpointClient.
type IndexEndpointCallOptions struct {
	CreateIndexEndpoint []gax.CallOption
	GetIndexEndpoint    []gax.CallOption
	ListIndexEndpoints  []gax.CallOption
	UpdateIndexEndpoint []gax.CallOption
	DeleteIndexEndpoint []gax.CallOption
	DeployIndex         []gax.CallOption
	UndeployIndex       []gax.CallOption
	MutateDeployedIndex []gax.CallOption
}

func defaultIndexEndpointGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("aiplatform.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("aiplatform.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://aiplatform.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultIndexEndpointCallOptions() *IndexEndpointCallOptions {
	return &IndexEndpointCallOptions{
		CreateIndexEndpoint: []gax.CallOption{},
		GetIndexEndpoint:    []gax.CallOption{},
		ListIndexEndpoints:  []gax.CallOption{},
		UpdateIndexEndpoint: []gax.CallOption{},
		DeleteIndexEndpoint: []gax.CallOption{},
		DeployIndex:         []gax.CallOption{},
		UndeployIndex:       []gax.CallOption{},
		MutateDeployedIndex: []gax.CallOption{},
	}
}

// internalIndexEndpointClient is an interface that defines the methods availaible from Vertex AI API.
type internalIndexEndpointClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateIndexEndpoint(context.Context, *aiplatformpb.CreateIndexEndpointRequest, ...gax.CallOption) (*CreateIndexEndpointOperation, error)
	CreateIndexEndpointOperation(name string) *CreateIndexEndpointOperation
	GetIndexEndpoint(context.Context, *aiplatformpb.GetIndexEndpointRequest, ...gax.CallOption) (*aiplatformpb.IndexEndpoint, error)
	ListIndexEndpoints(context.Context, *aiplatformpb.ListIndexEndpointsRequest, ...gax.CallOption) *IndexEndpointIterator
	UpdateIndexEndpoint(context.Context, *aiplatformpb.UpdateIndexEndpointRequest, ...gax.CallOption) (*aiplatformpb.IndexEndpoint, error)
	DeleteIndexEndpoint(context.Context, *aiplatformpb.DeleteIndexEndpointRequest, ...gax.CallOption) (*DeleteIndexEndpointOperation, error)
	DeleteIndexEndpointOperation(name string) *DeleteIndexEndpointOperation
	DeployIndex(context.Context, *aiplatformpb.DeployIndexRequest, ...gax.CallOption) (*DeployIndexOperation, error)
	DeployIndexOperation(name string) *DeployIndexOperation
	UndeployIndex(context.Context, *aiplatformpb.UndeployIndexRequest, ...gax.CallOption) (*UndeployIndexOperation, error)
	UndeployIndexOperation(name string) *UndeployIndexOperation
	MutateDeployedIndex(context.Context, *aiplatformpb.MutateDeployedIndexRequest, ...gax.CallOption) (*MutateDeployedIndexOperation, error)
	MutateDeployedIndexOperation(name string) *MutateDeployedIndexOperation
}

// IndexEndpointClient is a client for interacting with Vertex AI API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// A service for managing Vertex AI’s IndexEndpoints.
type IndexEndpointClient struct {
	// The internal transport-dependent client.
	internalClient internalIndexEndpointClient

	// The call options for this service.
	CallOptions *IndexEndpointCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *IndexEndpointClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *IndexEndpointClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *IndexEndpointClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateIndexEndpoint creates an IndexEndpoint.
func (c *IndexEndpointClient) CreateIndexEndpoint(ctx context.Context, req *aiplatformpb.CreateIndexEndpointRequest, opts ...gax.CallOption) (*CreateIndexEndpointOperation, error) {
	return c.internalClient.CreateIndexEndpoint(ctx, req, opts...)
}

// CreateIndexEndpointOperation returns a new CreateIndexEndpointOperation from a given name.
// The name must be that of a previously created CreateIndexEndpointOperation, possibly from a different process.
func (c *IndexEndpointClient) CreateIndexEndpointOperation(name string) *CreateIndexEndpointOperation {
	return c.internalClient.CreateIndexEndpointOperation(name)
}

// GetIndexEndpoint gets an IndexEndpoint.
func (c *IndexEndpointClient) GetIndexEndpoint(ctx context.Context, req *aiplatformpb.GetIndexEndpointRequest, opts ...gax.CallOption) (*aiplatformpb.IndexEndpoint, error) {
	return c.internalClient.GetIndexEndpoint(ctx, req, opts...)
}

// ListIndexEndpoints lists IndexEndpoints in a Location.
func (c *IndexEndpointClient) ListIndexEndpoints(ctx context.Context, req *aiplatformpb.ListIndexEndpointsRequest, opts ...gax.CallOption) *IndexEndpointIterator {
	return c.internalClient.ListIndexEndpoints(ctx, req, opts...)
}

// UpdateIndexEndpoint updates an IndexEndpoint.
func (c *IndexEndpointClient) UpdateIndexEndpoint(ctx context.Context, req *aiplatformpb.UpdateIndexEndpointRequest, opts ...gax.CallOption) (*aiplatformpb.IndexEndpoint, error) {
	return c.internalClient.UpdateIndexEndpoint(ctx, req, opts...)
}

// DeleteIndexEndpoint deletes an IndexEndpoint.
func (c *IndexEndpointClient) DeleteIndexEndpoint(ctx context.Context, req *aiplatformpb.DeleteIndexEndpointRequest, opts ...gax.CallOption) (*DeleteIndexEndpointOperation, error) {
	return c.internalClient.DeleteIndexEndpoint(ctx, req, opts...)
}

// DeleteIndexEndpointOperation returns a new DeleteIndexEndpointOperation from a given name.
// The name must be that of a previously created DeleteIndexEndpointOperation, possibly from a different process.
func (c *IndexEndpointClient) DeleteIndexEndpointOperation(name string) *DeleteIndexEndpointOperation {
	return c.internalClient.DeleteIndexEndpointOperation(name)
}

// DeployIndex deploys an Index into this IndexEndpoint, creating a DeployedIndex within
// it.
// Only non-empty Indexes can be deployed.
func (c *IndexEndpointClient) DeployIndex(ctx context.Context, req *aiplatformpb.DeployIndexRequest, opts ...gax.CallOption) (*DeployIndexOperation, error) {
	return c.internalClient.DeployIndex(ctx, req, opts...)
}

// DeployIndexOperation returns a new DeployIndexOperation from a given name.
// The name must be that of a previously created DeployIndexOperation, possibly from a different process.
func (c *IndexEndpointClient) DeployIndexOperation(name string) *DeployIndexOperation {
	return c.internalClient.DeployIndexOperation(name)
}

// UndeployIndex undeploys an Index from an IndexEndpoint, removing a DeployedIndex from it,
// and freeing all resources it’s using.
func (c *IndexEndpointClient) UndeployIndex(ctx context.Context, req *aiplatformpb.UndeployIndexRequest, opts ...gax.CallOption) (*UndeployIndexOperation, error) {
	return c.internalClient.UndeployIndex(ctx, req, opts...)
}

// UndeployIndexOperation returns a new UndeployIndexOperation from a given name.
// The name must be that of a previously created UndeployIndexOperation, possibly from a different process.
func (c *IndexEndpointClient) UndeployIndexOperation(name string) *UndeployIndexOperation {
	return c.internalClient.UndeployIndexOperation(name)
}

// MutateDeployedIndex update an existing DeployedIndex under an IndexEndpoint.
func (c *IndexEndpointClient) MutateDeployedIndex(ctx context.Context, req *aiplatformpb.MutateDeployedIndexRequest, opts ...gax.CallOption) (*MutateDeployedIndexOperation, error) {
	return c.internalClient.MutateDeployedIndex(ctx, req, opts...)
}

// MutateDeployedIndexOperation returns a new MutateDeployedIndexOperation from a given name.
// The name must be that of a previously created MutateDeployedIndexOperation, possibly from a different process.
func (c *IndexEndpointClient) MutateDeployedIndexOperation(name string) *MutateDeployedIndexOperation {
	return c.internalClient.MutateDeployedIndexOperation(name)
}

// indexEndpointGRPCClient is a client for interacting with Vertex AI API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type indexEndpointGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing IndexEndpointClient
	CallOptions **IndexEndpointCallOptions

	// The gRPC API client.
	indexEndpointClient aiplatformpb.IndexEndpointServiceClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewIndexEndpointClient creates a new index endpoint service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// A service for managing Vertex AI’s IndexEndpoints.
func NewIndexEndpointClient(ctx context.Context, opts ...option.ClientOption) (*IndexEndpointClient, error) {
	clientOpts := defaultIndexEndpointGRPCClientOptions()
	if newIndexEndpointClientHook != nil {
		hookOpts, err := newIndexEndpointClientHook(ctx, clientHookParams{})
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
	client := IndexEndpointClient{CallOptions: defaultIndexEndpointCallOptions()}

	c := &indexEndpointGRPCClient{
		connPool:            connPool,
		disableDeadlines:    disableDeadlines,
		indexEndpointClient: aiplatformpb.NewIndexEndpointServiceClient(connPool),
		CallOptions:         &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	client.LROClient, err = lroauto.NewOperationsClient(ctx, gtransport.WithConnPool(connPool))
	if err != nil {
		// This error "should not happen", since we are just reusing old connection pool
		// and never actually need to dial.
		// If this does happen, we could leak connp. However, we cannot close conn:
		// If the user invoked the constructor with option.WithGRPCConn,
		// we would close a connection that's still in use.
		// TODO: investigate error conditions.
		return nil, err
	}
	c.LROClient = &client.LROClient
	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *indexEndpointGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *indexEndpointGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *indexEndpointGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *indexEndpointGRPCClient) CreateIndexEndpoint(ctx context.Context, req *aiplatformpb.CreateIndexEndpointRequest, opts ...gax.CallOption) (*CreateIndexEndpointOperation, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateIndexEndpoint[0:len((*c.CallOptions).CreateIndexEndpoint):len((*c.CallOptions).CreateIndexEndpoint)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.indexEndpointClient.CreateIndexEndpoint(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &CreateIndexEndpointOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *indexEndpointGRPCClient) GetIndexEndpoint(ctx context.Context, req *aiplatformpb.GetIndexEndpointRequest, opts ...gax.CallOption) (*aiplatformpb.IndexEndpoint, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetIndexEndpoint[0:len((*c.CallOptions).GetIndexEndpoint):len((*c.CallOptions).GetIndexEndpoint)], opts...)
	var resp *aiplatformpb.IndexEndpoint
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.indexEndpointClient.GetIndexEndpoint(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *indexEndpointGRPCClient) ListIndexEndpoints(ctx context.Context, req *aiplatformpb.ListIndexEndpointsRequest, opts ...gax.CallOption) *IndexEndpointIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListIndexEndpoints[0:len((*c.CallOptions).ListIndexEndpoints):len((*c.CallOptions).ListIndexEndpoints)], opts...)
	it := &IndexEndpointIterator{}
	req = proto.Clone(req).(*aiplatformpb.ListIndexEndpointsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*aiplatformpb.IndexEndpoint, string, error) {
		resp := &aiplatformpb.ListIndexEndpointsResponse{}
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
			resp, err = c.indexEndpointClient.ListIndexEndpoints(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetIndexEndpoints(), resp.GetNextPageToken(), nil
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

func (c *indexEndpointGRPCClient) UpdateIndexEndpoint(ctx context.Context, req *aiplatformpb.UpdateIndexEndpointRequest, opts ...gax.CallOption) (*aiplatformpb.IndexEndpoint, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "index_endpoint.name", url.QueryEscape(req.GetIndexEndpoint().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateIndexEndpoint[0:len((*c.CallOptions).UpdateIndexEndpoint):len((*c.CallOptions).UpdateIndexEndpoint)], opts...)
	var resp *aiplatformpb.IndexEndpoint
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.indexEndpointClient.UpdateIndexEndpoint(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *indexEndpointGRPCClient) DeleteIndexEndpoint(ctx context.Context, req *aiplatformpb.DeleteIndexEndpointRequest, opts ...gax.CallOption) (*DeleteIndexEndpointOperation, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteIndexEndpoint[0:len((*c.CallOptions).DeleteIndexEndpoint):len((*c.CallOptions).DeleteIndexEndpoint)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.indexEndpointClient.DeleteIndexEndpoint(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &DeleteIndexEndpointOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *indexEndpointGRPCClient) DeployIndex(ctx context.Context, req *aiplatformpb.DeployIndexRequest, opts ...gax.CallOption) (*DeployIndexOperation, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "index_endpoint", url.QueryEscape(req.GetIndexEndpoint())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeployIndex[0:len((*c.CallOptions).DeployIndex):len((*c.CallOptions).DeployIndex)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.indexEndpointClient.DeployIndex(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &DeployIndexOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *indexEndpointGRPCClient) UndeployIndex(ctx context.Context, req *aiplatformpb.UndeployIndexRequest, opts ...gax.CallOption) (*UndeployIndexOperation, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "index_endpoint", url.QueryEscape(req.GetIndexEndpoint())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UndeployIndex[0:len((*c.CallOptions).UndeployIndex):len((*c.CallOptions).UndeployIndex)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.indexEndpointClient.UndeployIndex(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &UndeployIndexOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *indexEndpointGRPCClient) MutateDeployedIndex(ctx context.Context, req *aiplatformpb.MutateDeployedIndexRequest, opts ...gax.CallOption) (*MutateDeployedIndexOperation, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "index_endpoint", url.QueryEscape(req.GetIndexEndpoint())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).MutateDeployedIndex[0:len((*c.CallOptions).MutateDeployedIndex):len((*c.CallOptions).MutateDeployedIndex)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.indexEndpointClient.MutateDeployedIndex(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &MutateDeployedIndexOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

// CreateIndexEndpointOperation manages a long-running operation from CreateIndexEndpoint.
type CreateIndexEndpointOperation struct {
	lro *longrunning.Operation
}

// CreateIndexEndpointOperation returns a new CreateIndexEndpointOperation from a given name.
// The name must be that of a previously created CreateIndexEndpointOperation, possibly from a different process.
func (c *indexEndpointGRPCClient) CreateIndexEndpointOperation(name string) *CreateIndexEndpointOperation {
	return &CreateIndexEndpointOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *CreateIndexEndpointOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*aiplatformpb.IndexEndpoint, error) {
	var resp aiplatformpb.IndexEndpoint
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *CreateIndexEndpointOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*aiplatformpb.IndexEndpoint, error) {
	var resp aiplatformpb.IndexEndpoint
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *CreateIndexEndpointOperation) Metadata() (*aiplatformpb.CreateIndexEndpointOperationMetadata, error) {
	var meta aiplatformpb.CreateIndexEndpointOperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *CreateIndexEndpointOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *CreateIndexEndpointOperation) Name() string {
	return op.lro.Name()
}

// DeleteIndexEndpointOperation manages a long-running operation from DeleteIndexEndpoint.
type DeleteIndexEndpointOperation struct {
	lro *longrunning.Operation
}

// DeleteIndexEndpointOperation returns a new DeleteIndexEndpointOperation from a given name.
// The name must be that of a previously created DeleteIndexEndpointOperation, possibly from a different process.
func (c *indexEndpointGRPCClient) DeleteIndexEndpointOperation(name string) *DeleteIndexEndpointOperation {
	return &DeleteIndexEndpointOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *DeleteIndexEndpointOperation) Wait(ctx context.Context, opts ...gax.CallOption) error {
	return op.lro.WaitWithInterval(ctx, nil, time.Minute, opts...)
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *DeleteIndexEndpointOperation) Poll(ctx context.Context, opts ...gax.CallOption) error {
	return op.lro.Poll(ctx, nil, opts...)
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *DeleteIndexEndpointOperation) Metadata() (*aiplatformpb.DeleteOperationMetadata, error) {
	var meta aiplatformpb.DeleteOperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *DeleteIndexEndpointOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *DeleteIndexEndpointOperation) Name() string {
	return op.lro.Name()
}

// DeployIndexOperation manages a long-running operation from DeployIndex.
type DeployIndexOperation struct {
	lro *longrunning.Operation
}

// DeployIndexOperation returns a new DeployIndexOperation from a given name.
// The name must be that of a previously created DeployIndexOperation, possibly from a different process.
func (c *indexEndpointGRPCClient) DeployIndexOperation(name string) *DeployIndexOperation {
	return &DeployIndexOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *DeployIndexOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*aiplatformpb.DeployIndexResponse, error) {
	var resp aiplatformpb.DeployIndexResponse
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *DeployIndexOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*aiplatformpb.DeployIndexResponse, error) {
	var resp aiplatformpb.DeployIndexResponse
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *DeployIndexOperation) Metadata() (*aiplatformpb.DeployIndexOperationMetadata, error) {
	var meta aiplatformpb.DeployIndexOperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *DeployIndexOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *DeployIndexOperation) Name() string {
	return op.lro.Name()
}

// MutateDeployedIndexOperation manages a long-running operation from MutateDeployedIndex.
type MutateDeployedIndexOperation struct {
	lro *longrunning.Operation
}

// MutateDeployedIndexOperation returns a new MutateDeployedIndexOperation from a given name.
// The name must be that of a previously created MutateDeployedIndexOperation, possibly from a different process.
func (c *indexEndpointGRPCClient) MutateDeployedIndexOperation(name string) *MutateDeployedIndexOperation {
	return &MutateDeployedIndexOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *MutateDeployedIndexOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*aiplatformpb.MutateDeployedIndexResponse, error) {
	var resp aiplatformpb.MutateDeployedIndexResponse
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *MutateDeployedIndexOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*aiplatformpb.MutateDeployedIndexResponse, error) {
	var resp aiplatformpb.MutateDeployedIndexResponse
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *MutateDeployedIndexOperation) Metadata() (*aiplatformpb.MutateDeployedIndexOperationMetadata, error) {
	var meta aiplatformpb.MutateDeployedIndexOperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *MutateDeployedIndexOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *MutateDeployedIndexOperation) Name() string {
	return op.lro.Name()
}

// UndeployIndexOperation manages a long-running operation from UndeployIndex.
type UndeployIndexOperation struct {
	lro *longrunning.Operation
}

// UndeployIndexOperation returns a new UndeployIndexOperation from a given name.
// The name must be that of a previously created UndeployIndexOperation, possibly from a different process.
func (c *indexEndpointGRPCClient) UndeployIndexOperation(name string) *UndeployIndexOperation {
	return &UndeployIndexOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *UndeployIndexOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*aiplatformpb.UndeployIndexResponse, error) {
	var resp aiplatformpb.UndeployIndexResponse
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *UndeployIndexOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*aiplatformpb.UndeployIndexResponse, error) {
	var resp aiplatformpb.UndeployIndexResponse
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *UndeployIndexOperation) Metadata() (*aiplatformpb.UndeployIndexOperationMetadata, error) {
	var meta aiplatformpb.UndeployIndexOperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *UndeployIndexOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *UndeployIndexOperation) Name() string {
	return op.lro.Name()
}

// IndexEndpointIterator manages a stream of *aiplatformpb.IndexEndpoint.
type IndexEndpointIterator struct {
	items    []*aiplatformpb.IndexEndpoint
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
	InternalFetch func(pageSize int, pageToken string) (results []*aiplatformpb.IndexEndpoint, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *IndexEndpointIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *IndexEndpointIterator) Next() (*aiplatformpb.IndexEndpoint, error) {
	var item *aiplatformpb.IndexEndpoint
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *IndexEndpointIterator) bufLen() int {
	return len(it.items)
}

func (it *IndexEndpointIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
