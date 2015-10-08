// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"os"
)

func expSQL() {
	for _, p := range conf.projects {
		var dialect dialects
		switch p.database.driverName {
		case "mymysql", "mysql":
			dialect = &mysqls{conf.Settings.Statement.Seperator, conf.Settings.Statement.Case}
		case "oci8", "oracle":
			dialect = &oracles{conf.Settings.Statement.Seperator, conf.Settings.Statement.Case}
		case "postgres":
			dialect = &postgresql{conf.Settings.Statement.Seperator, conf.Settings.Statement.Case}
		case "sqlserver":
			dialect = &sqlservers{conf.Settings.Statement.Seperator, conf.Settings.Statement.Case}
		case "sqlite", "sqlite3":
			dialect = &sqlite{conf.Settings.Statement.Seperator, conf.Settings.Statement.Case}
		default:
			dialect = &mysqls{conf.Settings.Statement.Seperator, conf.Settings.Statement.Case}
		}
		filename := p.Id() + ".sql"
		os.Remove(filename)
		var content string
		f, ferr := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
		defer f.Close()
		if ferr != nil {
			log.Fatal(ferr)
		}
		for _, t := range conf.tables {
			if t.module.project.Id() == p.Id() && t.Generate() == "true" {
				content += dialect.DropTable(t)
			}
		}
		for _, t := range conf.tables {
			if t.module.project.Id() == p.Id() && t.Generate() == "true" {
				content += dialect.CreateTable(t)
			}
		}
		for _, t := range conf.tables {
			if t.module.project.Id() == p.Id() && t.Generate() == "true" {
				content += dialect.CreateIndex(t)
			}
		}
		f.WriteString(content)
	}
}
