// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xsql

import (
	"database/sql"
	"time"

	"gopkg.in/goyy/goyy.v0/data/domain"
	"gopkg.in/goyy/goyy.v0/data/dql"
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/util/sqls"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

// add space string:{before is add:yes, after is add:no}
type query struct {
	db        *sql.DB
	session   *session
	dql       string
	args      []interface{}
	isNamed   bool
	namedDql  string
	namedArgs map[string]interface{}
}

// Retrieve a list of mapped entities from the dql and args.
func (me *query) Rows(out entity.Interfaces) error {
	if isDebug() {
		now := time.Now()
		defer debugLog(now, me.dql, me.args...)
	}
	stmt, err := me.db.Prepare(me.dql)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	defer stmt.Close()
	rows, err := stmt.Query(me.args...)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	defer rows.Close()
	cols, err := rows.Columns()
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	for rows.Next() {
		ei := out.New()
		containers := make([]interface{}, len(cols))
		for i := 0; i < cap(containers); i++ {
			col := strings.ToLower(cols[i])
			if eit, ok := ei.Type(col); !ok || eit == nil || eit.Field() == nil {
				var v interface{}
				containers[i] = &v
			} else {
				if value := ei.GetPtr(col); value == nil {
					var v interface{}
					containers[i] = &v
				} else {
					containers[i] = value
				}
			}
		}
		err = rows.Scan(containers...)
		if err != nil {
			logger.Error(err.Error())
			return err
		}
		out.Append(ei)
	}
	return nil
}

// Retrieve a single row mapped from the dql and args.
func (me *query) Row(out entity.Interface) error {
	if isDebug() {
		now := time.Now()
		defer debugLog(now, me.dql, me.args...)
	}
	stmt, err := me.db.Prepare(me.dql)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	defer stmt.Close()
	rows, err := stmt.Query(me.args...)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	defer rows.Close()
	cols, err := rows.Columns()
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	for rows.Next() {
		containers := make([]interface{}, len(cols))
		for i := 0; i < cap(containers); i++ {
			col := strings.ToLower(cols[i])
			if outt, ok := out.Type(col); !ok || outt == nil || outt.Field() == nil {
				var v interface{}
				containers[i] = &v
			} else {
				if value := out.GetPtr(col); value == nil {
					var v interface{}
					containers[i] = &v
				} else {
					containers[i] = value
				}
			}
		}
		err = rows.Scan(containers...)
		if err != nil {
			logger.Error(err.Error())
		}
		return err
	}
	return nil
}

// Retrieve a int mapped from the dql and args.
func (me *query) Int() (int, error) {
	var out int
	if err := me.val(&out); err != nil {
		return 0, err
	}
	return out, nil
}

// Retrieve a float mapped from the dql and args.
func (me *query) Float() (float64, error) {
	var out float64
	if err := me.val(&out); err != nil {
		return 0, err
	}
	return out, nil
}

// Retrieve a string mapped from the dql and args.
func (me *query) Str() (string, error) {
	var out string
	if err := me.val(&out); err != nil {
		return "", err
	}
	return out, nil
}

// Retrieve a time.Time mapped from the dql and args.
func (me *query) Time() (time.Time, error) {
	var out time.Time
	if err := me.val(&out); err != nil {
		return time.Now(), err
	}
	return out, nil
}
func (me *query) val(out interface{}) error {
	if isDebug() {
		now := time.Now()
		defer debugLog(now, me.dql, me.args...)
	}
	stmt, err := me.db.Prepare(me.dql)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	defer stmt.Close()
	rows, err := stmt.Query(me.args...)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	defer rows.Close()
	if !rows.Next() {
		return sql.ErrNoRows
	}
	if err = rows.Scan(out); err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}

// Retrieve a single row mapped from the dql and args.
func (me *query) Page(content entity.Interfaces, pageable domain.Pageable) (domain.Page, error) {
	dqlCount := sqls.ParseCountSql(me.dql)
	queryCount := me.session.Query(dqlCount, me.args...)
	totalElements, err := queryCount.Int()
	if err != nil {
		return nil, err
	}
	d := dql.New(me.session.dialect)
	dqlPage := d.SelectPage(me.dql, pageable)
	queryPage := me.session.Query(dqlPage, me.args...)
	err = queryPage.Rows(content)
	if err != nil {
		return nil, err
	}
	return domain.NewPage(pageable, content, totalElements), nil
}

// Query To String
func (me *query) String() string {
	return me.dql
}
