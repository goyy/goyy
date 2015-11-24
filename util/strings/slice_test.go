// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package strings_test

import (
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"testing"
)

func TestSlice(t *testing.T) {
	s := []struct {
		s          string
		start, end int
		out        string
	}{
		{"", 0, 2, ""},
		{"abc", 0, 2, "ab"},
		{"abc", 2, 0, ""},
		{"abc", 2, 4, "c"},
		{"abc", 4, 6, ""},
		{"abc", 2, 2, ""},
		{"abc", -2, -1, "b"},
		{"abc", -4, 2, "ab"},
	}
	for _, v := range s {
		if out := strings.Slice(v.s, v.start, v.end); out != v.out {
			t.Errorf("Slice(%#q, %v, %v) = %#q, want %#q", v.s, v.start, v.end, out, v.out)
		}
	}
}

func TestLeft(t *testing.T) {
	s := []struct {
		s   string
		l   int
		out string
	}{
		{"", 2, ""},
		{"abc", 0, ""},
		{"abc", 2, "ab"},
		{"abc", 4, "abc"},
	}
	for _, v := range s {
		if out := strings.Left(v.s, v.l); out != v.out {
			t.Errorf("Left(%#q, %v) = %#q, want %#q", v.s, v.l, out, v.out)
		}
	}
}

func TestRight(t *testing.T) {
	s := []struct {
		s   string
		l   int
		out string
	}{
		{"", 2, ""},
		{"abc", 0, ""},
		{"abc", 2, "bc"},
		{"abc", 4, "abc"},
	}
	for _, v := range s {
		if out := strings.Right(v.s, v.l); out != v.out {
			t.Errorf("Right(%#q, %v) = %#q, want %#q", v.s, v.l, out, v.out)
		}
	}
}

func TestMid(t *testing.T) {
	s := []struct {
		s        string
		start, l int
		out      string
	}{
		{"", 0, 2, ""},
		{"abc", 0, 2, "ab"},
		{"abc", 0, 4, "abc"},
		{"abc", 2, 4, "c"},
		{"abc", 4, 2, ""},
		{"abc", -2, 2, "bc"},
	}
	for _, v := range s {
		if out := strings.Mid(v.s, v.start, v.l); out != v.out {
			t.Errorf("Mid(%#q, %v, %v) = %#q, want %#q", v.s, v.start, v.l, out, v.out)
		}
	}
}

func TestBefore(t *testing.T) {
	s := []struct {
		s, sep, out string
	}{
		{"", "", ""},
		{"", "a", ""},
		{"abc", "", "abc"},
		{"abc", "a", ""},
		{"abcba", "b", "a"},
		{"abc", "c", "ab"},
		{"abc", "d", ""},
	}
	for _, v := range s {
		if out := strings.Before(v.s, v.sep); out != v.out {
			t.Errorf("Before(%#q, %#q) = %#q, want %#q", v.s, v.sep, out, v.out)
		}
	}
}

func TestAfter(t *testing.T) {
	s := []struct {
		s, sep, out string
	}{
		{"", "", ""},
		{"", "a", ""},
		{"abc", "", "abc"},
		{"abc", "a", "bc"},
		{"abcba", "b", "cba"},
		{"abc", "c", ""},
		{"abc", "d", ""},
	}
	for _, v := range s {
		if out := strings.After(v.s, v.sep); out != v.out {
			t.Errorf("After(%#q, %#q) = %#q, want %#q", v.s, v.sep, out, v.out)
		}
	}
}

func TestBeforeLast(t *testing.T) {
	s := []struct {
		s, sep, out string
	}{
		{"", "", ""},
		{"", "a", ""},
		{"abc", "", "abc"},
		{"abc", "a", ""},
		{"abcba", "b", "abc"},
		{"abc", "c", "ab"},
		{"abc", "d", ""},
	}
	for _, v := range s {
		if out := strings.BeforeLast(v.s, v.sep); out != v.out {
			t.Errorf("BeforeLast(%#q, %#q) = %#q, want %#q", v.s, v.sep, out, v.out)
		}
	}
}

func TestAfterLast(t *testing.T) {
	s := []struct {
		s, sep, out string
	}{
		{"", "", ""},
		{"", "a", ""},
		{"abc", "", "abc"},
		{"abc", "a", "bc"},
		{"abcba", "b", "a"},
		{"abc", "c", ""},
		{"abc", "d", ""},
	}
	for _, v := range s {
		if out := strings.AfterLast(v.s, v.sep); out != v.out {
			t.Errorf("AfterLast(%#q, %#q) = %#q, want %#q", v.s, v.sep, out, v.out)
		}
	}
}

func TestBetween(t *testing.T) {
	s := []struct {
		s, start, end, out string
	}{
		{"yabcz", "", "", ""},
		{"yabcz", "y", "z", "abc"},
		{"yabczydefz", "y", "z", "abc"},
	}
	for _, v := range s {
		if out := strings.Between(v.s, v.start, v.end); out != v.out {
			t.Errorf("Between(%#q, %#q, %#q) = %#q, want %#q", v.s, v.start, v.end, out, v.out)
		}
	}
}

func TestBetweenSame(t *testing.T) {
	s := []struct {
		s, tag, out string
	}{
		{"", "", ""},
		{"", "tag", ""},
		{"tagabctag", "", ""},
		{"tagabctag", "tag", "abc"},
	}
	for _, v := range s {
		if out := strings.BetweenSame(v.s, v.tag); out != v.out {
			t.Errorf("BetweenSame(%#q, %#q) = %#q, want %#q", v.s, v.tag, out, v.out)
		}
	}
}

func TestBetweens(t *testing.T) {
	s := []struct {
		s, start, end, out string
	}{
		{"[a][b][c]", "[", "]", `[a b c]`},
		{"1(aa)2(bb)3(cc)4", "(", ")", `[aa bb cc]`},
		{" 1${id}2 3${name} ${pwd}4", "${", "}", `[id name pwd]`},
		{"[a][b][c]", "", "", "[]"},
		{"", "[", "]", "[]"},
		{"", "", "", "[]"},
	}
	for _, v := range s {
		if out := strings.Betweens(v.s, v.start, v.end); fmt.Sprintf("%s", out) != v.out {
			t.Errorf("Betweens(%#q, %#q, %#q) = %#q, want %#q", v.s, v.start, v.end, fmt.Sprintf("%s", out), v.out)
		}
	}
}
