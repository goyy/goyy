// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package unicodes implements unicode utility functions.
package unicodes

import (
	"unicode"
)

// Is reports whether the rune is in the specified table of ranges.
func Is(rangeTab *unicode.RangeTable, r rune) bool {
	return unicode.Is(rangeTab, r)
}

// IsUpper reports whether the rune is an upper case letter.
func IsUpper(r rune) bool {
	return unicode.IsUpper(r)
}

// IsLower reports whether the rune is a lower case letter.
func IsLower(r rune) bool {
	return unicode.IsLower(r)
}

// IsTitle reports whether the rune is a title case letter.
func IsTitle(r rune) bool {
	return unicode.IsTitle(r)
}

// To maps the rune to the specified case: UpperCase, LowerCase, or TitleCase.
func To(_case int, r rune) rune {
	return unicode.To(_case, r)
}

// ToUpper maps the rune to upper case.
func ToUpper(r rune) rune {
	return unicode.ToUpper(r)
}

// ToLower maps the rune to lower case.
func ToLower(r rune) rune {
	return unicode.ToLower(r)
}

// ToTitle maps the rune to title case.
func ToTitle(r rune) rune {
	return unicode.ToTitle(r)
}

// SimpleFold iterates over Unicode code points equivalent under
// the Unicode-defined simple case folding.  Among the code points
// equivalent to rune (including rune itself), SimpleFold returns the
// smallest rune > r if one exists, or else the smallest rune >= 0.
//
// For example:
//	SimpleFold('A') = 'a'
//	SimpleFold('a') = 'A'
//
//	SimpleFold('K') = 'k'
//	SimpleFold('k') = '\u212A' (Kelvin symbol, K)
//	SimpleFold('\u212A') = 'K'
//
//	SimpleFold('1') = '1'
//
func SimpleFold(r rune) rune {
	return unicode.SimpleFold(r)
}
