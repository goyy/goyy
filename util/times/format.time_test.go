// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package times_test

import (
	"gopkg.in/goyy/goyy.v0/util/times"
	"testing"
)

func TestFormat(t *testing.T) {
	expected := "2014-04-03 13:31:45"
	if out := times.Format(times.YYMDHMS, in); out != expected {
		t.Errorf(`times.Format(%v, %s) = %v, want %v`, times.YYMDHMS, in, out, expected)
	}
}

func TestFormatGMT(t *testing.T) {
	expected := "Thu, 03 Apr 2014 13:31:45 GMT"
	if out := times.FormatGMT(in); out != expected {
		t.Errorf(`times.FormatGMT(%v) = %v, want %v`, in, out, expected)
	}
}

func TestFormatYYMD(t *testing.T) {
	expected := "2014-04-03"
	if out := times.FormatYYMD(in); out != expected {
		t.Errorf(`times.FormatYYMD(%v) = %v, want %v`, in, out, expected)
	}
}

func TestFormatYYMDHMS(t *testing.T) {
	expected := "2014-04-03 13:31:45"
	if out := times.FormatYYMDHMS(in); out != expected {
		t.Errorf(`times.FormatYYMDHMS(%v) = %v, want %v`, in, out, expected)
	}
}

func TestFormatYYMDHM(t *testing.T) {
	expected := "2014-04-03 13:31"
	if out := times.FormatYYMDHM(in); out != expected {
		t.Errorf(`times.FormatYYMDHM(%v) = %v, want %v`, in, out, expected)
	}
}
