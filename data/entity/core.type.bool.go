// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package entity

type Bool struct {
	base
	value bool
}

func (me *Bool) Value() bool {
	return me.value
}

func (me *Bool) ValuePtr() *bool {
	return &me.value
}

func (me *Bool) SetValue(v bool) {
	me.value = v
	me.field.SetModified(true)
}

func (me *Bool) SetDefault(v string) error {
	if v == "true" {
		me.value = true
	} else {
		me.value = false
	}
	return nil
}

func (me *Bool) SetString(v string) error {
	if err := me.SetDefault(v); err != nil {
		return err
	} else {
		me.field.SetModified(true)
		return nil
	}
}

func (me *Bool) String() string {
	if me.value {
		return "true"
	} else {
		return "false"
	}
}

func (me *Bool) Name() string {
	return "bool"
}
