// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package sha256_test

import (
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/crypto/sha256"
	"testing"
)

func TestDigestHex(t *testing.T) {
	in := "goyy"
	expected := "294d5a5ab822caec5c8ef8db496a44a5f185d83feb0699ab69e79b51166a1116"
	if out, _ := sha256.DigestHex(in); out != expected {
		t.Errorf(`sha256.DigestHex("%s") = "%s", want "%s"`, in, out, expected)
	}
}

func TestDigest(t *testing.T) {
	in := []byte("goyy")
	expected := "294d5a5ab822caec5c8ef8db496a44a5f185d83feb0699ab69e79b51166a1116"
	if out, _ := sha256.Digest(in); fmt.Sprintf("%x", out) != expected {
		t.Errorf(`sha256.Digest("%s") = "%x", want "%s"`, in, out, expected)
	}
}

func TestDigestSumHex(t *testing.T) {
	in := "goyy"
	sum := "goyy:"
	expected := "676f79793a294d5a5ab822caec5c8ef8db496a44a5f185d83feb0699ab69e79b51166a1116"
	if out, _ := sha256.DigestSumHex(in, sum); out != expected {
		t.Errorf(`sha256.DigestSumHex("%s") = "%x", want "%s"`, in, out, expected)
	}
}

func TestDigestSum(t *testing.T) {
	in := []byte("goyy")
	sum := []byte("goyy:")
	expected := "676f79793a294d5a5ab822caec5c8ef8db496a44a5f185d83feb0699ab69e79b51166a1116"
	if out, _ := sha256.DigestSum(in, sum); fmt.Sprintf("%x", out) != expected {
		t.Errorf(`sha256.DigestSum("%s") = "%x", want "%s"`, in, out, expected)
	}
}
