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

package storage

import (
	"net/http"
	"reflect"

	"cloud.google.com/go/internal/trace"
	"golang.org/x/net/context"
	"google.golang.org/api/googleapi"
	raw "google.golang.org/api/storage/v1"
)

// ACLRole is the level of access to grant.
type ACLRole string

const (
	RoleOwner  ACLRole = "OWNER"
	RoleReader ACLRole = "READER"
	RoleWriter ACLRole = "WRITER"
)

// ACLEntity refers to a user or group.
// They are sometimes referred to as grantees.
//
// It could be in the form of:
// "user-<userId>", "user-<email>", "group-<groupId>", "group-<email>",
// "domain-<domain>" and "project-team-<projectId>".
//
// Or one of the predefined constants: AllUsers, AllAuthenticatedUsers.
type ACLEntity string

const (
	AllUsers              ACLEntity = "allUsers"
	AllAuthenticatedUsers ACLEntity = "allAuthenticatedUsers"
)

// ACLRule represents a grant for a role to an entity (user, group or team) for a Google Cloud Storage object or bucket.
type ACLRule struct {
	Entity ACLEntity
	Role   ACLRole
}

// ACLHandle provides operations on an access control list for a Google Cloud Storage bucket or object.
type ACLHandle struct {
	c           *Client
	bucket      string
	object      string
	isDefault   bool
	userProject string // for requester-pays buckets
}

// Delete permanently deletes the ACL entry for the given entity.
func (a *ACLHandle) Delete(ctx context.Context, entity ACLEntity) (err error) {
	ctx = trace.StartSpan(ctx, "cloud.google.com/go/storage.ACL.Delete")
	defer func() { trace.EndSpan(ctx, err) }()

	if a.object != "" {
		return a.objectDelete(ctx, entity)
	}
	if a.isDefault {
		return a.bucketDefaultDelete(ctx, entity)
	}
	return a.bucketDelete(ctx, entity)
}

// Set sets the permission level for the given entity.
func (a *ACLHandle) Set(ctx context.Context, entity ACLEntity, role ACLRole) (err error) {
	ctx = trace.StartSpan(ctx, "cloud.google.com/go/storage.ACL.Set")
	defer func() { trace.EndSpan(ctx, err) }()

	if a.object != "" {
		return a.objectSet(ctx, entity, role, false)
	}
	if a.isDefault {
		return a.objectSet(ctx, entity, role, true)
	}
	return a.bucketSet(ctx, entity, role)
}

// List retrieves ACL entries.
func (a *ACLHandle) List(ctx context.Context) (_ []ACLRule, err error) {
	ctx = trace.StartSpan(ctx, "cloud.google.com/go/storage.ACL.List")
	defer func() { trace.EndSpan(ctx, err) }()

	if a.object != "" {
		return a.objectList(ctx)
	}
	if a.isDefault {
		return a.bucketDefaultList(ctx)
	}
	return a.bucketList(ctx)
}

func (a *ACLHandle) bucketDefaultList(ctx context.Context) ([]ACLRule, error) {
	var acls *raw.ObjectAccessControls
	var err error
	err = runWithRetry(ctx, func() error {
		req := a.c.raw.DefaultObjectAccessControls.List(a.bucket)
		a.configureCall(req, ctx)
		acls, err = req.Do()
		return err
	})
	if err != nil {
		return nil, err
	}
	return toACLRules(acls.Items), nil
}

func (a *ACLHandle) bucketDefaultDelete(ctx context.Context, entity ACLEntity) error {
	return runWithRetry(ctx, func() error {
		req := a.c.raw.DefaultObjectAccessControls.Delete(a.bucket, string(entity))
		a.configureCall(req, ctx)
		return req.Do()
	})
}

func (a *ACLHandle) bucketList(ctx context.Context) ([]ACLRule, error) {
	var acls *raw.BucketAccessControls
	var err error
	err = runWithRetry(ctx, func() error {
		req := a.c.raw.BucketAccessControls.List(a.bucket)
		a.configureCall(req, ctx)
		acls, err = req.Do()
		return err
	})
	if err != nil {
		return nil, err
	}
	r := make([]ACLRule, len(acls.Items))
	for i, v := range acls.Items {
		r[i].Entity = ACLEntity(v.Entity)
		r[i].Role = ACLRole(v.Role)
	}
	return r, nil
}

func (a *ACLHandle) bucketSet(ctx context.Context, entity ACLEntity, role ACLRole) error {
	acl := &raw.BucketAccessControl{
		Bucket: a.bucket,
		Entity: string(entity),
		Role:   string(role),
	}
	err := runWithRetry(ctx, func() error {
		req := a.c.raw.BucketAccessControls.Update(a.bucket, string(entity), acl)
		a.configureCall(req, ctx)
		_, err := req.Do()
		return err
	})
	if err != nil {
		return err
	}
	return nil
}

func (a *ACLHandle) bucketDelete(ctx context.Context, entity ACLEntity) error {
	err := runWithRetry(ctx, func() error {
		req := a.c.raw.BucketAccessControls.Delete(a.bucket, string(entity))
		a.configureCall(req, ctx)
		return req.Do()
	})
	if err != nil {
		return err
	}
	return nil
}

func (a *ACLHandle) objectList(ctx context.Context) ([]ACLRule, error) {
	var acls *raw.ObjectAccessControls
	var err error
	err = runWithRetry(ctx, func() error {
		req := a.c.raw.ObjectAccessControls.List(a.bucket, a.object)
		a.configureCall(req, ctx)
		acls, err = req.Do()
		return err
	})
	if err != nil {
		return nil, err
	}
	return toACLRules(acls.Items), nil
}

func (a *ACLHandle) objectSet(ctx context.Context, entity ACLEntity, role ACLRole, isBucketDefault bool) error {
	type setRequest interface {
		Do(opts ...googleapi.CallOption) (*raw.ObjectAccessControl, error)
		Header() http.Header
	}

	acl := &raw.ObjectAccessControl{
		Bucket: a.bucket,
		Entity: string(entity),
		Role:   string(role),
	}
	var req setRequest
	if isBucketDefault {
		req = a.c.raw.DefaultObjectAccessControls.Update(a.bucket, string(entity), acl)
	} else {
		req = a.c.raw.ObjectAccessControls.Update(a.bucket, a.object, string(entity), acl)
	}
	a.configureCall(req, ctx)
	return runWithRetry(ctx, func() error {
		_, err := req.Do()
		return err
	})
}

func (a *ACLHandle) objectDelete(ctx context.Context, entity ACLEntity) error {
	return runWithRetry(ctx, func() error {
		req := a.c.raw.ObjectAccessControls.Delete(a.bucket, a.object, string(entity))
		a.configureCall(req, ctx)
		return req.Do()
	})
}

func (a *ACLHandle) configureCall(call interface {
	Header() http.Header
}, ctx context.Context) {
	vc := reflect.ValueOf(call)
	vc.MethodByName("Context").Call([]reflect.Value{reflect.ValueOf(ctx)})
	if a.userProject != "" {
		vc.MethodByName("UserProject").Call([]reflect.Value{reflect.ValueOf(a.userProject)})
	}
	setClientHeader(call.Header())
}

func toACLRules(items []*raw.ObjectAccessControl) []ACLRule {
	r := make([]ACLRule, 0, len(items))
	for _, item := range items {
		r = append(r, ACLRule{Entity: ACLEntity(item.Entity), Role: ACLRole(item.Role)})
	}
	return r
}
