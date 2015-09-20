// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rsa_test

import (
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/crypto/rsa"
)

func ExampleRsaHex() {
	in := "goyy"
	out, _ := rsa.EncryptHex(in, publicKey)
	dst, _ := rsa.DecryptHex(out, privateKey)
	fmt.Println(in == dst)

	// Output:true
}
