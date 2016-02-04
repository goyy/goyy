// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package repository

import (
	"database/sql"
	"os"

	"gopkg.in/goyy/goyy.v0/data/dialect"
	"gopkg.in/goyy/goyy.v0/data/domain"
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/xsql"
	"gopkg.in/goyy/goyy.v0/util/sqls"
)

type Repository struct {
	factory xsql.Factory
}

// New returns Repository.
func New(d dialect.Interface, name string) *Repository {
	f, err := xsql.New(d, name)
	if err != nil {
		logger.Error("env.DataSource failed", err)
		os.Exit(3)
	}
	return &Repository{factory: f}
}

// Get Dialect
func (me *Repository) Dialect() dialect.Interface {
	return me.factory.Dialect()
}

// Insert entity data into database.
func (me *Repository) Insert(e entity.Interface) (int64, error) {
	session, err := me.factory.Session()
	if err != nil {
		logger.Debug(err)
		return 0, err
	}
	defer session.Close()
	return session.Insert(e)
}

// Update entity data into database.
// Returns the number of records to be updated.
func (me *Repository) Update(e entity.Interface) (int64, error) {
	session, err := me.factory.Session()
	if err != nil {
		logger.Debug(err)
		return 0, err
	}
	defer session.Close()
	return session.Update(e)
}

// Delete entity data into database.
// Returns the number of records to be updated.
func (me *Repository) Delete(e entity.Interface) (int64, error) {
	session, err := me.factory.Session()
	if err != nil {
		logger.Debug(err)
		return 0, err
	}
	defer session.Close()
	return session.Delete(e)
}

// Logical delete entity data into database.
// Returns the number of records to be updated.
func (me *Repository) Disable(e entity.Interface) (int64, error) {
	session, err := me.factory.Session()
	if err != nil {
		logger.Debug(err)
		return 0, err
	}
	defer session.Close()
	return session.Disable(e)
}

// Get entity data by the primary key.
func (me *Repository) Get(out entity.Interface) error {
	session, err := me.factory.Session()
	if err != nil {
		logger.Debug(err)
		return err
	}
	defer session.Close()
	return session.Get(out)
}

// Query a single record.
func (me *Repository) SelectOne(out entity.Interface, dql string, args ...interface{}) error {
	session, err := me.factory.Session()
	if err != nil {
		logger.Debug(err)
		return err
	}
	defer session.Close()
	query := session.Query(dql, args...)
	err = query.Row(out)
	return err
}

// Query multiple records.
func (me *Repository) SelectList(out entity.Interfaces, dql string, args ...interface{}) error {
	session, err := me.factory.Session()
	if err != nil {
		logger.Debug(err)
		return err
	}
	defer session.Close()
	query := session.Query(dql, args...)
	err = query.Rows(out)
	return err
}

// Paging query record.
func (me *Repository) SelectPage(content entity.Interfaces, pageable domain.Pageable, dql string, args ...interface{}) (domain.Page, error) {
	session, err := me.factory.Session()
	if err != nil {
		logger.Debug(err)
		return nil, err
	}
	defer session.Close()
	query := session.Query(dql, args...)
	return query.Page(content, pageable)
}

// Query a single record.
// Return value of int type.
func (me *Repository) SelectInt(dql string, args ...interface{}) (int, error) {
	session, err := me.factory.Session()
	if err != nil {
		logger.Debug(err)
		return 0, err
	}
	defer session.Close()
	query := session.Query(dql, args...)
	return query.Int()
}

// Query a single record.
// Return value of float64 type.
func (me *Repository) SelectFloat(dql string, args ...interface{}) (float64, error) {
	session, err := me.factory.Session()
	if err != nil {
		logger.Debug(err)
		return 0, err
	}
	defer session.Close()
	query := session.Query(dql, args...)
	return query.Float()
}

// Query a single record.
// Return value of string type.
func (me *Repository) SelectStr(dql string, args ...interface{}) (string, error) {
	session, err := me.factory.Session()
	if err != nil {
		logger.Debug(err)
		return "", err
	}
	defer session.Close()
	query := session.Query(dql, args...)
	return query.Str()
}

// Query a single record by named query.
func (me *Repository) SelectOneByNamed(out entity.Interface, dql string, args map[string]interface{}) error {
	session, err := me.factory.Session()
	if err != nil {
		logger.Debug(err)
		return err
	}
	defer session.Close()
	query, err := session.NamedQuery(dql, args)
	if err != nil {
		logger.Debug(err)
		return err
	}
	err = query.Row(out)
	return err
}

// Query multiple records by named query.
func (me *Repository) SelectListByNamed(out entity.Interfaces, dql string, args map[string]interface{}) error {
	session, err := me.factory.Session()
	if err != nil {
		logger.Debug(err)
		return err
	}
	defer session.Close()
	query, err := session.NamedQuery(dql, args)
	if err != nil {
		logger.Debug(err)
		return err
	}
	err = query.Rows(out)
	return err
}

// Paging query record by named query.
func (me *Repository) SelectPageByNamed(content entity.Interfaces, pageable domain.Pageable, dql string, args map[string]interface{}) (domain.Page, error) {
	session, err := me.factory.Session()
	if err != nil {
		logger.Debug(err)
		return nil, err
	}
	defer session.Close()
	query, err := session.NamedQuery(dql, args)
	if err != nil {
		logger.Debug(err)
		return nil, err
	}
	return query.Page(content, pageable)
}

// select count(*) from...
// Return value of int type.
func (me *Repository) SelectCountByNamed(dql string, args map[string]interface{}) (int, error) {
	countSql := sqls.ParseCountSql(dql)
	return me.SelectIntByNamed(countSql, args)
}

// Query a single record by named query.
// Return value of int type.
func (me *Repository) SelectIntByNamed(dql string, args map[string]interface{}) (int, error) {
	session, err := me.factory.Session()
	if err != nil {
		logger.Debug(err)
		return 0, err
	}
	defer session.Close()
	query, err := session.NamedQuery(dql, args)
	if err != nil {
		logger.Debug(err)
		return 0, err
	}
	return query.Int()
}

// Query a single record by sifts.
func (me *Repository) SelectOneBySift(out entity.Interface, sifts ...domain.Sift) error {
	session, err := me.factory.Session()
	if err != nil {
		logger.Debug(err)
		return err
	}
	defer session.Close()
	return session.SelectOne(out, sifts...)
}

// Query multiple records by sifts.
func (me *Repository) SelectListBySift(out entity.Interfaces, sifts ...domain.Sift) error {
	session, err := me.factory.Session()
	if err != nil {
		logger.Debug(err)
		return err
	}
	defer session.Close()
	return session.SelectList(out, sifts...)
}

// Paging query records by sifts.
func (me *Repository) SelectPageBySift(content entity.Interfaces, pageable domain.Pageable, sifts ...domain.Sift) (domain.Page, error) {
	session, err := me.factory.Session()
	if err != nil {
		logger.Debug(err)
		return nil, err
	}
	defer session.Close()
	return session.SelectPage(content, pageable, sifts...)
}

// select count(*) from...
// Return value of int type.
func (me *Repository) SelectCountBySift(e entity.Interface, sifts ...domain.Sift) (int, error) {
	session, err := me.factory.Session()
	if err != nil {
		logger.Debug(err)
		return 0, err
	}
	defer session.Close()
	return session.SelectCount(e, sifts...)
}

// Exec executes a prepared statement with the given arguments and
// returns a Result summarizing the effect of the statement.
func (me *Repository) Exec(query string, args ...interface{}) (sql.Result, error) {
	session, err := me.factory.Session()
	if err != nil {
		logger.Debug(err)
		return nil, err
	}
	defer session.Close()
	return session.Exec(query, args...)
}

// Begin starts a transaction. The isolation level is dependent on
// the driver.
func (me *Repository) Begin() (xsql.Tx, error) {
	sesson, err := me.factory.Session()
	if err != nil {
		return nil, err
	}
	tx, err := sesson.Begin()
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// Ping verifies a connection to the database is still alive,
// establishing a connection if necessary.
func (me *Repository) Ping() error {
	return me.factory.Ping()
}
