// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cache_test

import (
	"fmt"
	"gopkg.in/goyy/goyy.v0/data/cache"
)

func ExampleString() {
	cache.Set("key-a", "value-a")
	v, _ := cache.Get("key-a")
	fmt.Println(v)

	// Output:
	// value-a
}
