// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xsql_test

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-oci8"
	"gopkg.in/goyy/goyy.v0/comm/log"
	"gopkg.in/goyy/goyy.v0/data/dialect"
	"gopkg.in/goyy/goyy.v0/data/xsql"
)

var db xsql.DB

func init() {
	v, err := xsql.New(&dialect.MySQL{}, "db")
	xsql.SetPriority(log.Perror) // Pdebug|Perror
	if err != nil {
		panic(err)
	} else {
		db = v
	}
}
