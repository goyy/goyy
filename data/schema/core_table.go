// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package schema

import (
	"gopkg.in/goyy/goyy.v0/util/strings"
)

type table struct {
	name     string
	comment  string
	columns  map[string]*column
	primary  *column
	version  *column
	deletion *column
	creater  *column
	created  *column
	modifier *column
	modified *column
}

//========================================
// Set
//========================================

func (me *table) setPrimary(column *column) {
	me.primary = column
}

func (me *table) setVersion(column *column) {
	me.version = column
}

func (me *table) setDeletion(column *column) {
	me.deletion = column
}

func (me *table) setCreater(column *column) {
	me.creater = column
}

func (me *table) setCreated(column *column) {
	me.created = column
}

func (me *table) setModifier(column *column) {
	me.modifier = column
}

func (me *table) setModified(column *column) {
	me.modified = column
}

//========================================
// Get
//========================================

func (me *table) Name() string {
	return me.name
}

func (me *table) Comment() string {
	return me.comment
}

func (me *table) Column(columnName string) Column {
	if c, ok := me.columns[columnName]; ok {
		return c
	} else {
		return nil
	}
}

func (me *table) Columns() []Column {
	cols := make([]Column, len(me.columns))
	i := 0
	for _, col := range me.columns {
		cols[i] = col
	}
	return cols
}

func (me *table) Primary() Column {
	return me.primary
}

func (me *table) Version() Column {
	return me.version
}

func (me *table) Deletion() Column {
	return me.deletion
}

func (me *table) Creater() Column {
	return me.creater
}

func (me *table) Created() Column {
	return me.created
}

func (me *table) Modifier() Column {
	return me.modifier
}

func (me *table) Modified() Column {
	return me.modified
}

//========================================
// ToString
//========================================

func (me *table) String() string {
	return me.name
}

//========================================
// New Column
//========================================

func (me *table) COLUMN(columnName string) Column {
	return me.newColumn(columnName)
}

func (me *table) PRIMARY(columnName string) Column {
	return me.newColumn(columnName).setPrimary()
}

func (me *table) VERSION(columnName string) Column {
	return me.newColumn(columnName).setVersion()
}

func (me *table) DELETION(columnName string) Column {
	return me.newColumn(columnName).setDeletion()
}

func (me *table) CREATER(columnName string) Column {
	return me.newColumn(columnName).setCreater()
}

func (me *table) CREATED(columnName string) Column {
	return me.newColumn(columnName).setCreated()
}

func (me *table) MODIFIER(columnName string) Column {
	return me.newColumn(columnName).setModifier()
}

func (me *table) MODIFIED(columnName string) Column {
	return me.newColumn(columnName).setModified()
}

func (me *table) TRANSIENT(columnName string) Column {
	return me.newColumn(columnName).setTransient()
}

func (me *table) newColumn(columnName string) *column {
	if me == nil {
		panic("table are not allowed to nil.")
	}
	if strings.IsBlank(columnName) {
		panic("Column name are not allowed to empty.")
	}
	name := strings.ToLower(columnName)
	if v, ok := me.columns[name]; ok {
		return v
	} else {
		c := new(column)
		c.name = name
		c.table = me
		me.columns[name] = c
		return c
	}
}
