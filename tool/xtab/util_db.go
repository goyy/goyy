// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"sync"

	"gopkg.in/goyy/goyy.v0/data/dialect"
	"gopkg.in/goyy/goyy.v0/data/xsql"
)

var once sync.Once
var xdb xsql.DB

func getDB(driverName, projId string) xsql.DB {
	once.Do(func() {
		d := getDialect(driverName)
		v, err := xsql.New(d, projId)
		if err != nil {
			panic(err)
		} else {
			xdb = v
		}
	})
	return xdb
}

func getDDL(driverName, projId string) string {
	dialect := getDialects(driverName)
	var content string
	for _, t := range conf.tables {
		if t.module.project.ID() == projId && t.Generate() == "true" {
			content += dialect.DropTable(t)
		}
	}
	for _, t := range conf.tables {
		if t.module.project.ID() == projId && t.Generate() == "true" {
			content += dialect.CreateTable(t)
		}
	}
	for _, t := range conf.tables {
		if t.module.project.ID() == projId && t.Generate() == "true" {
			content += dialect.CreateIndex(t)
		}
	}
	for _, t := range conf.tables {
		if t.module.project.ID() == projId && t.Generate() == "true" {
			content += dialect.CreateUniqueIndex(t)
		}
	}
	return content
}

func getDialect(driverName string) dialect.Interface {
	switch driverName {
	case "mymysql", "mysql":
		return &dialect.MySQL{}
	case "oci8", "oracle":
		return &dialect.Oracle{}
	case "postgres":
		return &dialect.PostgreSQL{}
	case "sqlserver":
		return &dialect.SQLServer{}
	case "sqlite", "sqlite3":
		return &dialect.Sqlite{}
	default:
		return &dialect.MySQL{}
	}
}

func getDialects(driverName string) dialects {
	switch driverName {
	case "mymysql", "mysql":
		return &mysqls{conf.Settings.Statement.Seperator, conf.Settings.Statement.Case}
	case "oci8", "oracle":
		return &oracles{conf.Settings.Statement.Seperator, conf.Settings.Statement.Case}
	case "postgres":
		return &postgresql{conf.Settings.Statement.Seperator, conf.Settings.Statement.Case}
	case "sqlserver":
		return &sqlservers{conf.Settings.Statement.Seperator, conf.Settings.Statement.Case}
	case "sqlite", "sqlite3":
		return &sqlite{conf.Settings.Statement.Seperator, conf.Settings.Statement.Case}
	default:
		return &mysqls{conf.Settings.Statement.Seperator, conf.Settings.Statement.Case}
	}
}
