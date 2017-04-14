// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package service

import (
	"database/sql"

	"gopkg.in/goyy/goyy.v0/comm/xtype"
	"gopkg.in/goyy/goyy.v0/data/domain"
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/result"
)

// Service service.Service.
type Service interface {
	NewEntity() entity.Interface
	NewEntities() entity.Interfaces
	NewEntityResult() *result.Entity
	NewEntitiesResult() *result.Entities
	NewPageResult() *result.Page
	Get(out entity.Interface) error
	SelectOne(out entity.Interface, query string, args ...interface{}) error
	SelectList(out entity.Interfaces, query string, args ...interface{}) error
	SelectPage(content entity.Interfaces, pageable domain.Pageable, dql string, args ...interface{}) (out domain.Page, err error)
	SelectCount(dql string, args ...interface{}) (int, error)
	SelectInt(dql string, args ...interface{}) (int, error)
	SelectFloat(dql string, args ...interface{}) (float64, error)
	SelectStr(dql string, args ...interface{}) (string, error)
	SelectOneByNamed(out entity.Interface, query string, args map[string]interface{}) error
	SelectListByNamed(out entity.Interfaces, query string, args map[string]interface{}) error
	SelectPageByNamed(content entity.Interfaces, pageable domain.Pageable, dql string, args map[string]interface{}) (out domain.Page, err error)
	SelectCountByNamed(dql string, args map[string]interface{}) (int, error)
	SelectIntByNamed(dql string, args map[string]interface{}) (int, error)
	SelectFloatByNamed(dql string, args map[string]interface{}) (float64, error)
	SelectStrByNamed(dql string, args map[string]interface{}) (string, error)
	SelectOneBySift(out entity.Interface, sifts ...domain.Sift) error
	SelectListBySift(out entity.Interfaces, sifts ...domain.Sift) error
	SelectPageBySift(content entity.Interfaces, pageable domain.Pageable, sifts ...domain.Sift) (out domain.Page, err error)
	SelectCountBySift(sifts ...domain.Sift) (int, error)
	Save(p xtype.Principal, e entity.Interface) error
	SaveAndTx(p xtype.Principal, e entity.Interface) error
	Disable(p xtype.Principal, e entity.Interface) (int64, error)
	DisableAndTx(p xtype.Principal, e entity.Interface) (int64, error)
	Exec(dml string, args ...interface{}) (sql.Result, error)
	ExecAndTx(dml string, args ...interface{}) (sql.Result, error)
}
