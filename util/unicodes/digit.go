// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package unicodes implements unicode utility functions.
package unicodes

import (
	"unicode"
)

// IsDigit reports whether the rune is a decimal digit.
func IsDigit(r rune) bool {
	return unicode.IsDigit(r)
}
