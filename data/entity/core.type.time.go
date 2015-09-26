// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package entity

import (
	"gopkg.in/goyy/goyy.v0/util/errors"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"gopkg.in/goyy/goyy.v0/util/times"
	"time"
)

type Time struct {
	base
	value time.Time
}

func (me *Time) Value() time.Time {
	return me.value
}

func (me *Time) ValuePtr() *time.Time {
	return &me.value
}

func (me *Time) SetValue(v time.Time) {
	me.value = v
	me.field.SetModified(true)
}

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

func (me *Time) SetString(layout, value string) error {
	if err := me.SetDefault(layout, value); err != nil {
		return err
	} else {
		me.field.SetModified(true)
		return nil
	}
}

func (me *Time) String() string {
	return times.FormatYYMDHMS(me.value)
}

func (me *Time) Name() string {
	return "time"
}
