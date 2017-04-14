// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package entity

import (
	"strconv"

	"gopkg.in/goyy/goyy.v0/util/errors"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

// Float64 float64 type.
type Float64 struct {
	base
	value float64
}

// Value gets the value.
func (me *Float64) Value() float64 {
	return me.value
}

// ValuePtr gets the value of the pointer type.
func (me *Float64) ValuePtr() *float64 {
	return &me.value
}

// SetValue sets the value.
func (me *Float64) SetValue(v float64) {
	me.value = v
	me.field.SetModified(true)
}

// SetDefault sets the default value.
func (me *Float64) SetDefault(v string) error {
	if strings.IsBlank(v) {
		return errors.NewNotBlank("v")
	}
	val, err := strconv.ParseFloat(v, 64)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	me.value = val
	return nil
}

// SetString sets the value of the string type.
func (me *Float64) SetString(v string) error {
	if err := me.SetDefault(v); err != nil {
		return err
	}
	me.field.SetModified(true)
	return nil
}

// String gets the value of the string type.
func (me *Float64) String() string {
	return strconv.FormatFloat(me.value, 'f', -1, 64)
}

// Name gets the name of the type.
func (me *Float64) Name() string {
	return "float64"
}
