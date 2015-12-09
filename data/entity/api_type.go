// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package entity

import (
	"gopkg.in/goyy/goyy.v0/data/schema"
)

type Type interface {
	Column() schema.Column
	SetColumn(v schema.Column)
	Field() Field
	SetField(v Field)
	SetDefault(v string) error
	HasUpdate() bool
	HasInsert() bool
	String() string
	Name() string
}
