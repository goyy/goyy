// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package sqls_test

import (
	"gopkg.in/goyy/goyy.v0/util/sqls"
	"testing"
)

func TestParseCountSql(t *testing.T) {
	in := "select id, name from table where id = 1"
	expected := "select count(*) from table where id = 1"
	if out := sqls.ParseCountSql(in); out != expected {
		t.Errorf(`sqls.ParseCountSql("%s") = "%s", want "%s"`, in, out, expected)
	}
}

func TestParseCountSqlUpper(t *testing.T) {
	in := "SELECT id, name FROM table where id = 1"
	expected := "select count(*) FROM table where id = 1"
	if out := sqls.ParseCountSql(in); out != expected {
		t.Errorf(`sqls.ParseCountSql("%s") = "%s", want "%s"`, in, out, expected)
	}
}

func TestParseCountSqlMore(t *testing.T) {
	in := "select id, name,(select count(*) from table2) count from table where id = 1 and exists (select 1 from table3)"
	expected := "select count(*) from table where id = 1 and exists (select 1 from table3)"
	if out := sqls.ParseCountSql(in); out != expected {
		t.Errorf(`sqls.ParseCountSql("%s") = "%s", want "%s"`, in, out, expected)
	}
}

func TestParseCountSqlMoreUpper(t *testing.T) {
	in := "SELECT id, name,(SELECT count(*) FROM table2) count FROM table where id = 1 and exists (SELECT 1 FROM table3)"
	expected := "select count(*) FROM table where id = 1 and exists (SELECT 1 FROM table3)"
	if out := sqls.ParseCountSql(in); out != expected {
		t.Errorf(`sqls.ParseCountSql("%s") = "%s", want "%s"`, in, out, expected)
	}
}
