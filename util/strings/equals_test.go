// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package strings_test

import (
	"gopkg.in/goyy/goyy.v0/util/strings"
	"testing"
)

func TestEqualFold(t *testing.T) {
	s := []struct {
		s, t string
		out  bool
	}{
		{"abc", "abc", true},
		{"ABcd", "ABcd", true},
		{"123abc", "123ABC", true},
		{"αβδ", "ΑΒΔ", true},
		{"abc", "xyz", false},
		{"abc", "XYZ", false},
		{"abcdefghijk", "abcdefghijX", false},
		{"abcdefghijk", "abcdefghij\u212A", true},
		{"abcdefghijK", "abcdefghij\u212A", true},
		{"abcdefghijkz", "abcdefghij\u212Ay", false},
		{"abcdefghijKz", "abcdefghij\u212Ay", false},
	}
	for _, v := range s {
		if out := strings.EqualFold(v.s, v.t); out != v.out {
			t.Errorf("EqualFold(%#q, %#q) = %v, want %v", v.s, v.t, out, v.out)
		}
		if out := strings.EqualFold(v.t, v.s); out != v.out {
			t.Errorf("EqualFold(%#q, %#q) = %v, want %v", v.t, v.s, out, v.out)
		}
	}
}
