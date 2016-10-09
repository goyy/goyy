// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplControllerTest = `package <%.PackageName%><%with $ := .%>

import (
	"testing"

	_ "<%.Tstpath%>"
	"gopkg.in/goyy/goyy.v0/test/assert"
)<%range $i, $e := .Entities%>

func Test<%tstname $e.Name $e.Relationship%>ControllerIndex(t *testing.T) {
	if !assert.HTTPSuccess(t, ctl.Index, "GET", ctl.ApiIndex(), nil) {
		t.Errorf(` + "`" + `GET: %s: Fail` + "`" + `, ctl.ApiIndex())
	}
}<%end%><%end%>
`
