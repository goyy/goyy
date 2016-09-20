// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"strings"
)

type xModules struct {
	Module []*xModule `xml:"module"`
}

type xModule struct {
	Id       string `xml:"id,attr"`
	Name     string `xml:"name,attr"`
	Prefix   string `xml:"prefix,attr"`
	Project  string `xml:"project,attr"`
	Generate string `xml:"generate,attr"`
	Menu     string `xml:"menu,attr"`
	Comment  string `xml:"comment,attr"`
	Clipath  string `xml:"clipath,attr"`
	Apipath  string `xml:"apipath,attr"`
	Tstpath  string `xml:"tstpath,attr"`
}

type module struct {
	id       string
	name     string
	prefix   string
	project  *project
	generate string
	menu     string
	comment  string
	clipath  string
	apipath  string
	tstpath  string
}

func (me *module) Id() string { // module.id: this
	return me.id
}

func (me *module) SetId(value string) {
	me.id = value
}

func (me *module) Name() string { // module.name: this
	return me.name
}

func (me *module) SetName(value string) {
	me.name = value
}

func (me *module) Prefix() string { // module.prefix: this -> project
	if strings.TrimSpace(me.prefix) == "" && me.project != nil {
		return me.project.Prefix()
	}
	return me.prefix
}

func (me *module) SetPrefix(value string) {
	me.prefix = value
}

func (me *module) Generate() string { // module.generate: this -> project
	if strings.TrimSpace(me.generate) == "" && me.project != nil {
		return me.project.Generate()
	}
	return me.generate
}

func (me *module) SetGenerate(value string) {
	me.generate = value
}

func (me *module) Menu() string { // module.menu: this -> project
	if strings.TrimSpace(me.menu) == "" && me.project != nil {
		return me.project.Menu()
	}
	return me.menu
}

func (me *module) SetMenu(value string) {
	me.menu = value
}

func (me *module) Comment() string { // module.comment: this
	return me.comment
}

func (me *module) SetComment(value string) {
	me.comment = value
}

func (me *module) Clipath() string { // module.clipath: this -> project
	if strings.TrimSpace(me.clipath) == "" && me.project != nil {
		return me.project.Clipath()
	}
	return me.clipath
}

func (me *module) SetClipath(value string) {
	me.clipath = value
}

func (me *module) Apipath() string { // module.apipath: this -> project
	if strings.TrimSpace(me.apipath) == "" && me.project != nil {
		return me.project.Apipath()
	}
	return me.apipath
}

func (me *module) SetApipath(value string) {
	me.apipath = value
}

func (me *module) Tstpath() string { // module.tstpath: this -> project
	if strings.TrimSpace(me.tstpath) == "" && me.project != nil {
		return me.project.Tstpath()
	}
	return me.tstpath
}

func (me *module) SetTstpath(value string) {
	me.tstpath = value
}
