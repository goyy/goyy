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
	Id       string `xml:"id,attr"`
	Name     string `xml:"name,attr"`
	Prefix   string `xml:"prefix,attr"`
	Database string `xml:"database,attr"`
	Generate string `xml:"generate,attr"`
	Comment  string `xml:"comment,attr"`
}

type project struct {
	id       string
	name     string
	prefix   string
	database *database
	generate string
	comment  string
}

func (me *project) Id() string {
	return me.id
}

func (me *project) SetId(value string) {
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

func (me *project) Comment() string {
	return me.comment
}

func (me *project) SetComment(value string) {
	me.comment = value
}
