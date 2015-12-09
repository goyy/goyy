// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplControllerClientMain = `package {{.PackageName}}

import (
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

func init() {{"{"}}{{range $i, $e := .Entities}}{{with $name := ctlvar $e.Name $e.Relationship}}
	xhttp.GET({{$name}}.PathIndex(), {{$name}}.Index)
	xhttp.POST({{$name}}.PathIndex(), {{$name}}.Index)
	xhttp.GET({{$name}}.PathShow(), {{$name}}.Show)
	xhttp.POST({{$name}}.PathAdd(), {{$name}}.Add)
	xhttp.POST({{$name}}.PathEdit(), {{$name}}.Edit)
	xhttp.POST({{$name}}.PathSave(), {{$name}}.Save)
	xhttp.POST({{$name}}.PathDisable(), {{$name}}.Disable){{if eq "tree" $e.Extend}}
	xhttp.GET({{$name}}.PathTree(), {{$name}}.Tree){{end}}{{end}}{{end}}
}
`
