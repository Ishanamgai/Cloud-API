// Copyright 2022 Google LLC
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

package oslogin_test

import (
	"context"

	oslogin "cloud.google.com/go/oslogin/apiv1beta"
	osloginpb "google.golang.org/genproto/googleapis/cloud/oslogin/v1beta"
)

func ExampleNewClient() {
	ctx := context.Background()
	c, err := oslogin.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleClient_DeletePosixAccount() {
	ctx := context.Background()
	c, err := oslogin.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &osloginpb.DeletePosixAccountRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/oslogin/v1beta#DeletePosixAccountRequest.
	}
	err = c.DeletePosixAccount(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_DeleteSshPublicKey() {
	ctx := context.Background()
	c, err := oslogin.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &osloginpb.DeleteSshPublicKeyRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/oslogin/v1beta#DeleteSshPublicKeyRequest.
	}
	err = c.DeleteSshPublicKey(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_GetLoginProfile() {
	ctx := context.Background()
	c, err := oslogin.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &osloginpb.GetLoginProfileRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/oslogin/v1beta#GetLoginProfileRequest.
	}
	resp, err := c.GetLoginProfile(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_GetSshPublicKey() {
	ctx := context.Background()
	c, err := oslogin.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &osloginpb.GetSshPublicKeyRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/oslogin/v1beta#GetSshPublicKeyRequest.
	}
	resp, err := c.GetSshPublicKey(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ImportSshPublicKey() {
	ctx := context.Background()
	c, err := oslogin.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &osloginpb.ImportSshPublicKeyRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/oslogin/v1beta#ImportSshPublicKeyRequest.
	}
	resp, err := c.ImportSshPublicKey(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_UpdateSshPublicKey() {
	ctx := context.Background()
	c, err := oslogin.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &osloginpb.UpdateSshPublicKeyRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/oslogin/v1beta#UpdateSshPublicKeyRequest.
	}
	resp, err := c.UpdateSshPublicKey(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
