// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xsql

import (
	"database/sql"
	"gopkg.in/goyy/goyy.v0/comm/env"
	"gopkg.in/goyy/goyy.v0/data/dialect"
	"gopkg.in/goyy/goyy.v0/data/dml"
	"gopkg.in/goyy/goyy.v0/data/dql"
)

// New Factory
func New(dialect dialect.Interface, name string) (f Factory, err error) {
	dbconf, err := env.Database(name)
	if err != nil {
		return
	}
	db, err := sql.Open(dbconf.DriverName, dbconf.DataSourceName)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(dbconf.MaxIdleConns)
	db.SetMaxOpenConns(dbconf.MaxOpenConns)
	f = &factory{
		driverName:     dbconf.DriverName,
		dataSourceName: dbconf.DataSourceName,
		maxIdleConns:   dbconf.MaxIdleConns,
		maxOpenConns:   dbconf.MaxOpenConns,
		dialect:        dialect,
		db:             db,
	}
	return
}

type factory struct {
	driverName, dataSourceName string
	maxIdleConns, maxOpenConns int
	dialect                    dialect.Interface
	db                         *sql.DB
}

// New Session
func (me *factory) Session() (Session, error) {
	return &session{
		db:      me.db,
		dialect: me.dialect,
		dml:     dml.New(me.dialect),
		dql:     dql.New(me.dialect),
	}, nil
}
