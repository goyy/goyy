// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package hmac_test

import (
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/crypto/hmac"
	"testing"
)

func TestSha512Hex(t *testing.T) {
	in := "goyy"
	key := "key"
	expected := "ecaefb579a65c35352558f88a6b691199884784bfd35deb997ec1e40d20cc99d49d586a59e7b4c6a6e8c250e8dd16262547e991bd9e32f5da722f071230308c3"
	if out, _ := hmac.Sha512Hex(in, key); out != expected {
		format := `hmac.Sha512Hex("%s", "%s") = "%s", want "%s"`
		t.Errorf(format, in, key, out, expected)
	}
}

func TestSha512(t *testing.T) {
	in := []byte("goyy")
	key := []byte("key")
	expected := "ecaefb579a65c35352558f88a6b691199884784bfd35deb997ec1e40d20cc99d49d586a59e7b4c6a6e8c250e8dd16262547e991bd9e32f5da722f071230308c3"
	if out, _ := hmac.Sha512(in, key); fmt.Sprintf("%x", out) != expected {
		format := `hmac.Sha512("%s", "%s") = "%x", want "%s"`
		t.Errorf(format, in, key, out, expected)
	}
}

func TestSha512SumHex(t *testing.T) {
	in := "goyy"
	key := "key"
	sum := "goyy:"
	expected := "676f79793aecaefb579a65c35352558f88a6b691199884784bfd35deb997ec1e40d20cc99d49d586a59e7b4c6a6e8c250e8dd16262547e991bd9e32f5da722f071230308c3"
	if out, _ := hmac.Sha512SumHex(in, key, sum); out != expected {
		format := `hmac.Sha512SumHex("%s", "%s", "%s") = "%x", want "%s"`
		t.Errorf(format, in, key, sum, out, expected)
	}
}

func TestSha512Sum(t *testing.T) {
	in := []byte("goyy")
	key := []byte("key")
	sum := []byte("goyy:")
	expected := "676f79793aecaefb579a65c35352558f88a6b691199884784bfd35deb997ec1e40d20cc99d49d586a59e7b4c6a6e8c250e8dd16262547e991bd9e32f5da722f071230308c3"
	if out, _ := hmac.Sha512Sum(in, key, sum); fmt.Sprintf("%x", out) != expected {
		format := `hmac.Sha512Sum("%s", "%s") = "%x", want "%s"`
		t.Errorf(format, in, key, out, expected)
	}
}
