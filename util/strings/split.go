// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package strings

import (
	"strings"
	"unicode"

	"gopkg.in/goyy/goyy.v0/util/unicodes"
)

// Split slices s into all substrings separated by sep and returns a slice of
// the substrings between those separators.
// If sep is empty, Split splits after each UTF-8 sequence.
// It is equivalent to SplitN with a count of -1.
func Split(s, sep string) []string { return strings.Split(s, sep) }

// SplitAfter slices s into all substrings after each instance of sep and
// returns a slice of those substrings.
// If sep is empty, SplitAfter splits after each UTF-8 sequence.
// It is equivalent to SplitAfterN with a count of -1.
func SplitAfter(s, sep string) []string { return strings.SplitAfter(s, sep) }

// SplitAfterN slices s into substrings after each instance of sep and
// returns a slice of those substrings.
// If sep is empty, SplitAfterN splits after each UTF-8 sequence.
// The count determines the number of substrings to return:
//   n > 0: at most n substrings; the last substring will be the unsplit remainder.
//   n == 0: the result is nil (zero substrings)
//   n < 0: all substrings
func SplitAfterN(s, sep string, n int) []string { return strings.SplitAfterN(s, sep, n) }

// SplitN slices s into substrings separated by sep and returns a slice of
// the substrings between those separators.
// If sep is empty, SplitN splits after each UTF-8 sequence.
// The count determines the number of substrings to return:
//   n > 0: at most n substrings; the last substring will be the unsplit remainder.
//   n == 0: the result is nil (zero substrings)
//   n < 0: all substrings
func SplitN(s, sep string, n int) []string { return strings.SplitN(s, sep, n) }

// Fields splits the string s around each instance of one or more consecutive white space
// characters, returning an array of substrings of s or an empty list if s contains only white space.
func Fields(s string) []string { return strings.Fields(s) }

// FieldsFunc splits the string s at each run of Unicode code points c satisfying f(c)
// and returns an array of slices of s. If all code points in s satisfy f(c) or the
// string is empty, an empty slice is returned.
func FieldsFunc(s string, f func(rune) bool) []string { return strings.FieldsFunc(s, f) }

// FieldsSpace splits the string s around each instance of one or more consecutive white space
// characters, returning an array of substrings of s or an empty list if s contains only white space.
func FieldsSpace(s string) []string {
	// First count the fields.
	n := 0
	inField := false
	isQuote := false
	var quote rune
	isContiue := func(r rune) bool {
		if unicodes.IsQuote(r) {
			if isQuote {
				if quote == r {
					isQuote = false
					return true
				}
			} else {
				isQuote = true
				quote = r
				return false
			}
		}
		if isQuote {
			return true
		}
		return false
	}
	for _, r := range s {
		if isContiue(r) {
			continue
		}
		wasInField := inField
		inField = !unicode.IsSpace(r)
		if inField && !wasInField {
			n++
		}
	}

	// Now create them.
	a := make([]string, n)
	na := 0
	fieldStart := -1 // Set to -1 when looking for start of field.
	isQuote = false
	quote = ' '
	for i, r := range s {
		if isContiue(r) {
			continue
		}
		if unicode.IsSpace(r) {
			if fieldStart >= 0 {
				a[na] = s[fieldStart:i]
				na++
				fieldStart = -1
			}
		} else if fieldStart == -1 {
			fieldStart = i
		}
	}
	if fieldStart >= 0 { // Last field might end at EOF.
		a[na] = s[fieldStart:]
	}
	return a
}
