// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package jsons_test

import (
	"gopkg.in/goyy/goyy.v0/util/jsons"
	"testing"
)

func TestFormat(t *testing.T) {
	in := `ab\"cd
		ef g`
	expected := `ab\\\"cd\n\t\tef g`
	if out := jsons.Format(in); out != expected {
		t.Errorf(`Format("%s") = %s, want %s`, in, out, expected)
	}
}

func TestParse(t *testing.T) {
	in := `ab\\\"cd\n\t\tef g`
	expected := `ab\"cd
		ef g`
	if out := jsons.Parse(in); out != expected {
		t.Errorf(`Parse("%s") = %s, want %s`, in, out, expected)
	}
}
