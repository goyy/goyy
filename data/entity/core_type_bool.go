// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package entity

// Bool bool type.
type Bool struct {
	base
	value bool
}

// Value gets the value.
func (me *Bool) Value() bool {
	return me.value
}

// ValuePtr gets the value of the pointer type.
func (me *Bool) ValuePtr() *bool {
	return &me.value
}

// SetValue sets the value.
func (me *Bool) SetValue(v bool) {
	me.value = v
	me.field.SetModified(true)
}

// SetDefault sets the default value.
func (me *Bool) SetDefault(v string) error {
	if v == "true" {
		me.value = true
	} else {
		me.value = false
	}
	return nil
}

// SetString sets the value of the string type.
func (me *Bool) SetString(v string) error {
	if err := me.SetDefault(v); err != nil {
		return err
	}
	me.field.SetModified(true)
	return nil
}

// String gets the value of the string type.
func (me *Bool) String() string {
	if me.value {
		return "true"
	}
	return "false"
}

// Name gets the name of the type.
func (me *Bool) Name() string {
	return "bool"
}
