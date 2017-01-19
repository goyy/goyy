// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package entity

import (
	"strconv"

	"gopkg.in/goyy/goyy.v0/util/errors"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

// Uint8 uint8 type.
type Uint8 struct {
	base
	value uint8
}

// Value gets the value.
func (me *Uint8) Value() uint8 {
	return me.value
}

// ValuePtr gets the value of the pointer type.
func (me *Uint8) ValuePtr() *uint8 {
	return &me.value
}

// SetValue sets the value.
func (me *Uint8) SetValue(v uint8) {
	me.value = v
	me.field.SetModified(true)
}

// SetDefault sets the default value.
func (me *Uint8) SetDefault(v string) error {
	if strings.IsBlank(v) {
		return errors.NewNotBlank("v")
	}
	val, err := strconv.ParseInt(v, 10, 8)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	me.value = uint8(val)
	return nil
}

// SetString sets the value of the string type.
func (me *Uint8) SetString(v string) error {
	if err := me.SetDefault(v); err != nil {
		return err
	}
	me.field.SetModified(true)
	return nil
}

// String gets the value of the string type.
func (me *Uint8) String() string {
	return strconv.Itoa(int(me.value))
}

// Name gets the name of the type.
func (me *Uint8) Name() string {
	return "uint8"
}
