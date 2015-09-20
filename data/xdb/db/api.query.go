// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package db

type Query interface {
	Select(es ...Expression) Query
	From(t Table) Query
	Where(c Condition) Query
	GroupBy(es ...Expression) Query
	Having(c Condition) Query
	OrderBy(es ...Expression) Query
	Limit(l uint8) Query
	Offset(o uint8) Query

	List(out interface{}) error
	One(out interface{}) error

	Reset()
	String() string
}
