// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/goyy/goyy.v0/util/files"
)

func expSQL() {
	expDataSQL()
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
		filename := "./sql/ddl/" + p.ID() + ".sql"
		dir := files.Dir(filename)
		if !files.IsExist(dir) {
			files.MkdirAll(dir, 0755)
		}
		files.Remove(filename)
		var content string
		f, ferr := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0755)
		defer f.Close()
		if ferr != nil {
			log.Fatal(ferr)
		}
		for _, t := range conf.tables {
			if t.module.project.ID() == p.ID() && t.Generate() == "true" {
				content += dialect.DropTable(t)
			}
		}
		for _, t := range conf.tables {
			if t.module.project.ID() == p.ID() && t.Generate() == "true" {
				content += dialect.CreateTable(t)
			}
		}
		for _, t := range conf.tables {
			if t.module.project.ID() == p.ID() && t.Generate() == "true" {
				content += dialect.CreateIndex(t)
			}
		}
		for _, t := range conf.tables {
			if t.module.project.ID() == p.ID() && t.Generate() == "true" {
				content += dialect.CreateUniqueIndex(t)
			}
		}
		f.WriteString(content)
	}
}

func expDataSQL() {
	var dir = "./sql/dml/"
	types := []string{"area", "org", "dict", "post", "post.menu", "role", "role.post", "user", "user.role"}
	for _, typ := range types {
		writeBy("data."+typ, dir)
	}
	mergeFile(dir, `^insert.[\S]+.sql$`, "merge-file.sql")
}

func writeBy(typ, dir string) error {
	var tmpl string
	var dstfile string
	switch typ {
	case "data.area":
		tmpl = tmplDataArea
		dstfile = "insert.sys_area.sql"
	case "data.org":
		tmpl = tmplDataOrg
		dstfile = "insert.sys_org.sql"
	case "data.dict":
		tmpl = tmplDataDict
		dstfile = "insert.sys_dict.sql"
	case "data.post":
		tmpl = tmplDataPost
		dstfile = "insert.sys_post.sql"
	case "data.post.menu":
		tmpl = tmplDataPostMenu
		dstfile = "insert.sys_post_menu.sql"
	case "data.role":
		tmpl = tmplDataRole
		dstfile = "insert.sys_role.sql"
	case "data.role.post":
		tmpl = tmplDataRolePost
		dstfile = "insert.sys_role_post.sql"
	case "data.user":
		tmpl = tmplDataUser
		dstfile = "insert.sys_user.sql"
	case "data.user.role":
		tmpl = tmplDataUserRole
		dstfile = "insert.sys_user_role.sql"
	}
	dstfile = filepath.Join(dir, dstfile)
	if !files.IsExist(dir) {
		files.MkdirAll(dir, 0755)
	}

	return write(tmpl, dstfile)
}
