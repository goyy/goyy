// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package strings_test

import (
	"gopkg.in/goyy/goyy.v0/util/strings"
	"testing"
)

func TestRemoveSpace(t *testing.T) {
	s := []struct {
		s, out string
	}{
		{" a\nb   c\td e \rf   g    ", "abcdefg"},
	}
	for _, v := range s {
		if out := strings.RemoveSpace(v.s); out != v.out {
			t.Errorf("RemoveSpace(%#q) = %v, want %v", v.s, out, v.out)
		}
	}
}

func TestRemoveBlank(t *testing.T) {
	s := []struct {
		s, out string
	}{
		{" a\nb   c\td e \rf   g    ", "ab cd e f g"},
	}
	for _, v := range s {
		if out := strings.RemoveBlank(v.s); out != v.out {
			t.Errorf("RemoveBlank(%#q) = %v, want %v", v.s, out, v.out)
		}
	}
}

func TestRemoveStart(t *testing.T) {
	s := []struct {
		s, r, out string
	}{
		{"abc", "", "abc"},
		{"www.domain.com", "www.", "domain.com"},
		{"domain.com", "www.", "domain.com"},
		{"www.domain.com", "domain", "www.domain.com"},
	}
	for _, v := range s {
		if out := strings.RemoveStart(v.s, v.r); out != v.out {
			t.Errorf("RemoveStart(%#q, %#q) = %v, want %v", v.s, v.r, out, v.out)
		}
	}
}

func TestRemoveEnd(t *testing.T) {
	s := []struct {
		s, r, out string
	}{
		{"abc", "", "abc"},
		{"www.domain.com", ".com", "www.domain"},
		{"www.domain.com", ".com.", "www.domain.com"},
		{"www.domain.com", "domain", "www.domain.com"},
	}
	for _, v := range s {
		if out := strings.RemoveEnd(v.s, v.r); out != v.out {
			t.Errorf("RemoveEnd(%#q, %#q) = %v, want %v", v.s, v.r, out, v.out)
		}
	}
}

func TestRemove(t *testing.T) {
	s := []struct {
		s, r, out string
	}{
		{"abc", "", "abc"},
		{"queued", "ue", "qd"},
		{"queued", "zz", "queued"},
	}
	for _, v := range s {
		if out := strings.Remove(v.s, v.r); out != v.out {
			t.Errorf("Remove(%#q, %#q) = %v, want %v", v.s, v.r, out, v.out)
		}
	}
}
