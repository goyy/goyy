// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package strings

import (
	"bytes"
	"strings"
)

// NewReader returns a new Reader reading from s.
// It is similar to bytes.NewBufferString but more efficient and read-only.
func NewReader(s string) *strings.Reader { return strings.NewReader(s) }

// NewReplacer returns a new Replacer from a list of old, new string pairs.
// Replacements are performed in order, without overlapping matches.
func NewReplacer(oldnew ...string) *strings.Replacer { return strings.NewReplacer(oldnew...) }

// Count counts the number of non-overlapping instances of sep in s.
func Count(s, sep string) int { return strings.Count(s, sep) }

// HasPrefix tests whether the string s begins with prefix.
func HasPrefix(s, prefix string) bool { return strings.HasPrefix(s, prefix) }

// HasSuffix tests whether the string s ends with suffix.
func HasSuffix(s, suffix string) bool { return strings.HasSuffix(s, suffix) }

// Join concatenates the elements of a to create a single string.   The separator string
// sep is placed between elements in the resulting string.
func Join(a []string, sep string) string { return strings.Join(a, sep) }

// Map returns a copy of the string s with all its characters modified
// according to the mapping function. If mapping returns a negative value, the character is
// dropped from the string with no replacement.
func Map(mapping func(rune) rune, s string) string { return strings.Map(mapping, s) }

// Replace returns a copy of the string s with the first n
// non-overlapping instances of old replaced by new.
// If n < 0, there is no limit on the number of replacements.
func Replace(s, old, new string, n int) string { return strings.Replace(s, old, new, n) }

// HasAnyPrefix tests whether the string s begins with prefix.
func HasAnyPrefix(s string, prefix ...string) (string, bool) {
	if len(prefix) == 0 {
		return "", true
	}
	for _, v := range prefix {
		if strings.HasPrefix(s, v) {
			return v, true
		}
	}
	return "", false
}

// HasAnySuffix tests whether the string s ends with suffix.
func HasAnySuffix(s string, suffix ...string) (string, bool) {
	if len(suffix) == 0 {
		return "", true
	}
	for _, v := range suffix {
		if strings.HasSuffix(s, v) {
			return v, true
		}
	}
	return "", false
}

// Join concatenates the elements of a to create a single string.   The separator string
// sep is placed between elements in the resulting string.
func JoinIgnoreBlank(a []string, sep string) string {
	var b bytes.Buffer
	var i int
	for _, v := range a {
		if IsNotBlank(v) {
			if i == 0 {
				b.WriteString(v)
			} else {
				b.WriteString(sep + v)
			}
			i++
		}
	}
	return b.String()
}

// Runes returns the []rune value represented by the string
func Runes(s string) []rune {
	rs := make([]rune, 0)
	for _, r := range s {
		rs = append(rs, r)
	}
	return rs
}

// Overlays part of a String with another String.
func Overlay(str, overlay string, start, end int) string {
	if IsBlank(overlay) {
		overlay = ""
	}
	if IsBlank(str) {
		return overlay
	}
	rs := Runes(str)
	rslen := len(rs)
	if start < 0 {
		start = 0
	}
	if start > rslen {
		start = rslen
	}
	if end < 0 {
		end = 0
	}
	if end > rslen {
		end = rslen
	}
	if start > end {
		start, end = end, start
	}
	var b bytes.Buffer
	for i := 0; i < start; i++ {
		b.WriteRune(rs[i])
	}
	b.WriteString(overlay)
	for i := end; i < rslen; i++ {
		b.WriteRune(rs[i])
	}
	return b.String()
}
