// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"html/template"

	"gopkg.in/goyy/goyy.v0/util/files"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"gopkg.in/goyy/goyy.v0/util/times"
)

var now = times.NowUnixStr()

func writeTmpl(tmpls, dstfile string, params map[string]string) error {
	if files.IsExist(dstfile) == false {
		content, err := parseTmpl(tmpls, params)
		if err != nil {
			return err
		}
		return files.Write(dstfile, content, 0755)
	}
	return nil
}

func parseTmpl(tmpls string, params map[string]string) (string, error) {
	buf := &bytes.Buffer{}
	tmpl := newTmpl(tmpls)
	err := tmpl.Execute(buf, params)
	if err != nil {
		return "", err
	}
	content := buf.String()
	content = strings.Replace(content, "&#39;", "'", -1)
	return content, nil
}

func newTmpl(s string) *template.Template {
	return template.Must(template.New("t-xtab").Funcs(funcs).Parse(s))
}

var (
	funcs = template.FuncMap{
		"add1": func(v int) int { return v + 1 },
		"sub1": func(v int) int { return v - 1 },
		"message": func(key string) string { // get message for i18n
			return i18N.Message(key)
		},
		"supercol": util.IsSuperCol,
		"padname": func(s string, size int) string {
			return strings.PadRight(s, size, " ")
		},
		"blank":     func(s string) bool { return strings.IsBlank(s) },
		"notblank":  func(s string) bool { return strings.IsNoneBlank(s) },
		"padright":  func(s string, size int) string { return strings.PadRight(s, size, " ") },
		"lower":     func(s string) string { return strings.ToLower(s) },
		"lower1":    func(s string) string { return strings.ToLowerFirst(s) },
		"upper":     func(s string) string { return strings.ToUpper(s) },
		"upper1":    func(s string) string { return strings.ToUpperFirst(s) },
		"camel":     func(s string) string { return strings.Camel(s) },
		"uncamel":   func(s string) string { return strings.UnCamel(s, "_") },
		"seperator": func() string { return conf.Settings.Statement.Seperator },
		"case": func(s string) string {
			if strings.ToLower(conf.Settings.Statement.Case) == "upper" {
				return strings.ToUpper(s)
			}
			return strings.ToLower(s)
		},
		"uuid": func(driverName string) string {
			switch driverName {
			case "mymysql", "mysql":
				return "replace(uuid(), '-', '')"
			case "oci8", "oracle":
				return "sys_guid()"
			case "postgres":
				return "replace(uuid_generate_v4(), '-', '')"
			case "sqlserver":
				return "replace(newId(), '-', '')"
			case "sqlite", "sqlite3":
				return "uuid()"
			default:
				return "replace(uuid(), '-', '')"
			}
		},
	}
)
