// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

/*
Package dml implements sql utility functions.

Usage

	user := NewUser()
	user.SetId("1")
	user.SetName("admin")
	d := dml.New(&dialect.MySQL{})
	v, _ := d.Insert(user)
	fmt.Println(v)
*/
package dml
