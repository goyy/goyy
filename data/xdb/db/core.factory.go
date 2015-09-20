// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package db

import (
	"database/sql"
	"gopkg.in/goyy/goyy.v0/comm/env"
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// New Factory
func New(name string) (f Factory, err error) {
	db, err := env.Database(name)
	if err != nil {
		return
	}
	f = &factory{driverName: db.DriverName, dataSourceName: db.DataSourceName}
	return
}

type factory struct {
	driverName, dataSourceName string
}

// New Session
func (me *factory) Session() (Session, error) {
	db, err := sql.Open(me.driverName, me.dataSourceName)
	if err != nil {
		return nil, err
	}
	switch me.driverName {
	case "mysql", "mymysql":
		return &session{Db: db, Dialect: NewMysql()}, nil
	case "postgres":
		return &session{Db: db, Dialect: NewPostgres()}, nil
	}
	return &session{Db: db, Dialect: NewMysql()}, nil
}
