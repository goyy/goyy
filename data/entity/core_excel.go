// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package entity

// Excel import and export information for the field of the entity struct.
type excel struct {
	value  string // Name of the way to get the value of a field
	title  string // The title of the entity field displayed in Excel
	format string // The format of the entity field displayed in Excel
	genre  int    // Field type(0:export import;1:export only;2:import only)
	align  int    // Excel alignment(0:automatic;1:left;2:middle;3:right)
	sort   int    // Display order of Excel field(asc)
	width  int    // Display width of Excel field
}

func (me *excel) Value() string {
	return me.value
}

func (me *excel) SetValue(v string) {
	me.value = v
}

func (me *excel) Title() string {
	return me.title
}

func (me *excel) SetTitle(v string) {
	me.title = v
}

func (me *excel) Format() string {
	return me.format
}

func (me *excel) SetFormat(v string) {
	me.format = v
}

func (me *excel) Genre() int {
	return me.genre
}

func (me *excel) SetGenre(v int) {
	me.genre = v
}

func (me *excel) Align() int {
	return me.align
}

func (me *excel) SetAlign(v int) {
	me.align = v
}

func (me *excel) Sort() int {
	return me.sort
}

func (me *excel) SetSort(v int) {
	me.sort = v
}

func (me *excel) Width() int {
	return me.width
}

func (me *excel) SetWidth(v int) {
	me.width = v
}
