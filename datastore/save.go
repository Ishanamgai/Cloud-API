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

package datastore

import (
	"errors"
	"fmt"
	"reflect"
	"time"

	timepb "github.com/golang/protobuf/ptypes/timestamp"
	pb "google.golang.org/genproto/googleapis/datastore/v1beta3"
	llpb "google.golang.org/genproto/googleapis/type/latlng"
)

// saveEntity saves an EntityProto into a PropertyLoadSaver or struct pointer.
func saveEntity(key *Key, src interface{}) (*pb.Entity, error) {
	var err error
	var props []Property
	if e, ok := src.(PropertyLoadSaver); ok {
		props, err = e.Save()
	} else {
		props, err = SaveStruct(src)
	}
	if err != nil {
		return nil, err
	}
	return propertiesToProto(key, props)
}

// TODO(djd): Convert this and below to return ([]Property, error).
func saveStructProperty(props *[]Property, name string, noIndex bool, v reflect.Value) error {
	p := Property{
		Name:    name,
		NoIndex: noIndex,
	}

	switch x := v.Interface().(type) {
	case *Key, time.Time, GeoPoint:
		p.Value = x
	default:
		switch v.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			p.Value = v.Int()
		case reflect.Bool:
			p.Value = v.Bool()
		case reflect.String:
			p.Value = v.String()
		case reflect.Float32, reflect.Float64:
			p.Value = v.Float()
		case reflect.Slice:
			if v.Type().Elem().Kind() == reflect.Uint8 {
				p.Value = v.Bytes()
			} else {
				return saveSliceProperty(props, name, noIndex, v)
			}
		case reflect.Struct:
			if !v.CanAddr() {
				return fmt.Errorf("datastore: unsupported struct field: value is unaddressable")
			}
			sub, err := newStructPLS(v.Addr().Interface())
			if err != nil {
				return fmt.Errorf("datastore: unsupported struct field: %v", err)
			}
			return sub.(structPLS).save(props, name, noIndex)
		}
	}
	if p.Value == nil {
		return fmt.Errorf("datastore: unsupported struct field type: %v", v.Type())
	}
	*props = append(*props, p)
	return nil
}

func saveSliceProperty(props *[]Property, name string, noIndex bool, v reflect.Value) error {
	// Easy case: if the slice is empty, we're done.
	if v.Len() == 0 {
		return nil
	}
	// Work out the properties generated by the first element in the slice. This will
	// usually be a single property, but will be more if this is a slice of structs.
	var headProps []Property
	if err := saveStructProperty(&headProps, name, noIndex, v.Index(0)); err != nil {
		return err
	}

	// Convert the first element's properties into slice properties, and
	// keep track of the values in a map.
	values := make(map[string][]interface{}, len(headProps))
	for _, p := range headProps {
		values[p.Name] = append(make([]interface{}, 0, v.Len()), p.Value)
	}

	// Find the elements for the subsequent elements.
	for i := 1; i < v.Len(); i++ {
		elemProps := make([]Property, 0, len(headProps))
		if err := saveStructProperty(&elemProps, name, noIndex, v.Index(i)); err != nil {
			return err
		}
		for _, p := range elemProps {
			v, ok := values[p.Name]
			if !ok {
				return fmt.Errorf("datastore: unexpected property %q in elem %d of slice", p.Name, i)
			}
			values[p.Name] = append(v, p.Value)
		}
	}

	// Convert to the final properties.
	for _, p := range headProps {
		p.Value = values[p.Name]
		*props = append(*props, p)
	}
	return nil
}

func (s structPLS) Save() ([]Property, error) {
	var props []Property
	if err := s.save(&props, "", false); err != nil {
		return nil, err
	}
	return props, nil
}

func (s structPLS) save(props *[]Property, prefix string, noIndex bool) error {
	for i, t := range s.codec.byIndex {
		if t.name == "-" {
			continue
		}
		name := t.name
		if prefix != "" {
			name = prefix + name
		}
		v := s.v.Field(i)
		if !v.IsValid() || !v.CanSet() {
			continue
		}
		noIndex1 := noIndex || t.noIndex
		if err := saveStructProperty(props, name, noIndex1, v); err != nil {
			return err
		}
	}
	return nil
}

func propertiesToProto(key *Key, props []Property) (*pb.Entity, error) {
	e := &pb.Entity{
		Key:        keyToProto(key),
		Properties: map[string]*pb.Value{},
	}
	indexedProps := 0
	for _, p := range props {
		val, err := interfaceToProto(p.Value, p.NoIndex)
		if err != nil {
			return nil, fmt.Errorf("datastore: %v for a Property with Name %q", err, p.Name)
		}
		if !p.NoIndex {
			rVal := reflect.ValueOf(p.Value)
			if rVal.Kind() == reflect.Slice && rVal.Type().Elem().Kind() != reflect.Uint8 {
				indexedProps += rVal.Len()
			} else {
				indexedProps++
			}
		}
		if indexedProps > maxIndexedProperties {
			return nil, errors.New("datastore: too many indexed properties")
		}

		if _, ok := e.Properties[p.Name]; ok {
			return nil, fmt.Errorf("datastore: duplicate Property with Name %q", p.Name)
		}
		e.Properties[p.Name] = val
	}
	return e, nil
}

func interfaceToProto(iv interface{}, noIndex bool) (*pb.Value, error) {
	val := &pb.Value{ExcludeFromIndexes: noIndex}
	switch v := iv.(type) {
	case int:
		val.ValueType = &pb.Value_IntegerValue{int64(v)}
	case int32:
		val.ValueType = &pb.Value_IntegerValue{int64(v)}
	case int64:
		val.ValueType = &pb.Value_IntegerValue{v}
	case bool:
		val.ValueType = &pb.Value_BooleanValue{v}
	case string:
		if len(v) > 1500 && !noIndex {
			return nil, errors.New("string property too long to index")
		}
		val.ValueType = &pb.Value_StringValue{v}
	case float32:
		val.ValueType = &pb.Value_DoubleValue{float64(v)}
	case float64:
		val.ValueType = &pb.Value_DoubleValue{v}
	case *Key:
		if v != nil {
			val.ValueType = &pb.Value_KeyValue{keyToProto(v)}
		}
	case GeoPoint:
		if !v.Valid() {
			return nil, errors.New("invalid GeoPoint value")
		}
		val.ValueType = &pb.Value_GeoPointValue{&llpb.LatLng{
			Latitude:  v.Lat,
			Longitude: v.Lng,
		}}
	case time.Time:
		if v.Before(minTime) || v.After(maxTime) {
			return nil, errors.New("time value out of range")
		}
		val.ValueType = &pb.Value_TimestampValue{&timepb.Timestamp{
			Seconds: v.Unix(),
			Nanos:   int32(v.Nanosecond()),
		}}
	case []byte:
		if len(v) > 1500 && !noIndex {
			return nil, errors.New("[]byte property too long to index")
		}
		val.ValueType = &pb.Value_BlobValue{v}
	case []interface{}:
		arr := make([]*pb.Value, 0, len(v))
		for i, v := range v {
			elem, err := interfaceToProto(v, noIndex)
			if err != nil {
				return nil, fmt.Errorf("%v at index %d", err, i)
			}
			arr = append(arr, elem)
		}
		val.ValueType = &pb.Value_ArrayValue{&pb.ArrayValue{arr}}
		// ArrayValues have ExcludeFromIndexes set on the individual items, rather
		// than the top-level value.
		val.ExcludeFromIndexes = false
	default:
		if iv != nil {
			return nil, fmt.Errorf("invalid Value type %t", iv)
		}
		val.ValueType = &pb.Value_NullValue{}
	}
	// TODO(jbd): Support EntityValue.
	return val, nil
}
