// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xsql

import (
	"gopkg.in/goyy/goyy.v0/data/dialect"
)

type Factory interface {
	Session() (Session, error)
	Dialect() dialect.Interface
	Ping() error
}
