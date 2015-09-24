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
	Rootdir  string `xml:"rootdir,attr"`
	Clidir   string `xml:"clidir,attr"`
	Clipath  string `xml:"clipath,attr"`
	Apidir   string `xml:"apidir,attr"`
	Apipath  string `xml:"apipath,attr"`
}

type project struct {
	id       string
	name     string
	prefix   string
	database *database
	generate string
	comment  string
	rootdir  string
	clidir   string
	clipath  string
	apidir   string
	apipath  string
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

func (me *project) Rootdir() string {
	return me.rootdir
}

func (me *project) SetRootdir(value string) {
	me.rootdir = value
}

func (me *project) Clidir() string {
	return me.clidir
}

func (me *project) SetClidir(value string) {
	me.clidir = value
}

func (me *project) Clipath() string {
	return me.clipath
}

func (me *project) SetClipath(value string) {
	me.clipath = value
}

func (me *project) Apidir() string {
	return me.apidir
}

func (me *project) SetApidir(value string) {
	me.apidir = value
}

func (me *project) Apipath() string {
	return me.apipath
}

func (me *project) SetApipath(value string) {
	me.apipath = value
}
