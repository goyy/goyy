// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package entity

// Excel import and export information for the field of the entity struct.
type Excel interface {
	Value() string // Name of the way to get the value of a field
	SetValue(v string)
	Title() string // The title of the entity field displayed in Excel
	SetTitle(v string)
	Genre() int // Field type(0:export import;1:export only;2:import only)
	SetGenre(v int)
	Align() int // Excel alignment(0:automatic;1:left;2:middle;3:right)
	SetAlign(v int)
	Sort() int // Display order of Excel field(asc)
	SetSort(v int)
	Width() int // Display width of Excel field
	SetWidth(v int)
}

func NewExcel() Excel {
	return &excel{width: 3000}
}

func NewExcelBy(value, title string, genre, align, sort, width int) Excel {
	return &excel{
		value: value,
		title: title,
		genre: genre,
		align: align,
		sort:  sort,
		width: width,
	}
}
