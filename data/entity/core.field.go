// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package entity

type field struct {
	insertable bool
	updateable bool
	modified   bool
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

func DefaultField() Field {
	return &field{insertable: true, updateable: true}
}
