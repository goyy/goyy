// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package times_test

import (
	"gopkg.in/goyy/goyy.v0/util/times"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	expected := time.Date(2014, 4, 4, 13, 31, 45, 1234454, time.Local)
	if out := times.Add(in, times.Day); out != expected {
		t.Errorf(`times.Add(%v, %s) = %v, want %v`, times.YYMDHMS, in, out, expected)
	}
}

func TestAddUnix(t *testing.T) {
	var expected int64 = 1396589505
	if out := times.AddUnix(i, times.Day); out != expected {
		t.Errorf(`times.AddUnix(%v, %s) = %v, want %v`, times.YYMDHMS, in, out, expected)
	}
}

func TestAddStr(t *testing.T) {
	expected := "2014-04-04 13:31:45"
	if out, _ := times.AddStr(times.YYMDHMS, "2014-04-03 13:31:45", times.Day); out != expected {
		t.Errorf(`times.AddStr(%v, %s) = %v, want %v`, times.YYMDHMS, in, out, expected)
	}
}

func TestAddGMT(t *testing.T) {
	expected := "Fri, 04 Apr 2014 13:31:45 GMT"
	if out, _ := times.AddGMT("Thu, 03 Apr 2014 13:31:45 GMT", times.Day); out != expected {
		t.Errorf(`times.AddGMT(%v) = %v, want %v`, in, out, expected)
	}
}

func TestAddYYMD(t *testing.T) {
	expected := "2014-04-04"
	if out, _ := times.AddYYMD("2014-04-03", times.Day); out != expected {
		t.Errorf(`times.AddYYMD(%v) = %v, want %v`, in, out, expected)
	}
}

func TestAddYYMDHMS(t *testing.T) {
	expected := "2014-04-04 13:31:45"
	if out, _ := times.AddYYMDHMS("2014-04-03 13:31:45", times.Day); out != expected {
		t.Errorf(`times.AddYYMDHMS(%v) = %v, want %v`, in, out, expected)
	}
}

func TestAddYYMDHM(t *testing.T) {
	expected := "2014-04-04 13:31"
	if out, _ := times.AddYYMDHM("2014-04-03 13:31", times.Day); out != expected {
		t.Errorf(`times.AddYYMDHM(%v) = %v, want %v`, in, out, expected)
	}
}
