// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/xml"
	"gopkg.in/goyy/goyy.v0/util/files"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"log"
	"os"
)

type utils struct{}

func (me *utils) Pad(in string) string {
	return strings.PadEnd(in, 30)
}

func (me *utils) IsSuperCol(column, extends string) bool {
	if extends == "pk" && column == "id" {
		return true
	}
	if extends == "sys" {
		switch column {
		case "id":
			return true
		case "memo", "creates", "creater", "created", "modifier", "modified", "version", "deletion", "artifical", "history":
			return true
		}
	}
	if extends == "tree" {
		switch column {
		case "id":
			return true
		case "memo", "creates", "creater", "created", "modifier", "modified", "version", "deletion", "artifical", "history":
			return true
		case "code", "name", "fullname", "genre", "ordinal", "parent_id", "parent_ids", "parent_codes", "parent_names", "leaf", "grade":
			return true
		}
	}
	return false
}

func (me *utils) Etype(in string) string {
	switch in {
	case "int":
		return "entity.Int"
	case "long":
		return "entity.Int64"
	case "float":
		return "entity.Float64"
	case "bool":
		return "entity.Bool"
	case "time":
		return "entity.Time"
	default:
		return "entity.String"
	}
}

func (me *utils) Case(in string) (out string) {
	if conf.Settings.Statement.Case == lower {
		out = strings.ToLower(in)
	} else {
		out = strings.ToUpper(in)
	}
	return
}

func (me *utils) WriteString(filename, content string) {
	f, ferr := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	defer f.Close()
	if ferr != nil {
		log.Fatal(ferr)
	}
	_, werr := f.WriteString(content)
	if werr != nil {
		log.Fatal(werr)
	}
}

func (me *utils) InitFile(filename, content string) bool {
	if files.IsExist(filename) == false {
		me.WriteString(filename, content)
		return false
	}
	return true
}

func (me *utils) DecodeXML(filename string) *xConfiguration {
	f, ferr := os.Open(filename)
	defer f.Close()
	if ferr != nil {
		log.Fatal(ferr)
	}
	decoder := xml.NewDecoder(f)
	var xconf *xConfiguration
	if derr := decoder.Decode(&xconf); derr != nil {
		log.Fatal(derr)
	}
	return xconf
}
