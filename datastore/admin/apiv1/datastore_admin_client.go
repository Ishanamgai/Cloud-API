// Copyright 2020 Google LLC
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

package admin

import (
	"context"
	"math"
	"time"

	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	"github.com/golang/protobuf/proto"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	gtransport "google.golang.org/api/transport/grpc"
	adminpb "google.golang.org/genproto/googleapis/datastore/admin/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newDatastoreAdminClientHook clientHook

// DatastoreAdminCallOptions contains the retry settings for each method of DatastoreAdminClient.
type DatastoreAdminCallOptions struct {
	ExportEntities []gax.CallOption
	ImportEntities []gax.CallOption
	GetIndex       []gax.CallOption
	ListIndexes    []gax.CallOption
}

func defaultDatastoreAdminClientOptions() []option.ClientOption {
	return []option.ClientOption{
		option.WithEndpoint("datastore.googleapis.com:443"),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultDatastoreAdminCallOptions() *DatastoreAdminCallOptions {
	return &DatastoreAdminCallOptions{
		ExportEntities: []gax.CallOption{},
		ImportEntities: []gax.CallOption{},
		GetIndex: []gax.CallOption{
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
		ListIndexes: []gax.CallOption{
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

// DatastoreAdminClient is a client for interacting with Cloud Datastore API.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type DatastoreAdminClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// The gRPC API client.
	datastoreAdminClient adminpb.DatastoreAdminClient

	// LROClient is used internally to handle longrunning operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient

	// The call options for this service.
	CallOptions *DatastoreAdminCallOptions

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewDatastoreAdminClient creates a new datastore admin client.
//
// Google Cloud Datastore Admin API
//
// The Datastore Admin API provides several admin services for Cloud Datastore.
//
// ConceptsProject, namespace, kind, and entity as defined in the Google Cloud Datastore
// API.
//
// Operation: An Operation represents work being performed in the background.
//
// EntityFilter: Allows specifying a subset of entities in a project. This is
// specified as a combination of kinds and namespaces (either or both of which
// may be all).
//
// ServicesExport/ImportThe Export/Import service provides the ability to copy all or a subset of
// entities to/from Google Cloud Storage.
//
// Exported data may be imported into Cloud Datastore for any Google Cloud
// Platform project. It is not restricted to the export source project. It is
// possible to export from one project and then import into another.
//
// Exported data can also be loaded into Google BigQuery for analysis.
//
// Exports and imports are performed asynchronously. An Operation resource is
// created for each export/import. The state (including any errors encountered)
// of the export/import may be queried via the Operation resource.
//
// IndexThe index service manages Cloud Datastore composite indexes.
//
// Index creation and deletion are performed asynchronously.
// An Operation resource is created for each such asynchronous operation.
// The state of the operation (including any errors encountered)
// may be queried via the Operation resource.
//
// OperationThe Operations collection provides a record of actions performed for the
// specified project (including any operations in progress). Operations are not
// created directly but through calls on other collections or resources.
//
// An operation that is not yet done may be cancelled. The request to cancel is
// asynchronous and the operation may continue to run for some time after the
// request to cancel is made.
//
// An operation that is done may be deleted so that it is no longer listed as
// part of the Operation collection.
//
// ListOperations returns all pending operations, but not completed operations.
//
// Operations are created by service DatastoreAdmin,
// but are accessed via service google.longrunning.Operations.
func NewDatastoreAdminClient(ctx context.Context, opts ...option.ClientOption) (*DatastoreAdminClient, error) {
	clientOpts := defaultDatastoreAdminClientOptions()

	if newDatastoreAdminClientHook != nil {
		hookOpts, err := newDatastoreAdminClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	c := &DatastoreAdminClient{
		connPool:    connPool,
		CallOptions: defaultDatastoreAdminCallOptions(),

		datastoreAdminClient: adminpb.NewDatastoreAdminClient(connPool),
	}
	c.setGoogleClientInfo()

	c.LROClient, err = lroauto.NewOperationsClient(ctx, gtransport.WithConnPool(connPool))
	if err != nil {
		// This error "should not happen", since we are just reusing old connection pool
		// and never actually need to dial.
		// If this does happen, we could leak connp. However, we cannot close conn:
		// If the user invoked the constructor with option.WithGRPCConn,
		// we would close a connection that's still in use.
		// TODO: investigate error conditions.
		return nil, err
	}
	return c, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *DatastoreAdminClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *DatastoreAdminClient) Close() error {
	return c.connPool.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *DatastoreAdminClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// ExportEntities exports a copy of all or a subset of entities from Google Cloud Datastore
// to another storage system, such as Google Cloud Storage. Recent updates to
// entities may not be reflected in the export. The export occurs in the
// background and its progress can be monitored and managed via the
// Operation resource that is created. The output of an export may only be
// used once the associated operation is done. If an export operation is
// cancelled before completion it may leave partial data behind in Google
// Cloud Storage.
func (c *DatastoreAdminClient) ExportEntities(ctx context.Context, req *adminpb.ExportEntitiesRequest, opts ...gax.CallOption) (*ExportEntitiesOperation, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.ExportEntities[0:len(c.CallOptions.ExportEntities):len(c.CallOptions.ExportEntities)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.datastoreAdminClient.ExportEntities(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &ExportEntitiesOperation{
		lro: longrunning.InternalNewOperation(c.LROClient, resp),
	}, nil
}

// ImportEntities imports entities into Google Cloud Datastore. Existing entities with the
// same key are overwritten. The import occurs in the background and its
// progress can be monitored and managed via the Operation resource that is
// created. If an ImportEntities operation is cancelled, it is possible
// that a subset of the data has already been imported to Cloud Datastore.
func (c *DatastoreAdminClient) ImportEntities(ctx context.Context, req *adminpb.ImportEntitiesRequest, opts ...gax.CallOption) (*ImportEntitiesOperation, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.ImportEntities[0:len(c.CallOptions.ImportEntities):len(c.CallOptions.ImportEntities)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.datastoreAdminClient.ImportEntities(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &ImportEntitiesOperation{
		lro: longrunning.InternalNewOperation(c.LROClient, resp),
	}, nil
}

// GetIndex gets an index.
func (c *DatastoreAdminClient) GetIndex(ctx context.Context, req *adminpb.GetIndexRequest, opts ...gax.CallOption) (*adminpb.Index, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.GetIndex[0:len(c.CallOptions.GetIndex):len(c.CallOptions.GetIndex)], opts...)
	var resp *adminpb.Index
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.datastoreAdminClient.GetIndex(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ListIndexes lists the indexes that match the specified filters.  Datastore uses an
// eventually consistent query to fetch the list of indexes and may
// occasionally return stale results.
func (c *DatastoreAdminClient) ListIndexes(ctx context.Context, req *adminpb.ListIndexesRequest, opts ...gax.CallOption) *IndexIterator {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.ListIndexes[0:len(c.CallOptions.ListIndexes):len(c.CallOptions.ListIndexes)], opts...)
	it := &IndexIterator{}
	req = proto.Clone(req).(*adminpb.ListIndexesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*adminpb.Index, string, error) {
		var resp *adminpb.ListIndexesResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.datastoreAdminClient.ListIndexes(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.Indexes, resp.NextPageToken, nil
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
	it.pageInfo.MaxSize = int(req.PageSize)
	it.pageInfo.Token = req.PageToken
	return it
}

// ExportEntitiesOperation manages a long-running operation from ExportEntities.
type ExportEntitiesOperation struct {
	lro *longrunning.Operation
}

// ExportEntitiesOperation returns a new ExportEntitiesOperation from a given name.
// The name must be that of a previously created ExportEntitiesOperation, possibly from a different process.
func (c *DatastoreAdminClient) ExportEntitiesOperation(name string) *ExportEntitiesOperation {
	return &ExportEntitiesOperation{
		lro: longrunning.InternalNewOperation(c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *ExportEntitiesOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*adminpb.ExportEntitiesResponse, error) {
	var resp adminpb.ExportEntitiesResponse
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
func (op *ExportEntitiesOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*adminpb.ExportEntitiesResponse, error) {
	var resp adminpb.ExportEntitiesResponse
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
func (op *ExportEntitiesOperation) Metadata() (*adminpb.ExportEntitiesMetadata, error) {
	var meta adminpb.ExportEntitiesMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *ExportEntitiesOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *ExportEntitiesOperation) Name() string {
	return op.lro.Name()
}

// ImportEntitiesOperation manages a long-running operation from ImportEntities.
type ImportEntitiesOperation struct {
	lro *longrunning.Operation
}

// ImportEntitiesOperation returns a new ImportEntitiesOperation from a given name.
// The name must be that of a previously created ImportEntitiesOperation, possibly from a different process.
func (c *DatastoreAdminClient) ImportEntitiesOperation(name string) *ImportEntitiesOperation {
	return &ImportEntitiesOperation{
		lro: longrunning.InternalNewOperation(c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *ImportEntitiesOperation) Wait(ctx context.Context, opts ...gax.CallOption) error {
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
func (op *ImportEntitiesOperation) Poll(ctx context.Context, opts ...gax.CallOption) error {
	return op.lro.Poll(ctx, nil, opts...)
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *ImportEntitiesOperation) Metadata() (*adminpb.ImportEntitiesMetadata, error) {
	var meta adminpb.ImportEntitiesMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *ImportEntitiesOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *ImportEntitiesOperation) Name() string {
	return op.lro.Name()
}

// IndexIterator manages a stream of *adminpb.Index.
type IndexIterator struct {
	items    []*adminpb.Index
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
	InternalFetch func(pageSize int, pageToken string) (results []*adminpb.Index, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *IndexIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *IndexIterator) Next() (*adminpb.Index, error) {
	var item *adminpb.Index
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *IndexIterator) bufLen() int {
	return len(it.items)
}

func (it *IndexIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
