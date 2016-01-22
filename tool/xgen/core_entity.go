// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"container/heap"

	"gopkg.in/goyy/goyy.v0/comm/xtype"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

// entity is an SQL table.
type entity struct {
	Project         string
	Extend          string // value:[tree,sys,pk]
	Relationship    string // value:[master,slave]
	Name            string
	Comment         string
	Table           string
	Fields          []*field
	PrimaryKeys     []*field
	FieldMaxLen     int
	ColumnMaxLen    int
	TypeMaxLen      int
	AllFieldMaxLen  int
	AllColumnMaxLen int
	AllTypeMaxLen   int
}

func (me entity) GetComment() string {
	if strings.IsBlank(me.Comment) {
		return strings.ToUpper(me.Name)
	}
	return me.Comment
}

// GetExcels gets the column names sorted using Excel.Sort
func (me entity) GetExcelColumns() []string {
	var columns []string
	m := make(map[int]string, 0)
	h := &xtype.IntHeap{}
	heap.Init(h)
	for _, f := range me.Fields {
		if f.IsExcel {
			m[f.Excel.Sort] = f.Column
			heap.Push(h, f.Excel.Sort)
		}
	}
	l := h.Len()
	if l > 0 {
		for i := 0; i < l; i++ {
			k := heap.Pop(h).(int)
			columns = append(columns, m[k])
		}
	}
	return columns
}

// IsExcelColumns to determine whether there is excel information
func (me entity) IsExcelColumns() bool {
	s := me.GetExcelColumns()
	if s == nil || len(s) == 0 {
		return false
	} else {
		return true
	}
}
