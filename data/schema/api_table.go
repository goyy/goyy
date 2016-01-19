// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package schema

type Table interface {
	Name() string
	Comment() string
	Column(columnName string) Column
	Columns() []Column
	Primary() Column
	Version() Column
	Deletion() Column
	Creater() Column
	Created() Column
	Modifier() Column
	Modified() Column
	String() string

	COLUMN(columnName string) Column
	PRIMARY(columnName string) Column
	VERSION(columnName string) Column
	DELETION(columnName string) Column
	CREATER(columnName string) Column
	CREATED(columnName string) Column
	MODIFIER(columnName string) Column
	MODIFIED(columnName string) Column
	TRANSIENT(columnName string) Column
}
