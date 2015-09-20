// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package strings

import (
	"strings"
)

// Trim returns a slice of the string s with all leading and
// trailing Unicode code points contained in cutset removed.
func Trim(s string, cutset string) string { return strings.Trim(s, cutset) }

// TrimFunc returns a slice of the string s with all leading
// and trailing Unicode code points c satisfying f(c) removed.
func TrimFunc(s string, f func(rune) bool) string { return strings.TrimFunc(s, f) }

// TrimLeft returns a slice of the string s with all leading
// Unicode code points contained in cutset removed.
func TrimLeft(s string, cutset string) string { return strings.TrimLeft(s, cutset) }

// TrimLeftFunc returns a slice of the string s with all leading
// Unicode code points c satisfying f(c) removed.
func TrimLeftFunc(s string, f func(rune) bool) string { return strings.TrimLeftFunc(s, f) }

// TrimRight returns a slice of the string s, with all trailing
// Unicode code points contained in cutset removed.
func TrimRight(s string, cutset string) string { return strings.TrimRight(s, cutset) }

// TrimRightFunc returns a slice of the string s with all trailing
// Unicode code points c satisfying f(c) removed.
func TrimRightFunc(s string, f func(rune) bool) string { return strings.TrimRightFunc(s, f) }

// TrimSpace returns a slice of the string s, with all leading
// and trailing white space removed, as defined by Unicode.
func TrimSpace(s string) string { return strings.TrimSpace(s) }

// TrimPrefix returns s without the provided leading prefix string.
// If s doesn't start with prefix, s is returned unchanged.
func TrimPrefix(s, prefix string) string { return strings.TrimPrefix(s, prefix) }

// TrimSuffix returns s without the provided trailing suffix string.
// If s doesn't end with suffix, s is returned unchanged.
func TrimSuffix(s, suffix string) string { return strings.TrimSuffix(s, suffix) }
