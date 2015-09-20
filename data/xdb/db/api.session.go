// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package db

type Session interface {
	Query() Query

	Insert(entity ...interface{}) (string, error)
	Update(entity ...interface{}) (string, error)
	Delete(entity ...interface{}) (string, error)

	Begin() (Tx, error)

	Close() error
}
