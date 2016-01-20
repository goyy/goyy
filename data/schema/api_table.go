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

	COLUMN(columnName, comment string) Column
	PRIMARY(columnName, comment string) Column
	VERSION(columnName, comment string) Column
	DELETION(columnName, comment string) Column
	CREATER(columnName, comment string) Column
	CREATED(columnName, comment string) Column
	MODIFIER(columnName, comment string) Column
	MODIFIED(columnName, comment string) Column
	TRANSIENT(columnName, comment string) Column
}
