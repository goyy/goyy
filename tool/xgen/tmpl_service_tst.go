// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplServiceTest = `package <%.PackageName%><%with $ := .%>

import (
	"testing"

	_ "<%.TstPath%>"
	"gopkg.in/goyy/goyy.v0/data/domain"
)<%range $i, $e := .Entities%>

func Test<%tstname $e.Name $e.Relationship%>SelectCountBySift(t *testing.T) {
	sIdEQ, _ := domain.NewSift("sIdEQ", "1")
	sDeletionEQ, _ := domain.NewSift("sDeletionEQ", "0")
	out, _ := Mgr.SelectCountBySift(sIdEQ, sDeletionEQ)
	expected := 0
	if out != expected {
		t.Errorf(` + "`" + `SelectCountBySift:"%v", want:"%v"` + "`" + `, out, expected)
	}
}<%end%><%end%>
`
