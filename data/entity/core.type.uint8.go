// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package entity

import (
	"gopkg.in/goyy/goyy.v0/util/errors"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"strconv"
)

type Uint8 struct {
	base
	value uint8
}

func (me *Uint8) Value() uint8 {
	return me.value
}

func (me *Uint8) ValuePtr() *uint8 {
	return &me.value
}

func (me *Uint8) SetValue(v uint8) {
	me.value = v
	me.field.SetModified(true)
}

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

func (me *Uint8) SetString(v string) error {
	if err := me.SetDefault(v); err != nil {
		return err
	} else {
		me.field.SetModified(true)
		return nil
	}
}

func (me *Uint8) String() string {
	return strconv.Itoa(int(me.value))
}

func (me *Uint8) Name() string {
	return "uint8"
}
