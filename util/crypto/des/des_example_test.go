// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package des_test

import (
	"bytes"
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/crypto/des"
)

func ExampleEncrypt() {
	tt := desTests[0]
	dst, _ := des.Encrypt(tt.in, tt.key)
	fmt.Println(bytes.Compare(tt.out, dst))

	// Output:0
}

func ExampleDecrypt() {
	tt := desTests[0]
	dst, _ := des.Decrypt(tt.out, tt.key)
	fmt.Println(bytes.Compare(tt.in, dst))

	// Output:0
}

func ExampleEncryptHex() {
	dst, _ := des.EncryptHex("goyy", "key")
	fmt.Println(dst)

	// Output:0595738e
}

func ExampleDecryptHex() {
	dst, _ := des.DecryptHex("0595738e", "key")
	fmt.Println(dst)

	// Output:goyy
}
