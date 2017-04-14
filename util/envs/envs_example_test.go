// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package envs_test

import (
	"fmt"

	"gopkg.in/goyy/goyy.v0/util/envs"
)

func ExampleParseURLQuery() {
	fmt.Println(envs.ParseGOPATH("%GOPATH%/src/gopkg.in/goyy/goyy.v0"))
	fmt.Println(envs.ParseGOPATH("/src/gopkg.in/goyy/goyy.v0"))
	fmt.Println(envs.ParseGOPATH(""))
	fmt.Println(envs.ParseGOPATH(" "))

	// Output:
	// e:\gopath/src/gopkg.in/goyy/goyy.v0
	// /src/gopkg.in/goyy/goyy.v0
	//
	//
}
