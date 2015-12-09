// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package db

import (
	"reflect"
)

type mysql struct {
	dialect
}

func (me *mysql) ParseBool(value reflect.Value) bool {
	return value.Int() != 0
}

func NewMysql() Dialect {
	return &mysql{}
}
