// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package strings_test

import (
	"fmt"

	"gopkg.in/goyy/goyy.v0/util/strings"
)

func ExampleContains() {
	fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.Contains("seafood", "bar"))
	fmt.Println(strings.Contains("seafood", ""))
	fmt.Println(strings.Contains("", ""))
	fmt.Println(strings.Contains("", "foo"))

	// Output:
	// true
	// false
	// true
	// true
	// false
}

func ExampleContainsAny() {
	fmt.Println(strings.ContainsAny("team", "i"))
	fmt.Println(strings.ContainsAny("failure", "u & i"))
	fmt.Println(strings.ContainsAny("foo", ""))
	fmt.Println(strings.ContainsAny("", ""))

	// Output:
	// false
	// true
	// false
	// false
}

func ExampleContainsRune() {
	fmt.Println(strings.ContainsRune("team", 'i'))
	fmt.Println(strings.ContainsRune("failure", 'u'))

	// Output:
	// false
	// true
}

func ExampleContainsSpace() {
	fmt.Println(strings.ContainsSpace("a"))
	fmt.Println(strings.ContainsSpace(" a "))
	fmt.Println(strings.ContainsSpace("ab c"))
	fmt.Println(strings.ContainsSpace("ab\tc"))
	fmt.Println(strings.ContainsSpace("ab\rc"))
	fmt.Println(strings.ContainsSpace("ab\nc"))

	// Output:
	// false
	// true
	// true
	// true
	// true
	// true
}

func ExampleContainsOnly() {
	fmt.Println(strings.ContainsOnly("abab", "abc"))
	fmt.Println(strings.ContainsOnly("ab1", "abc"))
	fmt.Println(strings.ContainsOnly("abz", "abc"))

	// Output:
	// true
	// false
	// false
}

func ExampleContainsNone() {
	fmt.Println(strings.ContainsNone("abab", "xyz"))
	fmt.Println(strings.ContainsNone("ab1", "xyz"))
	fmt.Println(strings.ContainsNone("abz", "xyz"))

	// Output:
	// true
	// true
	// false
}

func ExampleContainsSlice() {
	fmt.Println(strings.ContainsSlice("abcd", []string{"ab", "cd"}))
	fmt.Println(strings.ContainsSlice("ab1", []string{"ab", "12"}))
	fmt.Println(strings.ContainsSlice("abz", []string{"xy", "cd"}))
	fmt.Println(strings.ContainsSlice("abz", []string{"", ""}))
	fmt.Println(strings.ContainsSlice("", []string{"xy", "cd"}))

	// Output:
	// true
	// false
	// false
	// true
	// false
}

func ExampleContainsSliceAny() {
	fmt.Println(strings.ContainsSliceAny("abab", []string{"ab", "cd"}))
	fmt.Println(strings.ContainsSliceAny("ab1", []string{"abc", "12"}))
	fmt.Println(strings.ContainsSliceAny("abz", []string{"xy", "cd"}))
	fmt.Println(strings.ContainsSliceAny("abz", []string{"", ""}))
	fmt.Println(strings.ContainsSliceAny("", []string{"xy", "cd"}))

	// Output:
	// true
	// false
	// false
	// true
	// false
}
