// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package internal

import (
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/schema"
)

//go:generate xgen -entity=$GOFILE -service

// @entity
type Permission struct {
	table      schema.Table  `db:"table=sys_menu"`
	id         entity.String `db:"column=id&primary=true"`
	permission entity.String `db:"column=permission"`
}
