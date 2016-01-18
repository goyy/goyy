// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package repository_test

import (
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/schema"
)

//go:generate xgen -entity=$GOFILE

// @entity
type User struct {
	table    schema.Table  `db:"table=users"`
	id       entity.String `db:"column=id&primary=true"`
	code     entity.String `db:"column=code"`
	name     entity.String `db:"column=name"`
	password entity.String `db:"column=password"`
	memo     entity.String `db:"column=memo"`
	genre    entity.String `db:"column=genre"`
	status   entity.String `db:"column=status"`
	roles    entity.String `db:"column=roles"`
	posts    entity.String `db:"column=posts"`
	org      entity.String `db:"column=org"`
	area     entity.String `db:"column=area"`
	creater  entity.String `db:"column=creater&creater=true"`
	created  entity.Int64  `db:"column=created&created=true"`
	modifier entity.String `db:"column=modifier&modifier=true"`
	modified entity.Int64  `db:"column=modified&modified=true"`
	version  entity.Int    `db:"column=version&version=true"`
	deletion entity.Int    `db:"column=deletion&deletion=true"`
}
