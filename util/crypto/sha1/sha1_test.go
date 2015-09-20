// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package sha1_test

import (
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/crypto/sha1"
	"testing"
)

func TestDigestHex(t *testing.T) {
	in := "goyy"
	expected := "9a5de4d2e62e2c0f3018eeff35e09ab3d41781fb"
	if out, _ := sha1.DigestHex(in); out != expected {
		t.Errorf(`sha1.DigestHex("%s") = "%s", want "%s"`, in, out, expected)
	}
}

func TestDigest(t *testing.T) {
	in := []byte("goyy")
	expected := "9a5de4d2e62e2c0f3018eeff35e09ab3d41781fb"
	if out, _ := sha1.Digest(in); fmt.Sprintf("%x", out) != expected {
		t.Errorf(`sha1.Digest("%s") = "%x", want "%s"`, in, out, expected)
	}
}

func TestDigestSumHex(t *testing.T) {
	in := "goyy"
	sum := "goyy:"
	expected := "676f79793a9a5de4d2e62e2c0f3018eeff35e09ab3d41781fb"
	if out, _ := sha1.DigestSumHex(in, sum); out != expected {
		t.Errorf(`sha1.DigestSumHex("%s") = "%x", want "%s"`, in, out, expected)
	}
}

func TestDigestSum(t *testing.T) {
	in := []byte("goyy")
	sum := []byte("goyy:")
	expected := "676f79793a9a5de4d2e62e2c0f3018eeff35e09ab3d41781fb"
	if out, _ := sha1.DigestSum(in, sum); fmt.Sprintf("%x", out) != expected {
		t.Errorf(`sha1.DigestSum("%s") = "%x", want "%s"`, in, out, expected)
	}
}
