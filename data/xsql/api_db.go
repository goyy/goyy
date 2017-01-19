// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xsql

import (
	"database/sql"

	"gopkg.in/goyy/goyy.v0/data/dialect"
	"gopkg.in/goyy/goyy.v0/data/domain"
	"gopkg.in/goyy/goyy.v0/data/entity"
)

// DB xsql.DB.
type DB interface {
	Query(dql string, args ...interface{}) Query
	NamedQuery(dql string, args map[string]interface{}) (Query, error)
	Sifter(sifts ...domain.Sift) Sifter

	Get(out entity.Interface) error

	Insert(e entity.Interface) (int64, error)
	Update(e entity.Interface) (int64, error)
	Delete(e entity.Interface) (int64, error)
	Disable(e entity.Interface) (int64, error)

	Exec(dml string, args ...interface{}) (sql.Result, error)

	Begin() (Tx, error)

	Close() error

	Dialect() dialect.Interface

	Ping() error
}
