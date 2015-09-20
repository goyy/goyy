// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package hmac_test

import (
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/crypto/hmac"
	"testing"
)

func TestSha1Hex(t *testing.T) {
	in := "goyy"
	key := "key"
	expected := "221f50c7d0d8f13e160be5660a1a27d5bb18861c"
	if out, _ := hmac.Sha1Hex(in, key); out != expected {
		format := `hmac.Sha1Hex("%s", "%s") = "%s", want "%s"`
		t.Errorf(format, in, key, out, expected)
	}
}

func TestSha1(t *testing.T) {
	in := []byte("goyy")
	key := []byte("key")
	expected := "221f50c7d0d8f13e160be5660a1a27d5bb18861c"
	if out, _ := hmac.Sha1(in, key); fmt.Sprintf("%x", out) != expected {
		format := `hmac.Sha1("%s", "%s") = "%x", want "%s"`
		t.Errorf(format, in, key, out, expected)
	}
}

func TestSha1SumHex(t *testing.T) {
	in := "goyy"
	key := "key"
	sum := "goyy:"
	expected := "676f79793a221f50c7d0d8f13e160be5660a1a27d5bb18861c"
	if out, _ := hmac.Sha1SumHex(in, key, sum); out != expected {
		format := `hmac.Sha1SumHex("%s", "%s", "%s") = "%x", want "%s"`
		t.Errorf(format, in, key, sum, out, expected)
	}
}

func TestSha1Sum(t *testing.T) {
	in := []byte("goyy")
	key := []byte("key")
	sum := []byte("goyy:")
	expected := "676f79793a221f50c7d0d8f13e160be5660a1a27d5bb18861c"
	if out, _ := hmac.Sha1Sum(in, key, sum); fmt.Sprintf("%x", out) != expected {
		format := `hmac.Sha1Sum("%s", "%s") = "%x", want "%s"`
		t.Errorf(format, in, key, out, expected)
	}
}
