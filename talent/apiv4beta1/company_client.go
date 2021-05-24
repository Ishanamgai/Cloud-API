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

package talent

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
	talentpb "google.golang.org/genproto/googleapis/cloud/talent/v4beta1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newCompanyClientHook clientHook

// CompanyCallOptions contains the retry settings for each method of CompanyClient.
type CompanyCallOptions struct {
	CreateCompany []gax.CallOption
	GetCompany    []gax.CallOption
	UpdateCompany []gax.CallOption
	DeleteCompany []gax.CallOption
	ListCompanies []gax.CallOption
}

func defaultCompanyGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("jobs.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("jobs.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://jobs.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCompanyCallOptions() *CompanyCallOptions {
	return &CompanyCallOptions{
		CreateCompany: []gax.CallOption{},
		GetCompany: []gax.CallOption{
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
		UpdateCompany: []gax.CallOption{},
		DeleteCompany: []gax.CallOption{
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
		ListCompanies: []gax.CallOption{
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

// internalCompanyClient is an interface that defines the methods availaible from Cloud Talent Solution API.
type internalCompanyClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateCompany(context.Context, *talentpb.CreateCompanyRequest, ...gax.CallOption) (*talentpb.Company, error)
	GetCompany(context.Context, *talentpb.GetCompanyRequest, ...gax.CallOption) (*talentpb.Company, error)
	UpdateCompany(context.Context, *talentpb.UpdateCompanyRequest, ...gax.CallOption) (*talentpb.Company, error)
	DeleteCompany(context.Context, *talentpb.DeleteCompanyRequest, ...gax.CallOption) error
	ListCompanies(context.Context, *talentpb.ListCompaniesRequest, ...gax.CallOption) *CompanyIterator
}

// CompanyClient is a client for interacting with Cloud Talent Solution API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// A service that handles company management, including CRUD and enumeration.
type CompanyClient struct {
	// The internal transport-dependent client.
	internalClient internalCompanyClient

	// The call options for this service.
	CallOptions *CompanyCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *CompanyClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *CompanyClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *CompanyClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateCompany creates a new company entity.
func (c *CompanyClient) CreateCompany(ctx context.Context, req *talentpb.CreateCompanyRequest, opts ...gax.CallOption) (*talentpb.Company, error) {
	return c.internalClient.CreateCompany(ctx, req, opts...)
}

// GetCompany retrieves specified company.
func (c *CompanyClient) GetCompany(ctx context.Context, req *talentpb.GetCompanyRequest, opts ...gax.CallOption) (*talentpb.Company, error) {
	return c.internalClient.GetCompany(ctx, req, opts...)
}

// UpdateCompany updates specified company.
func (c *CompanyClient) UpdateCompany(ctx context.Context, req *talentpb.UpdateCompanyRequest, opts ...gax.CallOption) (*talentpb.Company, error) {
	return c.internalClient.UpdateCompany(ctx, req, opts...)
}

// DeleteCompany deletes specified company.
// Prerequisite: The company has no jobs associated with it.
func (c *CompanyClient) DeleteCompany(ctx context.Context, req *talentpb.DeleteCompanyRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteCompany(ctx, req, opts...)
}

// ListCompanies lists all companies associated with the project.
func (c *CompanyClient) ListCompanies(ctx context.Context, req *talentpb.ListCompaniesRequest, opts ...gax.CallOption) *CompanyIterator {
	return c.internalClient.ListCompanies(ctx, req, opts...)
}

// companyGRPCClient is a client for interacting with Cloud Talent Solution API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type companyGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing CompanyClient
	CallOptions **CompanyCallOptions

	// The gRPC API client.
	companyClient talentpb.CompanyServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewCompanyClient creates a new company service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// A service that handles company management, including CRUD and enumeration.
func NewCompanyClient(ctx context.Context, opts ...option.ClientOption) (*CompanyClient, error) {
	clientOpts := defaultCompanyGRPCClientOptions()
	if newCompanyClientHook != nil {
		hookOpts, err := newCompanyClientHook(ctx, clientHookParams{})
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
	client := CompanyClient{CallOptions: defaultCompanyCallOptions()}

	c := &companyGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		companyClient:    talentpb.NewCompanyServiceClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *companyGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *companyGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *companyGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *companyGRPCClient) CreateCompany(ctx context.Context, req *talentpb.CreateCompanyRequest, opts ...gax.CallOption) (*talentpb.Company, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateCompany[0:len((*c.CallOptions).CreateCompany):len((*c.CallOptions).CreateCompany)], opts...)
	var resp *talentpb.Company
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.companyClient.CreateCompany(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *companyGRPCClient) GetCompany(ctx context.Context, req *talentpb.GetCompanyRequest, opts ...gax.CallOption) (*talentpb.Company, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetCompany[0:len((*c.CallOptions).GetCompany):len((*c.CallOptions).GetCompany)], opts...)
	var resp *talentpb.Company
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.companyClient.GetCompany(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *companyGRPCClient) UpdateCompany(ctx context.Context, req *talentpb.UpdateCompanyRequest, opts ...gax.CallOption) (*talentpb.Company, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "company.name", url.QueryEscape(req.GetCompany().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateCompany[0:len((*c.CallOptions).UpdateCompany):len((*c.CallOptions).UpdateCompany)], opts...)
	var resp *talentpb.Company
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.companyClient.UpdateCompany(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *companyGRPCClient) DeleteCompany(ctx context.Context, req *talentpb.DeleteCompanyRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteCompany[0:len((*c.CallOptions).DeleteCompany):len((*c.CallOptions).DeleteCompany)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.companyClient.DeleteCompany(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *companyGRPCClient) ListCompanies(ctx context.Context, req *talentpb.ListCompaniesRequest, opts ...gax.CallOption) *CompanyIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListCompanies[0:len((*c.CallOptions).ListCompanies):len((*c.CallOptions).ListCompanies)], opts...)
	it := &CompanyIterator{}
	req = proto.Clone(req).(*talentpb.ListCompaniesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*talentpb.Company, string, error) {
		var resp *talentpb.ListCompaniesResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.companyClient.ListCompanies(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetCompanies(), resp.GetNextPageToken(), nil
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

// CompanyIterator manages a stream of *talentpb.Company.
type CompanyIterator struct {
	items    []*talentpb.Company
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
	InternalFetch func(pageSize int, pageToken string) (results []*talentpb.Company, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *CompanyIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *CompanyIterator) Next() (*talentpb.Company, error) {
	var item *talentpb.Company
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *CompanyIterator) bufLen() int {
	return len(it.items)
}

func (it *CompanyIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
