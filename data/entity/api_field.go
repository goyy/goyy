// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package entity

type Field interface {
	Insertable() bool
	SetInsertable(v bool)
	Updateable() bool
	SetUpdateable(v bool)
	Modified() bool
	SetModified(v bool)
	Excel() Excel
	SetExcel(v Excel)
	Json() Json
	SetJson(v Json)
	Xml() Xml
	SetXml(v Xml)
}
