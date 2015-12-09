// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"strings"
)

type xDomains struct {
	Domain []*xDomain `xml:"domain"`
}

type xDomain struct {
	Id        string `xml:"id,attr"`
	Name      string `xml:"name,attr"`
	Types     string `xml:"types,attr"`
	Length    int    `xml:"length,attr"`
	Precision int    `xml:"precision,attr"`
	Comment   string `xml:"comment,attr"`
	Defaults  string `xml:"default,attr"`
	Nullable  string `xml:"nullable,attr"`
}

type domain struct {
	id        string
	name      string
	types     string
	length    int
	precision int
	comment   string
	defaults  string
	nullable  string
	etype     string
}

func (me *domain) Id() string {
	return me.id
}

func (me *domain) SetId(value string) {
	me.id = value
}

func (me *domain) Name() string {
	return me.name
}

func (me *domain) SetName(value string) {
	me.name = value
}

func (me *domain) Types() string {
	if strings.TrimSpace(me.types) == "" {
		return "string"
	}
	return me.types
}

func (me *domain) SetTypes(value string) {
	me.types = value
}

func (me *domain) Length() int {
	return me.length
}

func (me *domain) SetLength(value int) {
	me.length = value
}

func (me *domain) Precision() int {
	return me.precision
}

func (me *domain) SetPrecision(value int) {
	me.precision = value
}

func (me *domain) Comment() string {
	return me.comment
}

func (me *domain) SetComment(value string) {
	me.comment = value
}

func (me *domain) Defaults() string {
	return me.defaults
}

func (me *domain) SetDefaults(value string) {
	me.defaults = value
}

func (me *domain) Nullable() string {
	return me.nullable
}

func (me *domain) SetNullable(value string) {
	me.nullable = value
}

func (me *domain) Etype() string {
	if strings.TrimSpace(me.etype) == "" {
		return "entity.String"
	}
	return me.etype
}

func (me *domain) SetEtype(value string) {
	me.etype = value
}
