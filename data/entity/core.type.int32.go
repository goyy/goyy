// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package entity

import (
	"gopkg.in/goyy/goyy.v0/util/errors"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"strconv"
)

type Int32 struct {
	base
	value int32
}

func (me *Int32) Value() int32 {
	return me.value
}

func (me *Int32) ValuePtr() *int32 {
	return &me.value
}

func (me *Int32) SetValue(v int32) {
	me.value = v
	me.field.SetModified(true)
}

func (me *Int32) SetDefault(v string) error {
	if strings.IsBlank(v) {
		return errors.NewNotBlank("v")
	}
	val, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	me.value = int32(val)
	return nil
}

func (me *Int32) SetString(v string) error {
	if err := me.SetDefault(v); err != nil {
		return err
	} else {
		me.field.SetModified(true)
		return nil
	}
}

func (me *Int32) String() string {
	return strconv.Itoa(int(me.value))
}

func (me *Int32) Name() string {
	return "int32"
}
