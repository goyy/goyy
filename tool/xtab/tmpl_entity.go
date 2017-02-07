// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplEntity = `package {{if blank .Master}}{{.ID}}{{else}}{{.Master}}{{end}}
{{$ := .}}
import (
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/schema"
)

//go:generate xgen -entity=$GOFILE -scaffold{{if notblank .Module.Admpath}} -admpath={{.Module.Admpath}}{{end}} -apipath={{.Module.Apipath}} -tstpath={{.Module.Tstpath}}

// {{.Name}}` + i18N.Message("domain.title") + `.
// @entity(module:"{{.ID}}" project:"{{.Module.ID}}"{{if notblank .Slave}} relationship:"slave"{{end}})
type {{if notblank .Master}}{{camel .Slave}}{{end}}Entity struct {
	{{if eq .Super "pk"}}entity.Pk{{end}}{{if eq .Super "sys"}}entity.Sys{{end}}{{if eq .Super "tree"}}entity.Tree{{end}}
	{{padname "table" $.FieldMaxLen}} {{padname "schema.Table" $.TypeMaxLen}} ` + "`" + `db:"table={{.Module.Prefix}}_{{.ID}}&comment={{.Name}}"` + "`" + `{{range $column := .Columns}}{{if not (supercol $column.ID $.Super)}}
	{{padname $column.Field $.FieldMaxLen}} {{padname $column.Etype $.TypeMaxLen}} ` + "`" + `db:"column={{$column.ID}}&comment={{$column.Name}}{{if notblank $column.Defaults}}&default={{$column.Defaults}}{{end}}"` + "`" + `{{end}}{{end}}
}
`
