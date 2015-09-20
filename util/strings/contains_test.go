// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package strings_test

import (
	"gopkg.in/goyy/goyy.v0/util/strings"
	"testing"
)

func TestContains(t *testing.T) {
	s := []struct {
		str, substr string
		expected    bool
	}{
		{"abc", "bc", true},
		{"abc", "bcd", false},
		{"abc", "", true},
		{"", "a", false},
	}
	for _, v := range s {
		if strings.Contains(v.str, v.substr) != v.expected {
			t.Errorf(`Contains("%s", "%s") = %v, want %v`, v.str, v.substr, !v.expected, v.expected)
		}
	}
}

func TestContainsAny(t *testing.T) {
	s := []struct {
		str, substr string
		expected    bool
	}{
		{"", "", false},
		{"", "a", false},
		{"", "abc", false},
		{"a", "", false},
		{"a", "a", true},
		{"aaa", "a", true},
		{"abc", "xyz", false},
		{"abc", "xcz", true},
		{"a☺b☻c☹d", "uvw☻xyz", true},
		{"aRegExp*", ".(|)*+?^$[]", true},
	}
	for _, v := range s {
		if strings.ContainsAny(v.str, v.substr) != v.expected {
			t.Errorf(`ContainsAny("%s", "%s") = %v, want %v`, v.str, v.substr, !v.expected, v.expected)
		}
	}
}

func TestContainsRune(t *testing.T) {
	s := []struct {
		str      string
		r        rune
		expected bool
	}{
		{"", 'a', false},
		{"a", 'a', true},
		{"aaa", 'a', true},
		{"abc", 'y', false},
		{"abc", 'c', true},
		{"a☺b☻c☹d", 'x', false},
		{"a☺b☻c☹d", '☻', true},
		{"aRegExp*", '*', true},
	}
	for _, v := range s {
		if strings.ContainsRune(v.str, v.r) != v.expected {
			t.Errorf(`ContainsRune("%s", "%s") = %v, want %v`, v.str, v.r, !v.expected, v.expected)
		}
	}
}

func TestContainsSpace(t *testing.T) {
	s := []struct {
		str      string
		expected bool
	}{
		{"", false},
		{"a", false},
		{" a", true},
		{"ab c", true},
		{"ab\tc", true},
		{"ab\rc", true},
		{"ab\nc", true},
	}
	for _, v := range s {
		if strings.ContainsSpace(v.str) != v.expected {
			t.Errorf(`ContainsSpace("%s") = %v, want %v`, v.str, !v.expected, v.expected)
		}
	}
}

func TestContainsOnly(t *testing.T) {
	s := []struct {
		str, substr string
		expected    bool
	}{
		{"", "", false},
		{"ab", "", false},
		{"abab", "abc", true},
		{"ab1", "abc", false},
		{"abz", "abc", false},
	}
	for _, v := range s {
		if strings.ContainsOnly(v.str, v.substr) != v.expected {
			t.Errorf(`ContainsOnly("%s", "%s") = %v, want %v`, v.str, v.substr, !v.expected, v.expected)
		}
	}
}

func TestContainsNone(t *testing.T) {
	s := []struct {
		str, substr string
		expected    bool
	}{
		{"", "", true},
		{"ab", "", true},
		{"abab", "xyz", true},
		{"ab1", "xyz", true},
		{"abz", "xyz", false},
	}
	for _, v := range s {
		if strings.ContainsNone(v.str, v.substr) != v.expected {
			t.Errorf(`ContainsNone("%s", "%s") = %v, want %v`, v.str, v.substr, !v.expected, v.expected)
		}
	}
}
