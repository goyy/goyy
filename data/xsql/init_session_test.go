// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xsql_test

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-oci8"
	"gopkg.in/goyy/goyy.v0/data/dialect"
	"gopkg.in/goyy/goyy.v0/data/xsql"
)

var factory xsql.Factory
var session xsql.Session

func init() {
	factory, err := xsql.New(&dialect.MySQL{}, "db")
	if err != nil {
		panic(err)
	}
	session, err = factory.Session()
	if err != nil {
		panic(err)
	}
}
