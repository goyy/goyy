// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package strings_test

import (
	"gopkg.in/goyy/goyy.v0/util/strings"
	"testing"
)

func TestAbbr(t *testing.T) {
	s := []struct {
		in, out string
		width   int
	}{
		{"", "", 6},
		{"a", "a", 6},
		{"abc", "abc", 6},
		{"AbC123", "AbC123", 6},
		{"azAZ09_", "azA...", 6},
		{"世界hello", "世界h...", 6},
	}
	for _, v := range s {
		if out := strings.Abbr(v.in, v.width); out != v.out {
			t.Errorf("Abbr(%#q, %d) = %#q, want %#q", v.in, v.width, out, v.out)
		}
	}
}

func TestAnon(t *testing.T) {
	s := []struct {
		in, out string
	}{
		{"", ""},
		{"a", "a"},
		{"ac", "a***c"},
		{"abc", "a***c"},
		{"AbC123", "A***3"},
		{"azAZ09_", "a***_"},
		{"世界hello", "世***o"},
	}
	for _, v := range s {
		if out := strings.Anon(v.in); out != v.out {
			t.Errorf("Anon(%#q) = %#q, want %#q", v.in, out, v.out)
		}
	}
}

func TestAnonymity(t *testing.T) {
	s := []struct {
		in, out string
	}{
		{"", ""},
		{"abc", "abc"},
		{"15566668888", "155****8888"},
		{"abcdefghijk", "abc****hijk"},
		{"abcdefghijklmn", "abc****klmn"},
	}
	for _, v := range s {
		if out := strings.Anonymity(v.in); out != v.out {
			t.Errorf("Anonymity(%#q) = %#q, want %#q", v.in, out, v.out)
		}
	}
}

func TestAnonymous(t *testing.T) {
	s := []struct {
		in, out           string
		left, right, star int
	}{
		{"", "", 1, 1, 0},
		{"a", "a", -1, -1, 10},
		{"abc", "a***c", -1, -1, 3},
		{"abcd", "abcd", 5, 5, 3},
		{"AbC123", "A****3", 1, 1, 0},
		{"azAZ09_", "a***_", 1, 1, 3},
		{"世界hello", "世***lo", 1, 2, 3},
	}
	for _, v := range s {
		if out := strings.Anonymous(v.in, v.left, v.right, v.star); out != v.out {
			t.Errorf("Anonymous(%#q, %d, %d, %d) = %#q, want %#q", v.in, v.left, v.right, v.star, out, v.out)
		}
	}
}
