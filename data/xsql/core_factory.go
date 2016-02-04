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
	"gopkg.in/goyy/goyy.v0/util/strings"
)

// New Factory
func New(dialect dialect.Interface, name string) (Factory, error) {
	f := &factory{
		name:    name,
		dialect: dialect,
	}
	err := f.init()
	return f, err
}

type factory struct {
	name                       string
	driverName, dataSourceName string
	maxIdleConns, maxOpenConns int
	dialect                    dialect.Interface
	db                         *sql.DB
}

// Get Dialect
func (me *factory) Dialect() dialect.Interface {
	return me.dialect
}

// init
func (me *factory) init() error {
	dbconf, err := env.Database(me.name)
	if err != nil {
		return err
	}
	db, err := sql.Open(dbconf.DriverName, dbconf.DataSourceName)
	if err != nil {
		return err
	}
	db.SetMaxIdleConns(dbconf.MaxIdleConns)
	db.SetMaxOpenConns(dbconf.MaxOpenConns)
	me.db = db
	me.driverName = dbconf.DriverName
	me.dataSourceName = dbconf.DataSourceName
	me.maxIdleConns = dbconf.MaxIdleConns
	me.maxOpenConns = dbconf.MaxOpenConns
	return nil
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

// Ping verifies a connection to the database is still alive,
// establishing a connection if necessary.
func (me *factory) Ping() error {
	err := me.db.Ping()
	if err != nil {
		return err
	}
	if me.dialect.Type() == dialect.ORACLE {
		s, err := me.Session()
		if err != nil {
			return err
		}
		_, err = s.Query("select 1 from dual").Str()
		s.Close()
		if err != nil {
			if strings.HasPrefix(err.Error(), "ORA-03114") {
				if me.db != nil {
					me.db.Close()
				}
				me.init()
			}
		}
	}
	return nil
}
