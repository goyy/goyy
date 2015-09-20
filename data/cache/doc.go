// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

/*
Package cache implements cache utility functions.

Usage

	cache.Init(cache.Conf{Address: "10.105.99.81:6379"})

	cache.Set("key-a", "value-a")
	v, _ := cache.Get("key-a")
	fmt.Println(v)
*/
package cache
