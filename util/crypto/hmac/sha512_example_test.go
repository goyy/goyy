// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package hmac_test

import (
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/crypto/hmac"
)

func ExampleSha512Hex() {
	dst, _ := hmac.Sha512Hex("goyy", "key")
	fmt.Println(dst)

	// Output:ecaefb579a65c35352558f88a6b691199884784bfd35deb997ec1e40d20cc99d49d586a59e7b4c6a6e8c250e8dd16262547e991bd9e32f5da722f071230308c3
}

func ExampleSha512() {
	dst, _ := hmac.Sha512([]byte("goyy"), []byte("key"))
	fmt.Println(fmt.Sprintf("%x", dst))

	// Output:ecaefb579a65c35352558f88a6b691199884784bfd35deb997ec1e40d20cc99d49d586a59e7b4c6a6e8c250e8dd16262547e991bd9e32f5da722f071230308c3
}

func ExampleSha512SumHex() {
	dst, _ := hmac.Sha512SumHex("goyy", "key", "goyy:")
	fmt.Println(dst)

	// Output:676f79793aecaefb579a65c35352558f88a6b691199884784bfd35deb997ec1e40d20cc99d49d586a59e7b4c6a6e8c250e8dd16262547e991bd9e32f5da722f071230308c3
}

func ExampleSha512Sum() {
	dst, _ := hmac.Sha512Sum([]byte("goyy"), []byte("key"), []byte("goyy:"))
	fmt.Println(fmt.Sprintf("%x", dst))

	// Output:676f79793aecaefb579a65c35352558f88a6b691199884784bfd35deb997ec1e40d20cc99d49d586a59e7b4c6a6e8c250e8dd16262547e991bd9e32f5da722f071230308c3
}
