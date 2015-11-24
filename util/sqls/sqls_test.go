// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package sqls_test

import (
	"fmt"
	"gopkg.in/goyy/goyy.v0/data/dialect"
	"gopkg.in/goyy/goyy.v0/util/sqls"
	"testing"
)

func TestParseCountSql(t *testing.T) {
	s := []struct {
		in, expected string
	}{
		{
			"select id, name from table where id = 1",
			"select count(*) from table where id = 1",
		},
		{
			"SELECT id, name FROM table where id = 1",
			"select count(*) FROM table where id = 1",
		},
		{
			"select id, name,(select count(*) from table2) count from table where id = 1 and exists (select 1 from table3)",
			"select count(*) from table where id = 1 and exists (select 1 from table3)",
		},
		{
			"SELECT id, name,(SELECT count(*) FROM table2) count FROM table where id = 1 and exists (SELECT 1 FROM table3)",
			"select count(*) FROM table where id = 1 and exists (SELECT 1 FROM table3)",
		},
	}
	for _, v := range s {
		if out := sqls.ParseCountSql(v.in); out != v.expected {
			t.Errorf(`sqls.ParseCountSql("%s") = "%s", want "%s"`, v.in, out, v.expected)
		}
	}
}

func TestParseNamedSql(t *testing.T) {
	s := []struct {
		dia             dialect.Interface
		sql             string
		args            map[string]interface{}
		sqlout, argsout string
	}{
		{
			&dialect.Oracle{},
			"select * from users where id = #{id} and name = #{name}",
			map[string]interface{}{"id": "1", "name": "goyy", "memo": "memo"},
			"select * from users where id = :0 and name = :1",
			"[`1` `goyy`]",
		},
		{
			&dialect.MySQL{},
			"select * from users where id = #{id} and name = #{name}",
			map[string]interface{}{"id": "1", "name": "goyy", "memo": "memo"},
			"select * from users where id = ? and name = ?",
			"[`1` `goyy`]",
		},
	}
	for _, v := range s {
		sqlout, argsout, _ := sqls.ParseNamedSql(v.dia, v.sql, v.args)
		if sqlout != v.sqlout {
			t.Errorf("sqls.ParseNamedSql(%#q, %#q, %#q) = %#q, want %#q", v.dia.Type(), v.sql, v.args, sqlout, v.sqlout)
		}
		if fmt.Sprintf("%#q", argsout) != v.argsout {
			t.Errorf("sqls.ParseNamedSql(%#q, %#q, %#q) = %#q, want %#q", v.dia.Type(), v.sql, v.args, argsout, v.argsout)
		}
	}
}
