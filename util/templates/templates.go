// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package templates

import (
	"bytes"
	"text/template"

	"gopkg.in/goyy/goyy.v0/util/strings"
)

// New allocates a new template with the given name.
func New(name string) *template.Template {
	return template.New(name).Funcs(funcMapText)
}

// Process compile template and output.
func Process(tmpl string, data interface{}) (r string, err error) {
	buf := new(bytes.Buffer)
	t, err := New("goyy-temples").Parse(tmpl)
	if err != nil {
		return
	}
	err = t.Execute(buf, data)
	if err != nil {
		return
	}
	r = strings.RemoveBlank(buf.String())
	return
}
