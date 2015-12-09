// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"strconv"
	"strings"
)

type oracles struct {
	Seperator, Case string
}

func (me *oracles) DropTable(t *table) string {
	return fmt.Sprintf("drop table %s cascade constraints%s\n\n", util.Case(t.Ids()), me.Seperator)
}

func (me *oracles) CreateTable(t *table) (sql string) {
	id := util.Case(t.Ids())
	sql = fmt.Sprintf("create table %s (\n%s\n)%s\n\n", id, me.CreateTableColumns(t), me.Seperator)
	if strings.TrimSpace(t.Comment()) != "" {
		sql += fmt.Sprintf("comment on table %s is '%s'%s\n", id, t.Comment(), me.Seperator)
		sql += me.CreateTableComment(t)
	}
	return
}

func (me *oracles) CreateIndex(t *table) (sql string) {
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

func (me *oracles) CreateTableColumns(t *table) (sql string) {
	var id, pk string
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
			id, pk = util.Case(c.Id()), util.Case("pk_"+t.Ids())
		}
	}
	if b {
		sql += fmt.Sprintf(",\n\tconstraint %s primary key (%s)", pk, id)
	}
	return
}

func (me *oracles) CreateTableColumn(c *column) (sql string) {
	var id, ids, types string
	//===========id============
	id = util.Case(c.Id())
	ids = "\t" + util.Pad(id)
	//==========types==========
	switch c.Types() {
	case "string":
		types = fmt.Sprintf(" varchar2(%s)", strconv.Itoa(c.Length()))
	case "int":
		types = " integer"
	case "long":
		types = " number(19)"
	case "float":
		if c.Precision() > 0 {
			types = fmt.Sprintf(" number(%s,%s)", strconv.Itoa(c.Length()), strconv.Itoa(c.Precision()))
		} else {
			types = fmt.Sprintf(" number(%s)", strconv.Itoa(c.Length()))
		}
	case "bool":
		types = " number(1)"
	case "time":
		types = " date"
	case "text":
		types = " clob"
	case "bytes":
		types = " blob"
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
	sql = ids + types
	return
}

func (me *oracles) CreateTableComment(t *table) (sql string) {
	var tid, cid string
	for _, c := range t.Columns() {
		tid, cid = util.Case(t.Ids()), c.Id()
		sql += fmt.Sprintf("comment on column %s.%s is '%s'%s\n", tid, cid, c.Comment(), me.Seperator)
	}
	sql += "\n"
	return
}
