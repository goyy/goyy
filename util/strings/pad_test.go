// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package strings_test

import (
	"gopkg.in/goyy/goyy.v0/util/strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	s := []struct {
		in, out string
		count   int
	}{
		{"", "", 0},
		{"", "", 1},
		{"", "", 2},
		{"-", "", 0},
		{"-", "-", 1},
		{"-", "----------", 10},
		{"abc ", "abc abc abc ", 3},
	}
	for _, v := range s {
		if out := strings.Repeat(v.in, v.count); out != v.out {
			t.Errorf("Repeat(%#q, %d) = %#q; want %#q", v.in, v.count, out, v.out)
			continue
		}
	}
}

func TestPadStart(t *testing.T) {
	s := []struct {
		s    string
		size int
		out  string
	}{
		{"", 3, "   "},
		{"bat", 3, "bat"},
		{"bat", 5, "  bat"},
		{"bat", 1, "bat"},
		{"bat", -1, "bat"},
	}
	for _, v := range s {
		if out := strings.PadStart(v.s, v.size); out != v.out {
			t.Errorf("PadStart(%#q, %v) = %#q, want %#q", v.s, v.size, out, v.out)
		}
	}
}

func TestPadEnd(t *testing.T) {
	s := []struct {
		s    string
		size int
		out  string
	}{
		{"", 3, "   "},
		{"bat", 3, "bat"},
		{"bat", 5, "bat  "},
		{"bat", 1, "bat"},
		{"bat", -1, "bat"},
	}
	for _, v := range s {
		if out := strings.PadEnd(v.s, v.size); out != v.out {
			t.Errorf("PadEnd(%#q, %v) = %#q, want %#q", v.s, v.size, out, v.out)
		}
	}
}

func TestPadLeft(t *testing.T) {
	s := []struct {
		s    string
		size int
		sep  string
		out  string
	}{
		{"", 3, "*", "***"},
		{"bat", 3, "*", "bat"},
		{"bat", 5, "*", "**bat"},
		{"bat", 1, "*", "bat"},
		{"bat", -1, "*", "bat"},
	}
	for _, v := range s {
		if out := strings.PadLeft(v.s, v.size, v.sep); out != v.out {
			t.Errorf("PadLeft(%#q, %v, %#q) = %#q, want %#q", v.s, v.size, v.sep, out, v.out)
		}
	}
}

func TestPadRight(t *testing.T) {
	s := []struct {
		s    string
		size int
		sep  string
		out  string
	}{
		{"", 3, "*", "***"},
		{"bat", 3, "*", "bat"},
		{"bat", 5, "*", "bat**"},
		{"bat", 1, "*", "bat"},
		{"bat", -1, "*", "bat"},
	}
	for _, v := range s {
		if out := strings.PadRight(v.s, v.size, v.sep); out != v.out {
			t.Errorf("PadRight(%#q, %v, %#q) = %#q, want %#q", v.s, v.size, v.sep, out, v.out)
		}
	}
}

func TestPad(t *testing.T) {
	s := []struct {
		s    string
		size int
		out  string
	}{
		{"", 3, "   "},
		{"ab", -1, "ab"},
		{"ab", 4, " ab "},
		{"abcd", 2, "abcd"},
		{"a", 4, " a  "},
	}
	for _, v := range s {
		if out := strings.Pad(v.s, v.size); out != v.out {
			t.Errorf("Pad(%#q, %v) = %#q, want %#q", v.s, v.size, out, v.out)
		}
	}
}

func TestCenter(t *testing.T) {
	s := []struct {
		s    string
		size int
		sep  string
		out  string
	}{
		{"", 3, "*", "***"},
		{"ab", -1, "*", "ab"},
		{"ab", 4, "*", "*ab*"},
		{"abcd", 2, "*", "abcd"},
		{"a", 4, "*", "*a**"},
		{"a", 4, "yz", "yayz"},
	}
	for _, v := range s {
		if out := strings.Center(v.s, v.size, v.sep); out != v.out {
			t.Errorf("Center(%#q, %v, %#q) = %#q, want %#q", v.s, v.size, v.sep, out, v.out)
		}
	}
}
