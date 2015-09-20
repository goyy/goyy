// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package strings

import (
	"strings"
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
