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

package redis

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	"github.com/golang/protobuf/proto"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	gtransport "google.golang.org/api/transport/grpc"
	redispb "google.golang.org/genproto/googleapis/cloud/redis/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var newCloudRedisClientHook clientHook

// CloudRedisCallOptions contains the retry settings for each method of CloudRedisClient.
type CloudRedisCallOptions struct {
	ListInstances    []gax.CallOption
	GetInstance      []gax.CallOption
	CreateInstance   []gax.CallOption
	UpdateInstance   []gax.CallOption
	UpgradeInstance  []gax.CallOption
	ImportInstance   []gax.CallOption
	ExportInstance   []gax.CallOption
	FailoverInstance []gax.CallOption
	DeleteInstance   []gax.CallOption
}

func defaultCloudRedisClientOptions() []option.ClientOption {
	return []option.ClientOption{
		option.WithEndpoint("redis.googleapis.com:443"),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCloudRedisCallOptions() *CloudRedisCallOptions {
	return &CloudRedisCallOptions{
		ListInstances:    []gax.CallOption{},
		GetInstance:      []gax.CallOption{},
		CreateInstance:   []gax.CallOption{},
		UpdateInstance:   []gax.CallOption{},
		UpgradeInstance:  []gax.CallOption{},
		ImportInstance:   []gax.CallOption{},
		ExportInstance:   []gax.CallOption{},
		FailoverInstance: []gax.CallOption{},
		DeleteInstance:   []gax.CallOption{},
	}
}

// CloudRedisClient is a client for interacting with Google Cloud Memorystore for Redis API.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type CloudRedisClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// The gRPC API client.
	cloudRedisClient redispb.CloudRedisClient

	// LROClient is used internally to handle longrunning operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient

	// The call options for this service.
	CallOptions *CloudRedisCallOptions

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewCloudRedisClient creates a new cloud redis client.
//
// Configures and manages Cloud Memorystore for Redis instances
//
// Google Cloud Memorystore for Redis v1
//
// The redis.googleapis.com service implements the Google Cloud Memorystore
// for Redis API and defines the following resource model for managing Redis
// instances:
//
//   The service works with a collection of cloud projects, named: /projects/*
//
//   Each project has a collection of available locations, named: /locations/*
//
//   Each location has a collection of Redis instances, named: /instances/*
//
//   As such, Redis instances are resources of the form:
//   /projects/{project_id}/locations/{location_id}/instances/{instance_id}
//
// Note that location_id must be referring to a GCP region; for example:
//
//   projects/redpepper-1290/locations/us-central1/instances/my-redis
func NewCloudRedisClient(ctx context.Context, opts ...option.ClientOption) (*CloudRedisClient, error) {
	clientOpts := defaultCloudRedisClientOptions()

	if newCloudRedisClientHook != nil {
		hookOpts, err := newCloudRedisClientHook(ctx, clientHookParams{})
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
	c := &CloudRedisClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		CallOptions:      defaultCloudRedisCallOptions(),

		cloudRedisClient: redispb.NewCloudRedisClient(connPool),
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
func (c *CloudRedisClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *CloudRedisClient) Close() error {
	return c.connPool.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *CloudRedisClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// ListInstances lists all Redis instances owned by a project in either the specified
// location (region) or all locations.
//
// The location should have the following format:
//
//   projects/{project_id}/locations/{location_id}
//
// If location_id is specified as - (wildcard), then all regions
// available to the project are queried, and the results are aggregated.
func (c *CloudRedisClient) ListInstances(ctx context.Context, req *redispb.ListInstancesRequest, opts ...gax.CallOption) *InstanceIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ListInstances[0:len(c.CallOptions.ListInstances):len(c.CallOptions.ListInstances)], opts...)
	it := &InstanceIterator{}
	req = proto.Clone(req).(*redispb.ListInstancesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*redispb.Instance, string, error) {
		var resp *redispb.ListInstancesResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.cloudRedisClient.ListInstances(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetInstances(), resp.GetNextPageToken(), nil
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

// GetInstance gets the details of a specific Redis instance.
func (c *CloudRedisClient) GetInstance(ctx context.Context, req *redispb.GetInstanceRequest, opts ...gax.CallOption) (*redispb.Instance, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.GetInstance[0:len(c.CallOptions.GetInstance):len(c.CallOptions.GetInstance)], opts...)
	var resp *redispb.Instance
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.cloudRedisClient.GetInstance(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreateInstance creates a Redis instance based on the specified tier and memory size.
//
// By default, the instance is accessible from the project’s
// default network (at https://cloud.google.com/vpc/docs/vpc).
//
// The creation is executed asynchronously and callers may check the returned
// operation to track its progress. Once the operation is completed the Redis
// instance will be fully functional. Completed longrunning.Operation will
// contain the new instance object in the response field.
//
// The returned operation is automatically deleted after a few hours, so there
// is no need to call DeleteOperation.
func (c *CloudRedisClient) CreateInstance(ctx context.Context, req *redispb.CreateInstanceRequest, opts ...gax.CallOption) (*CreateInstanceOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.CreateInstance[0:len(c.CallOptions.CreateInstance):len(c.CallOptions.CreateInstance)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.cloudRedisClient.CreateInstance(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &CreateInstanceOperation{
		lro: longrunning.InternalNewOperation(c.LROClient, resp),
	}, nil
}

// UpdateInstance updates the metadata and configuration of a specific Redis instance.
//
// Completed longrunning.Operation will contain the new instance object
// in the response field. The returned operation is automatically deleted
// after a few hours, so there is no need to call DeleteOperation.
func (c *CloudRedisClient) UpdateInstance(ctx context.Context, req *redispb.UpdateInstanceRequest, opts ...gax.CallOption) (*UpdateInstanceOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "instance.name", url.QueryEscape(req.GetInstance().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.UpdateInstance[0:len(c.CallOptions.UpdateInstance):len(c.CallOptions.UpdateInstance)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.cloudRedisClient.UpdateInstance(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &UpdateInstanceOperation{
		lro: longrunning.InternalNewOperation(c.LROClient, resp),
	}, nil
}

// UpgradeInstance upgrades Redis instance to the newer Redis version specified in the
// request.
func (c *CloudRedisClient) UpgradeInstance(ctx context.Context, req *redispb.UpgradeInstanceRequest, opts ...gax.CallOption) (*UpgradeInstanceOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.UpgradeInstance[0:len(c.CallOptions.UpgradeInstance):len(c.CallOptions.UpgradeInstance)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.cloudRedisClient.UpgradeInstance(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &UpgradeInstanceOperation{
		lro: longrunning.InternalNewOperation(c.LROClient, resp),
	}, nil
}

// ImportInstance import a Redis RDB snapshot file from Cloud Storage into a Redis instance.
//
// Redis may stop serving during this operation. Instance state will be
// IMPORTING for entire operation. When complete, the instance will contain
// only data from the imported file.
//
// The returned operation is automatically deleted after a few hours, so
// there is no need to call DeleteOperation.
func (c *CloudRedisClient) ImportInstance(ctx context.Context, req *redispb.ImportInstanceRequest, opts ...gax.CallOption) (*ImportInstanceOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ImportInstance[0:len(c.CallOptions.ImportInstance):len(c.CallOptions.ImportInstance)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.cloudRedisClient.ImportInstance(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &ImportInstanceOperation{
		lro: longrunning.InternalNewOperation(c.LROClient, resp),
	}, nil
}

// ExportInstance export Redis instance data into a Redis RDB format file in Cloud Storage.
//
// Redis will continue serving during this operation.
//
// The returned operation is automatically deleted after a few hours, so
// there is no need to call DeleteOperation.
func (c *CloudRedisClient) ExportInstance(ctx context.Context, req *redispb.ExportInstanceRequest, opts ...gax.CallOption) (*ExportInstanceOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ExportInstance[0:len(c.CallOptions.ExportInstance):len(c.CallOptions.ExportInstance)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.cloudRedisClient.ExportInstance(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &ExportInstanceOperation{
		lro: longrunning.InternalNewOperation(c.LROClient, resp),
	}, nil
}

// FailoverInstance initiates a failover of the master node to current replica node for a
// specific STANDARD tier Cloud Memorystore for Redis instance.
func (c *CloudRedisClient) FailoverInstance(ctx context.Context, req *redispb.FailoverInstanceRequest, opts ...gax.CallOption) (*FailoverInstanceOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.FailoverInstance[0:len(c.CallOptions.FailoverInstance):len(c.CallOptions.FailoverInstance)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.cloudRedisClient.FailoverInstance(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &FailoverInstanceOperation{
		lro: longrunning.InternalNewOperation(c.LROClient, resp),
	}, nil
}

// DeleteInstance deletes a specific Redis instance.  Instance stops serving and data is
// deleted.
func (c *CloudRedisClient) DeleteInstance(ctx context.Context, req *redispb.DeleteInstanceRequest, opts ...gax.CallOption) (*DeleteInstanceOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.DeleteInstance[0:len(c.CallOptions.DeleteInstance):len(c.CallOptions.DeleteInstance)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.cloudRedisClient.DeleteInstance(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &DeleteInstanceOperation{
		lro: longrunning.InternalNewOperation(c.LROClient, resp),
	}, nil
}

// CreateInstanceOperation manages a long-running operation from CreateInstance.
type CreateInstanceOperation struct {
	lro *longrunning.Operation
}

// CreateInstanceOperation returns a new CreateInstanceOperation from a given name.
// The name must be that of a previously created CreateInstanceOperation, possibly from a different process.
func (c *CloudRedisClient) CreateInstanceOperation(name string) *CreateInstanceOperation {
	return &CreateInstanceOperation{
		lro: longrunning.InternalNewOperation(c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *CreateInstanceOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*redispb.Instance, error) {
	var resp redispb.Instance
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
func (op *CreateInstanceOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*redispb.Instance, error) {
	var resp redispb.Instance
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
func (op *CreateInstanceOperation) Metadata() (*redispb.OperationMetadata, error) {
	var meta redispb.OperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *CreateInstanceOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *CreateInstanceOperation) Name() string {
	return op.lro.Name()
}

// DeleteInstanceOperation manages a long-running operation from DeleteInstance.
type DeleteInstanceOperation struct {
	lro *longrunning.Operation
}

// DeleteInstanceOperation returns a new DeleteInstanceOperation from a given name.
// The name must be that of a previously created DeleteInstanceOperation, possibly from a different process.
func (c *CloudRedisClient) DeleteInstanceOperation(name string) *DeleteInstanceOperation {
	return &DeleteInstanceOperation{
		lro: longrunning.InternalNewOperation(c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *DeleteInstanceOperation) Wait(ctx context.Context, opts ...gax.CallOption) error {
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
func (op *DeleteInstanceOperation) Poll(ctx context.Context, opts ...gax.CallOption) error {
	return op.lro.Poll(ctx, nil, opts...)
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *DeleteInstanceOperation) Metadata() (*redispb.OperationMetadata, error) {
	var meta redispb.OperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *DeleteInstanceOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *DeleteInstanceOperation) Name() string {
	return op.lro.Name()
}

// ExportInstanceOperation manages a long-running operation from ExportInstance.
type ExportInstanceOperation struct {
	lro *longrunning.Operation
}

// ExportInstanceOperation returns a new ExportInstanceOperation from a given name.
// The name must be that of a previously created ExportInstanceOperation, possibly from a different process.
func (c *CloudRedisClient) ExportInstanceOperation(name string) *ExportInstanceOperation {
	return &ExportInstanceOperation{
		lro: longrunning.InternalNewOperation(c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *ExportInstanceOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*redispb.Instance, error) {
	var resp redispb.Instance
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
func (op *ExportInstanceOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*redispb.Instance, error) {
	var resp redispb.Instance
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
func (op *ExportInstanceOperation) Metadata() (*redispb.OperationMetadata, error) {
	var meta redispb.OperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *ExportInstanceOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *ExportInstanceOperation) Name() string {
	return op.lro.Name()
}

// FailoverInstanceOperation manages a long-running operation from FailoverInstance.
type FailoverInstanceOperation struct {
	lro *longrunning.Operation
}

// FailoverInstanceOperation returns a new FailoverInstanceOperation from a given name.
// The name must be that of a previously created FailoverInstanceOperation, possibly from a different process.
func (c *CloudRedisClient) FailoverInstanceOperation(name string) *FailoverInstanceOperation {
	return &FailoverInstanceOperation{
		lro: longrunning.InternalNewOperation(c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *FailoverInstanceOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*redispb.Instance, error) {
	var resp redispb.Instance
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
func (op *FailoverInstanceOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*redispb.Instance, error) {
	var resp redispb.Instance
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
func (op *FailoverInstanceOperation) Metadata() (*redispb.OperationMetadata, error) {
	var meta redispb.OperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *FailoverInstanceOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *FailoverInstanceOperation) Name() string {
	return op.lro.Name()
}

// ImportInstanceOperation manages a long-running operation from ImportInstance.
type ImportInstanceOperation struct {
	lro *longrunning.Operation
}

// ImportInstanceOperation returns a new ImportInstanceOperation from a given name.
// The name must be that of a previously created ImportInstanceOperation, possibly from a different process.
func (c *CloudRedisClient) ImportInstanceOperation(name string) *ImportInstanceOperation {
	return &ImportInstanceOperation{
		lro: longrunning.InternalNewOperation(c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *ImportInstanceOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*redispb.Instance, error) {
	var resp redispb.Instance
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
func (op *ImportInstanceOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*redispb.Instance, error) {
	var resp redispb.Instance
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
func (op *ImportInstanceOperation) Metadata() (*redispb.OperationMetadata, error) {
	var meta redispb.OperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *ImportInstanceOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *ImportInstanceOperation) Name() string {
	return op.lro.Name()
}

// UpdateInstanceOperation manages a long-running operation from UpdateInstance.
type UpdateInstanceOperation struct {
	lro *longrunning.Operation
}

// UpdateInstanceOperation returns a new UpdateInstanceOperation from a given name.
// The name must be that of a previously created UpdateInstanceOperation, possibly from a different process.
func (c *CloudRedisClient) UpdateInstanceOperation(name string) *UpdateInstanceOperation {
	return &UpdateInstanceOperation{
		lro: longrunning.InternalNewOperation(c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *UpdateInstanceOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*redispb.Instance, error) {
	var resp redispb.Instance
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
func (op *UpdateInstanceOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*redispb.Instance, error) {
	var resp redispb.Instance
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
func (op *UpdateInstanceOperation) Metadata() (*redispb.OperationMetadata, error) {
	var meta redispb.OperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *UpdateInstanceOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *UpdateInstanceOperation) Name() string {
	return op.lro.Name()
}

// UpgradeInstanceOperation manages a long-running operation from UpgradeInstance.
type UpgradeInstanceOperation struct {
	lro *longrunning.Operation
}

// UpgradeInstanceOperation returns a new UpgradeInstanceOperation from a given name.
// The name must be that of a previously created UpgradeInstanceOperation, possibly from a different process.
func (c *CloudRedisClient) UpgradeInstanceOperation(name string) *UpgradeInstanceOperation {
	return &UpgradeInstanceOperation{
		lro: longrunning.InternalNewOperation(c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *UpgradeInstanceOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*redispb.Instance, error) {
	var resp redispb.Instance
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
func (op *UpgradeInstanceOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*redispb.Instance, error) {
	var resp redispb.Instance
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
func (op *UpgradeInstanceOperation) Metadata() (*redispb.OperationMetadata, error) {
	var meta redispb.OperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *UpgradeInstanceOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *UpgradeInstanceOperation) Name() string {
	return op.lro.Name()
}

// InstanceIterator manages a stream of *redispb.Instance.
type InstanceIterator struct {
	items    []*redispb.Instance
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
	InternalFetch func(pageSize int, pageToken string) (results []*redispb.Instance, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *InstanceIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *InstanceIterator) Next() (*redispb.Instance, error) {
	var item *redispb.Instance
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *InstanceIterator) bufLen() int {
	return len(it.items)
}

func (it *InstanceIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
