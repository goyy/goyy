// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package entity

import (
	"strconv"

	"gopkg.in/goyy/goyy.v0/util/errors"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

// Uint uint type.
type Uint struct {
	base
	value uint
}

// Value gets the value.
func (me *Uint) Value() uint {
	return me.value
}

// ValuePtr gets the value of the pointer type.
func (me *Uint) ValuePtr() *uint {
	return &me.value
}

// SetValue sets the value.
func (me *Uint) SetValue(v uint) {
	me.value = v
	me.field.SetModified(true)
}

// SetDefault sets the default value.
func (me *Uint) SetDefault(v string) error {
	if strings.IsBlank(v) {
		return errors.NewNotBlank("v")
	}
	val, err := strconv.ParseInt(v, 10, 0)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	me.value = uint(val)
	return nil
}

// SetString sets the value of the string type.
func (me *Uint) SetString(v string) error {
	if err := me.SetDefault(v); err != nil {
		return err
	}
	me.field.SetModified(true)
	return nil
}

// String gets the value of the string type.
func (me *Uint) String() string {
	return strconv.Itoa(int(me.value))
}

// Name gets the name of the type.
func (me *Uint) Name() string {
	return "uint"
}
