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
	"reflect"
	"testing"

	bq "google.golang.org/api/bigquery/v2"
)

func defaultLoadJob() *bq.Job {
	return &bq.Job{
		Configuration: &bq.JobConfiguration{
			Load: &bq.JobConfigurationLoad{
				DestinationTable: &bq.TableReference{
					ProjectId: "project-id",
					DatasetId: "dataset-id",
					TableId:   "table-id",
				},
				SourceUris: []string{"uri"},
			},
		},
	}
}

func stringFieldSchema() *FieldSchema {
	return &FieldSchema{Name: "fieldname", Type: StringFieldType}
}

func nestedFieldSchema() *FieldSchema {
	return &FieldSchema{
		Name:   "nested",
		Type:   RecordFieldType,
		Schema: Schema{stringFieldSchema()},
	}
}

func bqStringFieldSchema() *bq.TableFieldSchema {
	return &bq.TableFieldSchema{
		Name: "fieldname",
		Type: "STRING",
	}
}

func bqNestedFieldSchema() *bq.TableFieldSchema {
	return &bq.TableFieldSchema{
		Name:   "nested",
		Type:   "RECORD",
		Fields: []*bq.TableFieldSchema{bqStringFieldSchema()},
	}
}

func TestLoad(t *testing.T) {
	testCases := []struct {
		dst     Destination
		src     Source
		options []Option
		want    *bq.Job
	}{
		{
			dst:  defaultTable,
			src:  defaultGCS,
			want: defaultLoadJob(),
		},
		{
			dst: defaultTable,
			src: defaultGCS,
			options: []Option{
				MaxBadRecords(1),
				AllowJaggedRows(),
				AllowQuotedNewlines(),
				IgnoreUnknownValues(),
			},
			want: func() *bq.Job {
				j := defaultLoadJob()
				j.Configuration.Load.MaxBadRecords = 1
				j.Configuration.Load.AllowJaggedRows = true
				j.Configuration.Load.AllowQuotedNewlines = true
				j.Configuration.Load.IgnoreUnknownValues = true
				return j
			}(),
		},
		{
			dst: &Table{
				projectID:         "project-id",
				datasetID:         "dataset-id",
				tableID:           "table-id",
				CreateDisposition: "CREATE_NEVER",
				WriteDisposition:  "WRITE_TRUNCATE",
			},
			src: defaultGCS,
			want: func() *bq.Job {
				j := defaultLoadJob()
				j.Configuration.Load.CreateDisposition = "CREATE_NEVER"
				j.Configuration.Load.WriteDisposition = "WRITE_TRUNCATE"
				return j
			}(),
		},
		{
			dst: &Table{
				projectID: "project-id",
				datasetID: "dataset-id",
				tableID:   "table-id",
			},
			src: defaultGCS,
			options: []Option{
				DestinationSchema(Schema{
					stringFieldSchema(),
					nestedFieldSchema(),
				}),
			},
			want: func() *bq.Job {
				j := defaultLoadJob()
				j.Configuration.Load.Schema = &bq.TableSchema{
					Fields: []*bq.TableFieldSchema{
						bqStringFieldSchema(),
						bqNestedFieldSchema(),
					}}
				return j
			}(),
		},
		{
			dst: defaultTable,
			src: &GCSReference{
				uris:            []string{"uri"},
				SkipLeadingRows: 1,
				SourceFormat:    JSON,
				Encoding:        UTF_8,
				FieldDelimiter:  "\t",
				Quote:           "-",
			},
			want: func() *bq.Job {
				j := defaultLoadJob()
				j.Configuration.Load.SkipLeadingRows = 1
				j.Configuration.Load.SourceFormat = "NEWLINE_DELIMITED_JSON"
				j.Configuration.Load.Encoding = "UTF-8"
				j.Configuration.Load.FieldDelimiter = "\t"
				j.Configuration.Load.Quote = "-"
				return j
			}(),
		},
		{
			dst: defaultTable,
			src: &GCSReference{
				uris: []string{"uri"},
				// TODO(mcgreevy): Once the underlying API supports it, test that
				// a distinction is made between setting an empty Quote and not setting it at all.
				Quote: "",
			},
			want: func() *bq.Job {
				j := defaultLoadJob()
				j.Configuration.Load.Quote = ""
				return j
			}(),
		},
	}

	for _, tc := range testCases {
		jr := jobRecorder{}
		if _, err := load(&jr, tc.dst, tc.src, tc.options...); err != nil {
			t.Errorf("err calling load: %v", err)
			continue
		}
		if !reflect.DeepEqual(jr.Job, tc.want) {
			t.Errorf("insertJob got:\n%v\nwant:\n%v", jr.Job, tc.want)
		}
	}
}
