// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplEntity = `package {{.Module.Id}}
{{$ := .}}
import (
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/schema"
)

//go:generate xgen -entity=$GOFILE -scaffold -clidir={{.Module.Clidir}} -clipath={{.Module.Clipath}} -apipath={{.Module.Apipath}}

// {{.Name}}实体结构.
// @entity(project:"sys"{{if eq "slave" .Relationship}} relationship:"slave"{{end}})
type Entity struct {
	{{if eq .Parent.Id "pk"}}entity.Pk{{end}}{{if eq .Parent.Id "sys"}}entity.Sys{{end}}{{if eq .Parent.Id "tree"}}entity.Tree{{end}}
	table     schema.Table  ` + "`" + `orm:"table={{.Module.Prefix}}_{{.Id}}&comment={{.Name}}"` + "`" + `{{range $column := .Columns}}{{if not (extends $column.Id $.Parent.Id)}}
	{{padname $column.Field $.AllFieldMaxLen}} {{padname $column.Etype $.AllTypeMaxLen}} ` + "`" + `orm:"column={{$column.Id}}"` + "`" + `{{end}}{{end}}
}
`
