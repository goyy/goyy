// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package strings_test

import (
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

func ExampleIsBlank() {
	fmt.Println(strings.IsBlank(""))
	fmt.Println(strings.IsBlank("  "))
	fmt.Println(strings.IsBlank(" \t\r\n "))
	fmt.Println(strings.IsBlank("a"))
	fmt.Println(strings.IsBlank(" a  "))
	fmt.Println(strings.IsBlank(" \t\r\n a \t\r\n "))

	// Output:
	// true
	// true
	// true
	// false
	// false
	// false
}

func ExampleIsAnyBlank() {
	fmt.Println(strings.IsAnyBlank("", " "))
	fmt.Println(strings.IsAnyBlank("  ", " a "))
	fmt.Println(strings.IsAnyBlank(" \t\r\n ", " a "))
	fmt.Println(strings.IsAnyBlank("a", " a "))
	fmt.Println(strings.IsAnyBlank(" a ", " a "))
	fmt.Println(strings.IsAnyBlank(" \t\r\n a \t\r\n ", " a "))

	// Output:
	// true
	// true
	// true
	// false
	// false
	// false
}

func ExampleIsNoneBlank() {
	fmt.Println(strings.IsNoneBlank("", " "))
	fmt.Println(strings.IsNoneBlank("  ", " a "))
	fmt.Println(strings.IsNoneBlank(" \t\r\n ", " a "))
	fmt.Println(strings.IsNoneBlank("a", " a "))
	fmt.Println(strings.IsNoneBlank(" a ", " a "))
	fmt.Println(strings.IsNoneBlank(" \t\r\n a \t\r\n ", " a "))

	// Output:
	// false
	// false
	// false
	// true
	// true
	// true
}

func ExampleIsNotBlank() {
	fmt.Println(strings.IsNotBlank(""))
	fmt.Println(strings.IsNotBlank("  "))
	fmt.Println(strings.IsNotBlank(" \t\r\n "))
	fmt.Println(strings.IsNotBlank("a"))
	fmt.Println(strings.IsNotBlank(" a  "))
	fmt.Println(strings.IsNotBlank(" \t\r\n a \t\r\n "))

	// Output:
	// false
	// false
	// false
	// true
	// true
	// true
}
