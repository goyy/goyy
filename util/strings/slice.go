// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package strings

// The returned substring starts with the character in the start position
// and ends before the end position.
func Slice(s string, start, end int) string {
	if IsBlank(s) {
		return s
	}
	size := len(s)
	// handle negatives
	if end < 0 {
		end = size + end // remember end is negative
	}
	if start < 0 {
		start = size + start // remember start is negative
	}
	// check length next
	if end > size {
		end = size
	}
	// if start is greater than end, return ""
	if start > end {
		return ""
	}
	if start < 0 {
		start = 0
	}
	if end < 0 {
		end = 0
	}
	return s[start:end]
}

// Gets the leftmost len characters of a String.
func Left(s string, length int) string {
	if length < 0 || IsBlank(s) {
		return ""
	}
	if len(s) <= length {
		return s
	}
	return s[0:length]
}

// Gets the rightmost len characters of a String.
func Right(s string, length int) string {
	if length < 0 || IsBlank(s) {
		return ""
	}
	if len(s) <= length {
		return s
	}
	return s[len(s)-length:]
}

// Gets len characters from the middle of a String.
func Mid(s string, start, length int) string {
	if length < 0 || IsBlank(s) {
		return ""
	}
	size := len(s)
	if start < 0 {
		start = size + start
	}
	if start < 0 {
		start = 0
	}
	if start > size {
		return ""
	}
	if size <= start+length {
		return s[start:]
	}
	return s[start : start+length]
}

// Gets the substring before the first occurrence of a separator. The separator is not returned.
func Before(s, sep string) string {
	if IsBlank(s) || IsBlank(sep) {
		return s
	}
	pos := Index(s, sep)
	if pos == -1 {
		return ""
	}
	return s[0:pos]
}

// Gets the substring after the first occurrence of a separator. The separator is not returned.
func After(s, sep string) string {
	if IsBlank(s) || IsBlank(sep) {
		return s
	}
	pos := Index(s, sep)
	if pos == -1 {
		return ""
	}
	return s[pos+len(sep):]
}

// Gets the substring before the last occurrence of a separator. The separator is not returned.
func BeforeLast(s, sep string) string {
	if IsBlank(s) || IsBlank(sep) {
		return s
	}
	pos := IndexLast(s, sep)
	if pos == -1 {
		return ""
	}
	return s[0:pos]
}

// Gets the substring after the last occurrence of a separator. The separator is not returned.
func AfterLast(s, sep string) string {
	if IsBlank(s) || IsBlank(sep) {
		return s
	}
	pos := IndexLast(s, sep)
	if pos == -1 || pos == len(s)-len(sep) {
		return ""
	}
	return s[pos+len(sep):]
}

// Gets the String that is nested in between two Strings. Only the first match is returned.
func Between(s, start, end string) string {
	if IsBlank(s) || IsBlank(start) || IsBlank(end) {
		return ""
	}
	b := Index(s, start)
	if b != -1 {
		e := IndexStart(s, end, b+len(start))
		if e != -1 {
			return s[b+len(start) : e]
		}
	}
	return ""
}

// Gets the String that is nested in between two instances of the same String.
func BetweenSame(s, tag string) string {
	return Between(s, tag, tag)
}
