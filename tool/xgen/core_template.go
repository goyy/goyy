// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"gopkg.in/goyy/goyy.v0/util/strings"
	"text/template"
)

var (
	filters = template.FuncMap{
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
		"fname": func(s string, size int) string { // struct field name
			v := strings.ToUpperFirst(s)
			return strings.PadRight(v, size, " ")
		},
		"ctlvar": func(name, relationship string) string { // controller var name
			if relationship == "master" {
				return "ctl"
			} else {
				if strings.HasSuffix(name, "Entity") {
					name = strings.Before(name, "Entity")
				}
				return strings.ToLowerFirst(name) + "Ctl"
			}
		},
		"ctlname": func(name, relationship string) string { // controller struct name
			if relationship == "master" {
				return "Controller"
			} else {
				if strings.HasSuffix(name, "Entity") {
					name = strings.Before(name, "Entity")
				}
				return name + "Controller"
			}
		},
		"mgrvar": func(name, relationship string) string { // manager var name
			if relationship == "master" {
				return "Mgr"
			} else {
				if strings.HasSuffix(name, "Entity") {
					name = strings.Before(name, "Entity")
				}
				return name + "Mgr"
			}
		},
		"mgrname": func(name, relationship string) string { // manager struct name
			if relationship == "master" {
				return "Manager"
			} else {
				if strings.HasSuffix(name, "Entity") {
					name = strings.Before(name, "Entity")
				}
				return name + "Manager"
			}
		},
		"dtoname": func(name, relationship string) string { // DTO struct name
			if relationship == "master" {
				return "DTO"
			} else {
				if strings.HasSuffix(name, "Entity") {
					name = strings.Before(name, "Entity")
				}
				return name + "DTO"
			}
		},
		"dtosname": func(name, relationship string) string { // DTOs struct name
			if relationship == "master" {
				return "DTOs"
			} else {
				if strings.HasSuffix(name, "Entity") {
					name = strings.Before(name, "Entity")
				}
				return name + "DTOs"
			}
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
