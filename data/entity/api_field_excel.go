// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package entity

// Excel import and export information for the field of the entity struct.
type Excel interface {
	Value() string      // Name of the way to get the value of a field
	SetValue(v string)  // SetValue
	Title() string      // The title of the entity field displayed in the excel
	SetTitle(v string)  // SetTitle
	Format() string     // The format of the entity field displayed in the excel
	SetFormat(v string) // SetFormat
	Genre() int         // Field type(0:export or import,1:export only,2:import only)
	SetGenre(v int)     // SetGenre
	Align() int         // Alignment of columns(0:automatic,1:left,2:middle,3:right)
	SetAlign(v int)     // SetAlign
	Sort() int          // Display order of columns in the excel(0:asc,1:desc)
	SetSort(v int)      // SetSort
	Width() int         // Display width of columns in the excel
	SetWidth(v int)     // SetWidth
}

func NewExcel() Excel {
	return &excel{width: 20}
}

func NewExcelBy(value, title, format string, genre, align, sort, width int) Excel {
	return &excel{
		value:  value,
		title:  title,
		format: format,
		genre:  genre,
		align:  align,
		sort:   sort,
		width:  width,
	}
}
