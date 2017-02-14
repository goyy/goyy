// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"gopkg.in/goyy/goyy.v0/comm/log"
	"gopkg.in/goyy/goyy.v0/data/xsql"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

var dmlMap = map[string]string{}

func expDB() {
	isRunMenu := true
	for _, p := range conf.projects {
		dbName := p.database.name
		driverName := p.database.driverName
		proj := p.ID()
		params := map[string]string{
			"driverName": driverName,
		}
		if isExistTable(driverName, proj) {
			isRunMenu = false
			logger.Errorln(i18N.Message("valid.table.exist"))
			continue
		}
		content := getDDL(driverName, proj)
		if err := runSQL(driverName, proj, content); err != nil {
			logger.Error(err)
			return
		}
		if isRunMenu {
			genMenu()
			isRunMenu = false
		}
		if _, ok := dmlMap[dbName]; !ok {
			err := expData2DB(driverName, proj, params)
			if err != nil {
				logger.Errorln(err)
				return
			}
		}
		dmlMap[dbName] = proj
	}
}

func expData2DB(driverName, proj string, params map[string]string) error {
	for _, typ := range types {
		tmpl := typ2tmpl("data." + typ)
		content, err := parseTmpl(tmpl, params)
		if err != nil {
			return err
		}
		if err := runSQL(driverName, proj, content); err != nil {
			return err
		}
	}
	return nil
}

func runSQL(driverName, proj, stmt string) error {
	db := getDB(driverName, proj)
	stmts := strings.Split(stmt, conf.Settings.Statement.Seperator)
	for _, s := range stmts {
		if strings.IsBlank(s) {
			continue
		}
		_, err := db.Exec(s)
		if err != nil {
			logger.Error(err, s)
			return err
		}
	}
	return nil
}

func isExistTable(driverName, proj string) bool {
	tablename := "sys_user"
	if strings.ToLower(conf.Settings.Statement.Case) == "upper" {
		tablename = strings.ToUpper(tablename)
	}
	sql := "select count(1) from " + tablename
	db := getDB(driverName, proj)
	xsql.SetPriority(log.Pcritical)
	_, err := db.Exec(sql)
	xsql.SetPriority(log.Perror)
	if err != nil {
		msg := err.Error()
		if strings.Contains(msg, "Error 1146") {
			return false
		}
		if strings.Contains(msg, "ORA-00942") {
			return false
		}
		logger.Errorln(err)
	}
	return true
}
