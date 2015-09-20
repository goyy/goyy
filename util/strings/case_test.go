// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package strings_test

import (
	"gopkg.in/goyy/goyy.v0/util/strings"
	"testing"
)

func TestToLower(t *testing.T) {
	s := []struct {
		in, out string
	}{
		{"", ""},
		{"abc", "abc"},
		{"AbC123", "abc123"},
		{"azAZ09_", "azaz09_"},
		{"\u2C6D\u2C6D\u2C6D\u2C6D\u2C6D", "\u0251\u0251\u0251\u0251\u0251"},
	}
	for _, v := range s {
		if out := strings.ToLower(v.in); out != v.out {
			t.Errorf("ToLower(%#q) = %#q, want %#q", v.in, out, v.out)
		}
	}
}

func TestToLowerFirst(t *testing.T) {
	s := []struct {
		in, out string
	}{
		{"", ""},
		{"abc", "abc"},
		{"AbC123", "abC123"},
		{"azAZ09_", "azAZ09_"},
		{"\u2C6D\u2C6D\u2C6D\u2C6D\u2C6D", "\u0251\u2C6D\u2C6D\u2C6D\u2C6D"},
	}
	for _, v := range s {
		if out := strings.ToLowerFirst(v.in); out != v.out {
			t.Errorf("ToLowerFirst(%#q) = %#q, want %#q", v.in, out, v.out)
		}
	}
}

func TestToUpper(t *testing.T) {
	s := []struct {
		in, out string
	}{
		{"", ""},
		{"abc", "ABC"},
		{"AbC123", "ABC123"},
		{"azAZ09_", "AZAZ09_"},
		{"\u0250\u0250\u0250\u0250\u0250", "\u2C6F\u2C6F\u2C6F\u2C6F\u2C6F"},
	}
	for _, v := range s {
		if out := strings.ToUpper(v.in); out != v.out {
			t.Errorf("ToUpper(%#q) = %#q, want %#q", v.in, out, v.out)
		}
	}
}

func TestToUpperFirst(t *testing.T) {
	s := []struct {
		in, out string
	}{
		{"", ""},
		{"abc", "Abc"},
		{"AbC123", "AbC123"},
		{"azAZ09_", "AzAZ09_"},
		{"\u0250\u0250\u0250\u0250\u0250", "\u2C6F\u0250\u0250\u0250\u0250"},
	}
	for _, v := range s {
		if out := strings.ToUpperFirst(v.in); out != v.out {
			t.Errorf("ToUpperFirst(%#q) = %#q, want %#q", v.in, out, v.out)
		}
	}
}

func TestTitle(t *testing.T) {
	s := []struct {
		in, out string
	}{
		{"", ""},
		{"a", "A"},
		{"cat", "Cat"},
		{"cAt", "CAt"},
		{" aaa aaa aaa ", " Aaa Aaa Aaa "},
		{" Aaa Aaa Aaa ", " Aaa Aaa Aaa "},
		{"123a456", "123a456"},
		{"double-blind", "Double-Blind"},
		{"ÿøû", "Ÿøû"},
	}
	for _, v := range s {
		if out := strings.Title(v.in); out != v.out {
			t.Errorf("Title(%#q) = %#q, want %#q", v.in, out, v.out)
		}
	}
}

func TestToTitle(t *testing.T) {
	s := []struct {
		in, out string
	}{
		{"", ""},
		{"a", "A"},
		{" aaa aaa aaa ", " AAA AAA AAA "},
		{" Aaa Aaa Aaa ", " AAA AAA AAA "},
		{"123a456", "123A456"},
		{"double-blind", "DOUBLE-BLIND"},
		{"ÿøû", "ŸØÛ"},
	}
	for _, v := range s {
		if out := strings.ToTitle(v.in); out != v.out {
			t.Errorf("ToTitle(%#q) = %#q, want %#q", v.in, out, v.out)
		}
	}
}

func TestCamel(t *testing.T) {
	s := []struct {
		in, out string
	}{
		{"", ""},
		{"a", "A"},
		{"cat", "Cat"},
		{"cAt", "Cat"},
		{" aaa aaa aaa ", "AaaAaaAaa"},
		{"_Aaa_Aaa_Aaa ", "AaaAaaAaa"},
		{"123a456", "123a456"},
		{"douBle-blind", "DoubleBlind"},
		{"ÿøû", "Ÿøû"},
	}
	for _, v := range s {
		if out := strings.Camel(v.in); out != v.out {
			t.Errorf("Camel(%#q) = %#q, want %#q", v.in, out, v.out)
		}
	}
}

func TestUnCamel(t *testing.T) {
	s := []struct {
		in, out string
	}{
		{"", ""},
		{"A", "a"},
		{"Cat", "cat"},
		{"cAt", "c_at"},
		{"AaaAaaAaa", "aaa_aaa_aaa"},
		{" AaaAaaAaa ", "aaa_aaa_aaa"},
		{"123a456", "123a456"},
		{"DoubleBlind", "double_blind"},
		{"Ÿøû", "ÿøû"},
	}
	for _, v := range s {
		if out := strings.UnCamel(v.in, "_"); out != v.out {
			t.Errorf("UnCamel(%#q) = %#q, want %#q", v.in, out, v.out)
		}
	}
}
