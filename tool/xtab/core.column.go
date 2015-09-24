// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"strings"
)

type xColumns struct {
	Column []*xColumn `xml:"column"`
}

type xColumn struct {
	Extends  string `xml:"extends,attr"`
	Id       string `xml:"id,attr"`
	Name     string `xml:"name,attr"`
	Domain   string `xml:"domain,attr"`
	Index    string `xml:"index,attr"`
	Comment  string `xml:"comment,attr"`
	Defaults string `xml:"default,attr"`
	Nullable string `xml:"nullable,attr"`
}

type column struct {
	parent   *column
	id       string
	name     string
	domain   *domain
	index    string
	comment  string
	defaults string
	nullable string
	field    string
}

func (me *column) Id() string { // column.id: this -> parent
	if strings.TrimSpace(me.id) == "" && me.parent != nil {
		return me.parent.Id()
	}
	return me.id
}

func (me *column) SetId(value string) {
	me.id = value
}

func (me *column) Name() string { // column.name: this -> parent -> domain
	if strings.TrimSpace(me.name) == "" {
		if me.parent == nil {
			if me.domain != nil {
				return me.domain.Name()
			}
		} else {
			pn := me.parent.Name()
			if strings.TrimSpace(pn) == "" && me.domain != nil {
				return me.domain.Name()
			}
			return pn
		}
	}
	return me.name
}

func (me *column) SetName(value string) {
	me.name = value
}

func (me *column) Types() (types string) { // column.types: this.domain -> parent.domain
	if me.domain == nil {
		if me.parent == nil {
			types = "string"
		} else {
			types = me.parent.Types()
		}
	} else {
		types = me.domain.Types()
		if types == "" && me.parent != nil {
			types = me.parent.Types()
		}
	}
	return
}

func (me *column) Etype() (etype string) { // column.etype: this.domain -> parent.domain
	if me.domain == nil {
		if me.parent == nil {
			etype = "entity.String"
		} else {
			etype = me.parent.Etype()
		}
	} else {
		etype = me.domain.Etype()
		if etype == "" && me.parent != nil {
			etype = me.parent.Etype()
		}
	}
	return
}

func (me *column) Length() (length int) { // column.length: this.domain -> parent.domain
	if me.domain == nil {
		if me.parent == nil {
			length = 0
		} else {
			length = me.parent.Length()
		}
	} else {
		length = me.domain.Length()
		if length == 0 && me.parent != nil {
			length = me.parent.Length()
		}
	}
	return
}

func (me *column) Precision() (precision int) { // column.precision: this.domain -> parent.domain
	if me.domain == nil {
		if me.parent == nil {
			precision = 0
		} else {
			precision = me.parent.Precision()
		}
	} else {
		precision = me.domain.Precision()
		if precision == 0 && me.parent != nil {
			precision = me.parent.Precision()
		}
	}
	return
}

func (me *column) Index() string { // column.index: this -> parent
	if strings.TrimSpace(me.index) == "" && me.parent != nil {
		return me.parent.Index()
	}
	return me.index
}

func (me *column) SetIndex(value string) {
	me.index = value
}

func (me *column) Comment() string { // column.comment: this -> parent -> domain
	if strings.TrimSpace(me.comment) == "" {
		if me.parent == nil {
			if me.domain != nil {
				return me.domain.Comment()
			}
		} else {
			pc := me.parent.Comment()
			if strings.TrimSpace(pc) == "" && me.domain != nil {
				return me.domain.Name()
			}
			return pc
		}
	}
	return me.comment
}

func (me *column) SetComment(value string) {
	me.comment = value
}

func (me *column) Defaults() string { // column.defaults: this -> parent -> domain
	if strings.TrimSpace(me.defaults) == "" {
		if me.parent == nil {
			if me.domain != nil {
				return me.domain.Defaults()
			}
		} else {
			pc := me.parent.Defaults()
			if strings.TrimSpace(pc) == "" && me.domain != nil {
				return me.domain.Defaults()
			}
			return pc
		}
	}
	return me.defaults
}

func (me *column) SetDefaults(value string) {
	me.defaults = value
}

func (me *column) Nullable() string { // column.nullable: this -> parent -> domain
	if strings.TrimSpace(me.nullable) == "" {
		if me.parent == nil {
			if me.domain != nil {
				return me.domain.Nullable()
			}
		} else {
			pc := me.parent.Nullable()
			if strings.TrimSpace(pc) == "" && me.domain != nil {
				return me.domain.Nullable()
			}
			return pc
		}
	}
	return me.nullable
}

func (me *column) SetNullable(value string) {
	me.nullable = value
}

func (me *column) Field() string { // column.field: this -> parent
	if strings.TrimSpace(me.field) == "" && me.parent != nil {
		return me.parent.Field()
	}
	return me.field
}

func (me *column) SetField(value string) {
	me.field = value
}
