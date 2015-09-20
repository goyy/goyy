// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package sha512_test

import (
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/crypto/sha512"
	"testing"
)

func TestDigestHex(t *testing.T) {
	in := "goyy"
	expected := "1aa5692a60d265ec371f0f63e59deb0a3bfce9e07ec54de56b75aea9cb0a2ca96fd868d2866843bd0c9d82aa93a696ab557a6cbb57c22e16eab4bd2553198fad"
	if out, _ := sha512.DigestHex(in); out != expected {
		t.Errorf(`sha512.DigestHex("%s") = "%s", want "%s"`, in, out, expected)
	}
}

func TestDigest(t *testing.T) {
	in := []byte("goyy")
	expected := "1aa5692a60d265ec371f0f63e59deb0a3bfce9e07ec54de56b75aea9cb0a2ca96fd868d2866843bd0c9d82aa93a696ab557a6cbb57c22e16eab4bd2553198fad"
	if out, _ := sha512.Digest(in); fmt.Sprintf("%x", out) != expected {
		t.Errorf(`sha512.Digest("%s") = "%x", want "%s"`, in, out, expected)
	}
}

func TestDigestSumHex(t *testing.T) {
	in := "goyy"
	sum := "goyy:"
	expected := "676f79793a1aa5692a60d265ec371f0f63e59deb0a3bfce9e07ec54de56b75aea9cb0a2ca96fd868d2866843bd0c9d82aa93a696ab557a6cbb57c22e16eab4bd2553198fad"
	if out, _ := sha512.DigestSumHex(in, sum); out != expected {
		t.Errorf(`sha512.DigestSumHex("%s") = "%x", want "%s"`, in, out, expected)
	}
}

func TestDigestSum(t *testing.T) {
	in := []byte("goyy")
	sum := []byte("goyy:")
	expected := "676f79793a1aa5692a60d265ec371f0f63e59deb0a3bfce9e07ec54de56b75aea9cb0a2ca96fd868d2866843bd0c9d82aa93a696ab557a6cbb57c22e16eab4bd2553198fad"
	if out, _ := sha512.DigestSum(in, sum); fmt.Sprintf("%x", out) != expected {
		t.Errorf(`sha512.DigestSum("%s") = "%x", want "%s"`, in, out, expected)
	}
}
