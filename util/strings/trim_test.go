// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package strings_test

import (
	"gopkg.in/goyy/goyy.v0/util/strings"
	"testing"
)

func TestTrim(t *testing.T) {
	var s = []struct {
		f, in, cutset, out string
	}{
		{"Trim", "abba", "a", "bb"},
		{"Trim", "abba", "ab", ""},
		{"TrimLeft", "abba", "ab", ""},
		{"TrimRight", "abba", "ab", ""},
		{"TrimLeft", "abba", "a", "bba"},
		{"TrimRight", "abba", "a", "abb"},
		{"Trim", "<tag>", "<>", "tag"},
		{"Trim", "* listitem", " *", "listitem"},
		{"Trim", `"quote"`, `"`, "quote"},
		{"Trim", "\u2C6F\u2C6F\u0250\u0250\u2C6F\u2C6F", "\u2C6F", "\u0250\u0250"},
		//empty string tests
		{"Trim", "abba", "", "abba"},
		{"Trim", "", "123", ""},
		{"Trim", "", "", ""},
		{"TrimLeft", "abba", "", "abba"},
		{"TrimLeft", "", "123", ""},
		{"TrimLeft", "", "", ""},
		{"TrimRight", "abba", "", "abba"},
		{"TrimRight", "", "123", ""},
		{"TrimRight", "", "", ""},
		{"TrimRight", "☺\xc0", "☺", "☺\xc0"},
	}
	for _, v := range s {
		name := v.f
		var f func(string, string) string
		switch name {
		case "Trim":
			f = strings.Trim
		case "TrimLeft":
			f = strings.TrimLeft
		case "TrimRight":
			f = strings.TrimRight
		default:
			t.Errorf("Undefined trim function %s", name)
		}
		d := f(v.in, v.cutset)
		if d != v.out {
			t.Errorf("%s(%q, %q) = %q; want %q", name, v.in, v.cutset, d, v.out)
		}
	}
}

func TestTrimSpace(t *testing.T) {
	s := []struct{ in, out string }{
		{"", ""},
		{"abc", "abc"},
		{" abc ", "abc"},
		{" ", ""},
		{" \t\r\n \t\t\r\r\n\n ", ""},
		{" \t\r\n x\t\t\r\r\n\n ", "x"},
		{" \u2000\t\r\n x\t\t\r\r\ny\n \u3000", "x\t\t\r\r\ny"},
		{"1 \t\r\n2", "1 \t\r\n2"},
		{" x\x80", "x\x80"},
		{" x\xc0", "x\xc0"},
		{"x \xc0\xc0 ", "x \xc0\xc0"},
		{"x \xc0", "x \xc0"},
		{"x \xc0 ", "x \xc0"},
		{"x \xc0\xc0 ", "x \xc0\xc0"},
		{"x ☺\xc0\xc0 ", "x ☺\xc0\xc0"},
		{"x ☺ ", "x ☺"},
	}
	for _, v := range s {
		d := strings.TrimSpace(v.in)
		if d != v.out {
			t.Errorf(`TrimSpace("%v") = "%v"; want "%v"`, v.in, d, v.out)
		}
	}
}
