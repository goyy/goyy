// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package times_test

import (
	"gopkg.in/goyy/goyy.v0/util/times"
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	in := "2014-04-03 13:31:45"
	expected := time.Date(2014, 4, 3, 13, 31, 45, 0, time.Local)
	if out, _ := times.Parse(times.YYMDHMS, in); out != expected {
		t.Errorf(`times.Parse(%q, %q) = %v, want %v`, times.YYMDHMS, in, out, expected)
	}
}

func TestParseGMT(t *testing.T) {
	in := "Thu, 03 Apr 2014 13:31:45 GMT"
	expected := time.Date(2014, 4, 3, 13, 31, 45, 0, time.Local)
	if out, _ := times.ParseGMT(in); out != expected {
		t.Errorf(`times.ParseGMT(%q) = %v, want %v`, in, out, expected)
	}
}

func TestParseYYMD(t *testing.T) {
	in := "2014-04-03"
	expected := time.Date(2014, 4, 3, 0, 0, 0, 0, time.Local)
	if out, _ := times.ParseYYMD(in); out != expected {
		t.Errorf(`times.ParseYYMD(%v) = %v, want %v`, i, out, expected)
	}
}

func TestParseYYMDHMS(t *testing.T) {
	in := "2014-04-03 13:31:45"
	expected := time.Date(2014, 4, 3, 13, 31, 45, 0, time.Local)
	if out, _ := times.ParseYYMDHMS(in); out != expected {
		t.Errorf(`times.ParseYYMDHMS(%v) = %v, want %v`, i, out, expected)
	}
}

func TestParseYYMDHM(t *testing.T) {
	in := "2014-04-03 13:31"
	expected := time.Date(2014, 4, 3, 13, 31, 0, 0, time.Local)
	if out, _ := times.ParseYYMDHM(in); out != expected {
		t.Errorf(`times.ParseYYMDHM(%v) = %v, want %v`, i, out, expected)
	}
}
