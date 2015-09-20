// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package sha512_test

import (
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/crypto/sha512"
)

func ExampleDigestHex() {
	dst, _ := sha512.DigestHex("goyy")
	fmt.Println(dst)

	// Output:1aa5692a60d265ec371f0f63e59deb0a3bfce9e07ec54de56b75aea9cb0a2ca96fd868d2866843bd0c9d82aa93a696ab557a6cbb57c22e16eab4bd2553198fad
}

func ExampleDigest() {
	dst, _ := sha512.Digest([]byte("goyy"))
	fmt.Println(fmt.Sprintf("%x", dst))

	// Output:1aa5692a60d265ec371f0f63e59deb0a3bfce9e07ec54de56b75aea9cb0a2ca96fd868d2866843bd0c9d82aa93a696ab557a6cbb57c22e16eab4bd2553198fad
}

func ExampleDigestSumHex() {
	dst, _ := sha512.DigestSumHex("goyy", "goyy:")
	fmt.Println(dst)

	// Output:676f79793a1aa5692a60d265ec371f0f63e59deb0a3bfce9e07ec54de56b75aea9cb0a2ca96fd868d2866843bd0c9d82aa93a696ab557a6cbb57c22e16eab4bd2553198fad
}

func ExampleDigestSum() {
	dst, _ := sha512.DigestSum([]byte("goyy"), []byte("goyy:"))
	fmt.Println(fmt.Sprintf("%x", dst))

	// Output:676f79793a1aa5692a60d265ec371f0f63e59deb0a3bfce9e07ec54de56b75aea9cb0a2ca96fd868d2866843bd0c9d82aa93a696ab557a6cbb57c22e16eab4bd2553198fad
}
