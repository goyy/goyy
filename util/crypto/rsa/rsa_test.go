// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rsa_test

import (
	"gopkg.in/goyy/goyy.v0/util/crypto/rsa"
	"testing"
)

func TestRsaHex(t *testing.T) {
	in := "goyy"
	out, _ := rsa.EncryptHex(in, publicKey)
	expected, _ := rsa.DecryptHex(out, privateKey)
	if in != expected {
		format := `out = "%s", want "%s"`
		t.Errorf(format, expected, in)
	}
}
