// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dql_test

import (
	"fmt"
	"gopkg.in/goyy/goyy.v0/data/dialect"
	"gopkg.in/goyy/goyy.v0/data/dql"
)

func ExampleType() {
	user := NewUser()
	user.SetId("1")
	user.SetName("admin")
	d := dql.New(&dialect.MySQL{})
	v, _ := d.SelectOne(user)
	fmt.Println(v)

	// Output:select * from sys_user where id = ?
}
