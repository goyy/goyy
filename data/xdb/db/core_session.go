// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package db

import (
	"database/sql"
)

type session struct {
	Db      *sql.DB
	Dialect Dialect
}

// New Query
func (me *session) Query() Query {
	return &query{Db: me.Db, Dialect: me.Dialect}
}

// Insert SQL
func (me *session) Insert(entity ...interface{}) (string, error) {
	return "", nil
}

// Update SQL
func (me *session) Update(entity ...interface{}) (string, error) {
	return "", nil
}

// Delete SQL
func (me *session) Delete(entity ...interface{}) (string, error) {
	return "", nil
}

// Begin Transaction
func (me *session) Begin() (Tx, error) {
	return me.Db.Begin()
}

// Close Session
func (me *session) Close() error {
	return me.Db.Close()
}
