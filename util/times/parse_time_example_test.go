// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package times_test

import (
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/times"
)

func ExampleParse() {
	v, _ := times.Parse(times.YYMDHMS, "2014-04-03 13:31:45")
	fmt.Printf("%v", v)

	// Output: 2014-04-03 13:31:45 +0800 CST
}

func ExampleParseGMT() {
	v, _ := times.ParseGMT("Thu, 03 Apr 2014 13:31:45 GMT")
	fmt.Printf("%v", v)

	// Output: 2014-04-03 13:31:45 +0800 CST
}

func ExampleParseYYMD() {
	v, _ := times.ParseYYMD("2014-04-03")
	fmt.Printf("%v", v)

	// Output: 2014-04-03 00:00:00 +0800 CST
}

func ExampleParseYYMDHMS() {
	v, _ := times.ParseYYMDHMS("2014-04-03 13:31:45")
	fmt.Printf("%v", v)

	// Output: 2014-04-03 13:31:45 +0800 CST
}

func ExampleParseYYMDHM() {
	v, _ := times.ParseYYMDHM("2014-04-03 13:31")
	fmt.Printf("%v", v)

	// Output: 2014-04-03 13:31:00 +0800 CST
}
