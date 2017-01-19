// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xsql

import (
	"database/sql"
	"time"

	"gopkg.in/goyy/goyy.v0/comm/env"
	"gopkg.in/goyy/goyy.v0/data/dialect"
	"gopkg.in/goyy/goyy.v0/data/dml"
	"gopkg.in/goyy/goyy.v0/data/domain"
	"gopkg.in/goyy/goyy.v0/data/dql"
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/util/sqls"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"gopkg.in/goyy/goyy.v0/util/uuids"
)

// New DB
func New(dialect dialect.Interface, name string) (DB, error) {
	r := &db{
		name:    name,
		dialect: dialect,
	}
	if dialect != nil {
		r.dml = dml.New(dialect)
		r.dql = dql.New(dialect)
	}
	err := r.init()
	return r, err
}

// NewDB new DB.
func NewDB(name string) (DB, error) {
	r := &db{
		name: name,
	}
	err := r.init()
	return r, err
}

type db struct {
	name           string
	driverName     string
	dataSourceName string
	maxIdleConns   int
	maxOpenConns   int
	dialect        dialect.Interface
	sqlDB          *sql.DB
	dml            dml.Interface
	dql            dql.Interface
}

// init
func (me *db) init() error {
	dbconf, err := env.ParseDatabase(me.name)
	if err != nil {
		return err
	}
	sqlDB, err := sql.Open(dbconf.DriverName, dbconf.DataSourceName)
	if err != nil {
		return err
	}
	if me.dialect == nil {
		switch strings.ToLower(dbconf.DriverName) {
		case "oci8", "oracle":
			me.dialect = &dialect.Oracle{}
		case "postgres":
			me.dialect = &dialect.PostgreSQL{}
		case "mssql", "adodb":
			me.dialect = &dialect.SQLServer{}
		case "sqlite3":
			me.dialect = &dialect.Sqlite{}
		default:
			me.dialect = &dialect.MySQL{}
		}
	}
	if me.dml == nil {
		me.dml = dml.New(me.dialect)
	}
	if me.dql == nil {
		me.dql = dql.New(me.dialect)
	}
	sqlDB.SetMaxIdleConns(dbconf.MaxIdleConns)
	sqlDB.SetMaxOpenConns(dbconf.MaxOpenConns)
	me.sqlDB = sqlDB
	me.driverName = dbconf.DriverName
	me.dataSourceName = dbconf.DataSourceName
	me.maxIdleConns = dbconf.MaxIdleConns
	me.maxOpenConns = dbconf.MaxOpenConns
	return nil
}

// New Query
func (me *db) Query(dql string, args ...interface{}) Query {
	return &query{
		db:   me,
		dql:  sqls.FormatSpace(dql),
		args: args,
	}
}

// New NamedQuery
func (me *db) NamedQuery(dql string, args map[string]interface{}) (Query, error) {
	fdql := sqls.FormatSpace(dql)
	sql, err := sqls.ParseTemplateSql(fdql, args)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	d, a, err := sqls.ParseNamedSql(me.dialect, sql, args)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	return &query{
		db:        me,
		dql:       d,
		args:      a,
		isNamed:   true,
		namedDql:  fdql,
		namedArgs: args,
	}, nil
}

// New Sifter
func (me *db) Sifter(sifts ...domain.Sift) Sifter {
	return &sifter{
		db:    me,
		sifts: sifts,
	}
}

// Select one SQL
func (me *db) Get(out entity.Interface) error {
	dql, arg := me.dql.SelectOne(out)
	err := me.Query(dql, arg).Row(out)
	if err != nil {
		logger.Debug(err.Error())
		return err
	}
	return nil
}

// Insert SQL
func (me *db) Insert(e entity.Interface) (int64, error) {
	pk := e.Table().Primary().Name()
	if t, ok := e.Column(pk); ok {
		e.SetString(t.Name(), uuids.New())
	}
	dml, args := me.dml.Insert(e)
	res, err := me.Exec(dml, args...)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

// Update SQL
func (me *db) Update(e entity.Interface) (int64, error) {
	dml, args := me.dml.Update(e)
	res, err := me.Exec(dml, args...)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

// Delete SQL
func (me *db) Delete(e entity.Interface) (int64, error) {
	dml, arg := me.dml.Delete(e)
	res, err := me.Exec(dml, arg)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

// Logical Delete SQL
func (me *db) Disable(e entity.Interface) (int64, error) {
	dml, arg := me.dml.Disable(e)
	res, err := me.Exec(dml, arg)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (me *db) Exec(dml string, args ...interface{}) (sql.Result, error) {
	sql := sqls.FormatSpace(dml)
	if isDebug() {
		now := time.Now()
		defer debugLog(now, sql, args...)
	}
	stmt, err := me.sqlDB.Prepare(sql)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	defer stmt.Close()
	s, err := stmt.Exec(args...)
	if err != nil {
		logger.Error(err.Error())
	}
	return s, err
}

// Begin Transaction
func (me *db) Begin() (Tx, error) {
	return me.sqlDB.Begin()
}

// Close closes the database, releasing any open resources.
//
// It is rare to Close a DB, as the DB handle is meant to be
// long-lived and shared between many goroutines.
func (me *db) Close() error {
	return me.sqlDB.Close()
}

// Get Dialect
func (me *db) Dialect() dialect.Interface {
	return me.dialect
}

// Ping verifies a connection to the database is still alive,
// establishing a connection if necessary.
func (me *db) Ping() error {
	err := me.sqlDB.Ping()
	if err != nil {
		return err
	}
	if me.dialect.Type() == dialect.ORACLE {
		_, err = me.Query("select 1 from dual").Str()
		if err != nil {
			if strings.HasPrefix(err.Error(), "ORA-03114") {
				if me.sqlDB != nil {
					me.sqlDB.Close()
				}
				me.init()
			}
		}
	}
	return nil
}
