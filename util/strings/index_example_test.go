// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package strings_test

import (
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

func ExampleIndex() {
	fmt.Println(strings.Index("abABabABabAB", "ab"))
	fmt.Println(strings.Index("abcdABCDabcd", "cd"))
	fmt.Println(strings.Index("abcdABCDabcd", "de"))

	// Output:
	// 0
	// 2
	// -1
}

func ExampleIndexLast() {
	fmt.Println(strings.IndexLast("abABabABabAB", "ab"))
	fmt.Println(strings.IndexLast("abcdABCDabcd", "cd"))
	fmt.Println(strings.IndexLast("abcdABCDabcd", "de"))

	// Output:
	// 8
	// 10
	// -1
}

func ExampleIndexLastStart() {
	fmt.Println(strings.IndexLastStart("abABabABabAB", "ab", 6))
	fmt.Println(strings.IndexLastStart("abcdABCDabcd", "cd", 5))
	fmt.Println(strings.IndexLastStart("abcdABCDabcd", "de", 0))

	// Output:
	// 4
	// 2
	// -1
}

func ExampleIndexLastOrdinal() {
	fmt.Println(strings.IndexLastOrdinal("abABabABabAB", "ab", 2))
	fmt.Println(strings.IndexLastOrdinal("abcdABCDabcd", "cd", 2))
	fmt.Println(strings.IndexLastOrdinal("abcdABCDabcd", "de", 1))

	// Output:
	// 4
	// 2
	// -1
}

func ExampleIndexStart() {
	fmt.Println(strings.IndexStart("abABabABabAB", "ab", 6))
	fmt.Println(strings.IndexStart("abcdABCDabcd", "cd", 5))
	fmt.Println(strings.IndexStart("abcdABCDabcd", "de", 0))

	// Output:
	// 8
	// 10
	// -1
}

func ExampleIndexForward() {
	fmt.Println(strings.IndexForward("abABabABabAB", "ab", 6))
	fmt.Println(strings.IndexForward("abcdABCDabcd", "cd", 5))
	fmt.Println(strings.IndexForward("abcdABCDabcd", "de", 0))

	// Output:
	// 4
	// 2
	// -1
}

func ExampleIndexOrdinal() {
	fmt.Println(strings.IndexOrdinal("abABabABabAB", "ab", 3))
	fmt.Println(strings.IndexOrdinal("abcdABCDabcd", "cd", 2))
	fmt.Println(strings.IndexOrdinal("abcdABCDabcd", "de", 1))

	// Output:
	// 8
	// 10
	// -1
}
