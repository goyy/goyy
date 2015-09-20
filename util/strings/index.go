// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package strings

import (
	"strings"
)

// Index returns the index of the first instance of sep in s, or -1 if sep is not present in s.
func Index(s, sep string) int { return strings.Index(s, sep) }

// IndexAny returns the index of the first instance of any Unicode code point
// from chars in s, or -1 if no Unicode code point from chars is present in s.
func IndexAny(s, chars string) int { return strings.IndexAny(s, chars) }

// IndexByte returns the index of the first instance of c in s, or -1 if c is not present in s.
func IndexByte(s string, c byte) int { return strings.IndexByte(s, c) }

// IndexFunc returns the index into s of the first Unicode
// code point satisfying f(c), or -1 if none do.
func IndexFunc(s string, f func(rune) bool) int { return strings.IndexFunc(s, f) }

// IndexRune returns the index of the first instance of the Unicode code point
// r, or -1 if rune is not present in s.
func IndexRune(s string, r rune) int { return strings.IndexRune(s, r) }

// IndexLast returns the index of the last instance of sep in s, or -1 if sep is not present in s.
func IndexLast(s, sep string) int { return strings.LastIndex(s, sep) }

// IndexLastAny returns the index of the last instance of any Unicode code
// point from chars in s, or -1 if no Unicode code point from chars is
// present in s.
func IndexLastAny(s, chars string) int { return strings.LastIndexAny(s, chars) }

// IndexLastFunc returns the index into s of the last
// Unicode code point satisfying f(c), or -1 if none do.
func IndexLastFunc(s string, f func(rune) bool) int { return strings.LastIndexFunc(s, f) }

// IndexLastStart returns the index of the last instance of sep in s from a start position,
// or -1 if sep is not present in s.
func IndexLastStart(s, sep string, start int) int {
	if start < 0 {
		start = 0
	}
	n := len(sep)
	size := len(s)
	if n == 0 {
		return size
	}
	if start >= size {
		return -1
	}
	c := sep[0]
	if n == 1 {
		// special case worth making fast
		for i := size - 1 - start; i >= 0; i-- {
			if s[i] == c {
				return i
			}
		}
		return -1
	}
	// n > 1
	for i := size - 1 - start; i >= 0; i-- {
		if i+n >= size {
			if s[i] == c && s[i:] == sep {
				return i
			}
		} else {
			if s[i] == c && s[i:i+n] == sep {
				return i
			}
		}
	}
	return -1
}

// IndexLastOrdinal returns the index of the n-th last instance of sep in s,
// or -1 if sep is not present in s.
func IndexLastOrdinal(s, sep string, ordinal int) int {
	times := 0
	if ordinal < 1 {
		ordinal = 1
	}
	n := len(sep)
	size := len(s)
	if n == 0 {
		return size
	}
	c := sep[0]
	if n == 1 {
		// special case worth making fast
		for i := size - 1; i >= 0; i-- {
			if s[i] == c {
				times++
				if times == ordinal {
					return i
				}
			}
		}
		return -1
	}
	// n > 1
	for i := size - n; i >= 0; i-- {
		if s[i] == c && s[i:i+n] == sep {
			times++
			if times == ordinal {
				return i
			}
		}
	}
	return -1
}

// IndexStart returns the index of the first instance of sep in s from a start position,
// or -1 if sep is not present in s.
func IndexStart(s, sep string, start int) int {
	if start < 0 {
		start = 0
	}
	n := len(sep)
	size := len(s)
	if n == 0 {
		return 0
	}
	c := sep[0]
	if n == 1 {
		// special case worth making fast
		for i := start; i < size; i++ {
			if s[i] == c {
				return i
			}
		}
		return -1
	}
	// n > 1
	for i := start; i+n <= size; i++ {
		if s[i] == c && s[i:i+n] == sep {
			return i
		}
	}
	return -1
}

// IndexForward returns the forward index of the first instance of sep in s from a start position,
// or -1 if sep is not present in s.
func IndexForward(s, sep string, start int) int {
	size := len(s)
	n := len(sep)
	if size == 0 && n > 0 {
		return -1
	}
	if n == 0 {
		return 0
	}
	if start < 0 || start >= size {
		start = size - 1
	}
	c := sep[0]
	if n == 1 {
		// special case worth making fast
		for i := start; i >= 0; i-- {
			if s[i] == c {
				return i
			}
		}
		return -1
	}
	// n > 1
	if size < n {
		return -1
	}
	for i := start; i >= 0; i-- {
		if s[i] == c && s[i:i+n] == sep {
			return i
		}
	}
	return -1
}

// IndexOrdinal returns the index of the n-th instance of sep in s, or -1 if sep is not present in s.
func IndexOrdinal(s, sep string, ordinal int) int {
	times := 0
	if ordinal < 0 {
		ordinal = 1
	}
	n := len(sep)
	size := len(s)
	if n == 0 {
		return 0
	}
	c := sep[0]
	if n == 1 {
		// special case worth making fast
		for i := 0; i < size; i++ {
			if s[i] == c {
				times++
				if times == ordinal {
					return i
				}
			}
		}
		return -1
	}
	// n > 1
	for i := 0; i+n <= size; i++ {
		if s[i] == c && s[i:i+n] == sep {
			times++
			if times == ordinal {
				return i
			}
		}
	}
	return -1
}
