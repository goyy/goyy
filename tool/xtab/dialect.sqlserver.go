// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"strconv"
	"strings"
)

type sqlservers struct {
	Seperator, Case string
}

func (me *sqlservers) DropTable(t *table) string {
	ids := util.Case(t.Ids())
	return fmt.Sprintf("if exists (select TABLE_NAME from INFORMATION_SCHEMA.TABLES where TABLE_NAME = '%s') drop table %s%s\n\n", ids, ids, me.Seperator)
}

func (me *sqlservers) CreateTable(t *table) (sql string) {
	id := util.Case(t.Ids())
	sql = fmt.Sprintf("create table %s (\n%s\n)%s\n\n", id, me.CreateTableColumns(t), me.Seperator)
	return
}

func (me *sqlservers) CreateIndex(t *table) (sql string) {
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

func (me *sqlservers) CreateTableColumns(t *table) (sql string) {
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

func (me *sqlservers) CreateTableColumn(c *column) (sql string) {
	var id, ids, types string
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
		types = " numeric(19,0)"
	case "float":
		if c.Precision() > 0 {
			types = fmt.Sprintf(" numeric(%s,%s)", strconv.Itoa(c.Length()), strconv.Itoa(c.Precision()))
		} else {
			types = fmt.Sprintf(" numeric(%s)", strconv.Itoa(c.Length()))
		}
	case "bool":
		types = " tinyint"
	case "time":
		types = " datetime"
	case "text":
		types = " text"
	case "bytes":
		types = " image"
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
