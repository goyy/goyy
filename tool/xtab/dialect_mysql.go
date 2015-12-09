// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"strconv"
	"strings"
)

type mysqls struct {
	Seperator, Case string
}

func (me *mysqls) DropTable(t *table) string {
	return fmt.Sprintf("drop table if exists %s%s\n\n", util.Case(t.Ids()), me.Seperator)
}

func (me *mysqls) CreateTable(t *table) (sql string) {
	id := util.Case(t.Ids())
	sql = fmt.Sprintf("create table %s (\n%s\n)%s\n\n", id, me.CreateTableColumns(t), me.Seperator)
	if strings.TrimSpace(t.Comment()) != "" {
		sql += fmt.Sprintf("alter table %s comment '%s'%s\n\n", id, t.Comment(), me.Seperator)
	}
	return
}

func (me *mysqls) CreateIndex(t *table) (sql string) {
	tid := util.Case(t.Ids())
	for _, c := range t.Columns() {
		if c.Index() == "true" {
			cid := util.Case(c.Id())
			index := util.Case("idx_" + t.Ids() + "_" + cid)
			sql += fmt.Sprintf("create index %s on %s(%s)%s\n\n", index, tid, cid, me.Seperator)
		}
	}
	return
}

func (me *mysqls) CreateTableColumns(t *table) (sql string) {
	var id string
	var b bool
	var l int = len(t.Columns()) - 1
	for i, c := range t.Columns() {
		if i == l {
			sql += me.CreateTableColumn(c)
		} else {
			sql += me.CreateTableColumn(c) + ",\n"
		}
		if c.Id() == "id" {
			b = true
			id = util.Case(c.Id())
		}
	}
	if b {
		sql += fmt.Sprintf(",\n\tprimary key (%s)", id)
	}
	return
}

func (me *mysqls) CreateTableColumn(c *column) (sql string) {
	var id, ids, types, comment string
	//===========id============
	id = util.Case(c.Id())
	ids = "\t" + util.Pad(id)
	//==========types==========
	switch c.Types() {
	case "string":
		types = fmt.Sprintf(" varchar(%s)", strconv.Itoa(c.Length()))
	case "int":
		types = " int"
	case "long":
		types = " bigint"
	case "float":
		types = fmt.Sprintf(" decimal(%s,%s)", strconv.Itoa(c.Length()), strconv.Itoa(c.Precision()))
	case "bool":
		types = " tinyint"
	case "time":
		types = " datetime"
	case "text":
		types = " text"
	case "bytes":
		types = " longblob"
	}
	//========default=========
	if strings.TrimSpace(c.Defaults()) != "" {
		types += " default " + c.Defaults()
	}
	//========not null=========
	if c.Id() == "id" {
		types += " not null"
	} else {
		if strings.TrimSpace(c.Nullable()) == "false" {
			types += " not null"
		}
	}
	//=========comment=========
	if strings.TrimSpace(c.Comment()) == "" {
		sql = ids + types
	} else {
		comment = fmt.Sprintf(" comment '%s'", c.Comment())
		sql = ids + util.Pad(types) + comment
	}
	return
}
