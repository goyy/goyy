// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package strings_test

import (
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

func ExampleSlice() {
	fmt.Println(strings.Slice("abc", 0, 2))
	fmt.Println(strings.Slice("abc", 2, 4))
	fmt.Println(strings.Slice("abc", -2, -1))
	fmt.Println(strings.Slice("abc", -7, -6))

	// Output:
	// ab
	// c
	// b
	//
}

func ExampleLeft() {
	fmt.Println(strings.Left("abc", 2))
	fmt.Println(strings.Left("abc", 4))

	// Output:
	// ab
	// abc
}

func ExampleRight() {
	fmt.Println(strings.Right("abc", 2))
	fmt.Println(strings.Right("abc", 4))

	// Output:
	// bc
	// abc
}

func ExampleMid() {
	fmt.Println(strings.Mid("abc", 0, 2))
	fmt.Println(strings.Mid("abc", 0, 4))
	fmt.Println(strings.Mid("abc", -2, 2))

	// Output:
	// ab
	// abc
	// bc
}

func ExampleBefore() {
	fmt.Println(strings.Before("abc", "c"))
	fmt.Println(strings.Before("abcba", "b"))

	// Output:
	// ab
	// a
}

func ExampleAfter() {
	fmt.Println(strings.After("abc", "a"))
	fmt.Println(strings.After("abcba", "b"))
	fmt.Println(strings.After("abcba", "e"))

	// Output:
	// bc
	// cba
	//
}

func ExampleBeforeLast() {
	fmt.Println(strings.BeforeLast("abc", "c"))
	fmt.Println(strings.BeforeLast("abcba", "b"))

	// Output:
	// ab
	// abc
}

func ExampleAfterLast() {
	fmt.Println(strings.AfterLast("abc", "a"))
	fmt.Println(strings.AfterLast("abcba", "b"))

	// Output:
	// bc
	// a
}

func ExampleBetween() {
	fmt.Println(strings.Between("yabcz", "y", "z"))
	fmt.Println(strings.Between("yabczydefz", "y", "z"))

	// Output:
	// abc
	// abc
}

func ExampleBetweenSame() {
	fmt.Println(strings.BetweenSame("tagabctag", "tag"))

	// Output: abc
}

func ExampleBetweens() {
	fmt.Println(strings.Betweens("[a][b][c]", "[", "]"))
	fmt.Println(strings.Betweens("1(aa)2(bb)3(cc)4", "(", ")"))

	// Output:
	// [a b c]
	// [aa bb cc]
}
