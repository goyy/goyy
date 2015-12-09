// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package entity

import (
	"gopkg.in/goyy/goyy.v0/data/schema"
)

type base struct {
	column schema.Column
	field  Field
}

func (me *base) Column() schema.Column {
	return me.column
}

func (me *base) SetColumn(v schema.Column) {
	me.column = v
}

func (me *base) Field() Field {
	return me.field
}

func (me *base) SetField(v Field) {
	me.field = v
}

func (me *base) HasUpdate() bool {
	if me.field.Updateable() && me.field.Modified() && !me.column.IsTransient() {
		return true
	}
	return false
}

func (me *base) HasInsert() bool {
	if me.field.Insertable() && !me.column.IsTransient() {
		return true
	}
	return false
}
