// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package strings_test

import (
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

func ExampleOverlay() {
	fmt.Println(strings.Overlay("", "abc", 0, 0))
	fmt.Println(strings.Overlay("abcdef", "", 2, 4))
	fmt.Println(strings.Overlay("abcdef", "", 4, 2))
	fmt.Println(strings.Overlay("abcdef", "zzzz", 2, 4))
	fmt.Println(strings.Overlay("abcdef", "zzzz", 4, 2))
	fmt.Println(strings.Overlay("abcdef", "zzzz", -1, 4))
	fmt.Println(strings.Overlay("abcdef", "zzzz", 2, 8))
	fmt.Println(strings.Overlay("abcdef", "zzzz", -2, -3))
	fmt.Println(strings.Overlay("abcdef", "zzzz", 8, 10))
	fmt.Println(strings.Overlay("世界hello", "zzzz", 2, 4))
	// Output:
	// abc
	// abef
	// abef
	// abzzzzef
	// abzzzzef
	// zzzzef
	// abzzzz
	// zzzzabcdef
	// abcdefzzzz
	// 世界zzzzllo
}
