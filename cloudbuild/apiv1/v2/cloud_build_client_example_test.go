// Copyright 2019 Google LLC
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

package cloudbuild_test

import (
	"context"

	cloudbuild "cloud.google.com/go/cloudbuild/apiv1/v2"
	"google.golang.org/api/iterator"
	cloudbuildpb "google.golang.org/genproto/googleapis/devtools/cloudbuild/v1"
)

func ExampleNewClient() {
	ctx := context.Background()
	c, err := cloudbuild.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use client.
	_ = c
}

func ExampleClient_CreateBuild() {
	// import cloudbuildpb "google.golang.org/genproto/googleapis/devtools/cloudbuild/v1"

	ctx := context.Background()
	c, err := cloudbuild.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &cloudbuildpb.CreateBuildRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.CreateBuild(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_GetBuild() {
	// import cloudbuildpb "google.golang.org/genproto/googleapis/devtools/cloudbuild/v1"

	ctx := context.Background()
	c, err := cloudbuild.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &cloudbuildpb.GetBuildRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetBuild(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListBuilds() {
	// import cloudbuildpb "google.golang.org/genproto/googleapis/devtools/cloudbuild/v1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := cloudbuild.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &cloudbuildpb.ListBuildsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListBuilds(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_CancelBuild() {
	// import cloudbuildpb "google.golang.org/genproto/googleapis/devtools/cloudbuild/v1"

	ctx := context.Background()
	c, err := cloudbuild.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &cloudbuildpb.CancelBuildRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CancelBuild(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_RetryBuild() {
	// import cloudbuildpb "google.golang.org/genproto/googleapis/devtools/cloudbuild/v1"

	ctx := context.Background()
	c, err := cloudbuild.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &cloudbuildpb.RetryBuildRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.RetryBuild(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_CreateBuildTrigger() {
	// import cloudbuildpb "google.golang.org/genproto/googleapis/devtools/cloudbuild/v1"

	ctx := context.Background()
	c, err := cloudbuild.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &cloudbuildpb.CreateBuildTriggerRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateBuildTrigger(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_GetBuildTrigger() {
	// import cloudbuildpb "google.golang.org/genproto/googleapis/devtools/cloudbuild/v1"

	ctx := context.Background()
	c, err := cloudbuild.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &cloudbuildpb.GetBuildTriggerRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetBuildTrigger(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListBuildTriggers() {
	// import cloudbuildpb "google.golang.org/genproto/googleapis/devtools/cloudbuild/v1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := cloudbuild.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &cloudbuildpb.ListBuildTriggersRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListBuildTriggers(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_DeleteBuildTrigger() {
	ctx := context.Background()
	c, err := cloudbuild.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &cloudbuildpb.DeleteBuildTriggerRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteBuildTrigger(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_UpdateBuildTrigger() {
	// import cloudbuildpb "google.golang.org/genproto/googleapis/devtools/cloudbuild/v1"

	ctx := context.Background()
	c, err := cloudbuild.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &cloudbuildpb.UpdateBuildTriggerRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateBuildTrigger(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_RunBuildTrigger() {
	// import cloudbuildpb "google.golang.org/genproto/googleapis/devtools/cloudbuild/v1"

	ctx := context.Background()
	c, err := cloudbuild.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &cloudbuildpb.RunBuildTriggerRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.RunBuildTrigger(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_CreateWorkerPool() {
	// import cloudbuildpb "google.golang.org/genproto/googleapis/devtools/cloudbuild/v1"

	ctx := context.Background()
	c, err := cloudbuild.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &cloudbuildpb.CreateWorkerPoolRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateWorkerPool(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_GetWorkerPool() {
	// import cloudbuildpb "google.golang.org/genproto/googleapis/devtools/cloudbuild/v1"

	ctx := context.Background()
	c, err := cloudbuild.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &cloudbuildpb.GetWorkerPoolRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetWorkerPool(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_DeleteWorkerPool() {
	ctx := context.Background()
	c, err := cloudbuild.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &cloudbuildpb.DeleteWorkerPoolRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteWorkerPool(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_UpdateWorkerPool() {
	// import cloudbuildpb "google.golang.org/genproto/googleapis/devtools/cloudbuild/v1"

	ctx := context.Background()
	c, err := cloudbuild.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &cloudbuildpb.UpdateWorkerPoolRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateWorkerPool(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListWorkerPools() {
	// import cloudbuildpb "google.golang.org/genproto/googleapis/devtools/cloudbuild/v1"

	ctx := context.Background()
	c, err := cloudbuild.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &cloudbuildpb.ListWorkerPoolsRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.ListWorkerPools(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
