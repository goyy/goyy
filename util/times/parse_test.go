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
	var expected int64 = 1396503105
	if out, _ := times.ParseUnix(times.YYMDHMS, in); out != expected {
		t.Errorf(`times.ParseUnix(%q, %q) = %v, want %v`, times.YYMDHMS, in, out, expected)
	}
}

func TestParseUnixStr(t *testing.T) {
	in := "2014-04-03 13:31:45"
	expected := "1396503105"
	if out, _ := times.ParseUnixStr(times.YYMDHMS, in); out != expected {
		t.Errorf(`times.ParseUnixStr(%q, %q) = %v, want %v`, times.YYMDHMS, in, out, expected)
	}
}

func TestParseUnixNano(t *testing.T) {
	in := "2014-04-03 13:31:45"
	var expected int64 = 1396503105000000000
	if out, _ := times.ParseUnixNano(times.YYMDHMS, in); out != expected {
		t.Errorf(`times.ParseUnixNano(%q, %q) = %v, want %v`, times.YYMDHMS, in, out, expected)
	}
}

func TestParseGMT(t *testing.T) {
	in := "Thu, 03 Apr 2014 13:31:45 GMT"
	var expected int64 = 1396503105
	if out, _ := times.ParseGMT(in); out != expected {
		t.Errorf(`times.ParseGMT(%q) = %v, want %v`, in, out, expected)
	}
}

func TestParseYYMD(t *testing.T) {
	in := "2014-04-03"
	var expected int64 = 1396454400
	if out, _ := times.ParseYYMD(in); out != expected {
		t.Errorf(`times.ParseYYMD(%v) = %v, want %v`, i, out, expected)
	}
}

func TestParseYYMDHMS(t *testing.T) {
	in := "2014-04-03 13:31:45"
	var expected int64 = 1396503105
	if out, _ := times.ParseYYMDHMS(in); out != expected {
		t.Errorf(`times.ParseYYMDHMS(%v) = %v, want %v`, i, out, expected)
	}
}

func TestParseYYMDHM(t *testing.T) {
	in := "2014-04-03 13:31"
	var expected int64 = 1396503060
	if out, _ := times.ParseYYMDHM(in); out != expected {
		t.Errorf(`times.ParseYYMDHM(%v) = %v, want %v`, i, out, expected)
	}
}

func TestParseGmt(t *testing.T) {
	in := "Thu, 03 Apr 2014 13:31:45 GMT"
	expected := "1396503105"
	if out, _ := times.ParseGmt(in); out != expected {
		t.Errorf(`times.ParseGmt(%q) = %v, want %v`, in, out, expected)
	}
}

func TestParseYymd(t *testing.T) {
	in := "2014-04-03"
	expected := "1396454400"
	if out, _ := times.ParseYymd(in); out != expected {
		t.Errorf(`times.ParseYymd(%v) = %v, want %v`, in, out, expected)
	}
}

func TestParseYymdhms(t *testing.T) {
	in := "2014-04-03 13:31:45"
	expected := "1396503105"
	if out, _ := times.ParseYymdhms(in); out != expected {
		t.Errorf(`times.ParseYymdhms(%v) = %v, want %v`, in, out, expected)
	}
}

func TestParseYymdhm(t *testing.T) {
	in := "2014-04-03 13:31"
	expected := "1396503060"
	if out, _ := times.ParseYymdhm(in); out != expected {
		t.Errorf(`times.ParseYymdhm(%v) = %v, want %v`, in, out, expected)
	}
}
