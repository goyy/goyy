// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package md5_test

import (
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/crypto/md5"
	"testing"
)

func TestDigestHex(t *testing.T) {
	in := "goyy"
	expected := "db9e2a3e99dbace8332b3042a6beb793"
	if out, _ := md5.DigestHex(in); out != expected {
		t.Errorf(`md5.DigestHex("%s") = "%s", want "%s"`, in, out, expected)
	}
}

func TestDigest(t *testing.T) {
	in := []byte("goyy")
	expected := "db9e2a3e99dbace8332b3042a6beb793"
	if out, _ := md5.Digest(in); fmt.Sprintf("%x", out) != expected {
		t.Errorf(`md5.Digest("%s") = "%x", want "%s"`, in, out, expected)
	}
}

func TestDigestSumHex(t *testing.T) {
	in := "goyy"
	sum := "goyy:"
	expected := "676f79793adb9e2a3e99dbace8332b3042a6beb793"
	if out, _ := md5.DigestSumHex(in, sum); out != expected {
		t.Errorf(`md5.DigestSum("%s") = "%x", want "%s"`, in, out, expected)
	}
}

func TestDigestSum(t *testing.T) {
	in := []byte("goyy")
	sum := []byte("goyy:")
	expected := "676f79793adb9e2a3e99dbace8332b3042a6beb793"
	if out, _ := md5.DigestSum(in, sum); fmt.Sprintf("%x", out) != expected {
		t.Errorf(`md5.DigestSum("%s") = "%x", want "%s"`, in, out, expected)
	}
}
