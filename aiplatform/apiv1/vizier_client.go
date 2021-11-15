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

var newVizierClientHook clientHook

// VizierCallOptions contains the retry settings for each method of VizierClient.
type VizierCallOptions struct {
	CreateStudy                  []gax.CallOption
	GetStudy                     []gax.CallOption
	ListStudies                  []gax.CallOption
	DeleteStudy                  []gax.CallOption
	LookupStudy                  []gax.CallOption
	SuggestTrials                []gax.CallOption
	CreateTrial                  []gax.CallOption
	GetTrial                     []gax.CallOption
	ListTrials                   []gax.CallOption
	AddTrialMeasurement          []gax.CallOption
	CompleteTrial                []gax.CallOption
	DeleteTrial                  []gax.CallOption
	CheckTrialEarlyStoppingState []gax.CallOption
	StopTrial                    []gax.CallOption
	ListOptimalTrials            []gax.CallOption
}

func defaultVizierGRPCClientOptions() []option.ClientOption {
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

func defaultVizierCallOptions() *VizierCallOptions {
	return &VizierCallOptions{
		CreateStudy:                  []gax.CallOption{},
		GetStudy:                     []gax.CallOption{},
		ListStudies:                  []gax.CallOption{},
		DeleteStudy:                  []gax.CallOption{},
		LookupStudy:                  []gax.CallOption{},
		SuggestTrials:                []gax.CallOption{},
		CreateTrial:                  []gax.CallOption{},
		GetTrial:                     []gax.CallOption{},
		ListTrials:                   []gax.CallOption{},
		AddTrialMeasurement:          []gax.CallOption{},
		CompleteTrial:                []gax.CallOption{},
		DeleteTrial:                  []gax.CallOption{},
		CheckTrialEarlyStoppingState: []gax.CallOption{},
		StopTrial:                    []gax.CallOption{},
		ListOptimalTrials:            []gax.CallOption{},
	}
}

// internalVizierClient is an interface that defines the methods availaible from Vertex AI API.
type internalVizierClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateStudy(context.Context, *aiplatformpb.CreateStudyRequest, ...gax.CallOption) (*aiplatformpb.Study, error)
	GetStudy(context.Context, *aiplatformpb.GetStudyRequest, ...gax.CallOption) (*aiplatformpb.Study, error)
	ListStudies(context.Context, *aiplatformpb.ListStudiesRequest, ...gax.CallOption) *StudyIterator
	DeleteStudy(context.Context, *aiplatformpb.DeleteStudyRequest, ...gax.CallOption) error
	LookupStudy(context.Context, *aiplatformpb.LookupStudyRequest, ...gax.CallOption) (*aiplatformpb.Study, error)
	SuggestTrials(context.Context, *aiplatformpb.SuggestTrialsRequest, ...gax.CallOption) (*SuggestTrialsOperation, error)
	SuggestTrialsOperation(name string) *SuggestTrialsOperation
	CreateTrial(context.Context, *aiplatformpb.CreateTrialRequest, ...gax.CallOption) (*aiplatformpb.Trial, error)
	GetTrial(context.Context, *aiplatformpb.GetTrialRequest, ...gax.CallOption) (*aiplatformpb.Trial, error)
	ListTrials(context.Context, *aiplatformpb.ListTrialsRequest, ...gax.CallOption) *TrialIterator
	AddTrialMeasurement(context.Context, *aiplatformpb.AddTrialMeasurementRequest, ...gax.CallOption) (*aiplatformpb.Trial, error)
	CompleteTrial(context.Context, *aiplatformpb.CompleteTrialRequest, ...gax.CallOption) (*aiplatformpb.Trial, error)
	DeleteTrial(context.Context, *aiplatformpb.DeleteTrialRequest, ...gax.CallOption) error
	CheckTrialEarlyStoppingState(context.Context, *aiplatformpb.CheckTrialEarlyStoppingStateRequest, ...gax.CallOption) (*CheckTrialEarlyStoppingStateOperation, error)
	CheckTrialEarlyStoppingStateOperation(name string) *CheckTrialEarlyStoppingStateOperation
	StopTrial(context.Context, *aiplatformpb.StopTrialRequest, ...gax.CallOption) (*aiplatformpb.Trial, error)
	ListOptimalTrials(context.Context, *aiplatformpb.ListOptimalTrialsRequest, ...gax.CallOption) (*aiplatformpb.ListOptimalTrialsResponse, error)
}

// VizierClient is a client for interacting with Vertex AI API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Vertex Vizier API.
//
// Vizier service is a GCP service to solve blackbox optimization problems,
// such as tuning machine learning hyperparameters and searching over deep
// learning architectures.
type VizierClient struct {
	// The internal transport-dependent client.
	internalClient internalVizierClient

	// The call options for this service.
	CallOptions *VizierCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *VizierClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *VizierClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *VizierClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateStudy creates a Study. A resource name will be generated after creation of the
// Study.
func (c *VizierClient) CreateStudy(ctx context.Context, req *aiplatformpb.CreateStudyRequest, opts ...gax.CallOption) (*aiplatformpb.Study, error) {
	return c.internalClient.CreateStudy(ctx, req, opts...)
}

// GetStudy gets a Study by name.
func (c *VizierClient) GetStudy(ctx context.Context, req *aiplatformpb.GetStudyRequest, opts ...gax.CallOption) (*aiplatformpb.Study, error) {
	return c.internalClient.GetStudy(ctx, req, opts...)
}

// ListStudies lists all the studies in a region for an associated project.
func (c *VizierClient) ListStudies(ctx context.Context, req *aiplatformpb.ListStudiesRequest, opts ...gax.CallOption) *StudyIterator {
	return c.internalClient.ListStudies(ctx, req, opts...)
}

// DeleteStudy deletes a Study.
func (c *VizierClient) DeleteStudy(ctx context.Context, req *aiplatformpb.DeleteStudyRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteStudy(ctx, req, opts...)
}

// LookupStudy looks a study up using the user-defined display_name field instead of the
// fully qualified resource name.
func (c *VizierClient) LookupStudy(ctx context.Context, req *aiplatformpb.LookupStudyRequest, opts ...gax.CallOption) (*aiplatformpb.Study, error) {
	return c.internalClient.LookupStudy(ctx, req, opts...)
}

// SuggestTrials adds one or more Trials to a Study, with parameter values
// suggested by Vertex Vizier. Returns a long-running
// operation associated with the generation of Trial suggestions.
// When this long-running operation succeeds, it will contain
// a SuggestTrialsResponse.
func (c *VizierClient) SuggestTrials(ctx context.Context, req *aiplatformpb.SuggestTrialsRequest, opts ...gax.CallOption) (*SuggestTrialsOperation, error) {
	return c.internalClient.SuggestTrials(ctx, req, opts...)
}

// SuggestTrialsOperation returns a new SuggestTrialsOperation from a given name.
// The name must be that of a previously created SuggestTrialsOperation, possibly from a different process.
func (c *VizierClient) SuggestTrialsOperation(name string) *SuggestTrialsOperation {
	return c.internalClient.SuggestTrialsOperation(name)
}

// CreateTrial adds a user provided Trial to a Study.
func (c *VizierClient) CreateTrial(ctx context.Context, req *aiplatformpb.CreateTrialRequest, opts ...gax.CallOption) (*aiplatformpb.Trial, error) {
	return c.internalClient.CreateTrial(ctx, req, opts...)
}

// GetTrial gets a Trial.
func (c *VizierClient) GetTrial(ctx context.Context, req *aiplatformpb.GetTrialRequest, opts ...gax.CallOption) (*aiplatformpb.Trial, error) {
	return c.internalClient.GetTrial(ctx, req, opts...)
}

// ListTrials lists the Trials associated with a Study.
func (c *VizierClient) ListTrials(ctx context.Context, req *aiplatformpb.ListTrialsRequest, opts ...gax.CallOption) *TrialIterator {
	return c.internalClient.ListTrials(ctx, req, opts...)
}

// AddTrialMeasurement adds a measurement of the objective metrics to a Trial. This measurement
// is assumed to have been taken before the Trial is complete.
func (c *VizierClient) AddTrialMeasurement(ctx context.Context, req *aiplatformpb.AddTrialMeasurementRequest, opts ...gax.CallOption) (*aiplatformpb.Trial, error) {
	return c.internalClient.AddTrialMeasurement(ctx, req, opts...)
}

// CompleteTrial marks a Trial as complete.
func (c *VizierClient) CompleteTrial(ctx context.Context, req *aiplatformpb.CompleteTrialRequest, opts ...gax.CallOption) (*aiplatformpb.Trial, error) {
	return c.internalClient.CompleteTrial(ctx, req, opts...)
}

// DeleteTrial deletes a Trial.
func (c *VizierClient) DeleteTrial(ctx context.Context, req *aiplatformpb.DeleteTrialRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteTrial(ctx, req, opts...)
}

// CheckTrialEarlyStoppingState checks  whether a Trial should stop or not. Returns a
// long-running operation. When the operation is successful,
// it will contain a
// CheckTrialEarlyStoppingStateResponse.
func (c *VizierClient) CheckTrialEarlyStoppingState(ctx context.Context, req *aiplatformpb.CheckTrialEarlyStoppingStateRequest, opts ...gax.CallOption) (*CheckTrialEarlyStoppingStateOperation, error) {
	return c.internalClient.CheckTrialEarlyStoppingState(ctx, req, opts...)
}

// CheckTrialEarlyStoppingStateOperation returns a new CheckTrialEarlyStoppingStateOperation from a given name.
// The name must be that of a previously created CheckTrialEarlyStoppingStateOperation, possibly from a different process.
func (c *VizierClient) CheckTrialEarlyStoppingStateOperation(name string) *CheckTrialEarlyStoppingStateOperation {
	return c.internalClient.CheckTrialEarlyStoppingStateOperation(name)
}

// StopTrial stops a Trial.
func (c *VizierClient) StopTrial(ctx context.Context, req *aiplatformpb.StopTrialRequest, opts ...gax.CallOption) (*aiplatformpb.Trial, error) {
	return c.internalClient.StopTrial(ctx, req, opts...)
}

// ListOptimalTrials lists the pareto-optimal Trials for multi-objective Study or the
// optimal Trials for single-objective Study. The definition of
// pareto-optimal can be checked in wiki page.
// https://en.wikipedia.org/wiki/Pareto_efficiency (at https://en.wikipedia.org/wiki/Pareto_efficiency)
func (c *VizierClient) ListOptimalTrials(ctx context.Context, req *aiplatformpb.ListOptimalTrialsRequest, opts ...gax.CallOption) (*aiplatformpb.ListOptimalTrialsResponse, error) {
	return c.internalClient.ListOptimalTrials(ctx, req, opts...)
}

// vizierGRPCClient is a client for interacting with Vertex AI API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type vizierGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing VizierClient
	CallOptions **VizierCallOptions

	// The gRPC API client.
	vizierClient aiplatformpb.VizierServiceClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewVizierClient creates a new vizier service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Vertex Vizier API.
//
// Vizier service is a GCP service to solve blackbox optimization problems,
// such as tuning machine learning hyperparameters and searching over deep
// learning architectures.
func NewVizierClient(ctx context.Context, opts ...option.ClientOption) (*VizierClient, error) {
	clientOpts := defaultVizierGRPCClientOptions()
	if newVizierClientHook != nil {
		hookOpts, err := newVizierClientHook(ctx, clientHookParams{})
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
	client := VizierClient{CallOptions: defaultVizierCallOptions()}

	c := &vizierGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		vizierClient:     aiplatformpb.NewVizierServiceClient(connPool),
		CallOptions:      &client.CallOptions,
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
func (c *vizierGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *vizierGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *vizierGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *vizierGRPCClient) CreateStudy(ctx context.Context, req *aiplatformpb.CreateStudyRequest, opts ...gax.CallOption) (*aiplatformpb.Study, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateStudy[0:len((*c.CallOptions).CreateStudy):len((*c.CallOptions).CreateStudy)], opts...)
	var resp *aiplatformpb.Study
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.vizierClient.CreateStudy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *vizierGRPCClient) GetStudy(ctx context.Context, req *aiplatformpb.GetStudyRequest, opts ...gax.CallOption) (*aiplatformpb.Study, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetStudy[0:len((*c.CallOptions).GetStudy):len((*c.CallOptions).GetStudy)], opts...)
	var resp *aiplatformpb.Study
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.vizierClient.GetStudy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *vizierGRPCClient) ListStudies(ctx context.Context, req *aiplatformpb.ListStudiesRequest, opts ...gax.CallOption) *StudyIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListStudies[0:len((*c.CallOptions).ListStudies):len((*c.CallOptions).ListStudies)], opts...)
	it := &StudyIterator{}
	req = proto.Clone(req).(*aiplatformpb.ListStudiesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*aiplatformpb.Study, string, error) {
		resp := &aiplatformpb.ListStudiesResponse{}
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
			resp, err = c.vizierClient.ListStudies(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetStudies(), resp.GetNextPageToken(), nil
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

func (c *vizierGRPCClient) DeleteStudy(ctx context.Context, req *aiplatformpb.DeleteStudyRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteStudy[0:len((*c.CallOptions).DeleteStudy):len((*c.CallOptions).DeleteStudy)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.vizierClient.DeleteStudy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *vizierGRPCClient) LookupStudy(ctx context.Context, req *aiplatformpb.LookupStudyRequest, opts ...gax.CallOption) (*aiplatformpb.Study, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).LookupStudy[0:len((*c.CallOptions).LookupStudy):len((*c.CallOptions).LookupStudy)], opts...)
	var resp *aiplatformpb.Study
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.vizierClient.LookupStudy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *vizierGRPCClient) SuggestTrials(ctx context.Context, req *aiplatformpb.SuggestTrialsRequest, opts ...gax.CallOption) (*SuggestTrialsOperation, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).SuggestTrials[0:len((*c.CallOptions).SuggestTrials):len((*c.CallOptions).SuggestTrials)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.vizierClient.SuggestTrials(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &SuggestTrialsOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *vizierGRPCClient) CreateTrial(ctx context.Context, req *aiplatformpb.CreateTrialRequest, opts ...gax.CallOption) (*aiplatformpb.Trial, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateTrial[0:len((*c.CallOptions).CreateTrial):len((*c.CallOptions).CreateTrial)], opts...)
	var resp *aiplatformpb.Trial
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.vizierClient.CreateTrial(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *vizierGRPCClient) GetTrial(ctx context.Context, req *aiplatformpb.GetTrialRequest, opts ...gax.CallOption) (*aiplatformpb.Trial, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetTrial[0:len((*c.CallOptions).GetTrial):len((*c.CallOptions).GetTrial)], opts...)
	var resp *aiplatformpb.Trial
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.vizierClient.GetTrial(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *vizierGRPCClient) ListTrials(ctx context.Context, req *aiplatformpb.ListTrialsRequest, opts ...gax.CallOption) *TrialIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListTrials[0:len((*c.CallOptions).ListTrials):len((*c.CallOptions).ListTrials)], opts...)
	it := &TrialIterator{}
	req = proto.Clone(req).(*aiplatformpb.ListTrialsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*aiplatformpb.Trial, string, error) {
		resp := &aiplatformpb.ListTrialsResponse{}
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
			resp, err = c.vizierClient.ListTrials(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetTrials(), resp.GetNextPageToken(), nil
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

func (c *vizierGRPCClient) AddTrialMeasurement(ctx context.Context, req *aiplatformpb.AddTrialMeasurementRequest, opts ...gax.CallOption) (*aiplatformpb.Trial, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "trial_name", url.QueryEscape(req.GetTrialName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).AddTrialMeasurement[0:len((*c.CallOptions).AddTrialMeasurement):len((*c.CallOptions).AddTrialMeasurement)], opts...)
	var resp *aiplatformpb.Trial
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.vizierClient.AddTrialMeasurement(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *vizierGRPCClient) CompleteTrial(ctx context.Context, req *aiplatformpb.CompleteTrialRequest, opts ...gax.CallOption) (*aiplatformpb.Trial, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CompleteTrial[0:len((*c.CallOptions).CompleteTrial):len((*c.CallOptions).CompleteTrial)], opts...)
	var resp *aiplatformpb.Trial
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.vizierClient.CompleteTrial(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *vizierGRPCClient) DeleteTrial(ctx context.Context, req *aiplatformpb.DeleteTrialRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteTrial[0:len((*c.CallOptions).DeleteTrial):len((*c.CallOptions).DeleteTrial)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.vizierClient.DeleteTrial(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *vizierGRPCClient) CheckTrialEarlyStoppingState(ctx context.Context, req *aiplatformpb.CheckTrialEarlyStoppingStateRequest, opts ...gax.CallOption) (*CheckTrialEarlyStoppingStateOperation, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "trial_name", url.QueryEscape(req.GetTrialName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CheckTrialEarlyStoppingState[0:len((*c.CallOptions).CheckTrialEarlyStoppingState):len((*c.CallOptions).CheckTrialEarlyStoppingState)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.vizierClient.CheckTrialEarlyStoppingState(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &CheckTrialEarlyStoppingStateOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *vizierGRPCClient) StopTrial(ctx context.Context, req *aiplatformpb.StopTrialRequest, opts ...gax.CallOption) (*aiplatformpb.Trial, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).StopTrial[0:len((*c.CallOptions).StopTrial):len((*c.CallOptions).StopTrial)], opts...)
	var resp *aiplatformpb.Trial
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.vizierClient.StopTrial(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *vizierGRPCClient) ListOptimalTrials(ctx context.Context, req *aiplatformpb.ListOptimalTrialsRequest, opts ...gax.CallOption) (*aiplatformpb.ListOptimalTrialsResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListOptimalTrials[0:len((*c.CallOptions).ListOptimalTrials):len((*c.CallOptions).ListOptimalTrials)], opts...)
	var resp *aiplatformpb.ListOptimalTrialsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.vizierClient.ListOptimalTrials(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CheckTrialEarlyStoppingStateOperation manages a long-running operation from CheckTrialEarlyStoppingState.
type CheckTrialEarlyStoppingStateOperation struct {
	lro *longrunning.Operation
}

// CheckTrialEarlyStoppingStateOperation returns a new CheckTrialEarlyStoppingStateOperation from a given name.
// The name must be that of a previously created CheckTrialEarlyStoppingStateOperation, possibly from a different process.
func (c *vizierGRPCClient) CheckTrialEarlyStoppingStateOperation(name string) *CheckTrialEarlyStoppingStateOperation {
	return &CheckTrialEarlyStoppingStateOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *CheckTrialEarlyStoppingStateOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*aiplatformpb.CheckTrialEarlyStoppingStateResponse, error) {
	var resp aiplatformpb.CheckTrialEarlyStoppingStateResponse
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
func (op *CheckTrialEarlyStoppingStateOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*aiplatformpb.CheckTrialEarlyStoppingStateResponse, error) {
	var resp aiplatformpb.CheckTrialEarlyStoppingStateResponse
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
func (op *CheckTrialEarlyStoppingStateOperation) Metadata() (*aiplatformpb.CheckTrialEarlyStoppingStateMetatdata, error) {
	var meta aiplatformpb.CheckTrialEarlyStoppingStateMetatdata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *CheckTrialEarlyStoppingStateOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *CheckTrialEarlyStoppingStateOperation) Name() string {
	return op.lro.Name()
}

// SuggestTrialsOperation manages a long-running operation from SuggestTrials.
type SuggestTrialsOperation struct {
	lro *longrunning.Operation
}

// SuggestTrialsOperation returns a new SuggestTrialsOperation from a given name.
// The name must be that of a previously created SuggestTrialsOperation, possibly from a different process.
func (c *vizierGRPCClient) SuggestTrialsOperation(name string) *SuggestTrialsOperation {
	return &SuggestTrialsOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *SuggestTrialsOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*aiplatformpb.SuggestTrialsResponse, error) {
	var resp aiplatformpb.SuggestTrialsResponse
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
func (op *SuggestTrialsOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*aiplatformpb.SuggestTrialsResponse, error) {
	var resp aiplatformpb.SuggestTrialsResponse
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
func (op *SuggestTrialsOperation) Metadata() (*aiplatformpb.SuggestTrialsMetadata, error) {
	var meta aiplatformpb.SuggestTrialsMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *SuggestTrialsOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *SuggestTrialsOperation) Name() string {
	return op.lro.Name()
}

// StudyIterator manages a stream of *aiplatformpb.Study.
type StudyIterator struct {
	items    []*aiplatformpb.Study
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
	InternalFetch func(pageSize int, pageToken string) (results []*aiplatformpb.Study, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *StudyIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *StudyIterator) Next() (*aiplatformpb.Study, error) {
	var item *aiplatformpb.Study
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *StudyIterator) bufLen() int {
	return len(it.items)
}

func (it *StudyIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// TrialIterator manages a stream of *aiplatformpb.Trial.
type TrialIterator struct {
	items    []*aiplatformpb.Trial
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
	InternalFetch func(pageSize int, pageToken string) (results []*aiplatformpb.Trial, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *TrialIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *TrialIterator) Next() (*aiplatformpb.Trial, error) {
	var item *aiplatformpb.Trial
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *TrialIterator) bufLen() int {
	return len(it.items)
}

func (it *TrialIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
