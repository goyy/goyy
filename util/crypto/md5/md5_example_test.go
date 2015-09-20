// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package md5_test

import (
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/crypto/md5"
)

func ExampleDigestHex() {
	dst, _ := md5.DigestHex("goyy")
	fmt.Println(dst)

	// Output:db9e2a3e99dbace8332b3042a6beb793
}

func ExampleDigest() {
	dst, _ := md5.Digest([]byte("goyy"))
	fmt.Println(fmt.Sprintf("%x", dst))

	// Output:db9e2a3e99dbace8332b3042a6beb793
}

func ExampleDigestSumHex() {
	dst, _ := md5.DigestSumHex("goyy", "goyy:")
	fmt.Println(dst)

	// Output:676f79793adb9e2a3e99dbace8332b3042a6beb793
}

func ExampleDigestSum() {
	dst, _ := md5.DigestSum([]byte("goyy"), []byte("goyy:"))
	fmt.Println(fmt.Sprintf("%x", dst))

	// Output:676f79793adb9e2a3e99dbace8332b3042a6beb793
}
