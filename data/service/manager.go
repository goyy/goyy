// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package service

import (
	"database/sql"
	"errors"

	"gopkg.in/goyy/goyy.v0/data/domain"
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/repository"
	"gopkg.in/goyy/goyy.v0/data/result"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"gopkg.in/goyy/goyy.v0/util/times"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

type Manager struct {
	repository *repository.Repository
	Entity     func() entity.Interface
	Entities   func() entity.Interfaces
	Pre        func()
}

func (me *Manager) NewEntity() entity.Interface {
	if me.Entity != nil {
		return me.Entity()
	}
	return nil
}

func (me *Manager) NewEntities() entity.Interfaces {
	if me.Entities != nil {
		return me.Entities()
	}
	return nil
}

func (me *Manager) NewEntityResult() *result.Entity {
	return &result.Entity{Data: me.NewEntity()}
}

func (me *Manager) NewEntitiesResult() *result.Entities {
	return &result.Entities{Data: me.NewEntities()}
}

func (me *Manager) NewPageResult() *result.Page {
	out := domain.NewPageDefault(me.NewEntities())
	return &result.Page{Data: out}
}

func (me *Manager) Get(out entity.Interface) error {
	if me.Pre != nil {
		me.Pre()
	}
	return me.Repository().Get(out)
}

func (me *Manager) SelectOne(out entity.Interface, query string, args ...interface{}) (err error) {
	if me.Pre != nil {
		me.Pre()
	}
	err = me.Repository().SelectOne(out, query, args...)
	return
}

func (me *Manager) SelectList(out entity.Interfaces, query string, args ...interface{}) (err error) {
	if me.Pre != nil {
		me.Pre()
	}
	err = me.Repository().SelectList(out, query, args...)
	return
}

func (me *Manager) SelectPage(content entity.Interfaces, pageable domain.Pageable, dql string, args ...interface{}) (out domain.Page, err error) {
	if me.Pre != nil {
		me.Pre()
	}
	return me.Repository().SelectPage(content, pageable, dql, args...)
}

func (me *Manager) SelectInt(dql string, args ...interface{}) (int, error) {
	if me.Pre != nil {
		me.Pre()
	}
	return me.Repository().SelectInt(dql, args...)
}

func (me *Manager) SelectFloat(dql string, args ...interface{}) (float64, error) {
	if me.Pre != nil {
		me.Pre()
	}
	return me.Repository().SelectFloat(dql, args...)
}

func (me *Manager) SelectStr(dql string, args ...interface{}) (string, error) {
	if me.Pre != nil {
		me.Pre()
	}
	return me.Repository().SelectStr(dql, args...)
}

func (me *Manager) SelectOneByNamed(out entity.Interface, query string, args map[string]interface{}) (err error) {
	if me.Pre != nil {
		me.Pre()
	}
	err = me.Repository().SelectOneByNamed(out, query, args)
	return
}

func (me *Manager) SelectListByNamed(out entity.Interfaces, query string, args map[string]interface{}) (err error) {
	if me.Pre != nil {
		me.Pre()
	}
	err = me.Repository().SelectListByNamed(out, query, args)
	return
}

func (me *Manager) SelectPageByNamed(content entity.Interfaces, pageable domain.Pageable, dql string, args map[string]interface{}) (out domain.Page, err error) {
	if me.Pre != nil {
		me.Pre()
	}
	return me.Repository().SelectPageByNamed(content, pageable, dql, args)
}

func (me *Manager) SelectCountByNamed(dql string, args map[string]interface{}) (int, error) {
	if me.Pre != nil {
		me.Pre()
	}
	return me.Repository().SelectCountByNamed(dql, args)
}

func (me *Manager) SelectIntByNamed(dql string, args map[string]interface{}) (int, error) {
	if me.Pre != nil {
		me.Pre()
	}
	return me.Repository().SelectIntByNamed(dql, args)
}

func (me *Manager) SelectOneBySift(out entity.Interface, sifts ...domain.Sift) (err error) {
	if me.Pre != nil {
		me.Pre()
	}
	err = me.Repository().SelectOneBySift(out, sifts...)
	return
}

func (me *Manager) SelectListBySift(out entity.Interfaces, sifts ...domain.Sift) (err error) {
	if me.Pre != nil {
		me.Pre()
	}
	err = me.Repository().SelectListBySift(out, sifts...)
	return
}

func (me *Manager) SelectPageBySift(content entity.Interfaces, pageable domain.Pageable, sifts ...domain.Sift) (out domain.Page, err error) {
	if me.Pre != nil {
		me.Pre()
	}
	return me.Repository().SelectPageBySift(content, pageable, sifts...)
}

func (me *Manager) SelectCountBySift(sifts ...domain.Sift) (int, error) {
	if me.Pre != nil {
		me.Pre()
	}
	return me.Repository().SelectCountBySift(me.NewEntity(), sifts...)
}

func (me *Manager) Save(c xhttp.Context, e entity.Interface) error {
	if me.Pre != nil {
		me.Pre()
	}
	if err := e.Validate(); err != nil {
		return err
	}
	tx, err := me.Repository().Begin()
	if err != nil {
		return err
	}
	if strings.IsBlank(e.Get(e.Table().Primary().Name()).(string)) {
		if c != nil && c.Session().IsLogin() {
			if p, err := c.Session().Principal(); err == nil {
				e.SetString(creater, p.Id)
				e.SetString(modifier, p.Id)
			}
		}
		e.SetString(created, times.NowUnixStr())
		e.SetString(modified, times.NowUnixStr())
		_, err = me.Repository().Insert(e)
	} else {
		if c != nil && c.Session().IsLogin() {
			if p, err := c.Session().Principal(); err == nil {
				e.SetString(modifier, p.Id)
			}
		}
		e.SetString(modified, times.NowUnixStr())
		_, err = me.Repository().Update(e)
	}
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

func (me *Manager) Disable(c xhttp.Context, e entity.Interface) (int64, error) {
	if me.Pre != nil {
		me.Pre()
	}
	if strings.IsBlank(e.Get(e.Table().Primary().Name()).(string)) {
		return 0, errors.New("Gets the primary key value failed")
	}
	tx, err := me.Repository().Begin()
	if err != nil {
		return 0, err
	}
	if c != nil && c.Session().IsLogin() {
		if p, err := c.Session().Principal(); err == nil {
			e.SetString(modifier, p.Id)
			e.SetString(modified, times.NowUnixStr())
		}
	}
	r, err := me.Repository().Disable(e)
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

func (me *Manager) Exec(dml string, args ...interface{}) (sql.Result, error) {
	if me.Pre != nil {
		me.Pre()
	}
	tx, err := me.Repository().Begin()
	if err != nil {
		return nil, err
	}
	r, err := me.Repository().Exec(dml, args...)
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

func (me *Manager) SetRepository(val *repository.Repository) {
	me.repository = val
}

func (me *Manager) Repository() *repository.Repository {
	if me.repository == nil {
		me.SetRepository(Repository)
	}
	return me.repository
}
