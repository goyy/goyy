// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package entity

import (
	"strconv"

	"gopkg.in/goyy/goyy.v0/util/errors"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

// Float32 float32 type.
type Float32 struct {
	base
	value float32
}

// Value gets the value.
func (me *Float32) Value() float32 {
	return me.value
}

// ValuePtr gets the value of the pointer type.
func (me *Float32) ValuePtr() *float32 {
	return &me.value
}

// SetValue sets the value.
func (me *Float32) SetValue(v float32) {
	me.value = v
	me.field.SetModified(true)
}

// SetDefault sets the default value.
func (me *Float32) SetDefault(v string) error {
	if strings.IsBlank(v) {
		return errors.NewNotBlank("v")
	}
	val, err := strconv.ParseFloat(v, 32)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	me.value = float32(val)
	return nil
}

// SetString sets the value of the string type.
func (me *Float32) SetString(v string) error {
	if err := me.SetDefault(v); err != nil {
		return err
	}
	me.field.SetModified(true)
	return nil
}

// String gets the value of the string type.
func (me *Float32) String() string {
	return strconv.FormatFloat(float64(me.value), 'f', -1, 32)
}

// Name gets the name of the type.
func (me *Float32) Name() string {
	return "float32"
}
