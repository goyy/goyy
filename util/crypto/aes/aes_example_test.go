// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package aes_test

import (
	"bytes"
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/crypto/aes"
)

func ExampleEncrypt() {
	tt := aesTests[0]
	dst, _ := aes.Encrypt(tt.in, tt.key)
	fmt.Println(bytes.Compare(tt.out, dst))

	// Output:0
}

func ExampleDecrypt() {
	tt := aesTests[0]
	dst, _ := aes.Decrypt(tt.out, tt.key)
	fmt.Println(bytes.Compare(tt.in, dst))

	// Output:0
}

func ExampleEncryptHex() {
	dst, _ := aes.EncryptHex("goyy", "key")
	fmt.Println(dst)

	// Output:efe9f867
}

func ExampleDecryptHex() {
	dst, _ := aes.DecryptHex("efe9f867", "key")
	fmt.Println(dst)

	// Output:goyy
}
