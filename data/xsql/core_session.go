// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xsql

import (
	"database/sql"
	"time"

	"gopkg.in/goyy/goyy.v0/data/dialect"
	"gopkg.in/goyy/goyy.v0/data/dml"
	"gopkg.in/goyy/goyy.v0/data/domain"
	"gopkg.in/goyy/goyy.v0/data/dql"
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/util/sqls"
	"gopkg.in/goyy/goyy.v0/util/uuids"
)

type session struct {
	db      *sql.DB
	dialect dialect.Interface
	dml     dml.Interface
	dql     dql.Interface
}

// New Query
func (me *session) Query(dql string, args ...interface{}) Query {
	return &query{
		db:      me.db,
		session: me,
		dql:     sqls.FormatSpace(dql),
		args:    args,
	}
}

// New NamedQuery
func (me *session) NamedQuery(dql string, args map[string]interface{}) (Query, error) {
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
		db:        me.db,
		session:   me,
		dql:       d,
		args:      a,
		isNamed:   true,
		namedDql:  fdql,
		namedArgs: args,
	}, nil
}

// Select one SQL
func (me *session) Get(out entity.Interface) error {
	dql, arg := me.dql.SelectOne(out)
	err := me.Query(dql, arg).Row(out)
	if err != nil {
		logger.Debug(err.Error())
		return err
	}
	return nil
}

// Select one SQL
func (me *session) SelectOne(out entity.Interface, sifts ...domain.Sift) error {
	dql, args, err := me.dql.SelectListBySift(out, sifts...)
	if err != nil {
		logger.Debug(err.Error())
		return err
	}
	err = me.Query(dql, args...).Row(out)
	if err != nil {
		logger.Debug(err.Error())
		return err
	}
	return nil
}

// Select list SQL
func (me *session) SelectList(out entity.Interfaces, sifts ...domain.Sift) error {
	dql, args, err := me.dql.SelectListBySift(out.New(), sifts...)
	if err != nil {
		logger.Debug(err.Error())
		return err
	}
	err = me.Query(dql, args...).Rows(out)
	if err != nil {
		logger.Debug(err.Error())
		return err
	}
	return nil
}

// Select page SQL
func (me *session) SelectPage(content entity.Interfaces, pageable domain.Pageable, sifts ...domain.Sift) (domain.Page, error) {
	dql, args, err := me.dql.SelectListBySift(content.New(), sifts...)
	if err != nil {
		logger.Debug(err.Error())
		return nil, err
	}
	page, err := me.Query(dql, args...).Page(content, pageable)
	if err != nil {
		logger.Debug(err.Error())
		return nil, err
	}
	return page, nil
}

// Select count(*) SQL
func (me *session) SelectCount(e entity.Interface, sifts ...domain.Sift) (int, error) {
	dql, args, err := me.dql.SelectCountBySift(e, sifts...)
	if err != nil {
		logger.Debug(err.Error())
		return 0, err
	}
	return me.Query(dql, args...).Int()
}

// Insert SQL
func (me *session) Insert(e entity.Interface) (int64, error) {
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
func (me *session) Update(e entity.Interface) (int64, error) {
	dml, args := me.dml.Update(e)
	res, err := me.Exec(dml, args...)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

// Delete SQL
func (me *session) Delete(e entity.Interface) (int64, error) {
	dml, arg := me.dml.Delete(e)
	res, err := me.Exec(dml, arg)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

// Logical Delete SQL
func (me *session) Disable(e entity.Interface) (int64, error) {
	dml, arg := me.dml.Disable(e)
	res, err := me.Exec(dml, arg)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (me *session) Exec(dml string, args ...interface{}) (sql.Result, error) {
	sql := sqls.FormatSpace(dml)
	if isDebug() {
		now := time.Now()
		defer debugLog(now, sql, args...)
	}
	stmt, err := me.db.Prepare(sql)
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
func (me *session) Begin() (Tx, error) {
	return me.db.Begin()
}

// Close Session
func (me *session) Close() error {
	return nil
}

// Get Database Type
func (me *session) DBType() string {
	return me.dialect.Type()
}
