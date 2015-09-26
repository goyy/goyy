// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package times_test

import (
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/times"
)

func ExampleFormat() {
	fmt.Println(times.Format(times.YYMDHMS, in))

	// Output: 2014-04-03 13:31:45
}

func ExampleFormatGMT() {
	fmt.Println(times.FormatGMT(in))

	// Output: Thu, 03 Apr 2014 13:31:45 GMT
}

func ExampleFormatYYMD() {
	fmt.Println(times.FormatYYMD(in))

	// Output: 2014-04-03
}

func ExampleFormatYYMDHMS() {
	fmt.Println(times.FormatYYMDHMS(in))

	// Output: 2014-04-03 13:31:45
}

func ExampleFormatYYMDHM() {
	fmt.Println(times.FormatYYMDHM(in))

	// Output: 2014-04-03 13:31
}
