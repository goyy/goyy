// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package envs_test

import (
	"fmt"
	"testing"

	"gopkg.in/goyy/goyy.v0/util/envs"
)

func TestParseGOPATH(t *testing.T) {
	s := []struct {
		in       string
		expected string
	}{
		{"%GOPATH%/src/gopkg.in/goyy/goyy.v0", "e:\\gopath/src/gopkg.in/goyy/goyy.v0"},
		{"/src/gopkg.in/goyy/goyy.v0", "/src/gopkg.in/goyy/goyy.v0"},
		{"", ""},
		{" ", ""},
	}
	for _, v := range s {
		if out := envs.ParseGOPATH(v.in); fmt.Sprint(out) != v.expected {
			t.Errorf(`envs.ParseGOPATH("%s") = %v, want %v`, v.in, out, v.expected)
		}
	}
}
