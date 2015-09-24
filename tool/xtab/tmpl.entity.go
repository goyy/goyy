// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplEntity = `package {{.Id}}
{{$ := .}}
import (
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/schema"
)

//go:generate xgen -entity=$GOFILE -scaffold -clidir={{.Module.Clidir}} -clipath={{.Module.Clipath}} -apipath={{.Module.Apipath}}

// {{.Name}}实体结构.
// @entity(project:"{{.Module.Id}}"{{if eq "slave" .Relationship}} relationship:"slave"{{end}})
type Entity struct {
	{{if eq .Super "pk"}}entity.Pk{{end}}{{if eq .Super "sys"}}entity.Sys{{end}}{{if eq .Super "tree"}}entity.Tree{{end}}
	{{padname "table" $.FieldMaxLen}} {{padname "schema.Table" $.TypeMaxLen}} ` + "`" + `orm:"table={{.Module.Prefix}}_{{.Id}}&comment={{.Name}}"` + "`" + `{{range $column := .Columns}}{{if not (supercol $column.Id $.Super)}}
	{{padname $column.Field $.FieldMaxLen}} {{padname $column.Etype $.TypeMaxLen}} ` + "`" + `orm:"column={{$column.Id}}"` + "`" + `{{end}}{{end}}
}
`
