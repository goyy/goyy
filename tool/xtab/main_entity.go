// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/goyy/goyy.v0/util/files"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

func genEntity() {
	for _, p := range conf.projects {
		for _, t := range conf.tables {
			if t.module.project.Id() == p.Id() && t.module.project.generate == "true" && t.module.generate == "true" && t.generate == "true" {
				dir := t.module.rootdir + "/internal/" + t.id + "/"
				dstfile := filepath.Join(dir, t.id+"_entity.go")
				if strings.IsNotBlank(t.master) {
					dir = t.module.rootdir + "/internal/" + t.master + "/"
					dstfile = filepath.Join(dir, t.slave+"_entity.go")
				}
				if files.IsExist(dstfile) {
					continue
				} else {
					files.MkdirAll(dir, 0644)
				}
				buf := bytes.Buffer{}
				tmpl := newTmpl(tmplEntity)
				tmpl.Execute(&buf, t)
				ioutil.WriteFile(dstfile, buf.Bytes(), 0644)
				write(tmplEntityBat, dir+"generate.bat")
				write(tmplEntitySh, dir+"generate.sh")
			}
		}
	}
}

func write(tmpls, dstfile string) {
	if files.IsExist(dstfile) == false {
		buf := bytes.Buffer{}
		tmpl := newTmpl(tmpls)
		tmpl.Execute(&buf, nil)
		ioutil.WriteFile(dstfile, buf.Bytes(), 0644)
	}
}

func newTmpl(s string) *template.Template {
	return template.Must(template.New("T").Funcs(funcs).Parse(s))
}

var (
	funcs = template.FuncMap{
		"supercol": util.IsSuperCol,
		"padname": func(s string, size int) string {
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
	}
)
