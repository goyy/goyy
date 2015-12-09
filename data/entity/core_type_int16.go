// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package entity

import (
	"gopkg.in/goyy/goyy.v0/util/errors"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"strconv"
)

type Int16 struct {
	base
	value int16
}

func (me *Int16) Value() int16 {
	return me.value
}

func (me *Int16) ValuePtr() *int16 {
	return &me.value
}

func (me *Int16) SetValue(v int16) {
	me.value = v
	me.field.SetModified(true)
}

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

func (me *Int16) SetString(v string) error {
	if err := me.SetDefault(v); err != nil {
		return err
	} else {
		me.field.SetModified(true)
		return nil
	}
}

func (me *Int16) String() string {
	return strconv.Itoa(int(me.value))
}

func (me *Int16) Name() string {
	return "int16"
}
