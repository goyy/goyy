// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package times_test

import (
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/times"
)

func ExampleFormatUnix() {
	fmt.Println(times.FormatUnix(times.YYMDHMS, i))

	// Output: 2014-04-03 13:31:45
}

func ExampleFormatUnixGMT() {
	fmt.Println(times.FormatUnixGMT(i))

	// Output: Thu, 03 Apr 2014 13:31:45 GMT
}

func ExampleFormatUnixYYMD() {
	fmt.Println(times.FormatUnixYYMD(i))

	// Output: 2014-04-03
}

func ExampleFormatUnixYYMDHMS() {
	fmt.Println(times.FormatUnixYYMDHMS(i))

	// Output: 2014-04-03 13:31:45
}

func ExampleFormatUnixYYMDHM() {
	fmt.Println(times.FormatUnixYYMDHM(i))

	// Output: 2014-04-03 13:31
}
