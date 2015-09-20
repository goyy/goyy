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
}

type module struct {
	id       string
	name     string
	prefix   string
	project  *project
	generate string
	comment  string
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
