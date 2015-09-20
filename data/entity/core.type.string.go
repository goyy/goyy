// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package entity

type String struct {
	base
	value string
}

func (me *String) Value() string {
	return me.value
}

func (me *String) ValuePtr() *string {
	return &me.value
}

func (me *String) SetValue(v string) {
	me.value = v
	me.field.SetModified(true)
}

func (me *String) SetDefault(v string) error {
	me.value = v
	return nil
}

func (me *String) SetString(v string) error {
	if err := me.SetDefault(v); err != nil {
		return err
	} else {
		me.field.SetModified(true)
		return nil
	}
}

func (me *String) String() string {
	return me.value
}

func (me *String) Name() string {
	return "string"
}
