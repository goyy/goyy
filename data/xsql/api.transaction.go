// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xsql

import (
	"database/sql"
)

type Tx interface {
	Exec(dml string, args ...interface{}) (sql.Result, error)
	Commit() error
	Rollback() error
}
