// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package sha1_test

import (
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/crypto/sha1"
)

func ExampleDigestHex() {
	dst, _ := sha1.DigestHex("goyy")
	fmt.Println(dst)

	// Output:9a5de4d2e62e2c0f3018eeff35e09ab3d41781fb
}

func ExampleDigest() {
	dst, _ := sha1.Digest([]byte("goyy"))
	fmt.Println(fmt.Sprintf("%x", dst))

	// Output:9a5de4d2e62e2c0f3018eeff35e09ab3d41781fb
}

func ExampleDigestSumHex() {
	dst, _ := sha1.DigestSumHex("goyy", "goyy:")
	fmt.Println(dst)

	// Output:676f79793a9a5de4d2e62e2c0f3018eeff35e09ab3d41781fb
}

func ExampleDigestSum() {
	dst, _ := sha1.DigestSum([]byte("goyy"), []byte("goyy:"))
	fmt.Println(fmt.Sprintf("%x", dst))

	// Output:676f79793a9a5de4d2e62e2c0f3018eeff35e09ab3d41781fb
}
