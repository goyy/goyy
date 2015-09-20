// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package strings_test

import (
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

func ExampleRepeat() {
	fmt.Println(strings.Repeat("-", 10))

	// Output: ----------
}

func ExamplePadStart() {
	fmt.Printf("%#q\n", strings.PadStart("bat", 5))

	// Output: `  bat`
}

func ExamplePadEnd() {
	fmt.Printf("%#q\n", strings.PadEnd("bat", 5))

	// Output: `bat  `
}

func ExamplePadLeft() {
	fmt.Println(strings.PadLeft("bat", 5, "*"))

	// Output: **bat
}

func ExamplePadRight() {
	fmt.Println(strings.PadRight("bat", 5, "*"))

	// Output: bat**
}

func ExamplePad() {
	fmt.Printf("%#q\n", strings.Pad("bat", 5))

	// Output: ` bat `
}

func ExampleCenter() {
	fmt.Println(strings.Center("bat", 8, "tag"))

	// Output: tabattag
}
