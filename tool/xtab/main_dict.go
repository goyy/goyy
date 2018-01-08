// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/goyy/goyy.v0/comm/log"
	"gopkg.in/goyy/goyy.v0/data/xsql"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"gopkg.in/goyy/goyy.v0/util/times"
	"gopkg.in/goyy/goyy.v0/util/uuids"
)

func genDict(params string) {
	xsql.SetPriority(log.Perror)
	// insert into sys_dict
	for _, p := range conf.projects {
		for _, m := range conf.modules {
			if p.ID() == m.project.ID() {
				for _, t := range conf.tables {
					for _, c := range t.columns {
						addDicts(p.database.driverName, p.ID(), t.IDs(), c.ID(), c.Comment(), c.Dict())
					}
				}
			}
		}
	}
}

func addDicts(driverName, pid, table, column, comment, dict string) {
	var genre = table + "." + column
	var mkey, mval, filters, ordinal, memo string
	if strings.IsNotBlank(dict) {
		dicts := strings.Split(dict, ",")
		for o, d := range dicts {
			ordinal = strings.PadRight(strconv.Itoa(o+1), 2, "0")
			ds := strings.Split(d, ":")
			for i, v := range ds {
				switch i {
				case 0:
					mkey = strings.TrimSpace(v)
				case 1:
					mval = strings.TrimSpace(v)
				case 2:
					memo = strings.TrimSpace(v)
				case 3:
					filters = strings.TrimSpace(v)
				}
			}
			addDict(driverName, pid, genre, comment, mkey, mval, filters, ordinal, memo)
		}
	}
}

func addDict(driverName, pid, genre, descr, mkey, mval, filters, ordinal, memo string) {
	db := getDB(driverName, pid)
	id := uuids.New()
	now := times.NowUnix()
	d := &dict{
		id:      id,
		genre:   genre,
		descr:   descr,
		mkey:    mkey,
		mval:    mval,
		filters: filters,
		ordinal: ordinal,
		memo:    memo,
	}

	csql := "SELECT count(1) FROM sys_dict WHERE genre = ? and mkey = ?"

	sql := `INSERT INTO sys_dict
	(id, genre, descr, mkey, mval, filters, ordinal, memo, creates, creater, created, modifier, modified, version, deletion, artifical, history)
	VALUES
	(?, ?, ?, ?, ?, ?, ?, ?, '', '', ?, '', ?, 0, 0, 0, 0)`

	count, err := db.Query(csql, d.genre, d.mkey).Int()
	if err != nil {
		logger.Error(err)
		return
	}
	if count == 0 {
		_, err = db.Exec(sql, d.id, d.genre, d.descr, d.mkey, d.mval, d.filters, d.ordinal, d.memo, now, now)
		if err != nil {
			logger.Error(err)
			return
		}
	} else {
		usql := "UPDATE sys_dict SET descr = ?, mval = ?, memo = ? WHERE genre = ? and mkey = ?"
		_, err := db.Exec(usql, d.descr, d.mval, d.memo, d.genre, d.mkey)
		if err != nil {
			logger.Error(err)
			return
		}
	}
}

type dict struct {
	id, genre, descr, mkey, mval, filters, ordinal, memo string
}
