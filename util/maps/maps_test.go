// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package maps_test

import (
	"fmt"
	"testing"

	"gopkg.in/goyy/goyy.v0/util/maps"
)

func TestParseURLQuery(t *testing.T) {
	s := []struct {
		in       string
		expected string
	}{
		{"key1=value1&key2=value2", "map[key1:value1 key2:value2]"},
		{"key1=value1", "map[key1:value1]"},
		{"key1=", "map[key1:]"},
		{"=value1", "map[]"},
		{"=", "map[]"},
		{"", "map[]"},
		{" ", "map[]"},
	}
	for _, v := range s {
		if out := maps.ParseURLQuery(v.in); fmt.Sprint(out) != v.expected {
			t.Errorf(`maps.ParseURLQuery("%s") = %v, want %v`, v.in, out, v.expected)
		}
	}
}
