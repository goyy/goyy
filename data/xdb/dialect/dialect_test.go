// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dialect_test

import (
	"../dialect"
	"testing"
)

func TestCreateTable(t *testing.T) {
	table := &dialect.Table{Id: "user"}
	cid := &dialect.Column{Id: "id", Domain: dialect.Domains.String(50), Key: "PRI"}
	cname := &dialect.Column{Id: "name", Domain: dialect.Domains.String(200), Nullable: true}
	cage := &dialect.Column{Id: "age", Domain: dialect.Domains.Int(3), Default: "18"}
	expected := "CREATE TABLE USER (\r\tID VARCHAR(50) PRIMARY KEY,\r\tNAME VARCHAR(200) NOT NULL,\r\tAGE INT DEFAULT 18\r)"
	mysql := dialect.NewMysql()
	if out := mysql.CreateTable(table, cid, cname, cage); out != expected {
		t.Errorf(`dialect.CreateTable("%s", "%s", "%s") = "%s", want "%s"`, table, cid, cname, out, expected)
	}
}

func TestDropTable(t *testing.T) {
	in := "user"
	expected := "DROP TABLE IF EXISTS USER"
	mysql := dialect.NewMysql()
	if out := mysql.DropTable(in); out != expected {
		t.Errorf(`dialect.DropTable("%s") = "%s", want "%s"`, in, out, expected)
	}
}

func TestRenameTable(t *testing.T) {
	from, to := "user", "users"
	expected := "ALTER TABLE USER RENAME TO USERS"
	mysql := dialect.NewMysql()
	if out := mysql.RenameTable(from, to); out != expected {
		t.Errorf(`dialect.RenameTable("%s", "%s") = "%s", want "%s"`, from, to, out, expected)
	}
}

func TestAddColumn(t *testing.T) {
	table := "user"
	column := &dialect.Column{Id: "memo", Domain: dialect.Domains.String(200), Nullable: true}
	expected := "ALTER TABLE USER ADD COLUMN MEMO VARCHAR(200) NOT NULL"
	mysql := dialect.NewMysql()
	if out := mysql.AddColumn(table, column); out != expected {
		t.Errorf(
			`dialect.AddColumn("%s", "%s") = "%s", want "%s"`,
			table, column, out, expected)
	}
}

func TestRenameColumn(t *testing.T) {
	table, from, to := "user", "id", "uid"
	expected := "ALTER TABLE USER RENAME COLUMN ID TO UID"
	mysql := dialect.NewMysql()
	if out := mysql.RenameColumn(table, from, to); out != expected {
		t.Errorf(
			`dialect.RenameColumn("%s", "%s", "%s") = "%s", want "%s"`,
			table, from, to, out, expected)
	}
}

func TestChangeColumn(t *testing.T) {
	table := "user"
	column := &dialect.Column{Id: "memo", Domain: dialect.Domains.String(200), Nullable: true}
	expected := "ALTER TABLE USER ALTER COLUMN MEMO VARCHAR(200) NOT NULL"
	mysql := dialect.NewMysql()
	if out := mysql.ChangeColumn(table, column); out != expected {
		t.Errorf(
			`dialect.ChangeColumn("%s", "%s") = "%s", want "%s"`,
			table, column, out, expected)
	}
}

func TestDropColumn(t *testing.T) {
	table, column := "user", "age"
	expected := "ALTER TABLE USER DROP COLUMN AGE"
	mysql := dialect.NewMysql()
	if out := mysql.DropColumn(table, column); out != expected {
		t.Errorf(
			`dialect.DropColumn("%s", "%s") = "%s", want "%s"`,
			table, column, out, expected)
	}
}

func TestCreateIndex(t *testing.T) {
	name, table, unique, cname, cage := "idx_user", "user", false, "name", "age"
	expected := "CREATE INDEX IDX_USER ON USER (NAME, AGE)"
	mysql := dialect.NewMysql()
	if out := mysql.CreateIndex(name, table, unique, cname, cage); out != expected {
		t.Errorf(
			`dialect.CreateIndex("%s", "%s", "%s", "%s", "%s") = "%s", want "%s"`,
			name, table, unique, cname, cage, out, expected)
	}
}

func TestCreateIndexUnique(t *testing.T) {
	name, table, unique, cname, cage := "idx_user", "user", true, "name", "age"
	expected := "CREATE UNIQUE INDEX IDX_USER ON USER (NAME, AGE)"
	mysql := dialect.NewMysql()
	if out := mysql.CreateIndex(name, table, unique, cname, cage); out != expected {
		t.Errorf(
			`dialect.CreateIndex("%s", "%s", "%s", "%s", "%s") = "%s", want "%s"`,
			name, table, unique, cname, cage, out, expected)
	}
}

func TestDropIndex(t *testing.T) {
	in := "idx_user"
	expected := "DROP INDEX IDX_USER"
	mysql := dialect.NewMysql()
	if out := mysql.DropIndex(in); out != expected {
		t.Errorf(`dialect.DropIndex("%s") = "%s", want "%s"`, in, out, expected)
	}
}
