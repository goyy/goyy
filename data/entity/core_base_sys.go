// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package entity

import (
	"gopkg.in/goyy/goyy.v0/data/schema"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

// SysColumns default column for Sys structure.
var SysColumns = [...]string{
	"id",
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

// SysFields default field for Sys structure.
var SysFields = [...]string{
	"id",
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

// Sys sys structure used for inheritance.
type Sys struct {
	Pk
	memo      String `db:"column=memo"`
	creates   String `db:"column=creates"`
	creater   String `db:"column=creater&creater=true"`
	created   Int64  `db:"column=created&created=true&default=-62135596800"`
	modifier  String `db:"column=modifier&modifier=true"`
	modified  Int64  `db:"column=modified&modified=true&default=-62135596800"`
	version   Int    `db:"column=version&version=true"`
	deletion  Int    `db:"column=deletion&deletion=true"`
	artifical Int    `db:"column=artifical"`
	history   Int    `db:"column=history"`
}

// Memo gets the value of the memo attribute in the entity.Sys.
func (me *Sys) Memo() string {
	return me.memo.Value()
}

// SetMemo sets the value of the memo attribute in the entity.Sys.
func (me *Sys) SetMemo(v string) {
	me.memo.SetValue(v)
}

// Creates gets the value of the creates attribute in the entity.Sys.
func (me *Sys) Creates() string {
	return me.creates.Value()
}

// SetCreates sets the value of the creates attribute in the entity.Sys.
func (me *Sys) SetCreates(v string) {
	me.creates.SetValue(v)
}

// Creater gets the value of the creater attribute in the entity.Sys.
func (me *Sys) Creater() string {
	return me.creater.Value()
}

// SetCreater sets the value of the creater attribute in the entity.Sys.
func (me *Sys) SetCreater(v string) {
	me.creater.SetValue(v)
}

// Created gets the value of the created attribute in the entity.Sys.
func (me *Sys) Created() int64 {
	return me.created.Value()
}

// SetCreated sets the value of the created attribute in the entity.Sys.
func (me *Sys) SetCreated(v int64) {
	me.created.SetValue(v)
}

// Modifier gets the value of the nodifier attribute in the entity.Sys.
func (me *Sys) Modifier() string {
	return me.modifier.Value()
}

// SetModifier sets the value of the nodifier attribute in the entity.Sys.
func (me *Sys) SetModifier(v string) {
	me.modifier.SetValue(v)
}

// Modified gets the value of the modified attribute in the entity.Sys.
func (me *Sys) Modified() int64 {
	return me.modified.Value()
}

// SetModified sets the value of the modified attribute in the entity.Sys.
func (me *Sys) SetModified(v int64) {
	me.modified.SetValue(v)
}

// Version gets the value of the version attribute in the entity.Sys.
func (me *Sys) Version() int {
	return me.version.Value()
}

// SetVersion sets the value of the version attribute in the entity.Sys.
func (me *Sys) SetVersion(v int) {
	me.version.SetValue(v)
}

// Deletion gets the value of the deletion attribute in the entity.Sys.
func (me *Sys) Deletion() int {
	return me.deletion.Value()
}

// SetDeletion sets the value of the deletion attribute in the entity.Sys.
func (me *Sys) SetDeletion(v int) {
	me.deletion.SetValue(v)
}

// Artifical gets the value of the artifical attribute in the entity.Sys.
func (me *Sys) Artifical() int {
	return me.artifical.Value()
}

// SetArtifical sets the value of the artifical attribute in the entity.Sys.
func (me *Sys) SetArtifical(v int) {
	me.artifical.SetValue(v)
}

// History gets the value of the history attribute in the entity.Sys.
func (me *Sys) History() int {
	return me.history.Value()
}

// SetHistory sets the value of the history attribute in the entity.Sys.
func (me *Sys) SetHistory(v int) {
	me.history.SetValue(v)
}

// Get gets the value of the column in the database.
func (me *Sys) Get(column string) interface{} {
	switch column {
	case "memo":
		return me.memo.Value()
	case "creates":
		return me.creates.Value()
	case "creater":
		return me.creater.Value()
	case "created":
		return me.created.Value()
	case "modifier":
		return me.modifier.Value()
	case "modified":
		return me.modified.Value()
	case "version":
		return me.version.Value()
	case "deletion":
		return me.deletion.Value()
	case "artifical":
		return me.artifical.Value()
	case "history":
		return me.history.Value()
	}
	return me.Pk.Get(column)
}

// GetPtr gets the value of the pointer type through the database column.
func (me *Sys) GetPtr(column string) interface{} {
	switch column {
	case "memo":
		return me.memo.ValuePtr()
	case "creates":
		return me.creates.ValuePtr()
	case "creater":
		return me.creater.ValuePtr()
	case "created":
		return me.created.ValuePtr()
	case "modifier":
		return me.modifier.ValuePtr()
	case "modified":
		return me.modified.ValuePtr()
	case "version":
		return me.version.ValuePtr()
	case "deletion":
		return me.deletion.ValuePtr()
	case "artifical":
		return me.artifical.ValuePtr()
	case "history":
		return me.history.ValuePtr()
	}
	return me.Pk.GetPtr(column)
}

// GetString gets the value of the attribute in the struct.
func (me *Sys) GetString(field string) string {
	switch strings.ToLowerFirst(field) {
	case "memo":
		return me.memo.String()
	case "creates":
		return me.creates.String()
	case "creater":
		return me.creater.String()
	case "created":
		return me.created.String()
	case "modifier":
		return me.modifier.String()
	case "modified":
		return me.modified.String()
	case "version":
		return me.version.String()
	case "deletion":
		return me.deletion.String()
	case "artifical":
		return me.artifical.String()
	case "history":
		return me.history.String()
	}
	return me.Pk.GetString(field)
}

// SetString sets the value of the attribute in the struct.
func (me *Sys) SetString(field, value string) error {
	switch strings.ToLowerFirst(field) {
	case "memo":
		return me.memo.SetString(value)
	case "creates":
		return me.creates.SetString(value)
	case "creater":
		return me.creater.SetString(value)
	case "created":
		return me.created.SetString(value)
	case "modifier":
		return me.modifier.SetString(value)
	case "modified":
		return me.modified.SetString(value)
	case "version":
		return me.version.SetString(value)
	case "deletion":
		return me.deletion.SetString(value)
	case "artifical":
		return me.artifical.SetString(value)
	case "history":
		return me.history.SetString(value)
	}
	return me.Pk.SetString(field, value)
}

// Type get entity.Type through the column of the database.
func (me *Sys) Type(column string) (Type, bool) {
	switch column {
	case "memo":
		return &me.memo, true
	case "creates":
		return &me.creates, true
	case "creater":
		return &me.creater, true
	case "created":
		return &me.created, true
	case "modifier":
		return &me.modifier, true
	case "modified":
		return &me.modified, true
	case "version":
		return &me.version, true
	case "deletion":
		return &me.deletion, true
	case "artifical":
		return &me.artifical, true
	case "history":
		return &me.history, true
	}
	return me.Pk.Type(column)
}

// Column get schema.Column by the name of the entity attribute.
func (me *Sys) Column(field string) (schema.Column, bool) {
	switch strings.ToLowerFirst(field) {
	case "memo":
		return me.memo.Column(), true
	case "creates":
		return me.creates.Column(), true
	case "creater":
		return me.creater.Column(), true
	case "created":
		return me.created.Column(), true
	case "modifier":
		return me.modifier.Column(), true
	case "modified":
		return me.modified.Column(), true
	case "version":
		return me.version.Column(), true
	case "deletion":
		return me.deletion.Column(), true
	case "artifical":
		return me.artifical.Column(), true
	case "history":
		return me.history.Column(), true
	}
	return me.Pk.Column(field)
}
