// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package bytes implements functions for the manipulation of byte slices.
// It is analogous to the facilities of the strings package.
package bytes

import (
	"bytes"
)

// TrimRightNul trim \x00 character at the end
func TrimRightNul(src []byte) []byte {
	return bytes.TrimRight(src, "\x00")
}

// IsHex to determine whether the hex character
func IsHex(c byte) bool {
	switch {
	case '0' <= c && c <= '9':
		return true
	case 'a' <= c && c <= 'f':
		return true
	case 'A' <= c && c <= 'F':
		return true
	}
	return false
}

// NewBuffer creates a new Buffer.
func NewBuffer() bytes.Buffer {
	return bytes.Buffer{}
}
