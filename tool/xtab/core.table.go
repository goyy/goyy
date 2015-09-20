// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"strings"
)

type xTables struct {
	Table []*xTable `xml:"table"`
}

type xTable struct {
	Id       string     `xml:"id,attr"`
	Name     string     `xml:"name,attr"`
	Prefix   string     `xml:"prefix,attr"`
	Extends  string     `xml:"extends,attr"`
	Generate string     `xml:"generate,attr"`
	Comment  string     `xml:"comment,attr"`
	Columns  []*xColumn `xml:"column"`
}

type table struct {
	module   *module
	id       string
	name     string
	prefix   string
	parent   *table
	generate string
	comment  string
	columns  []*column
}

func (me *table) Id() string { // table.id: this
	return me.id
}

func (me *table) SetId(value string) {
	me.id = value
}

func (me *table) Ids() string {
	return me.Prefix() + "_" + me.Id()
}

func (me *table) Name() string { // table.name: this -> parent
	if strings.TrimSpace(me.name) == "" && me.parent != nil {
		return me.parent.Name()
	}
	return me.name
}

func (me *table) SetName(value string) {
	me.name = value
}

func (me *table) Prefix() string { // table.prefix: this -> module
	if strings.TrimSpace(me.prefix) == "" && me.module != nil {
		return me.module.Prefix()
	}
	return me.prefix
}

func (me *table) SetPrefix(value string) {
	me.prefix = value
}

func (me *table) Generate() string { // table.generate: this -> module
	if strings.TrimSpace(me.generate) == "" && me.module != nil {
		return me.module.Generate()
	}
	return me.generate
}

func (me *table) SetGenerate(value string) {
	me.generate = value
}

func (me *table) Comment() string { // table.comment: this -> parent
	if me.comment == "" && me.parent != nil {
		return me.parent.Comment()
	}
	return me.comment
}

func (me *table) SetComment(value string) {
	me.comment = value
}

func (me *table) Columns() []*column { // table.columns: this + parent
	var cs []*column
	if me.IsDefineId() {
		cs = append(cs, me.columns...)
		if me.parent != nil {
			for _, c := range me.parent.Columns() {
				if me.IsExists(me.columns, c.Id()) {
					continue
				}
				cs = append(cs, c)
			}
		}
	} else {
		if me.parent != nil {
			for _, c := range me.parent.Columns() {
				if c.Id() == "id" {
					cs = append(cs, c)
					break
				}
			}
		}
		cs = append(cs, me.columns...)
		if me.parent != nil {
			for _, c := range me.parent.Columns() {
				if me.IsExists(me.columns, c.Id()) || c.Id() == "id" {
					continue
				}
				cs = append(cs, c)
			}
		}
	}
	return cs
}

func (me *table) IsDefineId() bool {
	for _, c := range me.columns {
		if c.Id() == "id" {
			return true
		}
	}
	return false
}

func (me *table) IsExists(cs []*column, columnId string) bool {
	for _, c := range cs {
		if c.Id() == columnId {
			return true
		}
	}
	return false
}
