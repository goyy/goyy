// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package entity

import (
	"gopkg.in/goyy/goyy.v0/data/schema"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

var TreeColumns = [...]string{
	"id",
	"code",
	"name",
	"fullname",
	"genre",
	"leaf",
	"grade",
	"ordinal",
	"parent_id",
	"parent_ids",
	"parent_codes",
	"parent_names",
	"memo",
	"creates",
	"creater",
	"created",
	"modifier",
	"modified",
	"version",
	"deletion",
	"artifical",
	"history",
}

var TreeFields = [...]string{
	"id",
	"code",
	"name",
	"fullname",
	"genre",
	"leaf",
	"grade",
	"ordinal",
	"parentId",
	"parentIds",
	"parentCodes",
	"parentNames",
	"memo",
	"creates",
	"creater",
	"created",
	"modifier",
	"modified",
	"version",
	"deletion",
	"artifical",
	"history",
}

type Tree struct {
	Sys
	code        String `db:"column=code"`
	name        String `db:"column=name"`
	fullname    String `db:"column=fullname"`
	genre       String `db:"column=genre"`
	leaf        Int    `db:"column=leaf"`
	grade       Int    `db:"column=grade"`
	ordinal     String `db:"column=ordinal"`
	parentId    String `db:"column=parent_id"`
	parentIds   String `db:"column=parent_ids"`
	parentCodes String `db:"column=parent_codes"`
	parentNames String `db:"column=parent_mames"`
}

func (me *Tree) Code() string {
	return me.code.Value()
}

func (me *Tree) SetCode(v string) {
	me.code.SetValue(v)
}

func (me *Tree) Name() string {
	return me.name.Value()
}

func (me *Tree) SetName(v string) {
	me.name.SetValue(v)
}

func (me *Tree) Fullname() string {
	return me.fullname.Value()
}

func (me *Tree) SetFullname(v string) {
	me.fullname.SetValue(v)
}

func (me *Tree) Genre() string {
	return me.genre.Value()
}

func (me *Tree) SetGenre(v string) {
	me.genre.SetValue(v)
}

func (me *Tree) Leaf() int {
	return me.leaf.Value()
}

func (me *Tree) SetLeaf(v int) {
	me.leaf.SetValue(v)
}

func (me *Tree) Grade() int {
	return me.grade.Value()
}

func (me *Tree) SetGrade(v int) {
	me.grade.SetValue(v)
}

func (me *Tree) Ordinal() string {
	return me.ordinal.Value()
}

func (me *Tree) SetOrdinal(v string) {
	me.ordinal.SetValue(v)
}

func (me *Tree) ParentId() string {
	return me.parentId.Value()
}

func (me *Tree) SetParentId(v string) {
	me.parentId.SetValue(v)
}

func (me *Tree) ParentIds() string {
	return me.parentIds.Value()
}

func (me *Tree) SetParentIds(v string) {
	me.parentIds.SetValue(v)
}

func (me *Tree) ParentCodes() string {
	return me.parentCodes.Value()
}

func (me *Tree) SetParentCodes(v string) {
	me.parentCodes.SetValue(v)
}

func (me *Tree) ParentNames() string {
	return me.parentNames.Value()
}

func (me *Tree) SetParentNames(v string) {
	me.parentNames.SetValue(v)
}

func (me *Tree) Get(column string) interface{} {
	switch column {
	case "code":
		return me.code.Value()
	case "name":
		return me.name.Value()
	case "fullname":
		return me.fullname.Value()
	case "genre":
		return me.genre.Value()
	case "leaf":
		return me.leaf.Value()
	case "grade":
		return me.grade.Value()
	case "ordinal":
		return me.ordinal.Value()
	case "parent_id":
		return me.parentId.Value()
	case "parent_ids":
		return me.parentIds.Value()
	case "parent_codes":
		return me.parentCodes.Value()
	case "parent_names":
		return me.parentNames.Value()
	}
	return me.Sys.Get(column)
}

func (me *Tree) GetPtr(column string) interface{} {
	switch column {
	case "code":
		return me.code.ValuePtr()
	case "name":
		return me.name.ValuePtr()
	case "fullname":
		return me.fullname.ValuePtr()
	case "genre":
		return me.genre.ValuePtr()
	case "leaf":
		return me.leaf.ValuePtr()
	case "grade":
		return me.grade.ValuePtr()
	case "ordinal":
		return me.ordinal.ValuePtr()
	case "parent_id":
		return me.parentId.ValuePtr()
	case "parent_ids":
		return me.parentIds.ValuePtr()
	case "parent_codes":
		return me.parentCodes.ValuePtr()
	case "parent_names":
		return me.parentNames.ValuePtr()
	}
	return me.Sys.GetPtr(column)
}

func (me *Tree) GetString(field string) string {
	switch strings.ToLowerFirst(field) {
	case "code":
		return me.code.String()
	case "name":
		return me.name.String()
	case "fullname":
		return me.fullname.String()
	case "genre":
		return me.genre.String()
	case "leaf":
		return me.leaf.String()
	case "grade":
		return me.grade.String()
	case "ordinal":
		return me.ordinal.String()
	case "parentId":
		return me.parentId.String()
	case "parentIds":
		return me.parentIds.String()
	case "parentCodes":
		return me.parentCodes.String()
	case "parentNames":
		return me.parentNames.String()
	}
	return me.Sys.GetString(field)
}

func (me *Tree) SetString(field, value string) error {
	switch strings.ToLowerFirst(field) {
	case "code":
		return me.code.SetString(value)
	case "name":
		return me.name.SetString(value)
	case "fullname":
		return me.fullname.SetString(value)
	case "genre":
		return me.genre.SetString(value)
	case "leaf":
		return me.leaf.SetString(value)
	case "grade":
		return me.grade.SetString(value)
	case "ordinal":
		return me.ordinal.SetString(value)
	case "parentId":
		return me.parentId.SetString(value)
	case "parentIds":
		return me.parentIds.SetString(value)
	case "parentCodes":
		return me.parentCodes.SetString(value)
	case "parentNames":
		return me.parentNames.SetString(value)
	}
	return me.Sys.SetString(field, value)
}

func (me *Tree) Type(column string) (Type, bool) {
	switch column {
	case "code":
		return &me.code, true
	case "name":
		return &me.name, true
	case "fullname":
		return &me.fullname, true
	case "genre":
		return &me.genre, true
	case "leaf":
		return &me.leaf, true
	case "grade":
		return &me.grade, true
	case "ordinal":
		return &me.ordinal, true
	case "parent_id":
		return &me.parentId, true
	case "parent_ids":
		return &me.parentIds, true
	case "parent_codes":
		return &me.parentCodes, true
	case "parent_names":
		return &me.parentNames, true
	}
	return me.Sys.Type(column)
}

func (me *Tree) Column(field string) (schema.Column, bool) {
	switch strings.ToLowerFirst(field) {
	case "code":
		return me.code.Column(), true
	case "name":
		return me.name.Column(), true
	case "fullname":
		return me.fullname.Column(), true
	case "genre":
		return me.genre.Column(), true
	case "leaf":
		return me.leaf.Column(), true
	case "grade":
		return me.grade.Column(), true
	case "ordinal":
		return me.ordinal.Column(), true
	case "parentId":
		return me.parentId.Column(), true
	case "parentIds":
		return me.parentIds.Column(), true
	case "parentCodes":
		return me.parentCodes.Column(), true
	case "parentNames":
		return me.parentNames.Column(), true
	}
	return me.Sys.Column(field)
}
