// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package strings

import (
	"strings"
)

// IsBlank checks if a string is empty ("") or whitespace.
func IsBlank(s string) bool {
	return strings.TrimSpace(s) == ""
}

// IsAnyBlank checks if any one of the string are empty ("") or whitespace.
func IsAnyBlank(s ...string) bool {
	for _, c := range s {
		if IsBlank(c) {
			return true
		}
	}
	return false
}

// IsNoneBlank checks if none of the string are empty ("") or whitespace
func IsNoneBlank(s ...string) bool {
	for _, c := range s {
		if IsBlank(c) {
			return false
		}
	}
	return true
}

// IsNotBlank checks if a string is not empty ("") and not whitespace only.
func IsNotBlank(s string) bool {
	return strings.TrimSpace(s) != ""
}
