// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package db

import (
	"reflect"
)

type Dialect interface {
	// If database do not support boolean type this can be used to parse int
	// value to boolean value.
	ParseBool(value reflect.Value) bool

	// SetModelValue sets a model field from a db value.
	SetModelValue(value reflect.Value, field reflect.Value) error
}
