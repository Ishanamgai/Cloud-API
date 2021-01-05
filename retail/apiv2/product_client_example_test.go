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

package retail_test

import (
	"context"

	retail "cloud.google.com/go/retail/apiv2"
	retailpb "google.golang.org/genproto/googleapis/cloud/retail/v2"
)

func ExampleNewProductClient() {
	ctx := context.Background()
	c, err := retail.NewProductClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use client.
	_ = c
}

func ExampleProductClient_CreateProduct() {
	// import retailpb "google.golang.org/genproto/googleapis/cloud/retail/v2"

	ctx := context.Background()
	c, err := retail.NewProductClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &retailpb.CreateProductRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateProduct(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleProductClient_GetProduct() {
	// import retailpb "google.golang.org/genproto/googleapis/cloud/retail/v2"

	ctx := context.Background()
	c, err := retail.NewProductClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &retailpb.GetProductRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetProduct(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleProductClient_UpdateProduct() {
	// import retailpb "google.golang.org/genproto/googleapis/cloud/retail/v2"

	ctx := context.Background()
	c, err := retail.NewProductClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &retailpb.UpdateProductRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateProduct(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleProductClient_DeleteProduct() {
	ctx := context.Background()
	c, err := retail.NewProductClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &retailpb.DeleteProductRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteProduct(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleProductClient_ImportProducts() {
	// import retailpb "google.golang.org/genproto/googleapis/cloud/retail/v2"

	ctx := context.Background()
	c, err := retail.NewProductClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &retailpb.ImportProductsRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.ImportProducts(ctx, req)
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
