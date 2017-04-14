// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package entity

import (
	"strconv"

	"gopkg.in/goyy/goyy.v0/util/errors"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

// Int16 int16 type.
type Int16 struct {
	base
	value int16
}

// Value gets the value.
func (me *Int16) Value() int16 {
	return me.value
}

// ValuePtr gets the value of the pointer type.
func (me *Int16) ValuePtr() *int16 {
	return &me.value
}

// SetValue sets the value.
func (me *Int16) SetValue(v int16) {
	me.value = v
	me.field.SetModified(true)
}

// SetDefault sets the default value.
func (me *Int16) SetDefault(v string) error {
	if strings.IsBlank(v) {
		return errors.NewNotBlank("v")
	}
	val, err := strconv.ParseInt(v, 10, 16)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	me.value = int16(val)
	return nil
}

// SetString sets the value of the string type.
func (me *Int16) SetString(v string) error {
	if err := me.SetDefault(v); err != nil {
		return err
	}
	me.field.SetModified(true)
	return nil
}

// String gets the value of the string type.
func (me *Int16) String() string {
	return strconv.Itoa(int(me.value))
}

// Name gets the name of the type.
func (me *Int16) Name() string {
	return "int16"
}
