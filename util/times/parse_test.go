// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package times_test

import (
	"gopkg.in/goyy/goyy.v0/util/times"
	"testing"
)

func TestParseUnix(t *testing.T) {
	in := "2014-04-03 13:31:45"
	var expected int64 = 1396531905
	if out, _ := times.ParseUnix(times.YYMDHMS, in); out != expected {
		t.Errorf(`times.ParseUnix(%q, %q) = %v, want %v`, times.YYMDHMS, in, out, expected)
	}
}

func TestParseUnixNano(t *testing.T) {
	in := "2014-04-03 13:31:45"
	var expected int64 = 1396531905000000000
	if out, _ := times.ParseUnixNano(times.YYMDHMS, in); out != expected {
		t.Errorf(`times.ParseUnixNano(%q, %q) = %v, want %v`, times.YYMDHMS, in, out, expected)
	}
}
