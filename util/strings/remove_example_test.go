// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package strings_test

import (
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

func ExampleRemoveSpace() {
	fmt.Println(strings.RemoveSpace(" a\nb   c\td e \rf   g    "))

	// Output: abcdefg
}

func ExampleRemoveBlank() {
	fmt.Println(strings.RemoveBlank(" a\nb   c\td e \rf   g    "))

	// Output: ab cd e f g
}

func ExampleRemoveStart() {
	fmt.Println(strings.RemoveStart("abc", ""))
	fmt.Println(strings.RemoveStart("www.domain.com", "www."))
	fmt.Println(strings.RemoveStart("www.domain.com", "domain"))

	// Output:
	// abc
	// domain.com
	// www.domain.com
}

func ExampleRemoveEnd() {
	fmt.Println(strings.RemoveEnd("abc", ""))
	fmt.Println(strings.RemoveEnd("www.domain.com", ".com"))
	fmt.Println(strings.RemoveEnd("www.domain.com", "domain"))

	// Output:
	// abc
	// www.domain
	// www.domain.com
}

func ExampleRemove() {
	fmt.Println(strings.Remove("abc", ""))
	fmt.Println(strings.Remove("queued", "ue"))
	fmt.Println(strings.Remove("queued", "zz"))

	// Output:
	// abc
	// qd
	// queued
}
