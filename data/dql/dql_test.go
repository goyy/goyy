// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dql_test

import (
	"gopkg.in/goyy/goyy.v0/data/dialect"
	"gopkg.in/goyy/goyy.v0/data/dql"
	"testing"
)

func TestType(t *testing.T) {
	expected := "select * from sys_user where id = ?"
	user := NewUser()
	user.SetId("1")
	user.SetName("admin")
	d := dql.New(&dialect.MySQL{})
	if out, _ := d.SelectOne(user); out != expected {
		t.Errorf(`dql.SelectOneSql() = "%v", want "%v"`, out, expected)
	}
}
