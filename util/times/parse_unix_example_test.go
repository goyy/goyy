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

func ExampleParseUnixNano() {
	v, _ := times.ParseUnixNano(times.YYMDHMS, "2014-04-03 13:31:45")
	fmt.Printf("%v", v)

	// Output: 1396503105000000000
}

func ExampleParseUnixStr() {
	v, _ := times.ParseUnixStr(times.YYMDHMS, "2014-04-03 13:31:45")
	fmt.Printf("%v", v)

	// Output: 1396503105
}

func ExampleParseUnixGMT() {
	v, _ := times.ParseUnixGMT("Thu, 03 Apr 2014 13:31:45 GMT")
	fmt.Printf("%v", v)

	// Output: 1396503105
}

func ExampleParseUnixYYMD() {
	v, _ := times.ParseUnixYYMD("2014-04-03")
	fmt.Printf("%v", v)

	// Output: 1396454400
}

func ExampleParseUnixYYMDHMS() {
	v, _ := times.ParseUnixYYMDHMS("2014-04-03 13:31:45")
	fmt.Printf("%v", v)

	// Output: 1396503105
}

func ExampleParseUnixYYMDHM() {
	v, _ := times.ParseUnixYYMDHM("2014-04-03 13:31")
	fmt.Printf("%v", v)

	// Output: 1396503060
}

func ExampleParseUnixGmt() {
	v, _ := times.ParseUnixGmt("Thu, 03 Apr 2014 13:31:45 GMT")
	fmt.Printf("%v", v)

	// Output: 1396503105
}

func ExampleParseUnixYymd() {
	v, _ := times.ParseUnixYymd("2014-04-03")
	fmt.Printf("%v", v)

	// Output: 1396454400
}

func ExampleParseUnixYymdhms() {
	v, _ := times.ParseUnixYymdhms("2014-04-03 13:31:45")
	fmt.Printf("%v", v)

	// Output: 1396503105
}

func ExampleParseUnixYymdhm() {
	v, _ := times.ParseUnixYymdhm("2014-04-03 13:31")
	fmt.Printf("%v", v)

	// Output: 1396503060
}
