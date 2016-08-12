// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package unicodes implements unicode utility functions.
package unicodes

import (
	"unicode"
)

// IsGraphic reports whether the rune is defined as a Graphic by Unicode.
// Such characters include letters, marks, numbers, punctuation, symbols, and
// spaces, from categories L, M, N, P, S, Zs.
func IsGraphic(r rune) bool {
	return unicode.IsGraphic(r)
}

// IsPrint reports whether the rune is defined as printable by Go. Such
// characters include letters, marks, numbers, punctuation, symbols, and the
// ASCII space character, from categories L, M, N, P, S and the ASCII space
// character.  This categorization is the same as IsGraphic except that the
// only spacing character is ASCII space, U+0020.
func IsPrint(r rune) bool {
	return unicode.IsPrint(r)
}

// IsOneOf reports whether the rune is a member of one of the ranges.
// The function "In" provides a nicer signature and should be used in preference to IsOneOf.
func IsOneOf(ranges []*unicode.RangeTable, r rune) bool {
	return unicode.IsOneOf(ranges, r)
}

// In reports whether the rune is a member of one of the ranges.
func In(r rune, ranges ...*unicode.RangeTable) bool {
	return unicode.In(r, ranges...)
}

// IsControl reports whether the rune is a control character.
// The C (Other) Unicode category includes more code points
// such as surrogates; use Is(C, r) to test for them.
func IsControl(r rune) bool {
	return unicode.IsControl(r)
}

// IsLetter reports whether the rune is a letter (category L).
func IsLetter(r rune) bool {
	return unicode.IsLetter(r)
}

// IsMark reports whether the rune is a mark character (category M).
func IsMark(r rune) bool {
	return unicode.IsMark(r)
}

// IsNumber reports whether the rune is a number (category N).
func IsNumber(r rune) bool {
	return unicode.IsNumber(r)
}

// IsPunct reports whether the rune is a Unicode punctuation character
// (category P).
func IsPunct(r rune) bool {
	return unicode.IsPunct(r)
}

// IsSpace reports whether the rune is a space character as defined
// by Unicode's White Space property; in the Latin-1 space
// this is
//	'\t', '\n', '\v', '\f', '\r', ' ', U+0085 (NEL), U+00A0 (NBSP).
// Other definitions of spacing characters are set by category
// Z and property Pattern_White_Space.
func IsSpace(r rune) bool {
	return unicode.IsSpace(r)
}

// IsSymbol reports whether the rune is a symbolic character.
func IsSymbol(r rune) bool {
	return unicode.IsSymbol(r)
}
