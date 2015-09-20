// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package hmac_test

import (
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/crypto/hmac"
)

func ExampleSha256Hex() {
	dst, _ := hmac.Sha256Hex("goyy", "key")
	fmt.Println(dst)

	// Output:840f83ba7e39cc809ea5ff9829da11070b21aaf4f306a356d0eb782a7c09f54a
}

func ExampleSha256() {
	dst, _ := hmac.Sha256([]byte("goyy"), []byte("key"))
	fmt.Println(fmt.Sprintf("%x", dst))

	// Output:840f83ba7e39cc809ea5ff9829da11070b21aaf4f306a356d0eb782a7c09f54a
}

func ExampleSha256SumHex() {
	dst, _ := hmac.Sha256SumHex("goyy", "key", "goyy:")
	fmt.Println(dst)

	// Output:676f79793a840f83ba7e39cc809ea5ff9829da11070b21aaf4f306a356d0eb782a7c09f54a
}

func ExampleSha256Sum() {
	dst, _ := hmac.Sha256Sum([]byte("goyy"), []byte("key"), []byte("goyy:"))
	fmt.Println(fmt.Sprintf("%x", dst))

	// Output:676f79793a840f83ba7e39cc809ea5ff9829da11070b21aaf4f306a356d0eb782a7c09f54a
}
