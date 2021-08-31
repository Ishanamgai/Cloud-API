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

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httpbodypb "google.golang.org/genproto/googleapis/api/httpbody"
	aiplatformpb "google.golang.org/genproto/googleapis/cloud/aiplatform/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var newPredictionClientHook clientHook

// PredictionCallOptions contains the retry settings for each method of PredictionClient.
type PredictionCallOptions struct {
	Predict    []gax.CallOption
	RawPredict []gax.CallOption
	Explain    []gax.CallOption
}

func defaultPredictionGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("aiplatform.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("aiplatform.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://aiplatform.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultPredictionCallOptions() *PredictionCallOptions {
	return &PredictionCallOptions{
		Predict:    []gax.CallOption{},
		RawPredict: []gax.CallOption{},
		Explain:    []gax.CallOption{},
	}
}

// internalPredictionClient is an interface that defines the methods availaible from Vertex AI API.
type internalPredictionClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	Predict(context.Context, *aiplatformpb.PredictRequest, ...gax.CallOption) (*aiplatformpb.PredictResponse, error)
	RawPredict(context.Context, *aiplatformpb.RawPredictRequest, ...gax.CallOption) (*httpbodypb.HttpBody, error)
	Explain(context.Context, *aiplatformpb.ExplainRequest, ...gax.CallOption) (*aiplatformpb.ExplainResponse, error)
}

// PredictionClient is a client for interacting with Vertex AI API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// A service for online predictions and explanations.
type PredictionClient struct {
	// The internal transport-dependent client.
	internalClient internalPredictionClient

	// The call options for this service.
	CallOptions *PredictionCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *PredictionClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *PredictionClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *PredictionClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// Predict perform an online prediction.
func (c *PredictionClient) Predict(ctx context.Context, req *aiplatformpb.PredictRequest, opts ...gax.CallOption) (*aiplatformpb.PredictResponse, error) {
	return c.internalClient.Predict(ctx, req, opts...)
}

// RawPredict perform an online prediction with arbitrary http payload.
func (c *PredictionClient) RawPredict(ctx context.Context, req *aiplatformpb.RawPredictRequest, opts ...gax.CallOption) (*httpbodypb.HttpBody, error) {
	return c.internalClient.RawPredict(ctx, req, opts...)
}

// Explain perform an online explanation.
//
// If deployed_model_id is specified,
// the corresponding DeployModel must have
// explanation_spec
// populated. If deployed_model_id
// is not specified, all DeployedModels must have
// explanation_spec
// populated. Only deployed AutoML tabular Models have
// explanation_spec.
func (c *PredictionClient) Explain(ctx context.Context, req *aiplatformpb.ExplainRequest, opts ...gax.CallOption) (*aiplatformpb.ExplainResponse, error) {
	return c.internalClient.Explain(ctx, req, opts...)
}

// predictionGRPCClient is a client for interacting with Vertex AI API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type predictionGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing PredictionClient
	CallOptions **PredictionCallOptions

	// The gRPC API client.
	predictionClient aiplatformpb.PredictionServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewPredictionClient creates a new prediction service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// A service for online predictions and explanations.
func NewPredictionClient(ctx context.Context, opts ...option.ClientOption) (*PredictionClient, error) {
	clientOpts := defaultPredictionGRPCClientOptions()
	if newPredictionClientHook != nil {
		hookOpts, err := newPredictionClientHook(ctx, clientHookParams{})
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
	client := PredictionClient{CallOptions: defaultPredictionCallOptions()}

	c := &predictionGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		predictionClient: aiplatformpb.NewPredictionServiceClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *predictionGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *predictionGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *predictionGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *predictionGRPCClient) Predict(ctx context.Context, req *aiplatformpb.PredictRequest, opts ...gax.CallOption) (*aiplatformpb.PredictResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "endpoint", url.QueryEscape(req.GetEndpoint())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).Predict[0:len((*c.CallOptions).Predict):len((*c.CallOptions).Predict)], opts...)
	var resp *aiplatformpb.PredictResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.predictionClient.Predict(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *predictionGRPCClient) RawPredict(ctx context.Context, req *aiplatformpb.RawPredictRequest, opts ...gax.CallOption) (*httpbodypb.HttpBody, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "endpoint", url.QueryEscape(req.GetEndpoint())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).RawPredict[0:len((*c.CallOptions).RawPredict):len((*c.CallOptions).RawPredict)], opts...)
	var resp *httpbodypb.HttpBody
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.predictionClient.RawPredict(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *predictionGRPCClient) Explain(ctx context.Context, req *aiplatformpb.ExplainRequest, opts ...gax.CallOption) (*aiplatformpb.ExplainResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "endpoint", url.QueryEscape(req.GetEndpoint())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).Explain[0:len((*c.CallOptions).Explain):len((*c.CallOptions).Explain)], opts...)
	var resp *aiplatformpb.ExplainResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.predictionClient.Explain(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
