// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/goyy/goyy.v0/util/files"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

func genEntity() {
	for _, p := range conf.projects {
		for _, t := range conf.tables {
			if t.module.project.ID() == p.ID() && t.module.project.generate == "true" && t.module.generate == "true" && t.module.ID() != "sys" && t.generate == "true" {
				apidir := "../" + strings.AfterLast(t.module.Apipath(), "/")
				dir := apidir + "/internal/" + t.id + "/"
				dstfile := filepath.Join(dir, t.id+"_entity.go")
				if strings.IsNotBlank(t.master) {
					dir = apidir + "/internal/" + t.master + "/"
					dstfile = filepath.Join(dir, t.slave+"_entity.go")
				}
				if files.IsExist(dstfile) {
					continue
				} else {
					files.MkdirAll(dir, 0755)
				}
				buf := bytes.Buffer{}
				tmpl := newTmpl(tmplEntity)
				tmpl.Execute(&buf, t)
				ioutil.WriteFile(dstfile, buf.Bytes(), 0755)
				write(tmplEntityBat, dir+"generate.bat")
				write(tmplEntitySh, dir+"generate.sh")
			}
		}
	}
}
