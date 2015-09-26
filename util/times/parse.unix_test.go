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

func TestParseUnixNano(t *testing.T) {
	in := "2014-04-03 13:31:45"
	var expected int64 = 1396503105000000000
	if out, _ := times.ParseUnixNano(times.YYMDHMS, in); out != expected {
		t.Errorf(`times.ParseUnixNano(%q, %q) = %v, want %v`, times.YYMDHMS, in, out, expected)
	}
}

func TestParseUnixStr(t *testing.T) {
	in := "2014-04-03 13:31:45"
	expected := "1396503105"
	if out, _ := times.ParseUnixStr(times.YYMDHMS, in); out != expected {
		t.Errorf(`times.ParseUnixStr(%q, %q) = %v, want %v`, times.YYMDHMS, in, out, expected)
	}
}

func TestParseUnixGMT(t *testing.T) {
	in := "Thu, 03 Apr 2014 13:31:45 GMT"
	var expected int64 = 1396503105
	if out, _ := times.ParseUnixGMT(in); out != expected {
		t.Errorf(`times.ParseUnixGMT(%q) = %v, want %v`, in, out, expected)
	}
}

func TestParseUnixYYMD(t *testing.T) {
	in := "2014-04-03"
	var expected int64 = 1396454400
	if out, _ := times.ParseUnixYYMD(in); out != expected {
		t.Errorf(`times.ParseUnixYYMD(%v) = %v, want %v`, i, out, expected)
	}
}

func TestParseUnixYYMDHMS(t *testing.T) {
	in := "2014-04-03 13:31:45"
	var expected int64 = 1396503105
	if out, _ := times.ParseUnixYYMDHMS(in); out != expected {
		t.Errorf(`times.ParseUnixYYMDHMS(%v) = %v, want %v`, i, out, expected)
	}
}

func TestParseUnixYYMDHM(t *testing.T) {
	in := "2014-04-03 13:31"
	var expected int64 = 1396503060
	if out, _ := times.ParseUnixYYMDHM(in); out != expected {
		t.Errorf(`times.ParseUnixYYMDHM(%v) = %v, want %v`, i, out, expected)
	}
}

func TestParseUnixGmt(t *testing.T) {
	in := "Thu, 03 Apr 2014 13:31:45 GMT"
	expected := "1396503105"
	if out, _ := times.ParseUnixGmt(in); out != expected {
		t.Errorf(`times.ParseUnixGmt(%q) = %v, want %v`, in, out, expected)
	}
}

func TestParseUnixUnixYymd(t *testing.T) {
	in := "2014-04-03"
	expected := "1396454400"
	if out, _ := times.ParseUnixYymd(in); out != expected {
		t.Errorf(`times.ParseUnixYymd(%v) = %v, want %v`, in, out, expected)
	}
}

func TestParseUnixYymdhms(t *testing.T) {
	in := "2014-04-03 13:31:45"
	expected := "1396503105"
	if out, _ := times.ParseUnixYymdhms(in); out != expected {
		t.Errorf(`times.ParseUnixYymdhms(%v) = %v, want %v`, in, out, expected)
	}
}

func TestParseUnixYymdhm(t *testing.T) {
	in := "2014-04-03 13:31"
	expected := "1396503060"
	if out, _ := times.ParseUnixYymdhm(in); out != expected {
		t.Errorf(`times.ParseUnixYymdhm(%v) = %v, want %v`, in, out, expected)
	}
}
