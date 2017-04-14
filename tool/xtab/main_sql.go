// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"os"
	"path/filepath"

	"gopkg.in/goyy/goyy.v0/util/files"
)

var types = []string{"area", "org", "dict", "post", "post.menu", "role", "role.post", "user", "user.role"}

var ddlMap = map[string]string{}

func expSQL() {
	for _, p := range conf.projects {
		dbName := p.database.name
		driverName := p.database.driverName
		proj := p.ID()
		params := map[string]string{
			"driverName": driverName,
		}
		content := getDDL(driverName, proj)
		writeProj(p.ID(), content)
		if _, ok := ddlMap[dbName]; !ok {
			err := expDataSQL(params)
			if err != nil {
				logger.Errorln(err)
				return
			}
		}
		ddlMap[dbName] = proj
	}
}

func expDataSQL(params map[string]string) error {
	var dir = "./sql/dml/"
	for _, typ := range types {
		if err := writeBy("data."+typ, dir, params); err != nil {
			return err
		}
	}
	return mergeFile(dir, `^insert.[\S]+.sql$`, "merge-file.sql")
}

func writeProj(proj, content string) {
	filename := "./sql/ddl/" + proj + ".sql"
	dir := files.Dir(filename)
	if !files.IsExist(dir) {
		files.MkdirAll(dir, 0755)
	}
	files.Remove(filename)
	f, ferr := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0755)
	defer f.Close()
	if ferr != nil {
		logger.Error(ferr)
		return
	}
	f.WriteString(content)
}

func typ2tmpl(typ string) string {
	var tmpl string
	switch typ {
	case "data.area":
		tmpl = tmplDataArea
	case "data.org":
		tmpl = tmplDataOrg
	case "data.dict":
		tmpl = tmplDataDict
	case "data.post":
		tmpl = tmplDataPost
	case "data.post.menu":
		tmpl = tmplDataPostMenu
	case "data.role":
		tmpl = tmplDataRole
	case "data.role.post":
		tmpl = tmplDataRolePost
	case "data.user":
		tmpl = tmplDataUser
	case "data.user.role":
		tmpl = tmplDataUserRole
	}
	return tmpl
}

func typ2file(typ, dir string) string {
	var dstfile string
	switch typ {
	case "data.area":
		dstfile = "insert.sys_area.sql"
	case "data.org":
		dstfile = "insert.sys_org.sql"
	case "data.dict":
		dstfile = "insert.sys_dict.sql"
	case "data.post":
		dstfile = "insert.sys_post.sql"
	case "data.post.menu":
		dstfile = "insert.sys_post_menu.sql"
	case "data.role":
		dstfile = "insert.sys_role.sql"
	case "data.role.post":
		dstfile = "insert.sys_role_post.sql"
	case "data.user":
		dstfile = "insert.sys_user.sql"
	case "data.user.role":
		dstfile = "insert.sys_user_role.sql"
	}
	if !files.IsExist(dir) {
		files.MkdirAll(dir, 0755)
	}
	return filepath.Join(dir, dstfile)
}

func writeBy(typ, dir string, params map[string]string) error {
	tmpl := typ2tmpl(typ)
	dstfile := typ2file(typ, dir)
	return writeTmpl(tmpl, dstfile, params)
}
