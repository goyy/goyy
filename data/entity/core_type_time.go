// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package entity

import (
	"time"

	"gopkg.in/goyy/goyy.v0/util/errors"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"gopkg.in/goyy/goyy.v0/util/times"
)

// Time time type.
type Time struct {
	base
	value time.Time
}

// Value gets the value.
func (me *Time) Value() time.Time {
	return me.value
}

// ValuePtr gets the value of the pointer type.
func (me *Time) ValuePtr() *time.Time {
	return &me.value
}

// SetValue sets the value.
func (me *Time) SetValue(v time.Time) {
	me.value = v
	me.field.SetModified(true)
}

// SetDefault sets the default value.
func (me *Time) SetDefault(layout, value string) error {
	if strings.IsBlank(layout) {
		return errors.NewNotBlank("layout")
	}
	if strings.IsBlank(value) {
		return errors.NewNotBlank("value")
	}
	v, err := time.Parse(layout, value)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	me.value = v
	return nil
}

// SetString sets the value of the string type.
func (me *Time) SetString(layout, value string) error {
	if err := me.SetDefault(layout, value); err != nil {
		return err
	}
	me.field.SetModified(true)
	return nil
}

// String gets the value of the string type.
func (me *Time) String() string {
	return times.FormatYYMDHMS(me.value)
}

// Name gets the name of the type.
func (me *Time) Name() string {
	return "time"
}
