// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rsa_test

import (
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/crypto/rsa"
	"strings"
)

func ExampleGeneratingKey() {
	prikey, pubkey, _ := rsa.GeneratingKey(512)
	fmt.Println(strings.Contains(prikey, "-----BEGIN RSA PRIVATE KEY-----"))
	fmt.Println(strings.Contains(pubkey, "-----BEGIN RSA PUBLIC KEY-----"))

	// Output:
	// true
	// true
}
