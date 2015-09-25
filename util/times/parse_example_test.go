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

	// Output: 1396503105
}

func ExampleParseUnixStr() {
	v, _ := times.ParseUnixStr(times.YYMDHMS, "2014-04-03 13:31:45")
	fmt.Printf("%v", v)

	// Output: 1396503105
}

func ExampleParseUnixNano() {
	v, _ := times.ParseUnixNano(times.YYMDHMS, "2014-04-03 13:31:45")
	fmt.Printf("%v", v)

	// Output: 1396503105000000000
}

func ExampleParseGMT() {
	v, _ := times.ParseGMT("Thu, 03 Apr 2014 13:31:45 GMT")
	fmt.Printf("%v", v)

	// Output: 1396503105
}

func ExampleParseYYMD() {
	v, _ := times.ParseYYMD("2014-04-03")
	fmt.Printf("%v", v)

	// Output: 1396454400
}

func ExampleParseYYMDHMS() {
	v, _ := times.ParseYYMDHMS("2014-04-03 13:31:45")
	fmt.Printf("%v", v)

	// Output: 1396503105
}

func ExampleParseYYMDHM() {
	v, _ := times.ParseYYMDHM("2014-04-03 13:31")
	fmt.Printf("%v", v)

	// Output: 1396503060
}

func ExampleParseGmt() {
	v, _ := times.ParseGmt("Thu, 03 Apr 2014 13:31:45 GMT")
	fmt.Printf("%v", v)

	// Output: 1396503105
}

func ExampleParseYymd() {
	v, _ := times.ParseYymd("2014-04-03")
	fmt.Printf("%v", v)

	// Output: 1396454400
}

func ExampleParseYymdhms() {
	v, _ := times.ParseYymdhms("2014-04-03 13:31:45")
	fmt.Printf("%v", v)

	// Output: 1396503105
}

func ExampleParseYymdhm() {
	v, _ := times.ParseYymdhm("2014-04-03 13:31")
	fmt.Printf("%v", v)

	// Output: 1396503060
}
