// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package strings

import (
	"bytes"
	"strings"
	"unicode"
)

// ToLower returns a copy of the string s with all Unicode letters mapped to their lower case.
func ToLower(s string) string { return strings.ToLower(s) }

// ToLowerFirst returns a copy of the string s with first Unicode letters mapped to their lower case.
func ToLowerFirst(s string) string {
	if s == "" {
		return s
	}
	if len(s) == 1 {
		return strings.ToLower(s)
	}
	var b bytes.Buffer
	for i, v := range s {
		if i == 0 {
			b.WriteString(strings.ToLower(string(v)))
		} else {
			b.WriteRune(v)
		}
	}
	return b.String()
}

// ToLowerSpecial returns a copy of the string s with all Unicode letters mapped to their
// lower case, giving priority to the special casing rules.
func ToLowerSpecial(c unicode.SpecialCase, s string) string {
	return strings.ToLowerSpecial(c, s)
}

// ToUpper returns a copy of the string s with all Unicode letters mapped to their upper case.
func ToUpper(s string) string { return strings.ToUpper(s) }

// ToUpperFirst returns a copy of the string s with first Unicode letters mapped to their upper case.
func ToUpperFirst(s string) string {
	if s == "" {
		return s
	}
	if len(s) == 1 {
		return strings.ToUpper(s)
	}
	var b bytes.Buffer
	for i, v := range s {
		if i == 0 {
			b.WriteString(strings.ToUpper(string(v)))
		} else {
			b.WriteRune(v)
		}
	}
	return b.String()
}

// ToUpperSpecial returns a copy of the string s with all Unicode letters mapped to their
// upper case, giving priority to the special casing rules.
func ToUpperSpecial(c unicode.SpecialCase, s string) string {
	return strings.ToUpperSpecial(c, s)
}

// Title returns a copy of the string s with all Unicode letters that begin words
// mapped to their title case.
func Title(s string) string { return strings.Title(s) }

// ToTitle returns a copy of the string s with all Unicode letters mapped to their title case.
func ToTitle(s string) string { return strings.ToTitle(s) }

// ToTitleSpecial returns a copy of the string s with all Unicode letters mapped to their
// title case, giving priority to the special casing rules.
func ToTitleSpecial(c unicode.SpecialCase, s string) string {
	return strings.ToTitleSpecial(c, s)
}

// Camel returns a copy of the string s with all Unicode letters that begin words
// mapped to their camel case.
func Camel(s string) string {
	prev := true
	result := Map(
		func(r rune) rune {
			if isSeparator(r) {
				prev = true
				return '_'
			} else {
				if prev {
					prev = false
					return unicode.ToTitle(r)
				}
			}
			prev = false
			return unicode.ToLower(r)
		},
		s)
	return Remove(result, "_")
}

// UnCamel returns a copy of the string s with all Unicode letters that begin words
// mapped to their uncamel case.
func UnCamel(s, sep string) string {
	var b bytes.Buffer
	for i, r := range strings.TrimSpace(s) {
		if unicode.IsUpper(r) {
			if i > 0 {
				b.WriteString(sep)
			}
			b.WriteRune(unicode.ToLower(r))
		} else {
			b.WriteRune(r)
		}
		i++
	}
	return b.String()
}

// isSeparator reports whether the rune could mark a word boundary.
// TODO: update when package unicode captures more of the properties.
func isSeparator(r rune) bool {
	// ASCII alphanumerics and underscore are not separators
	if r <= 0x7F {
		switch {
		case '0' <= r && r <= '9':
			return false
		case 'a' <= r && r <= 'z':
			return false
		case 'A' <= r && r <= 'Z':
			return false
		case r == '_':
			return true
		case r == '-':
			return true
		}
		return true
	}
	// Letters and digits are not separators
	if unicode.IsLetter(r) || unicode.IsDigit(r) {
		return false
	}
	// Otherwise, all we can do for now is treat spaces as separators.
	return unicode.IsSpace(r)
}
