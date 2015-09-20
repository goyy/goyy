// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package properties_test

import (
	"fmt"
	"gopkg.in/goyy/goyy.v0/comm/properties"
)

func ExampleProperty() {
	p, _ := properties.New("./example.properties")
	fmt.Println(p.Property("say"))

	// Output:
	// Hello, world!
}

func ExamplePropertyf() {
	p, _ := properties.New("./example.properties")
	fmt.Println(p.Propertyf("sayf", "goyy"))

	// Output:
	// Hello, goyy!
}

func ExampleSetProperty() {
	p, _ := properties.New("./example.properties")
	// read
	fmt.Println(p.Property("eg"))
	// write
	p.SetProperty("eg", "Hello, goyy!")
	fmt.Println(p.Property("eg"))
	// revert
	p.SetProperty("eg", "Hello, world!")

	// Output:
	// Hello, world!
	// Hello, goyy!
}
