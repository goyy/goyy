// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"strconv"

	"gopkg.in/goyy/goyy.v0/util/strings"
)

// excelField import and export information for the field of the entity struct.
type excelField struct {
	Value  string // Name of the way to get the value of a field
	Title  string // The title of the entity field displayed in Excel
	Format string // The format of the entity field displayed in Excel
	Genre  int    // Field type(0:export import;1:export only;2:import only)
	Align  int    // Excel alignment(0:automatic;1:left;2:middle;3:right)
	Sort   int    // Display order of Excel field(asc)
	Width  int    // Display width of Excel field
}

// newExcelField analysis of tag to create a new excelField struct.
func newExcelField(f *field, tag string) (*excelField, bool) {
	if strings.IsBlank(tag) {
		return nil, false
	}
	fields := strings.Split(tag, "&")
	if fields == nil || len(fields) == 0 {
		return nil, false
	}
	ef := &excelField{
		Value: f.Name,
		Title: f.GetComment(),
		Width: 20,
	}
	ok := false
	for _, v := range fields {
		vs := strings.Split(v, "=")
		if len(vs) == 2 {
			name := strings.TrimSpace(vs[0])
			value := strings.TrimSpace(vs[1])
			switch strings.ToLower(name) {
			case "value":
				if strings.IsNotBlank(value) {
					ok = true
					ef.Value = value
				}
			case "title":
				if strings.IsNotBlank(value) {
					ok = true
					ef.Title = value
				}
			case "format":
				if strings.IsNotBlank(value) {
					ok = true
					ef.Format = value
				}
			case "genre":
				if strings.IsNotBlank(value) {
					if v, err := strconv.Atoi(value); err == nil {
						ok = true
						ef.Genre = v
					}
				}
			case "align":
				if strings.IsNotBlank(value) {
					if v, err := strconv.Atoi(value); err == nil {
						ok = true
						ef.Align = v
					}
				}
			case "sort":
				if strings.IsNotBlank(value) {
					if v, err := strconv.Atoi(value); err == nil {
						ok = true
						ef.Sort = v
					}
				}
			case "width":
				if strings.IsNotBlank(value) {
					if v, err := strconv.Atoi(value); err == nil {
						ok = true
						ef.Width = v
					}
				}
			}
		}
	}
	return ef, ok
}
