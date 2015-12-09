// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package times

import (
	"gopkg.in/goyy/goyy.v0/util/strings"
	"time"
)

// Parse parses a formatted string and returns the time value it represents.
// The layout defines the format by showing how the reference time.
func Parse(layout, value string) (out time.Time, err error) {
	if strings.IsBlank(value) || strings.IsBlank(layout) {
		return
	}
	out, err = time.ParseInLocation(layout, value, time.Local)
	return
}

// ParseGMT parses a formatted string and returns the time value it represents.
// The layout is "Mon, 02 Jan 2006 15:04:05 GMT"
func ParseGMT(value string) (time.Time, error) {
	return Parse(GMT, value)
}

// ParseYYMD parses a formatted string and returns the time value it represents.
// The layout is "2006-01-02"
func ParseYYMD(value string) (time.Time, error) {
	return Parse(YYMD, value)
}

// ParseGMT parses a formatted string and returns the time value it represents.
// The layout is "2006-01-02 15:04:05"
func ParseYYMDHMS(value string) (time.Time, error) {
	return Parse(YYMDHMS, value)
}

// ParseGMT parses a formatted string and returns the time value it represents.
// The layout is "2006-01-02 15:04"
func ParseYYMDHM(value string) (time.Time, error) {
	return Parse(YYMDHM, value)
}
