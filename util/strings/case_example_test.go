// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package strings_test

import (
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

func ExampleToLower() {
	s := []string{"",
		"abc",
		"AbC123",
		"azAZ09_"}
	for _, v := range s {
		fmt.Println(strings.ToLower(v))
	}

	// Output:
	//
	// abc
	// abc123
	// azaz09_
}

func ExampleToUpper() {
	s := []string{"",
		"abc",
		"AbC123",
		"azAZ09_"}
	for _, v := range s {
		fmt.Println(strings.ToUpper(v))
	}

	// Output:
	//
	// ABC
	// ABC123
	// AZAZ09_
}

func ExampleTitle() {
	s := []string{"",
		"a",
		"cat",
		"cAt",
		"aaa aaa aaa",
		"Aaa Aaa Aaa",
		"123a456",
		"double-blind",
		"ÿøû"}
	for _, v := range s {
		fmt.Println(strings.Title(v))
	}

	// Output:
	//
	// A
	// Cat
	// CAt
	// Aaa Aaa Aaa
	// Aaa Aaa Aaa
	// 123a456
	// Double-Blind
	// Ÿøû
}

func ExampleToTitle() {
	s := []string{"",
		"a",
		"cat",
		"cAt",
		"aaa aaa aaa",
		"Aaa Aaa Aaa",
		"123a456",
		"double-blind",
		"ÿøû"}
	for _, v := range s {
		fmt.Println(strings.ToTitle(v))
	}

	// Output:
	//
	// A
	// CAT
	// CAT
	// AAA AAA AAA
	// AAA AAA AAA
	// 123A456
	// DOUBLE-BLIND
	// ŸØÛ
}

func ExampleCamel() {
	s := []string{"",
		"a",
		"cat",
		"cAt",
		" aaa aaa aaa ",
		"_Aaa_Aaa_Aaa_",
		"123a456",
		"douBle-blind",
		"ÿøû"}
	for _, v := range s {
		fmt.Println(strings.Camel(v))
	}

	// Output:
	//
	// A
	// Cat
	// Cat
	// AaaAaaAaa
	// AaaAaaAaa
	// 123a456
	// DoubleBlind
	// Ÿøû
}

func ExampleUnCamel() {
	s := []string{"",
		"A",
		"Cat",
		"cAt",
		"AaaAaaAaa",
		" AaaAaaAaa ",
		"123a456",
		"DoubleBlind",
		"Ÿøû"}
	for _, v := range s {
		fmt.Println(strings.UnCamel(v, "_"))
	}

	// Output:
	//
	// a
	// cat
	// c_at
	// aaa_aaa_aaa
	// aaa_aaa_aaa
	// 123a456
	// double_blind
	// ÿøû
}
