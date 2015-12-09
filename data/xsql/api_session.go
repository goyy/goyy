// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xsql

import (
	"database/sql"
	"gopkg.in/goyy/goyy.v0/data/domain"
	"gopkg.in/goyy/goyy.v0/data/entity"
)

type Session interface {
	Query(dql string, args ...interface{}) Query
	NamedQuery(dql string, args map[string]interface{}) (Query, error)

	Get(out entity.Interface) error
	SelectOne(out entity.Interface, sifts ...domain.Sift) error
	SelectList(out entity.Interfaces, sifts ...domain.Sift) error
	SelectPage(content entity.Interfaces, pageable domain.Pageable, sifts ...domain.Sift) (domain.Page, error)
	SelectCount(e entity.Interface, sifts ...domain.Sift) (int, error)

	Insert(e entity.Interface) (int64, error)
	Update(e entity.Interface) (int64, error)
	Delete(e entity.Interface) (int64, error)
	Disable(e entity.Interface) (int64, error)

	Exec(dml string, args ...interface{}) (sql.Result, error)

	Begin() (Tx, error)

	Close() error

	DBType() string
}
