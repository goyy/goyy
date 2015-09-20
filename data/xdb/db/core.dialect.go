// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package db

import (
	"fmt"
	"reflect"
	"time"
)

type dialect struct{}

func (me *dialect) ParseBool(value reflect.Value) bool {
	return value.Bool()
}

func (me *dialect) SetModelValue(driverValue, fieldValue reflect.Value) error {
	// ignore zero types
	if !driverValue.Elem().IsValid() {
		return nil
	}
	fieldType := fieldValue.Type()
	switch fieldValue.Type().Kind() {
	case reflect.Bool:
		fieldValue.SetBool(me.ParseBool(driverValue.Elem()))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fieldValue.SetInt(driverValue.Elem().Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		// reading uint from int value causes panic
		switch driverValue.Elem().Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			fieldValue.SetUint(uint64(driverValue.Elem().Int()))
		default:
			fieldValue.SetUint(driverValue.Elem().Uint())
		}
	case reflect.Float32, reflect.Float64:
		fieldValue.SetFloat(driverValue.Elem().Float())
	case reflect.String:
		fieldValue.SetString(string(driverValue.Elem().Bytes()))
	case reflect.Slice:
		if reflect.TypeOf(driverValue.Interface()).Elem().Kind() == reflect.Uint8 {
			fieldValue.SetBytes(driverValue.Elem().Bytes())
		}
	case reflect.Struct:
		if fieldType == reflect.TypeOf(time.Time{}) {
			fieldValue.Set(driverValue.Elem())
		} else if fieldType == reflect.TypeOf(Modified{}) {
			if time, ok := driverValue.Elem().Interface().(time.Time); ok {
				fieldValue.Set(reflect.ValueOf(Modified{Value: time}))
			} else {
				panic(fmt.Sprintf("cannot set Modified value %T", driverValue.Elem().Interface()))
			}
		} else if fieldType == reflect.TypeOf(Created{}) {
			if time, ok := driverValue.Elem().Interface().(time.Time); ok {
				fieldValue.Set(reflect.ValueOf(Created{Value: time}))
			} else {
				panic(fmt.Sprintf("cannot set created value %T", driverValue.Elem().Interface()))
			}
		}
	}
	return nil
}
