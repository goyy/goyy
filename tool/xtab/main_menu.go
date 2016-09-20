// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"io/ioutil"

	"gopkg.in/goyy/goyy.v0/util/files"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

func genMenu() {
	for _, p := range conf.projects {
		clidir := "../" + strings.AfterLast(p.Clipath(), "/")
		dir := clidir + "/templates/core/include/"
		dstfile := dir + "header.html"
		if !files.IsExist(dstfile) {
			files.MkdirAll(dir, 0644)
		}
		data := map[string]interface{}{
			"Project": p,
			"Modules": conf.modules,
			"Tables":  conf.tables,
		}
		buf := bytes.Buffer{}
		tmpl := newTmpl(tmplMenu)
		tmpl.Execute(&buf, data)
		ioutil.WriteFile(dstfile, buf.Bytes(), 0644)
	}
}
