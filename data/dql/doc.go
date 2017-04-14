// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package dql implements sql utility functions.
//
// Usage
//
// 	user := NewUser()
// 	user.SetId("1")
// 	user.SetName("admin")
// 	d := dql.New(&dialect.MySQL{})
// 	v, _ := d.SelectOne(user)
// 	fmt.Println(v)
package dql
