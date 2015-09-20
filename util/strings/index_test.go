// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package strings_test

import (
	"gopkg.in/goyy/goyy.v0/util/strings"
	"testing"
)

func TestIndex(t *testing.T) {
	s := []struct {
		in, sep string
		out     int
	}{
		{"", "", 0},
		{"", "a", -1},
		{"", "foo", -1},
		{"fo", "foo", -1},
		{"foo", "foo", 0},
		{"oofofoofooo", "f", 2},
		{"oofofoofooo", "foo", 4},
		{"barfoobarfoo", "foo", 3},
		{"foo", "", 0},
		{"foo", "o", 1},
		{"abcABCabc", "A", 3},
		// cases with one byte strings - test special case in Index()
		{"", "a", -1},
		{"x", "a", -1},
		{"x", "x", 0},
		{"abc", "a", 0},
		{"abc", "b", 1},
		{"abc", "c", 2},
		{"abc", "x", -1},
	}
	for _, v := range s {
		d := strings.Index(v.in, v.sep)
		if d != v.out {
			t.Errorf(`Index("%v", "%v") = %v; want %v`, v.in, v.sep, d, v.out)
		}
	}
}

func TestIndexLast(t *testing.T) {
	s := []struct {
		in, sep string
		out     int
	}{
		{"", "", 0},
		{"", "a", -1},
		{"", "foo", -1},
		{"fo", "foo", -1},
		{"foo", "foo", 0},
		{"foo", "f", 0},
		{"oofofoofooo", "f", 7},
		{"oofofoofooo", "foo", 7},
		{"barfoobarfoo", "foo", 9},
		{"foo", "", 3},
		{"foo", "o", 2},
		{"abcABCabc", "A", 3},
		{"abcABCabc", "a", 6},
	}
	for _, v := range s {
		d := strings.IndexLast(v.in, v.sep)
		if d != v.out {
			t.Errorf(`IndexLast("%v", "%v") = %v; want %v`, v.in, v.sep, d, v.out)
		}
	}
}

func TestIndexLastStart(t *testing.T) {
	s := []struct {
		in, sep    string
		start, out int
	}{
		{"", "", 0, 0},
		{"", "a", 0, -1},
		{"", "foo", 0, -1},
		{"fo", "foo", 0, -1},
		{"foo", "foo", 0, 0},
		{"foo", "f", 0, 0},
		{"oofofoofooo", "f", 4, 4},
		{"oofofoofooo", "foo", 5, 4},
		{"barfoobarfoo", "foo", 3, 3},
		{"foo", "", 0, 3},
		{"foo", "o", 0, 2},
		{"abcABCabc", "A", 0, 3},
		{"abcABCabc", "a", 3, 0},
	}
	for _, v := range s {
		d := strings.IndexLastStart(v.in, v.sep, v.start)
		if d != v.out {
			t.Errorf(`IndexLastStart("%v", "%v", %v) = %v; want %v`, v.in, v.sep, v.start, d, v.out)
		}
	}
}

func TestIndexLastOrdinal(t *testing.T) {
	s := []struct {
		in, sep      string
		ordinal, out int
	}{
		{"", "", 1, 0},
		{"", "a", 1, -1},
		{"", "foo", 1, -1},
		{"fo", "foo", 1, -1},
		{"foo", "foo", 1, 0},
		{"foo", "f", 1, 0},
		{"oofofoofooo", "f", 2, 4},
		{"oofofoofooo", "foo", 2, 4},
		{"barfoobarfoo", "foo", 2, 3},
		{"foo", "", 1, 3},
		{"foo", "o", 1, 2},
		{"abcABCabc", "A", 1, 3},
		{"abcABCabc", "a", 2, 0},
	}
	for _, v := range s {
		d := strings.IndexLastOrdinal(v.in, v.sep, v.ordinal)
		if d != v.out {
			t.Errorf(`IndexLastOrdinal("%v", "%v", %v) = %v; want %v`, v.in, v.sep, v.ordinal, d, v.out)
		}
	}
}

func TestIndexStart(t *testing.T) {
	s := []struct {
		in, sep    string
		start, out int
	}{
		{"", "", 0, 0},
		{"", "a", 0, -1},
		{"", "foo", 0, -1},
		{"fo", "foo", 0, -1},
		{"foo", "foo", 0, 0},
		{"foo", "f", 0, 0},
		{"oofofoofooo", "f", 3, 4},
		{"oofofoofooo", "foo", 5, 7},
		{"barfoobarfoo", "foo", 4, 9},
		{"foo", "", 0, 0},
		{"foo", "o", 0, 1},
		{"abcABCabc", "A", 0, 3},
		{"abcABCabc", "a", 3, 6},
	}
	for _, v := range s {
		d := strings.IndexStart(v.in, v.sep, v.start)
		if d != v.out {
			t.Errorf(`IndexStart("%v", "%v", %v) = %v; want %v`, v.in, v.sep, v.start, d, v.out)
		}
	}
}

func TestIndexForward(t *testing.T) {
	s := []struct {
		in, sep    string
		start, out int
	}{
		{"", "", 9, 0},
		{"", "a", 9, -1},
		{"", "foo", 9, -1},
		{"fo", "foo", 9, -1},
		{"foo", "foo", 9, 0},
		{"foo", "f", 9, 0},
		{"oofofoofooo", "f", 6, 4},
		{"oofofoofooo", "foo", 10, 7},
		{"barfoobarfoo", "foo", 8, 3},
		{"foo", "", 0, 0},
		{"foo", "o", 2, 2},
		{"abcABCabc", "A", 8, 3},
		{"abcABCabc", "a", 8, 6},
	}
	for _, v := range s {
		d := strings.IndexForward(v.in, v.sep, v.start)
		if d != v.out {
			t.Errorf(`IndexForward("%v", "%v", %v) = %v; want %v`, v.in, v.sep, v.start, d, v.out)
		}
	}
}

func TestIndexOrdinal(t *testing.T) {
	s := []struct {
		in, sep      string
		ordinal, out int
	}{
		{"", "", 1, 0},
		{"", "a", 1, -1},
		{"", "foo", 1, -1},
		{"fo", "foo", 1, -1},
		{"foo", "foo", 1, 0},
		{"foo", "f", 1, 0},
		{"oofofoofooo", "f", 2, 4},
		{"oofofoofooo", "foo", 2, 7},
		{"barfoobarfoo", "foo", 2, 9},
		{"foo", "", 1, 0},
		{"foo", "o", 1, 1},
		{"abcABCabc", "A", 1, 3},
		{"abcABCabc", "a", 2, 6},
	}
	for _, v := range s {
		d := strings.IndexOrdinal(v.in, v.sep, v.ordinal)
		if d != v.out {
			t.Errorf(`IndexOrdinal("%v", "%v", %v) = %v; want %v`, v.in, v.sep, v.ordinal, d, v.out)
		}
	}
}
