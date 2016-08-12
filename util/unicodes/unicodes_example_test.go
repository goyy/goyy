// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package unicodes_test

import (
	"fmt"

	"gopkg.in/goyy/goyy.v0/util/unicodes"
)

func ExampleIsQuote() {
	fmt.Println(unicodes.IsQuote('"'))
	fmt.Println(unicodes.IsQuote('`'))
	fmt.Println(unicodes.IsQuote('\''))
	fmt.Println(unicodes.IsQuote('a'))
	fmt.Println(unicodes.IsQuote('1'))
	fmt.Println(unicodes.IsQuote('A'))
	fmt.Println(unicodes.IsQuote('一'))
	fmt.Println(unicodes.IsQuote(' '))

	// Output:
	// true
	// true
	// true
	// false
	// false
	// false
	// false
	// false
}

func ExampleIsHan() {
	fmt.Println(unicodes.IsHan('中'))
	fmt.Println(unicodes.IsHan('文'))
	fmt.Println(unicodes.IsHan('一'))
	fmt.Println(unicodes.IsHan('a'))
	fmt.Println(unicodes.IsHan('A'))
	fmt.Println(unicodes.IsHan('1'))
	fmt.Println(unicodes.IsHan(','))
	fmt.Println(unicodes.IsHan(' '))

	// Output:
	// true
	// true
	// true
	// false
	// false
	// false
	// false
	// false
}
