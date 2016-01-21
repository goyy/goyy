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

// Init sets the fields.
func (me *field) InitValidation(tag string) error {
	if strings.IsBlank(tag) {
		return nil
	}
	tag = strings.ToLower(tag)
	validations := strings.Split(tag, "&")
	if validations == nil || len(validations) == 0 {
		return nil
	}
	for _, v := range validations {
		vs := strings.Split(v, "=")
		if validations != nil && len(vs) == 2 {
			name := strings.TrimSpace(vs[0])
			value := strings.TrimSpace(vs[1])
			switch name {
			case "required":
				if strings.IsBlank(value) {
					continue
				}
				valid := &validation{
					Name: name,
				}
				if value == "true" {
					valid.Value = "true"
				} else {
					valid.Value = "false"
				}
				me.Validations = append(me.Validations, valid)
			case "min", "max", "range", "minlen", "maxlen", "rangelen",
				"email", "url", "ip", "mobile", "tel", "phone", "zipcode",
				"float", "integer", "alpha", "alrod", "alnum", "alnumrod",
				"alnumhan", "alnumhanrod", "alhan", "alhanrod", "han", "hanrod":
				if strings.IsBlank(value) {
					continue
				}
				valid := &validation{
					Name:  name,
					Value: value,
				}
				me.Validations = append(me.Validations, valid)
			}
		}
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
