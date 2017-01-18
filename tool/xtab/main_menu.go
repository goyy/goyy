// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"io/ioutil"
	"strconv"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/goyy/goyy.v0/data/dialect"
	"gopkg.in/goyy/goyy.v0/data/xsql"
	"gopkg.in/goyy/goyy.v0/util/files"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"gopkg.in/goyy/goyy.v0/util/uuids"
)

func genMenu() {
	for _, p := range conf.projects {
		clidir := "../" + strings.AfterLast(p.Clipath(), "/")
		dir := clidir + "/templates/core/include/"
		dstfile := dir + "header.html"
		if !files.IsExist(dstfile) {
			files.MkdirAll(dir, 0644)
		}
		data := map[string]interface{}{
			"Project": p,
			"Modules": conf.modules,
			"Tables":  conf.tables,
		}
		buf := bytes.Buffer{}
		tmpl := newTmpl(tmplMenu)
		tmpl.Execute(&buf, data)
		ioutil.WriteFile(dstfile, buf.Bytes(), 0644)
	}
	for _, p := range conf.projects {
		for mi, m := range conf.modules {
			if p.ID() == m.project.ID() && m.Menu() == "true" {
				root := &menu{
					id:      "root",
					code:    "00",
					name:    i18N.Message("tmpl.menu.data.root"),
					ordinal: "00",
				}
				po := strings.PadLeft(strconv.Itoa(mi+1), 2, "0")
				pp := addMenu(p.ID(), m.ID(), "", m.Name(), po, "10", root) // module
				for ti, t := range conf.tables {
					if m.ID() == t.module.ID() && t.Menu() == "true" && t.ID() != "-" {
						to := strings.PadLeft(strconv.Itoa(ti+1), 2, "0")
						mp := addMenu(p.ID(), m.ID(), t.ID(), t.Name(), to, "20", pp) // table
						addMenu(p.ID(), m.ID(), t.ID(), "view", "10", "30", mp)       // button:view
						addMenu(p.ID(), m.ID(), t.ID(), "add", "20", "30", mp)        // button:add
						addMenu(p.ID(), m.ID(), t.ID(), "edit", "30", "30", mp)       // button:edit
						addMenu(p.ID(), m.ID(), t.ID(), "disable", "40", "30", mp)    // button:disable
						addMenu(p.ID(), m.ID(), t.ID(), "export", "50", "30", mp)     // button:export
					}
				}
			}
		}
	}
}

func addMenu(pid, mid, tid, xname, ordinal, genre string, parent *menu) *menu {
	db := getDB(pid)
	id := uuids.New()
	m := &menu{
		id:       id,
		hidden:   "0",
		code:     parent.ordinal + ordinal,
		name:     xname,
		parentID: parent.id,
		leaf:     "0",
		genre:    genre,
		ordinal:  parent.ordinal + ordinal,
	}
	if parent.id == "root" {
		m.parentIDs = parent.id
		m.parentCodes = parent.code
	} else {
		m.parentIDs = parent.parentIDs + "," + parent.id
		m.parentCodes = parent.parentCodes + "," + parent.code
	}
	switch genre {
	case "10":
		if parent.id == "root" {
			m.fullname = xname
			m.parentNames = parent.name
		} else {
			m.fullname = parent.fullname + " - " + xname
			m.parentNames = parent.parentNames + "," + parent.name
		}
		m.grade = "2"
	case "20":
		m.href = "/" + mid + "/" + tid
		m.name = xname + i18N.Message("tmpl.menu.manage")
		if parent.id == "root" {
			m.fullname = m.name
			m.parentNames = parent.name
		} else {
			m.fullname = parent.fullname + " - " + m.name
			m.parentNames = parent.parentNames + "," + parent.name
		}
		m.grade = "3"
	case "30":
		m.hidden = "1"
		m.permission = mid + ":" + tid + ":" + xname
		m.name = i18N.Message("tmpl.menu.btn." + xname)
		if parent.id == "root" {
			m.fullname = m.name
			m.parentNames = parent.name
		} else {
			m.fullname = parent.fullname + " - " + m.name
			m.parentNames = parent.parentNames + "," + parent.name
		}
		m.leaf = "1"
		m.grade = "4"
	}

	csql := "SELECT count(1) FROM sys_menu WHERE code = ?"

	sql := `INSERT INTO sys_menu
	(id, href, target, icon, hidden, permission, code, name, fullname, genre, ordinal, parent_id, parent_ids, parent_codes, parent_names, leaf, grade, memo, creates, creater, created, modifier, modified, version, deletion, artifical, history)
	VALUES
	(?, ?, NULL, NULL, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NULL, NULL, NULL, 1447816484, NULL, 1447816484, 0, 0, 0, 0)`

	count, err := db.Query(csql, m.code).Int()
	if err != nil {
		logger.Error(err)
		return m
	}
	if count == 0 {
		_, err = db.Exec(sql, m.id, m.href, m.hidden, m.permission, m.code, m.name, m.fullname, m.genre, m.ordinal, m.parentID, m.parentIDs, m.parentCodes, m.parentNames, m.leaf, m.grade)
		if err != nil {
			logger.Error(err)
			return m
		}
	}
	return m
}

type menu struct {
	id, href, hidden, permission, code, name, fullname, genre, ordinal, parentID, parentIDs, parentCodes, parentNames, leaf, grade string
}

var once sync.Once
var xdb xsql.DB

func getDB(pid string) xsql.DB {
	once.Do(func() {
		v, err := xsql.New(&dialect.MySQL{}, pid)
		if err != nil {
			panic(err)
		} else {
			xdb = v
		}
	})
	return xdb
}
