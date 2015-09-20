// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package strings

import (
	"strings"
)

const padLimit = 8192

// Repeat returns a new string consisting of count copies of the string s.
func Repeat(s string, count int) string { return strings.Repeat(s, count) }

// Left pad a String with spaces.
func PadStart(s string, size int) string {
	return PadLeft(s, size, " ")
}

// Right pad a String with spaces.
func PadEnd(s string, size int) string {
	return PadRight(s, size, " ")
}

// Left pad a String with a specified character.
func PadLeft(in string, size int, sep string) string {
	if IsBlank(sep) {
		sep = " "
	}
	sepLen := len(sep)
	inLen := len(in)
	pads := size - inLen
	if pads <= 0 {
		return in
	}
	if sepLen == 1 && pads <= padLimit {
		return strings.Repeat(sep, pads) + in
	}
	if pads == sepLen {
		return sep + in
	} else if pads < sepLen {
		return sep[0:pads] + in
	} else {
		var padding string
		for i := 0; i < pads; i++ {
			padding += string(sep[i%sepLen])
		}
		return padding + in
	}
	return in
}

// Right pad a String with a specified character.
func PadRight(in string, size int, sep string) string {
	if IsBlank(sep) {
		sep = " "
	}
	sepLen := len(sep)
	inLen := len(in)
	pads := size - inLen
	if pads <= 0 {
		return in
	}
	if sepLen == 1 && pads <= padLimit {
		return in + strings.Repeat(sep, pads)
	}
	if pads == sepLen {
		return in + sep
	} else if pads < sepLen {
		return in + sep[0:pads]
	} else {
		var padding string
		for i := 0; i < pads; i++ {
			padding += string(sep[i%sepLen])
		}
		return in + padding
	}
	return in
}

// Pads a string in a larger string of size using the space character (' ').
func Pad(in string, size int) string {
	return Center(in, size, " ")
}

// Centers a string in a larger string of size.
func Center(in string, size int, sep string) (out string) {
	if IsBlank(sep) {
		sep = " "
	}
	inLen := len(in)
	pads := size - inLen
	if pads <= 0 {
		out = in
		return
	}
	out = PadLeft(in, inLen+pads/2, sep)
	out = PadRight(out, size, sep)
	return
}
