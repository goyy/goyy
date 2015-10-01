// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dml_test

import (
	"fmt"
	"gopkg.in/goyy/goyy.v0/data/dialect"
	"gopkg.in/goyy/goyy.v0/data/dml"
)

func ExampleType() {
	user := NewUser()
	user.SetId("1")
	user.SetName("admin")
	d := dml.New(&dialect.MySQL{})
	v, _ := d.Insert(user)
	fmt.Println(v)

	// Output:insert into sys_user (id,name,passwd,email) values (?,?,?,?)
}
