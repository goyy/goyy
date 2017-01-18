// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"gopkg.in/goyy/goyy.v0/util/strings"
)

type xTables struct {
	Table []*xTable `xml:"table"`
}

type xTable struct {
	ID          string     `xml:"id,attr"`
	Name        string     `xml:"name,attr"`
	Prefix      string     `xml:"prefix,attr"`
	Extends     string     `xml:"extends,attr"`
	Generate    string     `xml:"generate,attr"`
	Menu        string     `xml:"menu,attr"`
	Comment     string     `xml:"comment,attr"`
	Master      string     `xml:"master,attr"`
	Slave       string     `xml:"slave,attr"`
	Permissions string     `xml:"permissions,attr"`
	Href        string     `xml:"href,attr"`
	Buttons     string     `xml:"buttons,attr"`
	Columns     []*xColumn `xml:"column"`
}

type table struct {
	module          *module
	id              string
	name            string
	prefix          string
	parent          *table
	generate        string
	menu            string
	comment         string
	master          string
	slave           string
	permissions     string
	href            string
	buttons         string
	fieldMaxLen     int
	columnMaxLen    int
	typeMaxLen      int
	allFieldMaxLen  int
	allColumnMaxLen int
	allTypeMaxLen   int
	columns         []*column
}

func (me *table) Module() *module { // table.module: this
	return me.module
}

func (me *table) Parent() *table { // table.parent: this
	return me.parent
}

func (me *table) FieldMaxLen() int { // table.fieldMaxLen: this
	return me.fieldMaxLen
}

func (me *table) ColumnMaxLen() int { // table.columnMaxLen: this
	return me.columnMaxLen
}

func (me *table) TypeMaxLen() int { // table.typeMaxLen: this
	return me.typeMaxLen
}

func (me *table) AllFieldMaxLen() int { // table.allFieldMaxLen: this
	return me.allFieldMaxLen
}

func (me *table) AllColumnMaxLen() int { // table.allColumnMaxLen: this
	return me.allColumnMaxLen
}

func (me *table) AllTypeMaxLen() int { // table.allTypeMaxLen: this
	return me.allTypeMaxLen
}

func (me *table) Super() string { // table.super: this -> parent
	if me.parent != nil {
		super := me.parent.id
		if strings.IsNotBlank(super) {
			switch super {
			case "pk", "sys", "tree":
				return super
			default:
				return me.parent.Super()
			}
		}
	}
	return ""
}

func (me *table) ID() string { // table.id: this
	return me.id
}

func (me *table) SetID(value string) {
	me.id = value
}

func (me *table) IDs() string {
	return me.Prefix() + "_" + me.ID()
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

func (me *table) Menu() string { // table.menu: this -> module
	if strings.TrimSpace(me.menu) == "" && me.module != nil {
		return me.module.Menu()
	}
	return me.menu
}

func (me *table) SetMenu(value string) {
	me.menu = value
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

func (me *table) Master() string { // table.master: this
	return me.master
}

func (me *table) SetMaster(value string) {
	me.master = value
}

func (me *table) Slave() string { // table.slave: this
	return me.slave
}

func (me *table) SetSlave(value string) {
	me.slave = value
}

func (me *table) Permissions() string { // table.permissions: this
	if strings.TrimSpace(me.permissions) == "" {
		return fmt.Sprintf("%s:%s:view", me.module.ID(), me.ID())
	}
	return me.permissions
}

func (me *table) SetPermissions(value string) {
	me.permissions = value
}

func (me *table) Href() string { // table.href: this
	if strings.TrimSpace(me.permissions) == "" {
		return fmt.Sprintf("/%s/%s/%s.html", me.module.ID(), me.ID(), me.ID())
	}
	return me.href
}

func (me *table) SetHref(value string) {
	me.href = value
}

func (me *table) Buttons() string { // table.buttons: this
	return me.buttons
}

func (me *table) SetButtons(value string) {
	me.buttons = value
}

func (me *table) Columns() []*column { // table.columns: this + parent
	var cs []*column
	if me.IsDefineID() {
		cs = append(cs, me.columns...)
		if me.parent != nil {
			for _, c := range me.parent.Columns() {
				if me.IsExists(me.columns, c.ID()) {
					continue
				}
				cs = append(cs, c)
			}
		}
	} else {
		if me.parent != nil {
			for _, c := range me.parent.Columns() {
				if c.ID() == "id" {
					cs = append(cs, c)
					break
				}
			}
		}
		cs = append(cs, me.columns...)
		if me.parent != nil {
			for _, c := range me.parent.Columns() {
				if me.IsExists(me.columns, c.ID()) || c.ID() == "id" {
					continue
				}
				cs = append(cs, c)
			}
		}
	}
	return cs
}

func (me *table) IsDefineID() bool {
	for _, c := range me.columns {
		if c.ID() == "id" {
			return true
		}
	}
	return false
}

func (me *table) IsExists(cs []*column, columnID string) bool {
	for _, c := range cs {
		if c.ID() == columnID {
			return true
		}
	}
	return false
}
