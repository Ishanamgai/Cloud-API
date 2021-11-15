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

package appengine

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	appenginepb "google.golang.org/genproto/googleapis/appengine/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var newApplicationsClientHook clientHook

// ApplicationsCallOptions contains the retry settings for each method of ApplicationsClient.
type ApplicationsCallOptions struct {
	GetApplication    []gax.CallOption
	CreateApplication []gax.CallOption
	UpdateApplication []gax.CallOption
	RepairApplication []gax.CallOption
}

func defaultApplicationsGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("appengine.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("appengine.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://appengine.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultApplicationsCallOptions() *ApplicationsCallOptions {
	return &ApplicationsCallOptions{
		GetApplication:    []gax.CallOption{},
		CreateApplication: []gax.CallOption{},
		UpdateApplication: []gax.CallOption{},
		RepairApplication: []gax.CallOption{},
	}
}

// internalApplicationsClient is an interface that defines the methods availaible from App Engine Admin API.
type internalApplicationsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetApplication(context.Context, *appenginepb.GetApplicationRequest, ...gax.CallOption) (*appenginepb.Application, error)
	CreateApplication(context.Context, *appenginepb.CreateApplicationRequest, ...gax.CallOption) (*CreateApplicationOperation, error)
	CreateApplicationOperation(name string) *CreateApplicationOperation
	UpdateApplication(context.Context, *appenginepb.UpdateApplicationRequest, ...gax.CallOption) (*UpdateApplicationOperation, error)
	UpdateApplicationOperation(name string) *UpdateApplicationOperation
	RepairApplication(context.Context, *appenginepb.RepairApplicationRequest, ...gax.CallOption) (*RepairApplicationOperation, error)
	RepairApplicationOperation(name string) *RepairApplicationOperation
}

// ApplicationsClient is a client for interacting with App Engine Admin API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Manages App Engine applications.
type ApplicationsClient struct {
	// The internal transport-dependent client.
	internalClient internalApplicationsClient

	// The call options for this service.
	CallOptions *ApplicationsCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ApplicationsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ApplicationsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *ApplicationsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetApplication gets information about an application.
func (c *ApplicationsClient) GetApplication(ctx context.Context, req *appenginepb.GetApplicationRequest, opts ...gax.CallOption) (*appenginepb.Application, error) {
	return c.internalClient.GetApplication(ctx, req, opts...)
}

// CreateApplication creates an App Engine application for a Google Cloud Platform project.
// Required fields:
//
//   id - The ID of the target Cloud Platform project.
//
//   location - The region (at https://cloud.google.com/appengine/docs/locations) where you want the App Engine application located.
//
// For more information about App Engine applications, see Managing Projects, Applications, and Billing (at https://cloud.google.com/appengine/docs/standard/python/console/).
func (c *ApplicationsClient) CreateApplication(ctx context.Context, req *appenginepb.CreateApplicationRequest, opts ...gax.CallOption) (*CreateApplicationOperation, error) {
	return c.internalClient.CreateApplication(ctx, req, opts...)
}

// CreateApplicationOperation returns a new CreateApplicationOperation from a given name.
// The name must be that of a previously created CreateApplicationOperation, possibly from a different process.
func (c *ApplicationsClient) CreateApplicationOperation(name string) *CreateApplicationOperation {
	return c.internalClient.CreateApplicationOperation(name)
}

// UpdateApplication updates the specified Application resource.
// You can update the following fields:
//
//   auth_domain - Google authentication domain for controlling user access to the application.
//
//   default_cookie_expiration - Cookie expiration policy for the application.
func (c *ApplicationsClient) UpdateApplication(ctx context.Context, req *appenginepb.UpdateApplicationRequest, opts ...gax.CallOption) (*UpdateApplicationOperation, error) {
	return c.internalClient.UpdateApplication(ctx, req, opts...)
}

// UpdateApplicationOperation returns a new UpdateApplicationOperation from a given name.
// The name must be that of a previously created UpdateApplicationOperation, possibly from a different process.
func (c *ApplicationsClient) UpdateApplicationOperation(name string) *UpdateApplicationOperation {
	return c.internalClient.UpdateApplicationOperation(name)
}

// RepairApplication recreates the required App Engine features for the specified App Engine
// application, for example a Cloud Storage bucket or App Engine service
// account.
// Use this method if you receive an error message about a missing feature,
// for example, Error retrieving the App Engine service account.
// If you have deleted your App Engine service account, this will
// not be able to recreate it. Instead, you should attempt to use the
// IAM undelete API if possible at https://cloud.google.com/iam/reference/rest/v1/projects.serviceAccounts/undelete?apix_params={ (at https://cloud.google.com/iam/reference/rest/v1/projects.serviceAccounts/undelete?apix_params=%7B)“name”%3A"projects%2F-%2FserviceAccounts%2Funique_id"%2C"resource"%3A%7B%7D%7D .
// If the deletion was recent, the numeric ID can be found in the Cloud
// Console Activity Log.
func (c *ApplicationsClient) RepairApplication(ctx context.Context, req *appenginepb.RepairApplicationRequest, opts ...gax.CallOption) (*RepairApplicationOperation, error) {
	return c.internalClient.RepairApplication(ctx, req, opts...)
}

// RepairApplicationOperation returns a new RepairApplicationOperation from a given name.
// The name must be that of a previously created RepairApplicationOperation, possibly from a different process.
func (c *ApplicationsClient) RepairApplicationOperation(name string) *RepairApplicationOperation {
	return c.internalClient.RepairApplicationOperation(name)
}

// applicationsGRPCClient is a client for interacting with App Engine Admin API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type applicationsGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing ApplicationsClient
	CallOptions **ApplicationsCallOptions

	// The gRPC API client.
	applicationsClient appenginepb.ApplicationsClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewApplicationsClient creates a new applications client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Manages App Engine applications.
func NewApplicationsClient(ctx context.Context, opts ...option.ClientOption) (*ApplicationsClient, error) {
	clientOpts := defaultApplicationsGRPCClientOptions()
	if newApplicationsClientHook != nil {
		hookOpts, err := newApplicationsClientHook(ctx, clientHookParams{})
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
	client := ApplicationsClient{CallOptions: defaultApplicationsCallOptions()}

	c := &applicationsGRPCClient{
		connPool:           connPool,
		disableDeadlines:   disableDeadlines,
		applicationsClient: appenginepb.NewApplicationsClient(connPool),
		CallOptions:        &client.CallOptions,
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
func (c *applicationsGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *applicationsGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *applicationsGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *applicationsGRPCClient) GetApplication(ctx context.Context, req *appenginepb.GetApplicationRequest, opts ...gax.CallOption) (*appenginepb.Application, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetApplication[0:len((*c.CallOptions).GetApplication):len((*c.CallOptions).GetApplication)], opts...)
	var resp *appenginepb.Application
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.applicationsClient.GetApplication(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *applicationsGRPCClient) CreateApplication(ctx context.Context, req *appenginepb.CreateApplicationRequest, opts ...gax.CallOption) (*CreateApplicationOperation, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).CreateApplication[0:len((*c.CallOptions).CreateApplication):len((*c.CallOptions).CreateApplication)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.applicationsClient.CreateApplication(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &CreateApplicationOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *applicationsGRPCClient) UpdateApplication(ctx context.Context, req *appenginepb.UpdateApplicationRequest, opts ...gax.CallOption) (*UpdateApplicationOperation, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateApplication[0:len((*c.CallOptions).UpdateApplication):len((*c.CallOptions).UpdateApplication)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.applicationsClient.UpdateApplication(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &UpdateApplicationOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *applicationsGRPCClient) RepairApplication(ctx context.Context, req *appenginepb.RepairApplicationRequest, opts ...gax.CallOption) (*RepairApplicationOperation, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).RepairApplication[0:len((*c.CallOptions).RepairApplication):len((*c.CallOptions).RepairApplication)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.applicationsClient.RepairApplication(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &RepairApplicationOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

// CreateApplicationOperation manages a long-running operation from CreateApplication.
type CreateApplicationOperation struct {
	lro *longrunning.Operation
}

// CreateApplicationOperation returns a new CreateApplicationOperation from a given name.
// The name must be that of a previously created CreateApplicationOperation, possibly from a different process.
func (c *applicationsGRPCClient) CreateApplicationOperation(name string) *CreateApplicationOperation {
	return &CreateApplicationOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *CreateApplicationOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*appenginepb.Application, error) {
	var resp appenginepb.Application
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
func (op *CreateApplicationOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*appenginepb.Application, error) {
	var resp appenginepb.Application
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
func (op *CreateApplicationOperation) Metadata() (*appenginepb.OperationMetadataV1, error) {
	var meta appenginepb.OperationMetadataV1
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *CreateApplicationOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *CreateApplicationOperation) Name() string {
	return op.lro.Name()
}

// RepairApplicationOperation manages a long-running operation from RepairApplication.
type RepairApplicationOperation struct {
	lro *longrunning.Operation
}

// RepairApplicationOperation returns a new RepairApplicationOperation from a given name.
// The name must be that of a previously created RepairApplicationOperation, possibly from a different process.
func (c *applicationsGRPCClient) RepairApplicationOperation(name string) *RepairApplicationOperation {
	return &RepairApplicationOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *RepairApplicationOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*appenginepb.Application, error) {
	var resp appenginepb.Application
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
func (op *RepairApplicationOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*appenginepb.Application, error) {
	var resp appenginepb.Application
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
func (op *RepairApplicationOperation) Metadata() (*appenginepb.OperationMetadataV1, error) {
	var meta appenginepb.OperationMetadataV1
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *RepairApplicationOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *RepairApplicationOperation) Name() string {
	return op.lro.Name()
}

// UpdateApplicationOperation manages a long-running operation from UpdateApplication.
type UpdateApplicationOperation struct {
	lro *longrunning.Operation
}

// UpdateApplicationOperation returns a new UpdateApplicationOperation from a given name.
// The name must be that of a previously created UpdateApplicationOperation, possibly from a different process.
func (c *applicationsGRPCClient) UpdateApplicationOperation(name string) *UpdateApplicationOperation {
	return &UpdateApplicationOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *UpdateApplicationOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*appenginepb.Application, error) {
	var resp appenginepb.Application
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
func (op *UpdateApplicationOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*appenginepb.Application, error) {
	var resp appenginepb.Application
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
func (op *UpdateApplicationOperation) Metadata() (*appenginepb.OperationMetadataV1, error) {
	var meta appenginepb.OperationMetadataV1
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *UpdateApplicationOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *UpdateApplicationOperation) Name() string {
	return op.lro.Name()
}
