// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xsql

import (
	"gopkg.in/goyy/goyy.v0/data/domain"
	"gopkg.in/goyy/goyy.v0/data/entity"
	"time"
)

type Query interface {
	Rows(out entity.Interfaces) error
	Row(out entity.Interface) error
	Int() (int, error)
	Float() (float64, error)
	Str() (string, error)
	Time() (time.Time, error)
	Page(content entity.Interfaces, pageable domain.Pageable) (domain.Page, error)

	String() string
}
