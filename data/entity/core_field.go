// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package entity

type field struct {
	insertable bool
	updateable bool
	modified   bool
	excel      Excel
	json       Json
	xml        Xml
}

func (me *field) Insertable() bool {
	return me.insertable
}

func (me *field) SetInsertable(v bool) {
	me.insertable = v
}

func (me *field) Updateable() bool {
	return me.updateable
}

func (me *field) SetUpdateable(v bool) {
	me.updateable = v
}

func (me *field) Modified() bool {
	return me.modified
}

func (me *field) SetModified(v bool) {
	me.modified = v
}

func (me *field) Excel() Excel {
	return me.excel
}

func (me *field) SetExcel(v Excel) {
	me.excel = v
}

func (me *field) Json() Json {
	return me.json
}

func (me *field) SetJson(v Json) {
	me.json = v
}

func (me *field) Xml() Xml {
	return me.xml
}

func (me *field) SetXml(v Xml) {
	me.xml = v
}

func DefaultField() Field {
	return &field{insertable: true, updateable: true}
}
