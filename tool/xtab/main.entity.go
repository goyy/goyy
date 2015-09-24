// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"gopkg.in/goyy/goyy.v0/util/files"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"html/template"
	"io/ioutil"
	"path/filepath"
)

func genEntity() {
	for _, p := range conf.projects {
		for _, t := range conf.tables {
			if t.module.project.Id() == p.Id() && t.module.project.generate == "true" && t.module.generate == "true" && t.generate == "true" {
				dir := t.module.rootdir + "/internal/" + t.id + "/"
				dstfile := filepath.Join(dir, "main.domain.go")
				if files.IsExist(dstfile) {
					continue
				} else {
					files.MkdirAll(dir, 0644)
				}
				buf := bytes.Buffer{}
				tmpl := newTmpl(tmplEntity)
				tmpl.Execute(&buf, t)
				ioutil.WriteFile(dstfile, buf.Bytes(), 0644)
			}
		}
	}
}

var (
	filters = template.FuncMap{
		"extends": func(column, extends string) bool {
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
		},
		"padname": func(s string, size int) string { // struct field name
			return strings.PadRight(s, size, " ")
		},
		"blank":    func(s string) bool { return strings.IsBlank(s) },
		"notblank": func(s string) bool { return strings.IsNoneBlank(s) },
		"padright": func(s string, size int) string { return strings.PadRight(s, size, " ") },
		"lower":    func(s string) string { return strings.ToLower(s) },
		"lower1":   func(s string) string { return strings.ToLowerFirst(s) },
		"upper":    func(s string) string { return strings.ToUpper(s) },
		"upper1":   func(s string) string { return strings.ToUpperFirst(s) },
		"camel":    func(s string) string { return strings.Camel(s) },
		"uncamel":  func(s string) string { return strings.UnCamel(s, "_") },
		"tname": func(s string, size int) string { // table name
			v := strings.UnCamel(s, "_")
			v = strings.ToUpper(v)
			return strings.PadRight(v, size+len(v)+1, " ")
		},
		"cname": func(s string, size int) string { // column name
			v := strings.UnCamel(s, "_")
			v = strings.ToUpper(v)
			return strings.PadRight(v, size, " ")
		},
		"mname": func(name, pkg string) string { // settings module name
			if name == "Entity" || name == "entity" {
				return pkg
			} else {
				name = strings.ToLowerFirst(name)
				if strings.HasSuffix(name, "Entity") {
					name = strings.Before(name, "Entity")
				}
				return pkg + "." + name
			}
		},
		"entities": func(name string) string { // settings module name
			if name == "Entity" || name == "entity" {
				return "Entities"
			} else {
				if strings.HasSuffix(name, "Entity") {
					name = strings.Before(name, "Entity")
				}
				return name + "Entities"
			}
		},
	}
)

func newTmpl(s string) *template.Template {
	return template.Must(template.New("T").Funcs(filters).Parse(s))
}
