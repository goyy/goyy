// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"html/template"
	"io/ioutil"

	"gopkg.in/goyy/goyy.v0/util/files"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

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
		"message": func(key string) string { // get message for i18n
			return i18N.Message(key)
		},
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
