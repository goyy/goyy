// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dql_test

import (
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/schema"
)

//go:generate xgen -entity=$GOFILE

// @entity
type User struct {
	table  schema.Table  `db:"table=sys_user"`
	id     entity.String `db:"column=id&primary=true"`
	name   entity.String `db:"column=name"`
	passwd entity.String `db:"column=passwd"`
	email  entity.String `db:"column=email"`
}
