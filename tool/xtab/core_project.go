// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"strings"
)

type xProjects struct {
	Project []*xProject `xml:"project"`
}

type xProject struct {
	ID       string `xml:"id,attr"`
	Name     string `xml:"name,attr"`
	Prefix   string `xml:"prefix,attr"`
	Database string `xml:"database,attr"`
	Generate string `xml:"generate,attr"`
	Menu     string `xml:"menu,attr"`
	Comment  string `xml:"comment,attr"`
	Admpath  string `xml:"admpath,attr"`
	Apipath  string `xml:"apipath,attr"`
	Tstpath  string `xml:"tstpath,attr"`
}

type project struct {
	id       string
	name     string
	prefix   string
	database *database
	generate string
	menu     string
	comment  string
	admpath  string
	apipath  string
	tstpath  string
}

func (me *project) ID() string {
	return me.id
}

func (me *project) SetID(value string) {
	me.id = value
}

func (me *project) Name() string {
	return me.name
}

func (me *project) SetName(value string) {
	me.name = value
}

func (me *project) Prefix() string {
	return me.prefix
}

func (me *project) SetPrefix(value string) {
	me.prefix = value
}

func (me *project) Generate() string {
	if strings.TrimSpace(me.generate) == "" {
		return "true"
	}
	return me.generate
}

func (me *project) SetGenerate(value string) {
	me.generate = value
}

func (me *project) Menu() string {
	if strings.TrimSpace(me.menu) == "" {
		return "true"
	}
	return me.menu
}

func (me *project) SetMenu(value string) {
	me.menu = value
}

func (me *project) Comment() string {
	return me.comment
}

func (me *project) SetComment(value string) {
	me.comment = value
}

func (me *project) Admpath() string {
	return me.admpath
}

func (me *project) SetAdmpath(value string) {
	me.admpath = value
}

func (me *project) Apipath() string {
	return me.apipath
}

func (me *project) SetApipath(value string) {
	me.apipath = value
}

func (me *project) Tstpath() string {
	return me.tstpath
}

func (me *project) SetTstpath(value string) {
	me.tstpath = value
}
