// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package entity

import (
	"gopkg.in/goyy/goyy.v0/util/errors"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"strconv"
)

type Int8 struct {
	base
	value int8
}

func (me *Int8) Value() int8 {
	return me.value
}

func (me *Int8) ValuePtr() *int8 {
	return &me.value
}

func (me *Int8) SetValue(v int8) {
	me.value = v
	me.field.SetModified(true)
}

func (me *Int8) SetDefault(v string) error {
	if strings.IsBlank(v) {
		return errors.NewNotBlank("v")
	}
	val, err := strconv.ParseInt(v, 10, 8)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	me.value = int8(val)
	return nil
}

func (me *Int8) SetString(v string) error {
	if err := me.SetDefault(v); err != nil {
		return err
	} else {
		me.field.SetModified(true)
		return nil
	}
}

func (me *Int8) String() string {
	return strconv.Itoa(int(me.value))
}

func (me *Int8) Name() string {
	return "int8"
}
