// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package service

import (
	"database/sql"
	"errors"
	"os"

	"gopkg.in/goyy/goyy.v0/comm/xtype"
	"gopkg.in/goyy/goyy.v0/data/dialect"
	"gopkg.in/goyy/goyy.v0/data/domain"
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/result"
	"gopkg.in/goyy/goyy.v0/data/xsql"
	"gopkg.in/goyy/goyy.v0/util/sqls"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"gopkg.in/goyy/goyy.v0/util/times"
)

// DB xsql.DB.
var DB xsql.DB

// New returns xsql.DB.
func New(d dialect.Interface, name string) xsql.DB {
	db, err := xsql.New(d, name)
	if err != nil {
		logger.Error("env.DataSource failed", err)
		os.Exit(3)
	}
	return db
}

// NewDB returns xsql.DB.
func NewDB(name string) xsql.DB {
	db, err := xsql.NewDB(name)
	if err != nil {
		logger.Error("env.DataSource failed", err)
		os.Exit(3)
	}
	return db
}

// Manager service.Manager.
type Manager struct {
	db       xsql.DB
	Entity   func() entity.Interface
	Entities func() entity.Interfaces
	Pre      func()
}

// NewEntity returns the entity.Interface.
func (me *Manager) NewEntity() entity.Interface {
	if me.Entity != nil {
		return me.Entity()
	}
	return nil
}

// NewEntities returns the entity.Interfaces.
func (me *Manager) NewEntities() entity.Interfaces {
	if me.Entities != nil {
		return me.Entities()
	}
	return nil
}

// NewEntityResult returns the result.Entity.
func (me *Manager) NewEntityResult() *result.Entity {
	return &result.Entity{Data: me.NewEntity()}
}

// NewEntitiesResult returns the result.Entities.
func (me *Manager) NewEntitiesResult() *result.Entities {
	return &result.Entities{Data: me.NewEntities()}
}

// NewPageResult returns the result.Page.
func (me *Manager) NewPageResult() *result.Page {
	out := domain.NewPageDefault(me.NewEntities())
	return &result.Page{Data: out}
}

// Get get a record of the database through the primary key.
func (me *Manager) Get(out entity.Interface) error {
	if me.Pre != nil {
		me.Pre()
	}
	return me.DB().Get(out)
}

// SelectOne get a record of the database through the query filter conditions.
func (me *Manager) SelectOne(out entity.Interface, query string, args ...interface{}) (err error) {
	if me.Pre != nil {
		me.Pre()
	}
	return me.DB().Query(query, args...).Row(out)
}

// SelectList get multiple records of the database through the query filter conditions.
func (me *Manager) SelectList(out entity.Interfaces, query string, args ...interface{}) (err error) {
	if me.Pre != nil {
		me.Pre()
	}
	return me.DB().Query(query, args...).Rows(out)
}

// SelectPage paging query database records through query filtering conditions.
func (me *Manager) SelectPage(content entity.Interfaces, pageable domain.Pageable, dql string, args ...interface{}) (out domain.Page, err error) {
	if me.Pre != nil {
		me.Pre()
	}
	return me.DB().Query(dql, args...).Page(content, pageable)
}

// SelectCount gets the total number of records by querying the filter condition.
func (me *Manager) SelectCount(dql string, args ...interface{}) (int, error) {
	if me.Pre != nil {
		me.Pre()
	}
	countSql := sqls.ParseCountSql(dql)
	return me.SelectInt(countSql, args...)
}

// SelectInt gets the value of the int type by querying the filter condition.
func (me *Manager) SelectInt(dql string, args ...interface{}) (int, error) {
	if me.Pre != nil {
		me.Pre()
	}
	return me.DB().Query(dql, args...).Int()
}

// SelectFloat gets the value of the float64 type by querying the filter condition.
func (me *Manager) SelectFloat(dql string, args ...interface{}) (float64, error) {
	if me.Pre != nil {
		me.Pre()
	}
	return me.DB().Query(dql, args...).Float()
}

// SelectStr gets the value of the string type by querying the filter condition.
func (me *Manager) SelectStr(dql string, args ...interface{}) (string, error) {
	if me.Pre != nil {
		me.Pre()
	}
	return me.DB().Query(dql, args...).Str()
}

// SelectOneByNamed get a record of the database through the map type filter conditions.
func (me *Manager) SelectOneByNamed(out entity.Interface, dql string, args map[string]interface{}) (err error) {
	if me.Pre != nil {
		me.Pre()
	}
	query, err := me.DB().NamedQuery(dql, args)
	if err != nil {
		logger.Debug(err)
		return err
	}
	err = query.Row(out)
	return err
}

// SelectListByNamed get multiple records of the database through the map type filter conditions.
func (me *Manager) SelectListByNamed(out entity.Interfaces, dql string, args map[string]interface{}) (err error) {
	if me.Pre != nil {
		me.Pre()
	}
	query, err := me.DB().NamedQuery(dql, args)
	if err != nil {
		logger.Debug(err)
		return err
	}
	err = query.Rows(out)
	return err
}

// SelectPageByNamed paging query database records through map type filter conditions.
func (me *Manager) SelectPageByNamed(content entity.Interfaces, pageable domain.Pageable, dql string, args map[string]interface{}) (out domain.Page, err error) {
	if me.Pre != nil {
		me.Pre()
	}
	query, err := me.DB().NamedQuery(dql, args)
	if err != nil {
		logger.Debug(err)
		return nil, err
	}
	return query.Page(content, pageable)
}

// SelectCountByNamed gets the total number of records through the map type filter condition.
func (me *Manager) SelectCountByNamed(dql string, args map[string]interface{}) (int, error) {
	if me.Pre != nil {
		me.Pre()
	}
	countSql := sqls.ParseCountSql(dql)
	return me.SelectIntByNamed(countSql, args)
}

// SelectIntByNamed gets the value of the int type through the map type filter condition.
func (me *Manager) SelectIntByNamed(dql string, args map[string]interface{}) (int, error) {
	if me.Pre != nil {
		me.Pre()
	}
	query, err := me.DB().NamedQuery(dql, args)
	if err != nil {
		logger.Debug(err)
		return 0, err
	}
	return query.Int()
}

// SelectFloatByNamed gets the value of the float64 type through the map type filter condition.
func (me *Manager) SelectFloatByNamed(dql string, args map[string]interface{}) (float64, error) {
	if me.Pre != nil {
		me.Pre()
	}
	query, err := me.DB().NamedQuery(dql, args)
	if err != nil {
		logger.Debug(err)
		return 0.0, err
	}
	return query.Float()
}

// SelectStrByNamed gets the value of the string type through the map type filter condition.
func (me *Manager) SelectStrByNamed(dql string, args map[string]interface{}) (string, error) {
	if me.Pre != nil {
		me.Pre()
	}
	query, err := me.DB().NamedQuery(dql, args)
	if err != nil {
		logger.Debug(err)
		return "", err
	}
	return query.Str()
}

// SelectOneBySift get a record of the database through the domain.Sift type filter conditions.
func (me *Manager) SelectOneBySift(out entity.Interface, sifts ...domain.Sift) (err error) {
	if me.Pre != nil {
		me.Pre()
	}
	return me.DB().Sifter(sifts...).Row(out)
}

// SelectListBySift get multiple record of the database through the domain.Sift type filter conditions.
func (me *Manager) SelectListBySift(out entity.Interfaces, sifts ...domain.Sift) (err error) {
	if me.Pre != nil {
		me.Pre()
	}
	return me.DB().Sifter(sifts...).Rows(out)
}

// SelectPageBySift paging query records are filtered through the domain.Sift type.
func (me *Manager) SelectPageBySift(content entity.Interfaces, pageable domain.Pageable, sifts ...domain.Sift) (out domain.Page, err error) {
	if me.Pre != nil {
		me.Pre()
	}
	return me.DB().Sifter(sifts...).Page(content, pageable)
}

// SelectCountBySift gets the total number of records through the domain.Sift type filter condition.
func (me *Manager) SelectCountBySift(sifts ...domain.Sift) (int, error) {
	if me.Pre != nil {
		me.Pre()
	}
	return me.DB().Sifter(sifts...).Count(me.NewEntity())
}

func (me *Manager) save(p xtype.Principal, e entity.Interface) error {
	if strings.IsBlank(e.Get(e.Table().Primary().Name()).(string)) {
		if strings.IsNotBlank(p.Id) {
			e.SetString(creater, p.Id)
			e.SetString(modifier, p.Id)
		}
		e.SetString(created, times.NowUnixStr())
		e.SetString(modified, times.NowUnixStr())
		_, err := me.DB().Insert(e)
		if err != nil {
			return err
		}
	} else {
		if strings.IsNotBlank(p.Id) {
			e.SetString(modifier, p.Id)
		}
		e.SetString(modified, times.NowUnixStr())
		_, err := me.DB().Update(e)
		if err != nil {
			return err
		}
	}
	return nil
}

// Save save data, but do not commit transaction.
func (me *Manager) Save(p xtype.Principal, e entity.Interface) error {
	if me.Pre != nil {
		me.Pre()
	}
	if err := e.Validate(); err != nil {
		return err
	}
	return me.save(p, e)
}

// SaveAndTx save data and commit transactions.
func (me *Manager) SaveAndTx(p xtype.Principal, e entity.Interface) error {
	if me.Pre != nil {
		me.Pre()
	}
	if err := e.Validate(); err != nil {
		return err
	}
	tx, err := me.DB().Begin()
	if err != nil {
		return err
	}
	err = me.save(p, e)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (me *Manager) disable(p xtype.Principal, e entity.Interface) (int64, error) {
	if strings.IsBlank(e.Get(e.Table().Primary().Name()).(string)) {
		return 0, errors.New("Gets the primary key value failed")
	}
	if strings.IsNotBlank(p.Id) {
		e.SetString(modifier, p.Id)
		e.SetString(modified, times.NowUnixStr())
	}
	return me.DB().Disable(e)
}

// Disable delete data, but do not commit transaction.
func (me *Manager) Disable(p xtype.Principal, e entity.Interface) (int64, error) {
	if me.Pre != nil {
		me.Pre()
	}
	return me.disable(p, e)
}

// DisableAndTx delete data and commit transactions.
func (me *Manager) DisableAndTx(p xtype.Principal, e entity.Interface) (int64, error) {
	if me.Pre != nil {
		me.Pre()
	}
	tx, err := me.DB().Begin()
	if err != nil {
		return 0, err
	}
	r, err := me.disable(p, e)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return r, nil
}

func (me *Manager) exec(dml string, args ...interface{}) (sql.Result, error) {
	return me.DB().Exec(dml, args...)
}

// Exec execute Sql, but do not commit transaction.
func (me *Manager) Exec(dml string, args ...interface{}) (sql.Result, error) {
	if me.Pre != nil {
		me.Pre()
	}
	return me.exec(dml, args...)
}

// ExecAndTx execute Sql and commit transactions.
func (me *Manager) ExecAndTx(dml string, args ...interface{}) (sql.Result, error) {
	if me.Pre != nil {
		me.Pre()
	}
	tx, err := me.DB().Begin()
	if err != nil {
		return nil, err
	}
	r, err := me.exec(dml, args...)
	if err != nil {
		tx.Rollback()
		return r, err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return r, err
	}
	return r, nil
}

// SetDB sets the xsql.DB.
func (me *Manager) SetDB(val xsql.DB) {
	me.db = val
}

// DB gets the xsql.DB.
func (me *Manager) DB() xsql.DB {
	if me.db == nil {
		if DB == nil {
			logger.Errorln("The default DB cannot be empty.")
		} else {
			me.SetDB(DB)
		}
	}
	return me.db
}
