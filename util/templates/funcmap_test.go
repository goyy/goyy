// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package templates_test

import (
	"bytes"
	"gopkg.in/goyy/goyy.v0/util/templates"
	"testing"
)

const (
	text = `{{if exist . "name" }}true{{else}}false{{end}}`
	sql  = `select * from demo where 1=1 
		{{if exist . "name" }}and name like :name {{end}}
		{{if exist . "createdAt" }}and created_at > :createdAt {{end}}
		{{if eq .updatedAt "2014-03-19" }}and updated_at < :updatedAt {{end}}
		order by id`
)

func TestExist1(t *testing.T) {
	buf := new(bytes.Buffer)
	data := map[string]interface{}{"name": "goyy"}
	tmpl, _ := templates.New("tmpl").Parse(text)
	tmpl.Execute(buf, data)
	expected := "true"
	if out := buf.String(); out != expected {
		t.Errorf(`out = "%s", want "%s"`, out, expected)
	}
}

func TestExist2(t *testing.T) {
	buf := new(bytes.Buffer)
	data := map[string]interface{}{"memo": "goyy"}
	tmpl, _ := templates.New("tmpl").Parse(text)
	tmpl.Execute(buf, data)
	expected := "false"
	if out := buf.String(); out != expected {
		t.Errorf(`out = "%s", want "%s"`, out, expected)
	}
}
