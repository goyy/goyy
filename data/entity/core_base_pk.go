// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package entity

import (
	"gopkg.in/goyy/goyy.v0/data/schema"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

type Pk struct {
	id String `db:"column=id&primary=true"`
}

func (me *Pk) Id() string {
	return me.id.Value()
}

func (me *Pk) SetId(v string) {
	me.id.SetValue(v)
}

func (me *Pk) Get(column string) interface{} {
	switch column {
	case "id":
		return me.id.Value()
	default:
		return nil
	}
}

func (me *Pk) GetPtr(column string) interface{} {
	switch column {
	case "id":
		return me.id.ValuePtr()
	}
	return nil
}

func (me *Pk) GetString(field string) string {
	switch strings.ToLowerFirst(field) {
	case "id":
		return me.id.String()
	}
	return ""
}

func (me *Pk) SetString(field, value string) error {
	switch strings.ToLowerFirst(field) {
	case "id":
		return me.id.SetString(value)
	}
	return nil
}

func (me *Pk) Type(column string) (Type, bool) {
	switch column {
	case "id":
		return &me.id, true
	}
	return nil, false
}

func (me *Pk) Column(field string) (schema.Column, bool) {
	switch strings.ToLowerFirst(field) {
	case "id":
		return me.id.Column(), true
	}
	return nil, false
}
