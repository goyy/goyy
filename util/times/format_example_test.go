// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package times_test

import (
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/times"
)

func ExampleFormat() {
	fmt.Println(times.Format(in, times.YYMDHMS))

	// Output: 2014-04-03 13:31:45
}

func ExampleGmt() {
	fmt.Println(times.Gmt(in))

	// Output: Thu, 03 Apr 2014 13:31:45 GMT
}

func ExampleYymd() {
	fmt.Println(times.Yymd(in))

	// Output: 2014-04-03
}

func ExampleYymdhms() {
	fmt.Println(times.Yymdhms(in))

	// Output: 2014-04-03 13:31:45
}

func ExampleYymdhm() {
	fmt.Println(times.Yymdhm(in))

	// Output: 2014-04-03 13:31
}

func ExampleUformat() {
	fmt.Println(times.Uformat(i, times.YYMDHMS))

	// Output: 2014-04-03 13:31:45
}

func ExampleUgmt() {
	fmt.Println(times.Ugmt(i))

	// Output: Thu, 03 Apr 2014 13:31:45 GMT
}

func ExampleUyymd() {
	fmt.Println(times.Uyymd(i))

	// Output: 2014-04-03
}

func ExampleUyymdhms() {
	fmt.Println(times.Uyymdhms(i))

	// Output: 2014-04-03 13:31:45
}

func ExampleUyymdhm() {
	fmt.Println(times.Uyymdhm(i))

	// Output: 2014-04-03 13:31
}
