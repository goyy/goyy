// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package maps_test

import (
	"fmt"

	"gopkg.in/goyy/goyy.v0/util/maps"
)

func ExampleParseURLQuery() {
	fmt.Println(maps.ParseURLQuery("key1=value1&key2=value2"))
	fmt.Println(maps.ParseURLQuery("key1=value1"))
	fmt.Println(maps.ParseURLQuery("key1="))
	fmt.Println(maps.ParseURLQuery("=val1"))
	fmt.Println(maps.ParseURLQuery("="))
	fmt.Println(maps.ParseURLQuery(""))
	fmt.Println(maps.ParseURLQuery(" "))

	// Output:
	// map[key1:value1 key2:value2]
	// map[key1:value1]
	// map[key1:]
	// map[]
	// map[]
	// map[]
	// map[]
}
