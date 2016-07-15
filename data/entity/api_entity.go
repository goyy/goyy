// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package entity

import (
	"gopkg.in/goyy/goyy.v0/data/schema"
)

type Interface interface {
	New() Interface
	Get(column string) interface{}
	GetPtr(column string) interface{}
	GetString(field string) string
	SetString(field, value string) error
	Type(column string) (Type, bool)
	Table() schema.Table
	Column(field string) (schema.Column, bool)
	Columns() []schema.Column
	Names() []string
	Validate() error
	JSON() string
	ExcelColumns() []string
}

type Interfaces interface {
	Make(v int)
	New() Interface
	Append(v Interface)
	Len() int
	Cap() int
	Index(i int) Interface
	Slice() interface{}
	JSON() string
}
