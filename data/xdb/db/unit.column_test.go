// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package db_test

import (
	"testing"
)

func TestColumnEq(t *testing.T) {
	u := User.New(User{})
	query.Reset()
	query.From(u.Table)
	query.Where(u.Id.Eq("1"))
	expected := "select * from user where id = ?"
	if out := query.String(); out != expected {
		t.Errorf(`sql:"%s", want:"%s"`, out, expected)
	}
}

func TestColumnNe(t *testing.T) {
	u := User.New(User{})
	query.Reset()
	query.From(u.Table)
	query.Where(u.Id.Ne("1"))
	expected := "select * from user where id <> ?"
	if out := query.String(); out != expected {
		t.Errorf(`sql:"%s", want:"%s"`, out, expected)
	}
}

func TestColumnGt(t *testing.T) {
	u := User.New(User{})
	query.Reset()
	query.From(u.Table)
	query.Where(u.Id.Gt("1"))
	expected := "select * from user where id > ?"
	if out := query.String(); out != expected {
		t.Errorf(`sql:"%s", want:"%s"`, out, expected)
	}
}

func TestColumnLt(t *testing.T) {
	u := User.New(User{})
	query.Reset()
	query.From(u.Table)
	query.Where(u.Id.Lt("1"))
	expected := "select * from user where id < ?"
	if out := query.String(); out != expected {
		t.Errorf(`sql:"%s", want:"%s"`, out, expected)
	}
}

func TestColumnGe(t *testing.T) {
	u := User.New(User{})
	query.Reset()
	query.From(u.Table)
	query.Where(u.Id.Ge("1"))
	expected := "select * from user where id >= ?"
	if out := query.String(); out != expected {
		t.Errorf(`sql:"%s", want:"%s"`, out, expected)
	}
}

func TestColumnLe(t *testing.T) {
	u := User.New(User{})
	query.Reset()
	query.From(u.Table)
	query.Where(u.Id.Le("1"))
	expected := "select * from user where id <= ?"
	if out := query.String(); out != expected {
		t.Errorf(`sql:"%s", want:"%s"`, out, expected)
	}
}

func TestColumnLi(t *testing.T) {
	u := User.New(User{})
	query.Reset()
	query.From(u.Table)
	query.Where(u.Id.Li("1%"))
	expected := "select * from user where id like ?"
	if out := query.String(); out != expected {
		t.Errorf(`sql:"%s", want:"%s"`, out, expected)
	}
}

func TestColumnBe(t *testing.T) {
	u := User.New(User{})
	query.Reset()
	query.From(u.Table)
	query.Where(u.Id.Be("1", "2"))
	expected := "select * from user where id between ? and ?"
	if out := query.String(); out != expected {
		t.Errorf(`sql:"%s", want:"%s"`, out, expected)
	}
}

func TestColumnIn(t *testing.T) {
	u := User.New(User{})
	query.Reset()
	query.From(u.Table)
	query.Where(u.Id.In("1", "2", "3"))
	expected := "select * from user where id in ?"
	if out := query.String(); out != expected {
		t.Errorf(`sql:"%s", want:"%s"`, out, expected)
	}
}

func TestColumnNu(t *testing.T) {
	u := User.New(User{})
	query.Reset()
	query.From(u.Table)
	query.Where(u.Id.Nu())
	expected := "select * from user where id is null"
	if out := query.String(); out != expected {
		t.Errorf(`sql:"%s", want:"%s"`, out, expected)
	}
}

func TestColumnNn(t *testing.T) {
	u := User.New(User{})
	query.Reset()
	query.From(u.Table)
	query.Where(u.Id.Nn())
	expected := "select * from user where id is not null"
	if out := query.String(); out != expected {
		t.Errorf(`sql:"%s", want:"%s"`, out, expected)
	}
}

func TestColumnAsc(t *testing.T) {
	u := User.New(User{})
	query.Reset()
	query.From(u.Table)
	query.OrderBy(u.Id.Asc())
	expected := "select * from user order by id asc"
	if out := query.String(); out != expected {
		t.Errorf(`sql:"%s", want:"%s"`, out, expected)
	}
}

func TestColumnDesc(t *testing.T) {
	u := User.New(User{})
	query.Reset()
	query.From(u.Table)
	query.OrderBy(u.Id.Desc())
	expected := "select * from user order by id desc"
	if out := query.String(); out != expected {
		t.Errorf(`sql:"%s", want:"%s"`, out, expected)
	}
}

func TestColumnCount(t *testing.T) {
	u := User.New(User{})
	query.Reset()
	query.Select(u.Id.Count())
	query.From(u.Table)
	expected := "select count(id) from user"
	if out := query.String(); out != expected {
		t.Errorf(`sql:"%s", want:"%s"`, out, expected)
	}
}

func TestColumnMax(t *testing.T) {
	u := User.New(User{})
	query.Reset()
	query.Select(u.Id.Max())
	query.From(u.Table)
	expected := "select max(id) from user"
	if out := query.String(); out != expected {
		t.Errorf(`sql:"%s", want:"%s"`, out, expected)
	}
}

func TestColumnMin(t *testing.T) {
	u := User.New(User{})
	query.Reset()
	query.Select(u.Id.Min())
	query.From(u.Table)
	expected := "select min(id) from user"
	if out := query.String(); out != expected {
		t.Errorf(`sql:"%s", want:"%s"`, out, expected)
	}
}

func TestColumnSum(t *testing.T) {
	u := User.New(User{})
	query.Reset()
	query.Select(u.Id.Sum())
	query.From(u.Table)
	expected := "select sum(id) from user"
	if out := query.String(); out != expected {
		t.Errorf(`sql:"%s", want:"%s"`, out, expected)
	}
}

func TestColumnAvg(t *testing.T) {
	u := User.New(User{})
	query.Reset()
	query.Select(u.Id.Avg())
	query.From(u.Table)
	expected := "select avg(id) from user"
	if out := query.String(); out != expected {
		t.Errorf(`sql:"%s", want:"%s"`, out, expected)
	}
}
