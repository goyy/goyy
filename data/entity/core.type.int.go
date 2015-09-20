// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package entity

import (
	"gopkg.in/goyy/goyy.v0/util/errors"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"strconv"
)

type Int struct {
	base
	value int
}

func (me *Int) Value() int {
	return me.value
}

func (me *Int) ValuePtr() *int {
	return &me.value
}

func (me *Int) SetValue(v int) {
	me.value = v
	me.field.SetModified(true)
}

func (me *Int) SetDefault(v string) error {
	if strings.IsBlank(v) {
		return errors.NewNotBlank("v")
	}
	val, err := strconv.ParseInt(v, 10, 0)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	me.value = int(val)
	return nil
}

func (me *Int) SetString(v string) error {
	if err := me.SetDefault(v); err != nil {
		return err
	} else {
		me.field.SetModified(true)
		return nil
	}
}

func (me *Int) String() string {
	return strconv.Itoa(me.value)
}

func (me *Int) Name() string {
	return "int"
}
