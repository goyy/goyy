// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"io/ioutil"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/goyy/goyy.v0/comm/log"
	"gopkg.in/goyy/goyy.v0/data/xsql"
	"gopkg.in/goyy/goyy.v0/util/files"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"gopkg.in/goyy/goyy.v0/util/times"
	"gopkg.in/goyy/goyy.v0/util/uuids"
)

func genMenu() {
	xsql.SetPriority(log.Perror)
	// generate the header.html
	for _, p := range conf.projects {
		clidir := "../" + strings.AfterLast(p.Admpath(), "/")
		dir := clidir + "/templates/core/include/"
		dstfile := dir + "header.html"
		if !files.IsExist(dstfile) {
			files.MkdirAll(dir, 0755)
		}
		data := map[string]interface{}{
			"Project": p,
			"Modules": conf.modules,
			"Tables":  conf.tables,
		}
		buf := bytes.Buffer{}
		tmpl := newTmpl(tmplMenu)
		tmpl.Execute(&buf, data)
		ioutil.WriteFile(dstfile, buf.Bytes(), 0755)
	}
	// insert into sys_menu
	for _, p := range conf.projects {
		insertRootMenu(p.database.driverName, p.ID())
		for mi, m := range conf.modules {
			if p.ID() == m.project.ID() && m.Menu() == "true" {
				root := &menu{
					id:      "root",
					code:    "00",
					name:    i18N.Message("tmpl.menu.data.root"),
					ordinal: "00",
				}
				po := strings.PadLeft(strconv.Itoa(mi+1), 2, "0")
				pp := addMenu(p.database.driverName, p.ID(), m.ID(), "", m.Name(), po, "10", root) // module
				for ti, t := range conf.tables {
					if m.ID() == t.module.ID() && t.Menu() == "true" && t.ID() != "-" {
						to := strings.PadLeft(strconv.Itoa(ti+1), 2, "0")
						mp := addMenu(p.database.driverName, p.ID(), m.ID(), t.ID(), t.Name(), to, "20", pp) // table
						buttons := strings.Split(t.Buttons(), ",")
						for i, button := range buttons {
							if strings.IsBlank(button) {
								continue
							}
							addMenu(p.database.driverName, p.ID(), m.ID(), t.ID(), button, strconv.Itoa((i+2)*5), "30", mp)
						}
					}
				}
			}
		}
	}
}

func insertRootMenu(driverName, pid string) {
	db := getDB(driverName, pid)
	csql := "SELECT count(1) FROM sys_menu WHERE id = ?"

	sql := `INSERT INTO sys_menu
	(id, href, target, icon, hidden, permission, code, name, fullname, genre, ordinal, parent_id, parent_ids, parent_codes, parent_names, leaf, grade, memo, creates, creater, created, modifier, modified, version, deletion, artifical, history)
	VALUES
	('root', null, null, null, 0, null, '00', ?, null, '00', '00', null, null, null, null, 0, 1, null, null, null, ?, null, ?, 0, 0, 0, 0);`

	count, err := db.Query(csql, "root").Int()
	if err != nil {
		logger.Error(err)
		return
	}
	if count == 0 {
		name := i18N.Message("tmpl.menu.data.root")
		now := times.NowUnix()
		_, err = db.Exec(sql, name, now, now)
		if err != nil {
			logger.Error(err)
			return
		}
	}
}

func addMenu(driverName, pid, mid, tid, xname, ordinal, genre string, parent *menu) *menu {
	db := getDB(driverName, pid)
	id := uuids.New()
	now := times.NowUnix()
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
		var name string
		xconf := util.DecodeXML(xbuttons)
		for _, v := range xconf.Buttons.Button {
			if v.ID == xname {
				name = v.Name
			}
		}
		m.hidden = "1"
		m.permission = mid + ":" + tid + ":" + xname
		m.name = name
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
	(?, ?, NULL, NULL, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NULL, NULL, NULL, ?, NULL, ?, 0, 0, 0, 0)`

	count, err := db.Query(csql, m.code).Int()
	if err != nil {
		logger.Error(err)
		return m
	}
	if count == 0 {
		_, err = db.Exec(sql, m.id, m.href, m.hidden, m.permission, m.code, m.name, m.fullname, m.genre, m.ordinal, m.parentID, m.parentIDs, m.parentCodes, m.parentNames, m.leaf, m.grade, now, now)
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
