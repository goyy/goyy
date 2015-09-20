// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package sha256_test

import (
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/crypto/sha256"
)

func ExampleDigestHex() {
	dst, _ := sha256.DigestHex("goyy")
	fmt.Println(dst)

	// Output:294d5a5ab822caec5c8ef8db496a44a5f185d83feb0699ab69e79b51166a1116
}

func ExampleDigest() {
	dst, _ := sha256.Digest([]byte("goyy"))
	fmt.Println(fmt.Sprintf("%x", dst))

	// Output:294d5a5ab822caec5c8ef8db496a44a5f185d83feb0699ab69e79b51166a1116
}

func ExampleDigestSumHex() {
	dst, _ := sha256.DigestSumHex("goyy", "goyy:")
	fmt.Println(dst)

	// Output:676f79793a294d5a5ab822caec5c8ef8db496a44a5f185d83feb0699ab69e79b51166a1116
}

func ExampleDigestSum() {
	dst, _ := sha256.DigestSum([]byte("goyy"), []byte("goyy:"))
	fmt.Println(fmt.Sprintf("%x", dst))

	// Output:676f79793a294d5a5ab822caec5c8ef8db496a44a5f185d83feb0699ab69e79b51166a1116
}
