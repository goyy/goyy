// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dql

import (
	"gopkg.in/goyy/goyy.v0/data/dialect"
	"gopkg.in/goyy/goyy.v0/data/domain"
	"gopkg.in/goyy/goyy.v0/data/entity"
)

type Interface interface {
	SelectOne(e entity.Interface) (dql string, arg interface{})
	SelectPage(dql string, pageable domain.Pageable) string
	SelectListBySift(e entity.Interface, sifts ...domain.Sift) (dql string, args []interface{}, err error)
	SelectCountBySift(e entity.Interface, sifts ...domain.Sift) (dql string, args []interface{}, err error)
}

func New(i dialect.Interface) Interface {
	switch i.Type() {
	case dialect.ORACLE:
		return &oracle{}
	default:
		return &mysql{}
	}
}
