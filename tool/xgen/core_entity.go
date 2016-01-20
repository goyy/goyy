// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"gopkg.in/goyy/goyy.v0/util/strings"
)

// entity is an SQL table.
type entity struct {
	Project         string
	Extend          string // value:[tree,sys,pk]
	Relationship    string // value:[master,slave]
	Name            string
	Comment         string
	Table           string
	Fields          []*field
	PrimaryKeys     []*field
	FieldMaxLen     int
	ColumnMaxLen    int
	TypeMaxLen      int
	AllFieldMaxLen  int
	AllColumnMaxLen int
	AllTypeMaxLen   int
}

func (me entity) GetComment() string {
	if strings.IsBlank(me.Comment) {
		return strings.ToUpper(me.Name)
	}
	return me.Comment
}
