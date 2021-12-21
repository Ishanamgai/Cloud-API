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

package dialogflow_test

import (
	"context"

	dialogflow "cloud.google.com/go/dialogflow/apiv2"
	"google.golang.org/api/iterator"
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
)

func ExampleNewDocumentsClient() {
	ctx := context.Background()
	c, err := dialogflow.NewDocumentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleDocumentsClient_ListDocuments() {
	ctx := context.Background()
	c, err := dialogflow.NewDocumentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dialogflowpb.ListDocumentsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dialogflow/v2#ListDocumentsRequest.
	}
	it := c.ListDocuments(ctx, req)
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

func ExampleDocumentsClient_GetDocument() {
	ctx := context.Background()
	c, err := dialogflow.NewDocumentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dialogflowpb.GetDocumentRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dialogflow/v2#GetDocumentRequest.
	}
	resp, err := c.GetDocument(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDocumentsClient_CreateDocument() {
	ctx := context.Background()
	c, err := dialogflow.NewDocumentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dialogflowpb.CreateDocumentRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dialogflow/v2#CreateDocumentRequest.
	}
	op, err := c.CreateDocument(ctx, req)
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

func ExampleDocumentsClient_DeleteDocument() {
	ctx := context.Background()
	c, err := dialogflow.NewDocumentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dialogflowpb.DeleteDocumentRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dialogflow/v2#DeleteDocumentRequest.
	}
	op, err := c.DeleteDocument(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleDocumentsClient_UpdateDocument() {
	ctx := context.Background()
	c, err := dialogflow.NewDocumentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dialogflowpb.UpdateDocumentRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dialogflow/v2#UpdateDocumentRequest.
	}
	op, err := c.UpdateDocument(ctx, req)
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

func ExampleDocumentsClient_ReloadDocument() {
	ctx := context.Background()
	c, err := dialogflow.NewDocumentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dialogflowpb.ReloadDocumentRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dialogflow/v2#ReloadDocumentRequest.
	}
	op, err := c.ReloadDocument(ctx, req)
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

func ExampleDocumentsClient_ExportDocument() {
	ctx := context.Background()
	c, err := dialogflow.NewDocumentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dialogflowpb.ExportDocumentRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dialogflow/v2#ExportDocumentRequest.
	}
	op, err := c.ExportDocument(ctx, req)
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
