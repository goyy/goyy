// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplControllerJsonMain = `package {{.PackageName}}

import (
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

func init() {{"{"}}{{range $i, $e := .Entities}}{{with $name := ctlvar $e.Name $e.Relationship}}
	xhttp.GET({{$name}}.ApiIndex(), {{$name}}.Index)
	xhttp.POST({{$name}}.ApiIndex(), {{$name}}.Index)
	xhttp.GET({{$name}}.ApiShow(), {{$name}}.Show)
	xhttp.POST({{$name}}.ApiAdd(), {{$name}}.Add)
	xhttp.POST({{$name}}.ApiEdit(), {{$name}}.Edit)
	xhttp.POST({{$name}}.ApiSave(), {{$name}}.Save)
	xhttp.POST({{$name}}.ApiDisable(), {{$name}}.Disable){{if eq "tree" $e.Extend}}
	xhttp.GET({{$name}}.ApiTree(), {{$name}}.Tree){{end}}{{end}}{{end}}
}
`
