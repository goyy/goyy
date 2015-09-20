// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package times_test

import (
	"gopkg.in/goyy/goyy.v0/util/times"
	"testing"
	"time"
)

var in = time.Date(2014, 4, 3, 13, 31, 45, 1234454, time.Local)
var i int64 = in.Unix()

func TestYymd(t *testing.T) {
	expected := "2014-04-03"
	if out := times.Yymd(in); out != expected {
		t.Errorf(`times.Yymd(%v) = %v, want %v`, in, out, expected)
	}
}

func TestYymdhms(t *testing.T) {
	expected := "2014-04-03 13:31:45"
	if out := times.Yymdhms(in); out != expected {
		t.Errorf(`times.Yymdhms(%v) = %v, want %v`, in, out, expected)
	}
}

func TestYymdhm(t *testing.T) {
	expected := "2014-04-03 13:31"
	if out := times.Yymdhm(in); out != expected {
		t.Errorf(`times.Yymdhm(%v) = %v, want %v`, in, out, expected)
	}
}

func TestUyymd(t *testing.T) {
	expected := "2014-04-03"
	if out := times.Uyymd(i); out != expected {
		t.Errorf(`times.Uyymd(%v) = %v, want %v`, i, out, expected)
	}
}

func TestUyymdhms(t *testing.T) {
	expected := "2014-04-03 13:31:45"
	if out := times.Uyymdhms(i); out != expected {
		t.Errorf(`times.Uyymdhms(%v) = %v, want %v`, i, out, expected)
	}
}

func TestUyymdhm(t *testing.T) {
	expected := "2014-04-03 13:31"
	if out := times.Uyymdhm(i); out != expected {
		t.Errorf(`times.Uyymdhm(%v) = %v, want %v`, i, out, expected)
	}
}
