// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package db_test

import (
	"testing"
)

func TestSessionInsert(t *testing.T) {
	u := User.New(User{})
	query.Reset()
	query.From(u.Table)
	query.Where(u.Id.Gt("demo-i-2").And(u.Id.Lr("demo")))
	var values []User
	err := query.List(&values)
	if err != nil {
		t.Error(err)
	}
	expected := 2
	if out := len(values); out != expected {
		t.Errorf(`query:"%v", want:"%v"`, out, expected)
	}
}

func TestSessionUpdate(t *testing.T) {
	u := User.New(User{})
	query.Reset()
	query.From(u.Table)
	query.Where(u.Id.Eq("demo-i-2"))
	var value User
	err := query.One(&value)
	if err != nil {
		t.Error(err)
	}
	expected := "admin"
	if out := string(value.Creater.Value); out != expected {
		t.Errorf(`query:"%v", want:"%v"`, out, expected)
	}
}

func TestSessionDelete(t *testing.T) {
	u := User.New(User{})
	query.Reset()
	query.From(u.Table)
	query.Where(u.Id.Eq("demo-i-2"))
	var value User
	err := query.One(&value)
	if err != nil {
		t.Error(err)
	}
	expected := "admin"
	if out := string(value.Creater.Value); out != expected {
		t.Errorf(`query:"%v", want:"%v"`, out, expected)
	}
}
