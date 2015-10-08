// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"strconv"
	"strings"
)

type sqlite struct {
	Seperator, Case string
}

func (me *sqlite) DropTable(t *table) string {
	return fmt.Sprintf("drop table %s%s\n\n", util.Case(t.Ids()), me.Seperator)
}

func (me *sqlite) CreateTable(t *table) (sql string) {
	id := util.Case(t.Ids())
	sql = fmt.Sprintf("create table if not exists %s (\n%s\n)%s\n\n", id, me.CreateTableColumns(t), me.Seperator)
	return
}

func (me *sqlite) CreateIndex(t *table) (sql string) {
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

func (me *sqlite) CreateTableColumns(t *table) (sql string) {
	var l int = len(t.Columns()) - 1
	for i, c := range t.Columns() {
		if i == l {
			sql += me.CreateTableColumn(c)
		} else {
			sql += me.CreateTableColumn(c) + ",\n"
		}
	}
	return
}

func (me *sqlite) CreateTableColumn(c *column) (sql string) {
	var id, ids, types string
	//===========id============
	id = util.Case(c.Id())
	ids = "\t" + util.Pad(id)
	//==========types==========
	switch c.Types() {
	case "string":
		types = fmt.Sprintf(" varchar(%s)", strconv.Itoa(c.Length()))
	case "int":
		types = " integer"
	case "long":
		types = " bigint"
	case "float":
		if c.Precision() > 0 {
			types = fmt.Sprintf(" numeric(%s,%s)", strconv.Itoa(c.Length()), strconv.Itoa(c.Precision()))
		} else {
			types = fmt.Sprintf(" numeric(%s)", strconv.Itoa(c.Length()))
		}
	case "bool":
		types = " boolean"
	case "time":
		types = " timestamp"
	case "text":
		types = " text"
	case "bytes":
		types = " bytea"
	}
	//========default=========
	if strings.TrimSpace(c.Defaults()) != "" {
		types += " default " + c.Defaults()
	}
	//========not null=========
	if c.Id() == "id" {
		types += " not null" + " primary key"
	} else {
		if strings.TrimSpace(c.Nullable()) == "false" {
			types += " not null"
		}
	}
	sql = ids + types
	return
}

func (me *sqlite) CreateTableComment(t *table) (sql string) {
	var tid, cid string
	for _, c := range t.Columns() {
		tid, cid = util.Case(t.Ids()), c.Id()
		sql += fmt.Sprintf("comment on column %s.%s is '%s'%s\n\n", tid, cid, c.Comment(), me.Seperator)
	}
	return
}
