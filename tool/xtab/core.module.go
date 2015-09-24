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
	Comment  string `xml:"comment,attr"`
	Rootdir  string `xml:"rootdir,attr"`
	Clidir   string `xml:"clidir,attr"`
	Clipath  string `xml:"clipath,attr"`
	Apidir   string `xml:"apidir,attr"`
	Apipath  string `xml:"apipath,attr"`
}

type module struct {
	id       string
	name     string
	prefix   string
	project  *project
	generate string
	comment  string
	rootdir  string
	clidir   string
	clipath  string
	apidir   string
	apipath  string
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

func (me *module) Comment() string { // module.comment: this
	return me.comment
}

func (me *module) SetComment(value string) {
	me.comment = value
}

func (me *module) Rootdir() string { // module.rootdir: this -> project
	if strings.TrimSpace(me.rootdir) == "" && me.project != nil {
		return me.project.Rootdir()
	}
	return me.rootdir
}

func (me *module) SetRootdir(value string) {
	me.rootdir = value
}

func (me *module) Clidir() string { // module.clidir: this -> project
	if strings.TrimSpace(me.clidir) == "" && me.project != nil {
		return me.project.Clidir()
	}
	return me.clidir
}

func (me *module) SetClidir(value string) {
	me.clidir = value
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

func (me *module) Apidir() string { // module.apidir: this -> project
	if strings.TrimSpace(me.apidir) == "" && me.project != nil {
		return me.project.Apidir()
	}
	return me.apidir
}

func (me *module) SetApidir(value string) {
	me.apidir = value
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
