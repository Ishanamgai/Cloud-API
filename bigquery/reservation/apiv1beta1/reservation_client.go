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

package reservation

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
	gtransport "google.golang.org/api/transport/grpc"
	reservationpb "google.golang.org/genproto/googleapis/cloud/bigquery/reservation/v1beta1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

// CallOptions contains the retry settings for each method of Client.
type CallOptions struct {
	CreateReservation        []gax.CallOption
	ListReservations         []gax.CallOption
	GetReservation           []gax.CallOption
	DeleteReservation        []gax.CallOption
	UpdateReservation        []gax.CallOption
	ListCapacityCommitments  []gax.CallOption
	GetCapacityCommitment    []gax.CallOption
	DeleteCapacityCommitment []gax.CallOption
	UpdateCapacityCommitment []gax.CallOption
	SplitCapacityCommitment  []gax.CallOption
	MergeCapacityCommitments []gax.CallOption
	CreateAssignment         []gax.CallOption
	ListAssignments          []gax.CallOption
	DeleteAssignment         []gax.CallOption
	SearchAssignments        []gax.CallOption
	MoveAssignment           []gax.CallOption
	GetBiReservation         []gax.CallOption
	UpdateBiReservation      []gax.CallOption
}

func defaultClientOptions() []option.ClientOption {
	return []option.ClientOption{
		option.WithEndpoint("bigqueryreservation.googleapis.com:443"),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCallOptions() *CallOptions {
	return &CallOptions{
		CreateReservation: []gax.CallOption{},
		ListReservations: []gax.CallOption{
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
		GetReservation: []gax.CallOption{
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
		DeleteReservation: []gax.CallOption{
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
		UpdateReservation: []gax.CallOption{},
		ListCapacityCommitments: []gax.CallOption{
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
		GetCapacityCommitment: []gax.CallOption{
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
		DeleteCapacityCommitment: []gax.CallOption{
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
		UpdateCapacityCommitment: []gax.CallOption{},
		SplitCapacityCommitment:  []gax.CallOption{},
		MergeCapacityCommitments: []gax.CallOption{},
		CreateAssignment:         []gax.CallOption{},
		ListAssignments: []gax.CallOption{
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
		DeleteAssignment: []gax.CallOption{
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
		SearchAssignments: []gax.CallOption{
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
		MoveAssignment: []gax.CallOption{},
		GetBiReservation: []gax.CallOption{
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
		UpdateBiReservation: []gax.CallOption{},
	}
}

// Client is a client for interacting with BigQuery Reservation API.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type Client struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// The gRPC API client.
	client reservationpb.ReservationServiceClient

	// The call options for this service.
	CallOptions *CallOptions

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewClient creates a new reservation service client.
//
// This API allows users to manage their flat-rate BigQuery reservations.
//
// A reservation provides computational resource guarantees, in the form of
// slots (at https://cloud.google.com/bigquery/docs/slots), to users. A slot is a
// unit of computational power in BigQuery, and serves as the basic unit of
// parallelism. In a scan of a multi-partitioned table, a single slot operates
// on a single partition of the table. A reservation resource exists as a child
// resource of the admin project and location, e.g.:
// projects/myproject/locations/US/reservations/reservationName.
//
// A capacity commitment is a way to purchase compute capacity for BigQuery jobs
// (in the form of slots) with some committed period of usage. A capacity
// commitment resource exists as a child resource of the admin project and
// location, e.g.:
// projects/myproject/locations/US/capacityCommitments/id.
func NewClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
	connPool, err := gtransport.DialPool(ctx, append(defaultClientOptions(), opts...)...)
	if err != nil {
		return nil, err
	}
	c := &Client{
		connPool:    connPool,
		CallOptions: defaultCallOptions(),

		client: reservationpb.NewReservationServiceClient(connPool),
	}
	c.setGoogleClientInfo()

	return c, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *Client) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *Client) Close() error {
	return c.connPool.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *Client) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// CreateReservation creates a new reservation resource.
func (c *Client) CreateReservation(ctx context.Context, req *reservationpb.CreateReservationRequest, opts ...gax.CallOption) (*reservationpb.Reservation, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.CreateReservation[0:len(c.CallOptions.CreateReservation):len(c.CallOptions.CreateReservation)], opts...)
	var resp *reservationpb.Reservation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.CreateReservation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ListReservations lists all the reservations for the project in the specified location.
func (c *Client) ListReservations(ctx context.Context, req *reservationpb.ListReservationsRequest, opts ...gax.CallOption) *ReservationIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ListReservations[0:len(c.CallOptions.ListReservations):len(c.CallOptions.ListReservations)], opts...)
	it := &ReservationIterator{}
	req = proto.Clone(req).(*reservationpb.ListReservationsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*reservationpb.Reservation, string, error) {
		var resp *reservationpb.ListReservationsResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.client.ListReservations(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.Reservations, resp.NextPageToken, nil
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

// GetReservation returns information about the reservation.
func (c *Client) GetReservation(ctx context.Context, req *reservationpb.GetReservationRequest, opts ...gax.CallOption) (*reservationpb.Reservation, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.GetReservation[0:len(c.CallOptions.GetReservation):len(c.CallOptions.GetReservation)], opts...)
	var resp *reservationpb.Reservation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetReservation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DeleteReservation deletes a reservation.
// Returns google.rpc.Code.FAILED_PRECONDITION when reservation has
// assignments.
func (c *Client) DeleteReservation(ctx context.Context, req *reservationpb.DeleteReservationRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.DeleteReservation[0:len(c.CallOptions.DeleteReservation):len(c.CallOptions.DeleteReservation)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.client.DeleteReservation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// UpdateReservation updates an existing reservation resource.
func (c *Client) UpdateReservation(ctx context.Context, req *reservationpb.UpdateReservationRequest, opts ...gax.CallOption) (*reservationpb.Reservation, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "reservation.name", url.QueryEscape(req.GetReservation().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.UpdateReservation[0:len(c.CallOptions.UpdateReservation):len(c.CallOptions.UpdateReservation)], opts...)
	var resp *reservationpb.Reservation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.UpdateReservation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ListCapacityCommitments lists all the capacity commitments for the admin project.
func (c *Client) ListCapacityCommitments(ctx context.Context, req *reservationpb.ListCapacityCommitmentsRequest, opts ...gax.CallOption) *CapacityCommitmentIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ListCapacityCommitments[0:len(c.CallOptions.ListCapacityCommitments):len(c.CallOptions.ListCapacityCommitments)], opts...)
	it := &CapacityCommitmentIterator{}
	req = proto.Clone(req).(*reservationpb.ListCapacityCommitmentsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*reservationpb.CapacityCommitment, string, error) {
		var resp *reservationpb.ListCapacityCommitmentsResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.client.ListCapacityCommitments(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.CapacityCommitments, resp.NextPageToken, nil
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

// GetCapacityCommitment returns information about the capacity commitment.
func (c *Client) GetCapacityCommitment(ctx context.Context, req *reservationpb.GetCapacityCommitmentRequest, opts ...gax.CallOption) (*reservationpb.CapacityCommitment, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.GetCapacityCommitment[0:len(c.CallOptions.GetCapacityCommitment):len(c.CallOptions.GetCapacityCommitment)], opts...)
	var resp *reservationpb.CapacityCommitment
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetCapacityCommitment(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DeleteCapacityCommitment deletes a capacity commitment. Attempting to delete capacity commitment
// before its commitment_end_time will fail with the error code
// google.rpc.Code.FAILED_PRECONDITION.
func (c *Client) DeleteCapacityCommitment(ctx context.Context, req *reservationpb.DeleteCapacityCommitmentRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.DeleteCapacityCommitment[0:len(c.CallOptions.DeleteCapacityCommitment):len(c.CallOptions.DeleteCapacityCommitment)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.client.DeleteCapacityCommitment(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// UpdateCapacityCommitment updates an existing capacity commitment.
//
// Only renewal_plan field can be updated.
func (c *Client) UpdateCapacityCommitment(ctx context.Context, req *reservationpb.UpdateCapacityCommitmentRequest, opts ...gax.CallOption) (*reservationpb.CapacityCommitment, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "capacity_commitment.name", url.QueryEscape(req.GetCapacityCommitment().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.UpdateCapacityCommitment[0:len(c.CallOptions.UpdateCapacityCommitment):len(c.CallOptions.UpdateCapacityCommitment)], opts...)
	var resp *reservationpb.CapacityCommitment
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.UpdateCapacityCommitment(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SplitCapacityCommitment splits capacity commitment to two commitments of the same plan and
// commitment_end_time. A common use case to do that is to perform a downgrade
// e.g., in order to downgrade from 10000 slots to 8000, one might split 10000
// capacity commitment to 2000 and 8000, change the plan of the first one to
// flex and then delete it.
func (c *Client) SplitCapacityCommitment(ctx context.Context, req *reservationpb.SplitCapacityCommitmentRequest, opts ...gax.CallOption) (*reservationpb.SplitCapacityCommitmentResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.SplitCapacityCommitment[0:len(c.CallOptions.SplitCapacityCommitment):len(c.CallOptions.SplitCapacityCommitment)], opts...)
	var resp *reservationpb.SplitCapacityCommitmentResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.SplitCapacityCommitment(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// MergeCapacityCommitments merges capacity commitments of the same plan into one. Resulting capacity
// commitment has the longer commitment_end_time out of the two. Attempting to
// merge capacity commitments of different plan will fail with the error code
// google.rpc.Code.FAILED_PRECONDITION.
func (c *Client) MergeCapacityCommitments(ctx context.Context, req *reservationpb.MergeCapacityCommitmentsRequest, opts ...gax.CallOption) (*reservationpb.CapacityCommitment, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.MergeCapacityCommitments[0:len(c.CallOptions.MergeCapacityCommitments):len(c.CallOptions.MergeCapacityCommitments)], opts...)
	var resp *reservationpb.CapacityCommitment
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.MergeCapacityCommitments(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreateAssignment returns google.rpc.Code.PERMISSION_DENIED if user does not have
// ‘bigquery.admin’ permissions on the project using the reservation
// and the project that owns this reservation.
// Returns google.rpc.Code.INVALID_ARGUMENT when location of the assignment
// does not match location of the reservation.
func (c *Client) CreateAssignment(ctx context.Context, req *reservationpb.CreateAssignmentRequest, opts ...gax.CallOption) (*reservationpb.Assignment, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.CreateAssignment[0:len(c.CallOptions.CreateAssignment):len(c.CallOptions.CreateAssignment)], opts...)
	var resp *reservationpb.Assignment
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.CreateAssignment(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ListAssignments lists assignments.
// Only explicitly created assignments will be returned. E.g:
// organizationA contains project1 and project2. Reservation res1 exists.
// CreateAssignment was invoked previously and following assignments were
// created explicitly:
// <organizationA, res1>
// <project1, res1>
// Then this API will just return the above two assignments for reservation
// res1, and no expansion/merge will happen. Wildcard “-” can be used for
// reservations in the request. In that case all assignments belongs to the
// specified project and location will be listed. Note
// “-” cannot be used for projects nor locations.
func (c *Client) ListAssignments(ctx context.Context, req *reservationpb.ListAssignmentsRequest, opts ...gax.CallOption) *AssignmentIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ListAssignments[0:len(c.CallOptions.ListAssignments):len(c.CallOptions.ListAssignments)], opts...)
	it := &AssignmentIterator{}
	req = proto.Clone(req).(*reservationpb.ListAssignmentsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*reservationpb.Assignment, string, error) {
		var resp *reservationpb.ListAssignmentsResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.client.ListAssignments(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.Assignments, resp.NextPageToken, nil
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

// DeleteAssignment deletes a assignment. No expansion will happen.
// E.g:
// organizationA contains project1 and project2. Reservation res1 exists.
// CreateAssignment was invoked previously and following assignments were
// created explicitly:
// <organizationA, res1>
// <project1, res1>
// Then deletion of <organizationA, res1> won’t affect <project1, res1>. After
// deletion of <organizationA, res1>, queries from project1 will still use
// res1, while queries from project2 will use on-demand mode.
func (c *Client) DeleteAssignment(ctx context.Context, req *reservationpb.DeleteAssignmentRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.DeleteAssignment[0:len(c.CallOptions.DeleteAssignment):len(c.CallOptions.DeleteAssignment)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.client.DeleteAssignment(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// SearchAssignments looks up assignments for a specified resource for a particular region.
// If the request is about a project:
//
// Assignments created on the project will be returned if they exist.
//
// Otherwise assignments created on the closest ancestor will be
// returned. 3) Assignments for different JobTypes will all be returned.
// Same logic applies if the request is about a folder.
// If the request is about an organization, then assignments created on the
// organization will be returned (organization doesn’t have ancestors).
// Comparing to ListAssignments, there are some behavior
// differences:
//
// permission on the assignee will be verified in this API.
//
// Hierarchy lookup (project->folder->organization) happens in this API.
//
// Parent here is projects//locations/, instead of
// projects/*/locations/reservations/.
// Note “-” cannot be used for projects
// nor locations.
func (c *Client) SearchAssignments(ctx context.Context, req *reservationpb.SearchAssignmentsRequest, opts ...gax.CallOption) *AssignmentIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.SearchAssignments[0:len(c.CallOptions.SearchAssignments):len(c.CallOptions.SearchAssignments)], opts...)
	it := &AssignmentIterator{}
	req = proto.Clone(req).(*reservationpb.SearchAssignmentsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*reservationpb.Assignment, string, error) {
		var resp *reservationpb.SearchAssignmentsResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.client.SearchAssignments(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.Assignments, resp.NextPageToken, nil
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

// MoveAssignment moves a assignment under a new reservation. Customers can do this by
// deleting the existing assignment followed by creating another assignment
// under the new reservation, but this method provides a transactional way to
// do so, to make sure the assignee always has an associated reservation.
// Without the method customers might see some queries run on-demand which
// might be unexpected.
func (c *Client) MoveAssignment(ctx context.Context, req *reservationpb.MoveAssignmentRequest, opts ...gax.CallOption) (*reservationpb.Assignment, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.MoveAssignment[0:len(c.CallOptions.MoveAssignment):len(c.CallOptions.MoveAssignment)], opts...)
	var resp *reservationpb.Assignment
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.MoveAssignment(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetBiReservation retrieves a BI reservation.
func (c *Client) GetBiReservation(ctx context.Context, req *reservationpb.GetBiReservationRequest, opts ...gax.CallOption) (*reservationpb.BiReservation, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.GetBiReservation[0:len(c.CallOptions.GetBiReservation):len(c.CallOptions.GetBiReservation)], opts...)
	var resp *reservationpb.BiReservation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetBiReservation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UpdateBiReservation updates a BI reservation.
// Only fields specified in the field_mask are updated.
// Singleton BI reservation always exists with default size 0.
// In order to reserve BI capacity it needs to be updated to an amount
// greater than 0. In order to release BI capacity reservation size
// must be set to 0.
func (c *Client) UpdateBiReservation(ctx context.Context, req *reservationpb.UpdateBiReservationRequest, opts ...gax.CallOption) (*reservationpb.BiReservation, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "reservation.name", url.QueryEscape(req.GetReservation().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.UpdateBiReservation[0:len(c.CallOptions.UpdateBiReservation):len(c.CallOptions.UpdateBiReservation)], opts...)
	var resp *reservationpb.BiReservation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.UpdateBiReservation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// AssignmentIterator manages a stream of *reservationpb.Assignment.
type AssignmentIterator struct {
	items    []*reservationpb.Assignment
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
	InternalFetch func(pageSize int, pageToken string) (results []*reservationpb.Assignment, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *AssignmentIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *AssignmentIterator) Next() (*reservationpb.Assignment, error) {
	var item *reservationpb.Assignment
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *AssignmentIterator) bufLen() int {
	return len(it.items)
}

func (it *AssignmentIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// CapacityCommitmentIterator manages a stream of *reservationpb.CapacityCommitment.
type CapacityCommitmentIterator struct {
	items    []*reservationpb.CapacityCommitment
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
	InternalFetch func(pageSize int, pageToken string) (results []*reservationpb.CapacityCommitment, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *CapacityCommitmentIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *CapacityCommitmentIterator) Next() (*reservationpb.CapacityCommitment, error) {
	var item *reservationpb.CapacityCommitment
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *CapacityCommitmentIterator) bufLen() int {
	return len(it.items)
}

func (it *CapacityCommitmentIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// ReservationIterator manages a stream of *reservationpb.Reservation.
type ReservationIterator struct {
	items    []*reservationpb.Reservation
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
	InternalFetch func(pageSize int, pageToken string) (results []*reservationpb.Reservation, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *ReservationIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ReservationIterator) Next() (*reservationpb.Reservation, error) {
	var item *reservationpb.Reservation
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ReservationIterator) bufLen() int {
	return len(it.items)
}

func (it *ReservationIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
