// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package strings

import (
	"strings"
)

// Contains returns true if substr is within s.
func Contains(s, substr string) bool { return strings.Contains(s, substr) }

// ContainsAny returns true if any Unicode code points in chars are within s.
func ContainsAny(s, chars string) bool { return strings.ContainsAny(s, chars) }

// ContainsRune returns true if the Unicode code point r is within s.
func ContainsRune(s string, r rune) bool { return strings.ContainsRune(s, r) }

// Contains returns true if whitespace is within s.
func ContainsSpace(s string) bool {
	return strings.ContainsAny(s, " \t\r\n")
}

// Checks if the string contains only certain chars.
func ContainsOnly(s, chars string) bool {
	if s == "" || chars == "" {
		return false
	}
	for _, v := range s {
		if ContainsRune(chars, v) == false {
			return false
		}
	}
	return true
}

// Checks that the string does not contain certain chars.
func ContainsNone(s, chars string) bool {
	if s == "" || chars == "" {
		return true
	}
	for _, v := range s {
		if ContainsRune(chars, v) {
			return false
		}
	}
	return true
}

// ContainsSlice returns true if chars is within s.
func ContainsSlice(s string, chars []string) bool {
	if s == "" {
		return false
	}
	if len(chars) == 0 {
		return true
	}
	for _, v := range chars {
		if !Contains(s, v) {
			return false
		}
	}
	return true
}

// ContainsSliceAny returns true if any Unicode code points in chars are within s.
func ContainsSliceAny(s string, chars []string) bool {
	if s == "" {
		return false
	}
	if len(chars) == 0 {
		return true
	}
	for _, v := range chars {
		if Contains(s, v) {
			return true
		}
	}
	return false
}
