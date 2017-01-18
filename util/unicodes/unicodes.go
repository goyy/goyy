// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package unicodes implements unicode utility functions.
package unicodes

import (
	"unicode"
)

// IsQuote reports whether the rune is a quote.
// Returns true when the value of r is ' or " or `, otherwise it returns false.
func IsQuote(r rune) bool {
	switch r {
	case '\'':
		return true
	case '"':
		return true
	case '`':
		return true
	}
	return false
}

// IsHan reports whether the rune is a Chinese characters.
func IsHan(r rune) bool {
	return unicode.Is(unicode.Scripts["Han"], r)
}
