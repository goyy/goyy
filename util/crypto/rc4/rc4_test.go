// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rc4_test

import (
	"bytes"
	"gopkg.in/goyy/goyy.v0/util/crypto/rc4"
	"testing"
)

type rc4Test struct {
	in  []byte
	key []byte
	out []byte
}

var rc4Tests = []rc4Test{
	{
		[]byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa,
			0xbb, 0xcc, 0xdd},
		[]byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a,
			0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15,
			0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
		},
		[]byte{0x10, 0x35, 0xd3, 0xfa, 0xee, 0xfa, 0xcf, 0x4a, 0xfe, 0xa5, 0x34,
			0x3b, 0xc4, 0xe8},
	},
}

func TestEncrypt(t *testing.T) {
	tt := rc4Tests[0]
	if out, _ := rc4.Encrypt(tt.in, tt.key); bytes.Compare(out, tt.out) != 0 {
		format := `rc4.Encrypt("%#v", "%#v") = "%#v", want "%#v"`
		t.Errorf(format, tt.in, tt.key, out, tt.out)
	}
}

func TestDecrypt(t *testing.T) {
	tt := rc4Tests[0]
	if out, _ := rc4.Decrypt(tt.out, tt.key); bytes.Compare(out, tt.in) != 0 {
		format := `rc4.Decrypt("%#v", "%#v") = "%#v", want "%#v"`
		t.Errorf(format, tt.out, tt.key, out, tt.in)
	}
}

func TestEncryptHex(t *testing.T) {
	in := "goyy"
	key := "key"
	expected := "6c034d94"
	if out, _ := rc4.EncryptHex(in, key); out != expected {
		format := `rc4.EncryptHex("%s", "%s") = "%s", want "%s"`
		t.Errorf(format, in, key, out, expected)
	}
}

func TestDecryptHex(t *testing.T) {
	in := "6c034d94"
	key := "key"
	expected := "goyy"
	if out, _ := rc4.DecryptHex(in, key); out != expected {
		format := `rc4.DecryptHex("%s", "%s") = "%s", want "%s"`
		t.Errorf(format, in, key, out, expected)
	}
}
