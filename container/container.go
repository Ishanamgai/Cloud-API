// Copyright 2014 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package container contains a Google Container Engine client.
//
// For more information about the API,
// see https://cloud.google.com/container-engine/docs
package container // import "google.golang.org/cloud/container"

import (
	"errors"
	"time"

	raw "code.google.com/p/google-api-go-client/container/v1beta1"
	"golang.org/x/net/context"
	"google.golang.org/cloud/internal"
)

type Type string

var (
	TypeCreate Type = Type("createCluster")
	TypeDelete Type = Type("deleteCluster")
)

type Status string

var (
	Done         = Status("done")
	Pending      = Status("pending")
	Running      = Status("running")
	Error        = Status("error")
	Provisioning = Status("provisioning")
	Stopping     = Status("stopping")
)

// Resource is a Google Container Engine cluster resource.
type Resource struct {
	// Name is the name of this cluster. The name must be unique
	// within this project and zone, and can be up to 40 characters.
	Name string

	// Description is the description of the cluster. Optional.
	Description string

	// Zone is the Google Compute Engine zone in which the cluster resides.
	Zone string

	// Status is the current status of the cluster. It could either be
	// StatusError, StatusProvisioning, StatusRunning or StatusStopping.
	Status Status

	// Num is the number of the nodes in this cluster resource.
	Num int64

	// APIVersion is the version of the Kubernetes master and kubelets running
	// in this cluster. Allowed value is 0.4.2, or leave blank to
	// pick up the latest stable release.
	APIVersion string

	// Endpoint is the IP address of this cluster's Kubernetes master.
	// The endpoint can be accessed at https://username:password@endpoint/.
	// See Username and Password fields for the username and password information.
	Endpoint string

	// Username is the username to use when accessing the Kubernetes master endpoint.
	Username string

	// Password is the password to use when accessing the Kubernetes master endpoint.
	Password string

	// ContainerIPv4CIDR is the IP addresses of the container pods in
	// this cluster, in CIDR notation (e.g. 1.2.3.4/29).
	ContainerIPv4CIDR string

	// ServicesIPv4CIDR is the IP addresses of the Kubernetes services in this
	// cluster, in CIDR notation (e.g. 1.2.3.4/29). Service addresses are
	// always in the 10.0.0.0/16 range.
	ServicesIPv4CIDR string

	// MachineType is a Google Compute Engine machine type (e.g. n1-standard-1).
	// If none set, the default type is used while creating a new cluster.
	MachineType string

	// SourceImage is the fully-specified name of a Google Compute Engine image.
	// For example: https://www.googleapis.com/compute/v1/projects/debian-cloud/global/images/backports-debian-7-wheezy-vYYYYMMDD.
	SourceImage string

	// Created is the creation time of this cluster.
	Created time.Time
}

func resourceFromRaw(c *raw.Cluster) *Resource {
	if c == nil {
		return nil
	}
	r := &Resource{
		Name:              c.Name,
		Description:       c.Description,
		Zone:              c.Zone,
		Status:            Status(c.Status),
		Num:               c.NumNodes,
		APIVersion:        c.ClusterApiVersion,
		Endpoint:          c.Endpoint,
		Username:          c.MasterAuth.User,
		Password:          c.MasterAuth.Password,
		ContainerIPv4CIDR: c.ContainerIpv4Cidr,
		ServicesIPv4CIDR:  c.ServicesIpv4Cidr,
		MachineType:       c.NodeConfig.MachineType,
		SourceImage:       c.NodeConfig.SourceImage,
	}
	r.Created, _ = time.Parse(time.RFC3339, c.CreationTimestamp)
	return r
}

func resourcesFromRaw(c []*raw.Cluster) []*Resource {
	r := make([]*Resource, len(c))
	for i, val := range c {
		r[i] = resourceFromRaw(val)
	}
	return r
}

// Op represents a Google Container Engine API operation.
type Op struct {
	// Name is the name of the operation.
	Name string

	// Zone is the Google Compute Engine zone.
	Zone string

	// TargetURL is the URL of the cluster resource
	// that this operation is associated with.
	TargetURL string

	// Type is the operation type. It could be either be TypeCreate or TypeDelete.
	Type Type

	// Status is the current status of this operation. It could be either
	// OpDone or OpPending.
	Status Status
}

func opFromRaw(o *raw.Operation) *Op {
	if o == nil {
		return nil
	}
	return &Op{
		Name:      o.Name,
		Zone:      o.Zone,
		TargetURL: o.Target,
		Type:      Type(o.OperationType),
		Status:    Status(o.Status),
	}
}

func opsFromRaw(o []*raw.Operation) []*Op {
	ops := make([]*Op, len(o))
	for i, val := range o {
		ops[i] = opFromRaw(val)
	}
	return ops
}

// Clusters returns a list of cluster resources from the specified zone.
// If no zone is specified, it returns all clusters under the user project.
func Clusters(ctx context.Context, zone string) ([]*Resource, error) {
	s := rawService(ctx)
	if zone == "" {
		resp, err := s.Projects.Clusters.List(internal.ProjID(ctx)).Do()
		if err != nil {
			return nil, err
		}
		return resourcesFromRaw(resp.Clusters), nil
	}
	resp, err := s.Projects.Zones.Clusters.List(internal.ProjID(ctx), zone).Do()
	if err != nil {
		return nil, err
	}
	return resourcesFromRaw(resp.Clusters), nil
}

// Cluster returns metadata about the specified cluster.
func Cluster(ctx context.Context, zone, name string) (*Resource, error) {
	s := rawService(ctx)
	resp, err := s.Projects.Zones.Clusters.Get(internal.ProjID(ctx), zone, name).Do()
	if err != nil {
		return nil, err
	}
	return resourceFromRaw(resp), nil
}

// CreateCluster creates a new cluster with the provided metadata
// in the specified zone.
func CreateCluster(ctx context.Context, zone string, resource *Resource) (*Resource, error) {
	panic("not implemented")
}

// DeleteCluster deletes a cluster.
func DeleteCluster(ctx context.Context, zone, name string) error {
	s := rawService(ctx)
	_, err := s.Projects.Zones.Clusters.Delete(internal.ProjID(ctx), zone, name).Do()
	return err
}

// Operations returns a list of operations from the specified zone.
// If no zone is specified, it looks up for all of the operations
// that are running under the user's project.
func Operations(ctx context.Context, zone string) ([]*Op, error) {
	s := rawService(ctx)
	if zone == "" {
		resp, err := s.Projects.Operations.List(internal.ProjID(ctx)).Do()
		if err != nil {
			return nil, err
		}
		return opsFromRaw(resp.Operations), nil
	}
	resp, err := s.Projects.Zones.Operations.List(internal.ProjID(ctx), zone).Do()
	if err != nil {
		return nil, err
	}
	return opsFromRaw(resp.Operations), nil
}

// Operation returns an operation.
func Operation(ctx context.Context, zone, name string) (*Op, error) {
	s := rawService(ctx)
	resp, err := s.Projects.Zones.Operations.Get(internal.ProjID(ctx), zone, name).Do()
	if err != nil {
		return nil, err
	}
	if resp.ErrorMessage != "" {
		return nil, errors.New(resp.ErrorMessage)
	}
	return opFromRaw(resp), nil
}

func rawService(ctx context.Context) *raw.Service {
	return ctx.Value(internal.ContextKey("base")).(map[string]interface{})["container_service"].(*raw.Service)
}
