// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dialect

import (
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

func NewMysql() Dialect {
	return &mysql{}
}

type mysql struct {
}

func (me *mysql) Quote(s string) string {
	//return fmt.Sprintf(`"%s"`, s)
	return strings.ToUpper(s)
}

func (me *mysql) SqlType(domain Domain) string {
	switch domain.Type {
	case DTString:
		if domain.Length > 0 && domain.Length < 65532 {
			return me.Quote(fmt.Sprintf("varchar(%d)", domain.Length))
		}
		return me.Quote("longtext")
	case DTBool:
		return me.Quote("boolean")
	case DTInt:
		if domain.Length > 0 && domain.Length < 10 {
			return me.Quote("int")
		}
		return me.Quote("bigint")
	case DTFloat:
		return me.Quote("double")
	case DTTime:
		return me.Quote("timestamp")
	case DTByte:
		if domain.Length > 0 && domain.Length < 65532 {
			return me.Quote(fmt.Sprintf("varbinary(%d)", domain.Length))
		}
		return me.Quote("longblob")
	}
	panic("invalid sql type")
}

func (me *mysql) CreateTable(table *Table, columns ...*Column) string {
	ddl := []string{"CREATE TABLE "}
	ddl = append(ddl, me.Quote(table.Id), " (\r")
	for i, column := range columns {
		c := "\t" + me.columnSql(column)
		ddl = append(ddl, c)
		if i < len(columns)-1 {
			ddl = append(ddl, ",\r")
		}
	}
	ddl = append(ddl, "\r)")
	return strings.Join(ddl, "")
}

func (me *mysql) DropTable(table string) string {
	return fmt.Sprintf("DROP TABLE IF EXISTS %v", me.Quote(table))
}

func (me *mysql) RenameTable(from, to string) string {
	return fmt.Sprintf(
		"ALTER TABLE %v RENAME TO %v",
		me.Quote(from),
		me.Quote(to),
	)
}

func (me *mysql) AddColumn(table string, column *Column) string {
	return fmt.Sprintf(
		"ALTER TABLE %v ADD COLUMN %v",
		me.Quote(table),
		me.columnSql(column),
	)
}

func (me *mysql) RenameColumn(table, from, to string) string {
	return fmt.Sprintf(
		"ALTER TABLE %v RENAME COLUMN %v TO %v",
		me.Quote(table),
		me.Quote(from),
		me.Quote(to),
	)
}

func (me *mysql) ChangeColumn(table string, column *Column) string {
	return fmt.Sprintf(
		"ALTER TABLE %v ALTER COLUMN %v",
		me.Quote(table),
		me.columnSql(column),
	)
}

func (me *mysql) DropColumn(table, column string) string {
	return fmt.Sprintf(
		"ALTER TABLE %v DROP COLUMN %v",
		me.Quote(table),
		me.Quote(column),
	)
}

func (me *mysql) CreateIndex(name, table string, unique bool, columns ...string) string {
	a := []string{"CREATE"}
	if unique {
		a = append(a, "UNIQUE")
	}
	quotedColumns := make([]string, 0, len(columns))
	for _, c := range columns {
		quotedColumns = append(quotedColumns, me.Quote(c))
	}
	a = append(a, fmt.Sprintf(
		"INDEX %v ON %v (%v)",
		me.Quote(name),
		me.Quote(table),
		strings.Join(quotedColumns, ", "),
	))
	return strings.Join(a, " ")
}

func (me *mysql) DropIndex(name string) string {
	return fmt.Sprintf("DROP INDEX %v", me.Quote(name))
}

func (me *mysql) columnSql(column *Column) string {
	c := []string{
		me.Quote(column.Id),
		me.SqlType(column.Domain),
	}
	if column.Nullable {
		c = append(c, "NOT NULL")
	}
	if d := column.Default; d != "" {
		c = append(c, fmt.Sprintf("DEFAULT %v", d))
	}
	if d := column.Key; d == "PRI" {
		c = append(c, "PRIMARY KEY")
	}
	return strings.Join(c, " ")
}
