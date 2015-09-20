// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package result_test

import (
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/schema"
)

//go:generate xgen -entity=$GOFILE

// @entity
type User struct {
	table   schema.Table  `orm:"table=sys_user"`
	id      entity.String `orm:"column=id&primary=true"`
	name    entity.String `orm:"column=name"`
	passwd  entity.String `orm:"column=passwd"`
	age     entity.Int    `orm:"column=age"`
	email   entity.String `orm:"column=email"`
	version entity.Int    `orm:"column=version"`
}
