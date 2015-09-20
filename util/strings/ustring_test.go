// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package strings_test

import (
	"gopkg.in/goyy/goyy.v0/util/strings"
	"testing"
)

func TestJoinIgnoreBlank(t *testing.T) {
	s := []struct {
		in       []string
		sep, out string
	}{
		{nil, "a", ""},
		{[]string{"", "bcd"}, "a", "bcd"},
		{[]string{"abcd"}, "z", "abcd"},
		{[]string{"a", "b", "c", "d"}, "", "abcd"},
		{[]string{"1", "2", "3", "4"}, ",", "1,2,3,4"},
		{[]string{"1", ".2", ".3", ".4"}, "...", "1....2....3....4"},
		{[]string{"☺☻", ""}, "☹", "☺☻"},
		{[]string{"☺☻☹"}, "~", "☺☻☹"},
		{[]string{"☺", "☻", "☹"}, "", "☺☻☹"},
		{[]string{"1", "2", "3 4"}, " ", "1 2 3 4"},
		{[]string{"1", "2"}, " ", "1 2"},
		{[]string{"1", "23"}, "", "123"},
		{[]string{"1", "2", "3"}, "", "123"},
	}
	for _, v := range s {
		got := strings.JoinIgnoreBlank(v.in, v.sep)
		if got != v.out {
			t.Errorf(`JoinIgnoreBlank("%v","%v") = "%v"; want "%v"`, v.in, v.sep, got, v.out)
		}
	}
}

func TestOverlay(t *testing.T) {
	s := []struct {
		in, overlay, out string
		start, end       int
	}{
		{"", "abc", "abc", 0, 0},
		{"abcdef", "", "abef", 2, 4},
		{"abcdef", "", "abef", 4, 2},
		{"abcdef", "zzzz", "abzzzzef", 2, 4},
		{"abcdef", "zzzz", "abzzzzef", 4, 2},
		{"abcdef", "zzzz", "zzzzef", -1, 4},
		{"abcdef", "zzzz", "abzzzz", 2, 8},
		{"abcdef", "zzzz", "zzzzabcdef", -2, -3},
		{"abcdef", "zzzz", "abcdefzzzz", 8, 10},
		{"世界hello", "zzzz", "世界zzzzllo", 2, 4},
	}
	for _, v := range s {
		got := strings.Overlay(v.in, v.overlay, v.start, v.end)
		if got != v.out {
			t.Errorf(`Overlays("%v","%v","%v","%v") = "%v"; want "%v"`, v.in, v.overlay, v.start, v.end, got, v.out)
		}
	}
}
