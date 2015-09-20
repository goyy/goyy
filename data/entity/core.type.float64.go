// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package entity

import (
	"gopkg.in/goyy/goyy.v0/util/errors"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"strconv"
)

type Float64 struct {
	base
	value float64
}

func (me *Float64) Value() float64 {
	return me.value
}

func (me *Float64) ValuePtr() *float64 {
	return &me.value
}

func (me *Float64) SetValue(v float64) {
	me.value = v
	me.field.SetModified(true)
}

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

func (me *Float64) SetString(v string) error {
	if err := me.SetDefault(v); err != nil {
		return err
	} else {
		me.field.SetModified(true)
		return nil
	}
}

func (me *Float64) String() string {
	return strconv.FormatFloat(me.value, 'f', -1, 64)
}

func (me *Float64) Name() string {
	return "float64"
}
