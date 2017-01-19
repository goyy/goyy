// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package entity

import (
	"gopkg.in/goyy/goyy.v0/util/bytes"
)

// String string type.
type String struct {
	base
	value []byte
}

// Value gets the value.
func (me *String) Value() string {
	if me.value == nil || len(me.value) == 0 {
		return ""
	}
	return string(me.value)
}

// ValuePtr gets the value of the pointer type.
func (me *String) ValuePtr() *[]byte {
	return &me.value
}

// SetValue sets the value.
func (me *String) SetValue(v string) {
	me.value = []byte(v)
	me.field.SetModified(true)
}

// SetDefault sets the default value.
func (me *String) SetDefault(v string) error {
	me.value = []byte(v)
	return nil
}

// SetString sets the value of the string type.
func (me *String) SetString(v string) error {
	if err := me.SetDefault(v); err != nil {
		return err
	}
	me.field.SetModified(true)
	return nil
}

// String gets the value of the string type.
func (me *String) String() string {
	out := bytes.TrimRightNul(me.value)
	return string(out)
}

// Name gets the name of the type.
func (me *String) Name() string {
	return "string"
}
