// Copyright 2015 Google Inc. All Rights Reserved.
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

package bigquery

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"sort"
	"strings"
	"testing"
	"time"

	"cloud.google.com/go/internal/pretty"
	"cloud.google.com/go/internal/testutil"
	"golang.org/x/net/context"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

var (
	client  *Client
	dataset *Dataset
	schema  = Schema([]*FieldSchema{
		{Name: "name", Type: StringFieldType},
		{Name: "num", Type: IntegerFieldType},
	})
)

func TestMain(m *testing.M) {
	initIntegrationTest()
	os.Exit(m.Run())
}

// If integration tests will be run, create a unique bucket for them.
func initIntegrationTest() {
	flag.Parse() // needed for testing.Short()
	if testing.Short() {
		return
	}
	ctx := context.Background()
	ts := testutil.TokenSource(ctx, Scope)
	if ts == nil {
		log.Println("Integration tests skipped. See CONTRIBUTING.md for details")
		return
	}
	projID := testutil.ProjID()
	var err error
	client, err = NewClient(ctx, projID, option.WithTokenSource(ts))
	if err != nil {
		log.Fatalf("NewClient: %v", err)
	}
	dataset = client.Dataset("bigquery_integration_test")
	if err := dataset.Create(ctx); err != nil && !hasStatusCode(err, http.StatusConflict) { // AlreadyExists is 409
		log.Fatalf("creating dataset: %v", err)
	}
}

func TestIntegration_TableMetadata(t *testing.T) {
	if client == nil {
		t.Skip("Integration tests skipped")
	}
	ctx := context.Background()
	table := newTable(t, schema)
	defer table.Delete(ctx)
	// Check table metadata.
	md, err := table.Metadata(ctx)
	if err != nil {
		t.Fatal(err)
	}
	// TODO(jba): check md more thorougly.
	if got, want := md.ID, fmt.Sprintf("%s:%s.%s", dataset.ProjectID, dataset.DatasetID, table.TableID); got != want {
		t.Errorf("metadata.ID: got %q, want %q", got, want)
	}
	if got, want := md.Type, RegularTable; got != want {
		t.Errorf("metadata.Type: got %v, want %v", got, want)
	}
}

func TestIntegration_Tables(t *testing.T) {
	if client == nil {
		t.Skip("Integration tests skipped")
	}
	ctx := context.Background()
	table := newTable(t, schema)
	defer table.Delete(ctx)

	// Iterate over tables in the dataset.
	it := dataset.Tables(ctx)
	var tables []*Table
	for {
		tbl, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
		tables = append(tables, tbl)
	}
	if got, want := tables, []*Table{table}; !reflect.DeepEqual(got, want) {
		t.Errorf("Tables: got %v, want %v", pretty.Value(got), pretty.Value(want))
	}
}

func TestIntegration_UploadAndRead(t *testing.T) {
	if client == nil {
		t.Skip("Integration tests skipped")
	}
	ctx := context.Background()
	table := newTable(t, schema)
	defer table.Delete(ctx)

	// Populate the table.
	upl := table.Uploader()
	var (
		wantRows  [][]Value
		saverRows []*ValuesSaver
	)
	for i, name := range []string{"a", "b", "c"} {
		row := []Value{name, i}
		wantRows = append(wantRows, row)
		saverRows = append(saverRows, &ValuesSaver{
			Schema:   schema,
			InsertID: name,
			Row:      row,
		})
	}
	if err := upl.Put(ctx, saverRows); err != nil {
		t.Fatal(err)
	}

	// Wait until the data has been uploaded. This can take a few seconds, according
	// to https://cloud.google.com/bigquery/streaming-data-into-bigquery.
	for {
		it := table.Read(ctx)
		var v ValueList
		err := it.Next(&v)
		if err == nil {
			break
		}
		if err != iterator.Done {
			t.Fatal(err)
		}
		time.Sleep(1 * time.Second)
	}

	// Read the table.
	checkRead(t, "upload", table.Read(ctx), wantRows)

	// Query the table.
	q := client.Query(fmt.Sprintf("select name, num from %s", table.TableID))
	q.DefaultProjectID = dataset.ProjectID
	q.DefaultDatasetID = dataset.DatasetID
	rit, err := q.Read(ctx)
	if err != nil {
		t.Fatal(err)
	}
	checkRead(t, "query", rit, wantRows)

	// Query the long way.
	job1, err := q.Run(ctx)
	if err != nil {
		t.Fatal(err)
	}
	job2, err := client.JobFromID(ctx, job1.ID())
	if err != nil {
		t.Fatal(err)
	}
	// TODO(jba): poll status until job is done
	_, err = job2.Status(ctx)
	if err != nil {
		t.Fatal(err)
	}

	rit, err = job2.Read(ctx)
	if err != nil {
		t.Fatal(err)
	}
	checkRead(t, "job.Read", rit, wantRows)
}

func TestIntegration_Update(t *testing.T) {
	if client == nil {
		t.Skip("Integration tests skipped")
	}
	ctx := context.Background()
	table := newTable(t, schema)
	defer table.Delete(ctx)

	// Test Update.
	tm, err := table.Metadata(ctx)
	if err != nil {
		t.Fatal(err)
	}
	wantDescription := tm.Description + "more"
	wantName := tm.Name + "more"
	got, err := table.Update(ctx, TableMetadataToUpdate{
		Description: wantDescription,
		Name:        wantName,
	})
	if err != nil {
		t.Fatal(err)
	}
	if got.Description != wantDescription {
		t.Errorf("Description: got %q, want %q", got.Description, wantDescription)
	}
	if got.Name != wantName {
		t.Errorf("Name: got %q, want %q", got.Name, wantName)
	}
}

func TestIntegration_Load(t *testing.T) {
	if client == nil {
		t.Skip("Integration tests skipped")
	}
	ctx := context.Background()
	table := newTable(t, schema)
	defer table.Delete(ctx)

	// Load the table from a reader.
	r := strings.NewReader("a,0\nb,1\nc,2\n")
	wantRows := [][]Value{
		[]Value{"a", 0},
		[]Value{"b", 1},
		[]Value{"c", 2},
	}
	rs := NewReaderSource(r)
	loader := table.LoaderFrom(rs)
	loader.WriteDisposition = WriteTruncate
	job, err := loader.Run(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if err := wait(ctx, job); err != nil {
		t.Fatal(err)
	}
	checkRead(t, "reader load", table.Read(ctx), wantRows)
}

func TestIntegration_DML(t *testing.T) {
	if client == nil {
		t.Skip("Integration tests skipped")
	}
	ctx := context.Background()
	table := newTable(t, schema)
	defer table.Delete(ctx)

	// Use DML to insert.
	wantRows := [][]Value{
		[]Value{"a", 0},
		[]Value{"b", 1},
		[]Value{"c", 2},
	}
	query := fmt.Sprintf("INSERT bigquery_integration_test.%s (name, num) "+
		"VALUES ('a', 0), ('b', 1), ('c', 2)",
		table.TableID)
	q := client.Query(query)
	q.UseStandardSQL = true // necessary for DML
	job, err := q.Run(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if err := wait(ctx, job); err != nil {
		t.Fatal(err)
	}
	checkRead(t, "INSERT", table.Read(ctx), wantRows)
}

// Creates a new, temporary table with a unique name and the given schema.
func newTable(t *testing.T, s Schema) *Table {
	name := fmt.Sprintf("t%d", time.Now().UnixNano())
	table := dataset.Table(name)
	err := table.Create(context.Background(), s, TableExpiration(time.Now().Add(5*time.Minute)))
	if err != nil {
		t.Fatal(err)
	}
	return table
}

func checkRead(t *testing.T, msg string, it *RowIterator, want [][]Value) {
	got, err := readAll(it)
	if err != nil {
		t.Fatalf("%s: %v", msg, err)
	}
	if len(got) != len(want) {
		t.Errorf("%s: got %d rows, want %d", msg, len(got), len(want))
	}
	sort.Sort(byCol0(got))
	for i, r := range got {
		gotRow := []Value(r)
		wantRow := want[i]
		if !reflect.DeepEqual(gotRow, wantRow) {
			t.Errorf("%s #%d: got %v, want %v", msg, i, gotRow, wantRow)
		}
	}
}

func readAll(it *RowIterator) ([]ValueList, error) {
	var rows []ValueList
	for {
		var vals ValueList
		err := it.Next(&vals)
		if err == iterator.Done {
			return rows, nil
		}
		if err != nil {
			return nil, err
		}
		rows = append(rows, vals)
	}
}

type byCol0 []ValueList

func (b byCol0) Len() int           { return len(b) }
func (b byCol0) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b byCol0) Less(i, j int) bool { return b[i][0].(string) < b[j][0].(string) }

func hasStatusCode(err error, code int) bool {
	if e, ok := err.(*googleapi.Error); ok && e.Code == code {
		return true
	}
	return false
}

// wait polls the job until it is complete or an error is returned.
func wait(ctx context.Context, job *Job) error {
	for {
		status, err := job.Status(ctx)
		if err != nil {
			return fmt.Errorf("getting job status: %v", err)
		}
		if status.Done() {
			if status.Err() != nil {
				return fmt.Errorf("job status: %#v", status.Err())
			}
			return nil
		}
		time.Sleep(1 * time.Second)
	}
}
