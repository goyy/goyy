// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package repository_test

import (
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/goyy/goyy.v0/data/dialect"
	"gopkg.in/goyy/goyy.v0/data/repository"
)

var repo *repository.Repository

func init() {
	repo = repository.New(&dialect.MySQL{}, "db")
}
