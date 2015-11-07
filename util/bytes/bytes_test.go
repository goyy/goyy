// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package bytes_test

import (
	"gopkg.in/goyy/goyy.v0/util/bytes"
	"testing"
)

func TestTrimRightNul(t *testing.T) {
	a := [10]byte{1, 2, 3}
	src := a[:]
	expected := "\x01\x02\x03\x00\x00\x00\x00\x00\x00\x00"
	if out := string(src); out != expected {
		t.Errorf(`bytes.TrimRightNul("%s") = %v, want %v`, src, out, expected)
	}
	expected = "\x01\x02\x03"
	if out := string(bytes.TrimRightNul(src)); out != expected {
		t.Errorf(`bytes.TrimRightNul("%s") = %v, want %v`, src, out, expected)
	}
}
