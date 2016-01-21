// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package schema

import (
	"gopkg.in/goyy/goyy.v0/util/strings"
)

type column struct {
	table     *table
	name      string
	comment   string
	dict      string
	primary   bool
	version   bool
	deletion  bool
	creater   bool
	created   bool
	modifier  bool
	modified  bool
	transient bool
}

//========================================
// Set
//========================================

func (me *column) setPrimary() *column {
	me.primary = true
	me.table.setPrimary(me)
	return me
}

func (me *column) setVersion() *column {
	me.version = true
	me.table.setVersion(me)
	return me
}

func (me *column) setDeletion() *column {
	me.deletion = true
	me.table.setDeletion(me)
	return me
}

func (me *column) setCreater() *column {
	me.creater = true
	me.table.setCreater(me)
	return me
}

func (me *column) setCreated() *column {
	me.created = true
	me.table.setCreated(me)
	return me
}

func (me *column) setModifier() *column {
	me.modifier = true
	me.table.setModifier(me)
	return me
}

func (me *column) setModified() *column {
	me.modified = true
	me.table.setModified(me)
	return me
}

func (me *column) setTransient() *column {
	me.transient = true
	return me
}

//========================================
// Get
//========================================

func (me *column) Table() Table {
	return me.table
}

func (me *column) Name() string {
	return me.name
}

func (me *column) Comment() string {
	if strings.IsBlank(me.comment) {
		return strings.ToUpper(me.name)
	}
	return me.comment
}

func (me *column) Dict() string {
	return me.dict
}

func (me *column) SetDict(value string) {
	me.dict = value
}

func (me *column) IsPrimary() bool {
	return me.primary
}

func (me *column) IsVersion() bool {
	return me.version
}

func (me *column) IsDeletion() bool {
	return me.deletion
}

func (me *column) IsCreater() bool {
	return me.creater
}

func (me *column) IsCreated() bool {
	return me.created
}

func (me *column) IsModifier() bool {
	return me.modifier
}

func (me *column) IsModified() bool {
	return me.modified
}

func (me *column) IsTransient() bool {
	return me.transient
}

//========================================
// ToString
//========================================

func (me *column) String() string {
	return me.name
}
