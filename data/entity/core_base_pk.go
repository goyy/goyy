// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package entity

import (
	"gopkg.in/goyy/goyy.v0/data/schema"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

// Pk database primary key.
type Pk struct {
	id String `db:"column=id&primary=true"`
}

// Id get Pk.id.
func (me *Pk) Id() string {
	return me.id.Value()
}

// SetId set Pk.id.
func (me *Pk) SetId(v string) {
	me.id.SetValue(v)
}

// Get according to the database column to obtain values.
func (me *Pk) Get(column string) interface{} {
	switch column {
	case "id":
		return me.id.Value()
	default:
		return nil
	}
}

// GetPtr gets the value of a pointer type according to the database column name.
func (me *Pk) GetPtr(column string) interface{} {
	switch column {
	case "id":
		return me.id.ValuePtr()
	}
	return nil
}

// GetString gets the value according to the entity property name.
func (me *Pk) GetString(field string) string {
	switch strings.ToLowerFirst(field) {
	case "id":
		return me.id.String()
	}
	return ""
}

// SetString sets the value according to the entity property name.
func (me *Pk) SetString(field, value string) error {
	switch strings.ToLowerFirst(field) {
	case "id":
		return me.id.SetString(value)
	}
	return nil
}

// Type get entity.Type through the column of the database.
func (me *Pk) Type(column string) (Type, bool) {
	switch column {
	case "id":
		return &me.id, true
	}
	return nil, false
}

// Column get schema.Column by the name of the entity attribute.
func (me *Pk) Column(field string) (schema.Column, bool) {
	switch strings.ToLowerFirst(field) {
	case "id":
		return me.id.Column(), true
	}
	return nil, false
}
