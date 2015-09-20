// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package hmac_test

import (
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/crypto/hmac"
)

func ExampleMd5Hex() {
	dst, _ := hmac.Md5Hex("goyy", "key")
	fmt.Println(dst)

	// Output:aa35a3db5523705942e0d117faf76ed5
}

func ExampleMd5() {
	dst, _ := hmac.Md5([]byte("goyy"), []byte("key"))
	fmt.Println(fmt.Sprintf("%x", dst))

	// Output:aa35a3db5523705942e0d117faf76ed5
}

func ExampleMd5SumHex() {
	dst, _ := hmac.Md5SumHex("goyy", "key", "goyy:")
	fmt.Println(dst)

	// Output:676f79793aaa35a3db5523705942e0d117faf76ed5
}

func ExampleMd5Sum() {
	dst, _ := hmac.Md5Sum([]byte("goyy"), []byte("key"), []byte("goyy:"))
	fmt.Println(fmt.Sprintf("%x", dst))

	// Output:676f79793aaa35a3db5523705942e0d117faf76ed5
}
