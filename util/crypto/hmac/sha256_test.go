// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package hmac_test

import (
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/crypto/hmac"
	"testing"
)

func TestSha256Hex(t *testing.T) {
	in := "goyy"
	key := "key"
	expected := "840f83ba7e39cc809ea5ff9829da11070b21aaf4f306a356d0eb782a7c09f54a"
	if out, _ := hmac.Sha256Hex(in, key); out != expected {
		format := `hmac.Sha256Hex("%s", "%s") = "%s", want "%s"`
		t.Errorf(format, in, key, out, expected)
	}
}

func TestSha256(t *testing.T) {
	in := []byte("goyy")
	key := []byte("key")
	expected := "840f83ba7e39cc809ea5ff9829da11070b21aaf4f306a356d0eb782a7c09f54a"
	if out, _ := hmac.Sha256(in, key); fmt.Sprintf("%x", out) != expected {
		format := `hmac.Sha256("%s", "%s") = "%x", want "%s"`
		t.Errorf(format, in, key, out, expected)
	}
}

func TestSha256SumHex(t *testing.T) {
	in := "goyy"
	key := "key"
	sum := "goyy:"
	expected := "676f79793a840f83ba7e39cc809ea5ff9829da11070b21aaf4f306a356d0eb782a7c09f54a"
	if out, _ := hmac.Sha256SumHex(in, key, sum); out != expected {
		format := `hmac.Sha256SumHex("%s", "%s", "%s") = "%x", want "%s"`
		t.Errorf(format, in, key, sum, out, expected)
	}
}

func TestSha256Sum(t *testing.T) {
	in := []byte("goyy")
	key := []byte("key")
	sum := []byte("goyy:")
	expected := "676f79793a840f83ba7e39cc809ea5ff9829da11070b21aaf4f306a356d0eb782a7c09f54a"
	if out, _ := hmac.Sha256Sum(in, key, sum); fmt.Sprintf("%x", out) != expected {
		format := `hmac.Sha256Sum("%s", "%s") = "%x", want "%s"`
		t.Errorf(format, in, key, out, expected)
	}
}
