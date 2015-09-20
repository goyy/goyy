// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package hmac_test

import (
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/crypto/hmac"
	"testing"
)

func TestMd5Hex(t *testing.T) {
	in := "goyy"
	key := "key"
	expected := "aa35a3db5523705942e0d117faf76ed5"
	if out, _ := hmac.Md5Hex(in, key); out != expected {
		format := `hmac.Md5Hex("%s", "%s") = "%s", want "%s"`
		t.Errorf(format, in, key, out, expected)
	}
}

func TestMd5(t *testing.T) {
	in := []byte("goyy")
	key := []byte("key")
	expected := "aa35a3db5523705942e0d117faf76ed5"
	if out, _ := hmac.Md5(in, key); fmt.Sprintf("%x", out) != expected {
		format := `hmac.Md5("%s", "%s") = "%x", want "%s"`
		t.Errorf(format, in, key, out, expected)
	}
}

func TestMd5SumHex(t *testing.T) {
	in := "goyy"
	key := "key"
	sum := "goyy:"
	expected := "676f79793aaa35a3db5523705942e0d117faf76ed5"
	if out, _ := hmac.Md5SumHex(in, key, sum); out != expected {
		format := `hmac.Md5SumHex("%s", "%s", "%s") = "%x", want "%s"`
		t.Errorf(format, in, key, sum, out, expected)
	}
}

func TestMd5Sum(t *testing.T) {
	in := []byte("goyy")
	key := []byte("key")
	sum := []byte("goyy:")
	expected := "676f79793aaa35a3db5523705942e0d117faf76ed5"
	if out, _ := hmac.Md5Sum(in, key, sum); fmt.Sprintf("%x", out) != expected {
		format := `hmac.Md5Sum("%s", "%s") = "%x", want "%s"`
		t.Errorf(format, in, key, out, expected)
	}
}
