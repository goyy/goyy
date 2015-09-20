// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package reflects

import (
	"errors"
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"reflect"
)

const (
	fNotUse   = "Cannot use %s on a non-struct interface"
	fNoSuch   = "No such %s name: %s in i"
	fEmpty    = "%s can not be empty"
	flagType  = "type"
	flagValue = "value"
)

// Call calls the function i with the input arguments.
func Call(i interface{}, methodName string, methodArgs ...interface{}) (result []interface{}, err error) {
	m, err := Method(i, methodName)
	if err != nil {
		return
	}
	if len(methodArgs) != m.Type().NumIn() {
		err = errors.New("The number of methodArgs is not adapted.")
		return
	}
	in := make([]reflect.Value, len(methodArgs))
	for k, arg := range methodArgs {
		in[k] = reflect.ValueOf(arg)
	}
	vs := m.Call(in)
	result = make([]interface{}, len(vs))
	for k, v := range vs {
		result[k] = v.Interface()
	}
	return
}

// Set assigns value to the interface{} i. It error if CanSet returns false.
// As in Go, value's interface{} must be assignable to i's type.
func Set(i interface{}, fieldName string, value interface{}) error {
	if field, err := Field(i, fieldName); err != nil {
		return err
	} else {
		if !field.CanSet() {
			return fmt.Errorf("Cannot set %s field value", fieldName)
		}
		t := field.Type()
		v := reflect.ValueOf(value)
		if t != v.Type() {
			return errors.New("Assigns value type didn't match i." + fieldName + " field type")
		}
		field.Set(v)
		return nil
	}
}

// Interface returns i's current value as an interface{}.
// It error if the Value was obtained by accessing unexported struct fields.
func Interface(i interface{}, fieldName string) (interface{}, error) {
	if field, err := Field(i, fieldName); err != nil {
		return nil, err
	} else {
		return field.Interface(), nil
	}
}

// Tag returns the value associated with tagKey in the tag string.
// If there is no such tagKey in the tag, Get returns the empty string.
// If the tag does not have the conventional format,
// the value returned by Get is unspecified.
func Tag(i interface{}, fieldName, tagKey string) (string, error) {
	if strings.IsBlank(tagKey) {
		return "", fmt.Errorf(fEmpty, "tagKey")
	}
	if field, err := StructField(i, fieldName); err != nil {
		return "", err
	} else {
		return field.Tag.Get(tagKey), nil
	}
}

// Indirect returns the value that i points to.
// If i is a nil pointer, Indirect returns a zero Value.
// If i is not a pointer, Indirect returns i.
func Indirect(i interface{}) (v reflect.Value) {
	return reflect.Indirect(reflect.ValueOf(i))
}

// Kind returns i's Kind.
// If i is the zero Value (IsValid returns false), Kind returns Invalid.
func Kind(i interface{}, fieldName string) (reflect.Kind, error) {
	if field, err := Field(i, fieldName); err != nil {
		return reflect.Invalid, err
	} else {
		return field.Type().Kind(), nil
	}
}

// StructField returns i's reflect.StructField.
// It error if i's Kind is not Struct or Ptr.
func StructField(i interface{}, fieldName string) (field reflect.StructField, err error) {
	if strings.IsBlank(fieldName) {
		err = fmt.Errorf(fEmpty, "fieldName")
		return
	}
	if !IsStructOrStructPtr(i) {
		err = fmt.Errorf(fNotUse, "StructField")
		return
	}
	v := Indirect(i)
	t := v.Type()
	f, ok := t.FieldByName(fieldName)
	if !ok {
		err = fmt.Errorf(fNoSuch, "field", fieldName)
		return
	}
	field = f
	return
}

// Field returns i's reflect.Value.
// It error if i's Kind is not Struct or Ptr.
func Field(i interface{}, fieldName string) (field reflect.Value, err error) {
	if strings.IsBlank(fieldName) {
		err = fmt.Errorf(fEmpty, "fieldName")
		return
	}
	if !IsStructOrStructPtr(i) {
		err = fmt.Errorf(fNotUse, "Field")
		return
	}
	v := Indirect(i)
	field = v.FieldByName(fieldName)
	if !field.IsValid() {
		err = fmt.Errorf(fNoSuch, "field", fieldName)
		return
	}
	return
}

// Method returns i's reflect.Value.
// It error if i's Kind is not Struct or Ptr.
func Method(i interface{}, methodName string) (method reflect.Value, err error) {
	if strings.IsBlank(methodName) {
		err = fmt.Errorf(fEmpty, "methodName")
		return
	}
	if !IsStructOrStructPtr(i) {
		err = fmt.Errorf(fNotUse, "Method")
		return
	}
	v := reflect.ValueOf(i)
	method = v.MethodByName(methodName)
	if !method.IsValid() {
		err = fmt.Errorf(fNoSuch, "method", methodName)
		return
	}
	return
}

// Names returns the struct fields names list.
// It error if i's Kind is not Struct or Ptr.
// If anonymous is true that contains the field names anonymous.
func Names(i interface{}, anonymous bool) ([]string, error) {
	if !IsStructOrStructPtr(i) {
		return nil, fmt.Errorf(fNotUse, "Names")
	}
	v := Indirect(i)
	t := v.Type()
	return FieldNames(t, anonymous)
}

// FieldNames returns the struct fields names list.
// If anonymous is true that contains the field names anonymous.
func FieldNames(t reflect.Type, anonymous bool) (names []string, err error) {
	numField := t.NumField()
	for i := 0; i < numField; i++ {
		f := t.Field(i)
		if f.Anonymous && f.Type.Kind() == reflect.Struct {
			subNames, e := FieldNames(f.Type, anonymous)
			if e != nil {
				err = e
				return
			}
			for _, subName := range subNames {
				shouldAppend := true
				for _, name := range names {
					if subName == name {
						shouldAppend = false
						break
					}
				}
				if shouldAppend {
					names = append(names, subName)
				}
			}
		}
		if f.Anonymous && !anonymous {
			continue
		}
		names = append(names, f.Name)
	}
	return
}

// Interfaces returns the field - value struct pairs as a map.
// It error if i's Kind is not Struct or Ptr.
func Interfaces(i interface{}) (map[string]interface{}, error) {
	if !IsStructOrStructPtr(i) {
		return nil, fmt.Errorf(fNotUse, "Values")
	}
	v := Indirect(i)
	return FieldInterfaces(v)
}

// FieldInterfaces returns the field - value struct pairs as a map.
func FieldInterfaces(v reflect.Value) (map[string]interface{}, error) {
	values := make(map[string]interface{})
	t := v.Type()
	names, err := FieldNames(t, false)
	if err != nil {
		return nil, err
	}
	for _, name := range names {
		if f, ok := t.FieldByName(name); ok {
			if !f.Anonymous {
				vf := v.FieldByName(name)
				values[name] = vf.Interface()
			}
		}
	}
	return values, nil
}

// Tags lists the struct tag fields.
// It error if i's Kind is not Struct or Ptr.
func Tags(i interface{}, tagKey string) (map[string]string, error) {
	if !IsStructOrStructPtr(i) {
		return nil, fmt.Errorf(fNotUse, "Tags")
	}
	v := Indirect(i)
	t := v.Type()
	return FieldTags(t, tagKey)
}

// FieldTags lists the struct tag fields.
func FieldTags(t reflect.Type, tagKey string) (map[string]string, error) {
	tags := make(map[string]string)
	if strings.IsBlank(tagKey) {
		return nil, fmt.Errorf(fEmpty, "tagKey")
	}
	names, err := FieldNames(t, true)
	if err != nil {
		return nil, err
	}
	for _, name := range names {
		if f, ok := t.FieldByName(name); ok {
			tags[name] = f.Tag.Get(tagKey)
		}
	}
	return tags, nil
}

// MakeSlice creates a new zero-initialized slice value for the specified slice type,
// length, and capacity.
func MakeSlice(i interface{}, len, cap int) interface{} {
	st := reflect.SliceOf(Indirect(i).Type())
	return reflect.MakeSlice(st, len, cap).Interface()
}

// IsField checks if i field name is part of a struct.
// It error if i's Kind is not Struct or Ptr.
func IsField(i interface{}, fieldName string) (bool, error) {
	if strings.IsBlank(fieldName) {
		return false, fmt.Errorf(fEmpty, "fieldName")
	}
	if !IsStructOrStructPtr(i) {
		return false, fmt.Errorf(fNotUse, "IsField")
	}
	v := Indirect(i)
	t := v.Type()
	if _, ok := t.FieldByName(fieldName); ok {
		return true, nil
	} else {
		return false, fmt.Errorf(fNoSuch, "field", fieldName)
	}
}

// IsTag checks if i field tag is part of a struct.
// It error if i's Kind is not Struct or Ptr.
func IsTag(i interface{}, fieldName, tagKey string) (bool, error) {
	if strings.IsBlank(tagKey) {
		return false, fmt.Errorf(fEmpty, "tagKey")
	}
	if field, err := StructField(i, fieldName); err != nil {
		return false, err
	} else {
		tag := field.Tag.Get(tagKey)
		if strings.IsBlank(tag) {
			return false, nil
		}
		return true, nil
	}
}

// IsAnyType returns true if any type in types are within i.
func IsAnyType(i interface{}, types []reflect.Kind) bool {
	return HasAnyType(reflect.TypeOf(i), types)
}

// IsStruct returns true if i is of type struct.
func IsStruct(i interface{}) bool {
	return reflect.TypeOf(i).Kind() == reflect.Struct
}

// IsStructPtr returns true if i is of type struct pointer.
func IsStructPtr(i interface{}) bool {
	if !IsPtr(i) {
		return false
	}
	return HasStruct(Indirect(i))
}

// IsStructOrStructPtr returns true if i is of type struct or struct pointer.
func IsStructOrStructPtr(i interface{}) bool {
	if IsStruct(i) {
		return true
	}
	if IsStructPtr(i) {
		return true
	}
	return false
}

// IsSlice returns true if i is of type slice.
func IsSlice(i interface{}) bool {
	return reflect.TypeOf(i).Kind() == reflect.Slice
}

// IsPtr returns true if i is of type pointer.
func IsPtr(i interface{}) bool {
	return reflect.TypeOf(i).Kind() == reflect.Ptr
}

// IsString returns true if i is of type string.
func IsString(i interface{}) bool {
	return reflect.TypeOf(i).Kind() == reflect.String
}

// HasAnyType returns true if any type in types are within v.
func HasAnyType(v reflect.Type, types []reflect.Kind) bool {
	for _, t := range types {
		if v.Kind() == t {
			return true
		}
	}
	return false
}

// HasStruct returns true if i is of type struct.
func HasStruct(v reflect.Value) bool {
	return v.Kind() == reflect.Struct
}

// HasSlice returns true if i is of type slice.
func HasSlice(v reflect.Value) bool {
	return v.Kind() == reflect.Slice
}

// HasPtr returns true if i is of type pointer.
func HasPtr(v reflect.Value) bool {
	return v.Kind() == reflect.Ptr
}

// HasString returns true if i is of type string.
func HasString(v reflect.Value) bool {
	return v.Kind() == reflect.String
}
