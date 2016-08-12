// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package unicodes_test

import (
	"testing"

	"gopkg.in/goyy/goyy.v0/util/unicodes"
)

func TestIsQuote(t *testing.T) {
	s := []struct {
		in       rune
		expected bool
	}{
		{'"', true},
		{'`', true},
		{'\'', true},
		{'a', false},
		{'1', false},
		{'A', false},
		{'一', false},
		{' ', false},
	}
	for _, v := range s {
		if out := unicodes.IsQuote(v.in); out != v.expected {
			t.Errorf(`unicodes.IsQuote("%s") = %v, want %v`, v.in, out, v.expected)
		}
	}
}

func TestIsHan(t *testing.T) {
	s := []struct {
		in       rune
		expected bool
	}{
		{'中', true},
		{'文', true},
		{'一', true},
		{'a', false},
		{'A', false},
		{'1', false},
		{' ', false},
	}
	for _, v := range s {
		if out := unicodes.IsHan(v.in); out != v.expected {
			t.Errorf(`unicodes.IsHan("%s") = %v, want %v`, v.in, out, v.expected)
		}
	}
}
