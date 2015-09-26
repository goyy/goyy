// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package times_test

import (
	"gopkg.in/goyy/goyy.v0/util/times"
	"testing"
)

func TestFormatUnix(t *testing.T) {
	expected := "2014-04-03 13:31:45"
	if out := times.FormatUnix(times.YYMDHMS, i); out != expected {
		t.Errorf(`times.FormatUnix(%v, %s) = %v, want %v`, times.YYMDHMS, i, out, expected)
	}
}

func TestFormatUnixGMT(t *testing.T) {
	expected := "Thu, 03 Apr 2014 13:31:45 GMT"
	if out := times.FormatUnixGMT(i); out != expected {
		t.Errorf(`times.FormatUnixGMT(%v) = %v, want %v`, i, out, expected)
	}
}

func TestFormatUnixYYMD(t *testing.T) {
	expected := "2014-04-03"
	if out := times.FormatUnixYYMD(i); out != expected {
		t.Errorf(`times.FormatUnixYYMD(%v) = %v, want %v`, i, out, expected)
	}
}

func TestFormatUnixYYMDHMS(t *testing.T) {
	expected := "2014-04-03 13:31:45"
	if out := times.FormatUnixYYMDHMS(i); out != expected {
		t.Errorf(`times.FormatUnixYYMDHMS(%v) = %v, want %v`, i, out, expected)
	}
}

func TestFormatUnixYYMDHM(t *testing.T) {
	expected := "2014-04-03 13:31"
	if out := times.FormatUnixYYMDHM(i); out != expected {
		t.Errorf(`times.FormatUnixYYMDHM(%v) = %v, want %v`, i, out, expected)
	}
}
