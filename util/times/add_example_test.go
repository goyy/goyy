// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package times_test

import (
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/times"
)

func ExampleAdd() {
	fmt.Println(times.Add(in, times.Day))

	// Output: 2014-04-04 13:31:45.001234454 +0800 CST
}

func ExampleAddUnix() {
	fmt.Println(times.AddUnix(i, times.Day))

	// Output: 1396589505
}

func ExampleAddStr() {
	out, _ := times.AddStr(times.YYMDHMS, "2014-04-03 13:31:45", times.Day)
	fmt.Println(out)

	// Output: 2014-04-04 13:31:45
}

func ExampleAddGMT() {
	out, _ := times.AddGMT("Thu, 03 Apr 2014 13:31:45 GMT", times.Day)
	fmt.Println(out)

	// Output: Fri, 04 Apr 2014 13:31:45 GMT
}

func ExampleAddYYMD() {
	out, _ := times.AddYYMD("2014-04-03", times.Day)
	fmt.Println(out)

	// Output: 2014-04-04
}

func ExampleAddYYMDHMS() {
	out, _ := times.AddYYMDHMS("2014-04-03 13:31:45", times.Day)
	fmt.Println(out)

	// Output: 2014-04-04 13:31:45
}

func ExampleAddYYMDHM() {
	out, _ := times.AddYYMDHM("2014-04-03 13:31", times.Day)
	fmt.Println(out)

	// Output: 2014-04-04 13:31
}
