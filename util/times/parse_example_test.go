// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package times_test

import (
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/times"
)

func ExampleParseUnix() {
	v, _ := times.ParseUnix(times.YYMDHMS, "2014-04-03 13:31:45")
	fmt.Printf("%v", v)

	// Output: 1396531905
}

func ExampleParseUnixNano() {
	v, _ := times.ParseUnixNano(times.YYMDHMS, "2014-04-03 13:31:45")
	fmt.Printf("%v", v)

	// Output: 1396531905000000000
}
