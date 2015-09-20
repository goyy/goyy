// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package strings

import (
	"strings"
)

// Removes all whitespaces from a string.
func RemoveSpace(s string) (d string) {
	if IsBlank(s) {
		return
	}
	d = s
	sps := [...]string{"\t", "\n", "\v", "\f", "\r", " "}
	for _, sp := range sps {
		d = strings.Replace(d, sp, "", -1)
	}
	return
}

// Removes all whitespaces from a string.
func RemoveBlank(s string) (d string) {
	if IsBlank(s) {
		return
	}
	d = strings.TrimSpace(s)
	sps := [...]string{"\t", "\n", "\v", "\f", "\r", "  "}
	for _, sp := range sps {
		d = strings.Replace(d, sp, "", -1)
	}
	return
}

// Removes a substring only if it is at the beginning of a source string,
// otherwise returns the source string.
func RemoveStart(s, remove string) string {
	if IsBlank(s) || IsBlank(remove) {
		return s
	}
	if strings.HasPrefix(s, remove) {
		return s[len(remove):]
	}
	return s
}

// Removes a substring only if it is at the end of a source string,
// otherwise returns the source string.
func RemoveEnd(s, remove string) string {
	if IsBlank(s) || IsBlank(remove) {
		return s
	}
	if strings.HasSuffix(s, remove) {
		return s[0 : len(s)-len(remove)]
	}
	return s
}

// Removes all occurrences of a substring from within the source string.
func Remove(s, remove string) string {
	if IsBlank(s) || IsBlank(remove) {
		return s
	}
	return strings.Replace(s, remove, "", -1)
}
