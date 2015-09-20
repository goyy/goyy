// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dml

import (
	"gopkg.in/goyy/goyy.v0/data/dialect"
	"gopkg.in/goyy/goyy.v0/data/entity"
)

type Interface interface {
	Update(e entity.Interface) (dml string, args []interface{})
	Insert(e entity.Interface) (dml string, args []interface{})
	Delete(e entity.Interface) (dml string, arg interface{})
	Disable(e entity.Interface) (dml string, arg interface{})
}

func New(i dialect.Interface) Interface {
	switch i.Type() {
	case dialect.ORACLE:
		return &oracle{}
	default:
		return &mysql{}
	}
}
