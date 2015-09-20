// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package files_test

import (
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/files"
	"log"
)

func ExampleIsExist() {
	fmt.Println(files.IsExist("./example.txt"))
	fmt.Println(files.IsExist("./README"))

	// Output:
	// true
	// false
}

func ExampleRead() {
	s, _ := files.Read("./example.txt")
	fmt.Println(s)

	// Output: Hello world!
}

func ExampleWrite() {
	filename := "./example.txt"
	data := "Hello goyy!"
	if err := files.Write(filename, data, 0644); err != nil {
		log.Fatalf("Write %s: %v", filename, err)
	}
	s, _ := files.Read(filename)
	fmt.Println(s)
	files.Write(filename, "Hello world!", 0644) // recover

	// Output: Hello goyy!
}
