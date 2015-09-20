// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package strings_test

import (
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

func ExampleAbbr() {
	s := []struct {
		in    string
		width int
	}{
		{"", 6},
		{"a", 6},
		{"abc", 6},
		{"AbC123", 6},
		{"azAZ09_", 6},
		{"世界hello", 6},
	}
	for _, v := range s {
		fmt.Println(strings.Abbr(v.in, v.width))
	}

	// Output:
	//
	// a
	// abc
	// AbC123
	// azA...
	// 世界h...
}

func ExampleAnon() {
	s := []string{"", "a", "ac", "abc", "AbC123", "azAZ09_", "世界hello"}
	for _, v := range s {
		fmt.Println(strings.Anon(v))
	}

	// Output:
	//
	// a
	// a***c
	// a***c
	// A***3
	// a***_
	// 世***o
}

func ExampleAnonymity() {
	s := []string{"", "abc", "15566668888", "abcdefghijk", "abcdefghijklmn"}
	for _, v := range s {
		fmt.Println(strings.Anonymity(v))
	}

	// Output:
	//
	// abc
	// 155****8888
	// abc****hijk
	// abc****klmn
}

func ExampleAnonymous() {
	s := []struct {
		in                string
		left, right, star int
	}{
		{"", 1, 1, 0},
		{"a", -1, -1, 10},
		{"abc", -1, -1, 3},
		{"abcd", 5, 5, 3},
		{"AbC123", 1, 1, 0},
		{"azAZ09_", 1, 1, 3},
		{"世界hello", 1, 2, 3},
	}
	for _, v := range s {
		fmt.Println(strings.Anonymous(v.in, v.left, v.right, v.star))
	}

	// Output:
	//
	// a
	// a***c
	// abcd
	// A****3
	// a***_
	// 世***lo
}
