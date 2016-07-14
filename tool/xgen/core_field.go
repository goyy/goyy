// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"fmt"

	"gopkg.in/goyy/goyy.v0/util/strings"
)

// field is a field in an entity struct.
type field struct {
	Name        string
	Type        string
	Column      string
	Comment     string
	Dict        string
	Default     string
	Validations validations
	Excel       *excelField
	Json        *jsonField
	Xml         *xmlField
	IsExcel     bool
	IsJson      bool
	IsXml       bool
	IsPrimary   bool
	IsForeign   bool
	IsVersion   bool
	IsDeletion  bool
	IsCreater   bool
	IsCreated   bool
	IsModifier  bool
	IsModified  bool
	IsTransient bool
}

// Init sets the fields.
func (me *field) Init(name, typ, tag string) error {
	me.Name = name
	me.Type = typ

	// auto-detect foreign key
	if len(name) > 2 && name[len(name)-2:] == "Id" {
		me.IsForeign = true
	}

	var isSetColumn bool

	// parse attributes
	attributes := strings.Split(tag, "&")
	for _, attr := range attributes {
		pair := strings.Split(attr, "=")
		if len(pair) != 2 {
			return fmt.Errorf("Malformed tag: '%s'", attr)
		}

		switch strings.ToLower(pair[0]) {
		case "column":
			isSetColumn = true
			me.Column = pair[1]
		case "comment":
			me.Comment = pair[1]
		case "dict":
			me.Dict = pair[1]
		case "default":
			me.Default = pair[1]
		case "primary":
			if pair[1] == "true" {
				me.IsPrimary = true
			}
		case "version":
			if pair[1] == "true" {
				me.IsVersion = true
			}
		case "deletion":
			if pair[1] == "true" {
				me.IsDeletion = true
			}
		case "creater":
			if pair[1] == "true" {
				me.IsCreater = true
			}
		case "created":
			if pair[1] == "true" {
				me.IsCreated = true
			}
		case "modifier":
			if pair[1] == "true" {
				me.IsModifier = true
			}
		case "modified":
			if pair[1] == "true" {
				me.IsModified = true
			}
		case "transient":
			if pair[1] == "true" {
				me.IsTransient = true
			}
		default:
			return fmt.Errorf("Unknown attribute: '%s'", pair[0])
		}
		if name == "id" {
			me.IsPrimary = true
		}
	}

	if !isSetColumn {
		me.Column = strings.UnCamel(name, "_")
	}

	return nil
}

func (me field) String() string {
	buf := bytes.Buffer{}
	tmpl := newTmpl(`{{.Name}} {{.Type}}  {{.Column}} {{if .IsPrimary}} {{if .IsForeign}}{{end}}`)
	tmpl.Execute(&buf, me)
	return buf.String()
}

func (me field) GetComment() string {
	if strings.IsBlank(me.Comment) {
		return strings.ToUpper(me.Column)
	}
	return me.Comment
}
