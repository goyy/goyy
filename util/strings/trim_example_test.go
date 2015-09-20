// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package strings_test

import (
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

func ExampleTrim() {
	fmt.Printf("[%q]\n", strings.Trim("abba", ""))
	fmt.Printf("[%q]\n", strings.Trim(" !!! Achtung !!! ", "! "))
	// Output:
	// ["abba"]
	// ["Achtung"]
}

func ExampleTrimLeft() {
	fmt.Printf("[%q]\n", strings.TrimLeft("abba", ""))
	fmt.Printf("[%q]\n", strings.TrimLeft(" !!! Achtung !!! ", "! "))
	// Output:
	// ["abba"]
	// ["Achtung !!! "]
}

func ExampleTrimRight() {
	fmt.Printf("[%q]\n", strings.TrimRight("abba", ""))
	fmt.Printf("[%q]\n", strings.TrimRight(" !!! Achtung !!! ", "! "))
	// Output:
	// ["abba"]
	// [" !!! Achtung"]
}

func ExampleTrimSpace() {
	fmt.Println(strings.TrimSpace(" \t\n a lone gopher \n\t\r\n"))
	// Output: a lone gopher
}
