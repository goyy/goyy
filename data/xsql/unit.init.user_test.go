// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xsql_test

import (
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/schema"
)

//go:generate xgen -entity=$GOFILE

// @entity
type User struct {
	table    schema.Table  `orm:"table=users"`
	id       entity.String `orm:"column=id&primary=true"`
	code     entity.String `orm:"column=code"`
	name     entity.String `orm:"column=name"`
	password entity.String `orm:"column=password"`
	memo     entity.String `orm:"column=memo"`
	genre    entity.String `orm:"column=genre"`
	status   entity.String `orm:"column=status"`
	roles    entity.String `orm:"column=roles"`
	posts    entity.String `orm:"column=posts"`
	org      entity.String `orm:"column=org"`
	area     entity.String `orm:"column=area"`
	creater  entity.String `orm:"column=creater&creater=true"`
	created  entity.Int64  `orm:"column=created&created=true"`
	modifier entity.String `orm:"column=modifier&modifier=true"`
	modified entity.Int64  `orm:"column=modified&modified=true"`
	version  entity.Int    `orm:"column=version&version=true"`
	deletion entity.Int    `orm:"column=deletion&deletion=true"`
}
