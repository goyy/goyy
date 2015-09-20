// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package db_test

import (
	"testing"
)

func TestWhereId(t *testing.T) {
	u := User.New(User{})
	query.Reset()
	query.From(u.Table)
	query.Where(u.Id.Eq("1"))
	expected := "select * from user where id = ?"
	if out := query.String(); out != expected {
		t.Errorf(`sql:"%s", want:"%s"`, out, expected)
	}
}

func TestWhereAnd(t *testing.T) {
	u := User.New(User{})
	query.Reset()
	query.From(u.Table)
	query.Where(u.Id.Ne("1").And(u.Name.Li("sa%")))
	expected := "select * from user where id <> ? and name like ?"
	if out := query.String(); out != expected {
		t.Errorf(`sql:"%s", want:"%s"`, out, expected)
	}
}

func TestWhereOr(t *testing.T) {
	u := User.New(User{})
	query.Reset()
	query.From(u.Table)
	query.Where(u.Id.Eq("1").Or(u.Name.Li("sa%")))
	expected := "select * from user where id = ? or name like ?"
	if out := query.String(); out != expected {
		t.Errorf(`sql:"%s", want:"%s"`, out, expected)
	}
}

func TestWhereAndAnd(t *testing.T) {
	u := User.New(User{})
	query.Reset()
	query.From(u.Table)
	query.Where(u.Id.Ne("1").And(u.Name.Li("sa%")).And(u.Password.Nu()))
	expected := "select * from user where id <> ? and name like ? and password is null"
	if out := query.String(); out != expected {
		t.Errorf(`sql:"%s", want:"%s"`, out, expected)
	}
}

func TestWhereAndGroup(t *testing.T) {
	u := User.New(User{})
	query.Reset()
	query.From(u.Table)
	query.Where(u.Id.Ne("1").And(u.Name.Li("sa%").Or(u.Password.Nu())))
	expected := "select * from user where id <> ? and (name like ? or password is null)"
	if out := query.String(); out != expected {
		t.Errorf(`sql:"%s", want:"%s"`, out, expected)
	}
}

func TestWhereAndAndGroup(t *testing.T) {
	u := User.New(User{})
	e := u.Id.Ne("1")
	e = e.And(u.Memo.Li("%is%"))
	e = e.And(u.Name.Li("sa%").Or(u.Password.Nu()))
	query.Reset()
	query.From(u.Table)
	query.Where(e)
	expected := "select * from user where id <> ? and memo like ? and (name like ? or password is null)"
	if out := query.String(); out != expected {
		t.Errorf(`sql:"%s", want:"%s"`, out, expected)
	}
}
