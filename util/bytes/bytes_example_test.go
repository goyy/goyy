// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package bytes_test

import (
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/bytes"
)

func ExampleTrimRightNul() {
	a := [10]byte{1, 2, 3}
	src := a[:]
	fmt.Printf("%q\n", src)
	fmt.Printf("%q\n", bytes.TrimRightNul(src))

	// Output:
	// "\x01\x02\x03\x00\x00\x00\x00\x00\x00\x00"
	// "\x01\x02\x03"
}
