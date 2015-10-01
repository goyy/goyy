// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dml_test

import (
	"gopkg.in/goyy/goyy.v0/data/dialect"
	"gopkg.in/goyy/goyy.v0/data/dml"
	"testing"
)

func TestType(t *testing.T) {
	expected := "insert into sys_user (id,name,passwd,email) values (?,?,?,?)"
	user := NewUser()
	user.SetId("1")
	user.SetName("admin")
	d := dml.New(&dialect.MySQL{})
	if out, _ := d.Insert(user); out != expected {
		t.Errorf(`dml.InsertSql() = "%v", want "%v"`, out, expected)
	}
}
