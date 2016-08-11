// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplControllerHtmlMain = `package {{.PackageName}}

import (
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

func init() {{"{"}}{{range $i, $e := .Entities}}{{with $name := ctlvar $e.Name $e.Relationship}}
	xhttp.GET({{$name}}.PathIndex(), {{$name}}.Index, {{$name}}.PermitView())
	xhttp.POST({{$name}}.PathIndex(), {{$name}}.Index, {{$name}}.PermitView())
	xhttp.GET({{$name}}.PathShow(), {{$name}}.Show, {{$name}}.PermitView())
	xhttp.POST({{$name}}.PathShow(), {{$name}}.Show, {{$name}}.PermitView())
	xhttp.POST({{$name}}.PathAdd(), {{$name}}.Add, {{$name}}.PermitAdd())
	xhttp.POST({{$name}}.PathEdit(), {{$name}}.Edit, {{$name}}.PermitEdit())
	xhttp.POST({{$name}}.PathSave(), {{$name}}.Save, {{$name}}.PermitAdd(), {{$name}}.PermitEdit())
	xhttp.POST({{$name}}.PathDisable(), {{$name}}.Disable, {{$name}}.PermitDisable()){{if eq "tree" $e.Extend}}
	xhttp.GET({{$name}}.PathTree(), {{$name}}.Tree, {{$name}}.PermitView()){{end}}{{end}}{{end}}
}
`
