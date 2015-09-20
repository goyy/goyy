// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rc4_test

import (
	"bytes"
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/crypto/rc4"
)

func ExampleEncrypt() {
	tt := rc4Tests[0]
	dst, _ := rc4.Encrypt(tt.in, tt.key)
	fmt.Println(bytes.Compare(tt.out, dst))

	// Output:0
}

func ExampleDecrypt() {
	tt := rc4Tests[0]
	dst, _ := rc4.Decrypt(tt.out, tt.key)
	fmt.Println(bytes.Compare(tt.in, dst))

	// Output:0
}

func ExampleEncryptHex() {
	dst, _ := rc4.EncryptHex("goyy", "key")
	fmt.Println(dst)

	// Output:6c034d94
}

func ExampleDecryptHex() {
	dst, _ := rc4.DecryptHex("6c034d94", "key")
	fmt.Println(dst)

	// Output:goyy
}
