// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package hmac_test

import (
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/crypto/hmac"
)

func ExampleSha1Hex() {
	dst, _ := hmac.Sha1Hex("goyy", "key")
	fmt.Println(dst)

	// Output:221f50c7d0d8f13e160be5660a1a27d5bb18861c
}

func ExampleSha1() {
	dst, _ := hmac.Sha1([]byte("goyy"), []byte("key"))
	fmt.Println(fmt.Sprintf("%x", dst))

	// Output:221f50c7d0d8f13e160be5660a1a27d5bb18861c
}

func ExampleSha1SumHex() {
	dst, _ := hmac.Sha1SumHex("goyy", "key", "goyy:")
	fmt.Println(dst)

	// Output:676f79793a221f50c7d0d8f13e160be5660a1a27d5bb18861c
}

func ExampleSha1Sum() {
	dst, _ := hmac.Sha1Sum([]byte("goyy"), []byte("key"), []byte("goyy:"))
	fmt.Println(fmt.Sprintf("%x", dst))

	// Output:676f79793a221f50c7d0d8f13e160be5660a1a27d5bb18861c
}
