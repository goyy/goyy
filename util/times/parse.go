// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package times

import (
	"gopkg.in/goyy/goyy.v0/util/strings"
	"time"
)

// ParseUnix parses a formatted string and returns the time unix value it represents.
// The layout defines the format by showing how the reference time.
func ParseUnix(layout, value string) (int64, error) {
	if strings.IsBlank(value) {
		return Default, nil
	}
	t, err := time.ParseInLocation(layout, value, time.Local)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

// ParseUnixNano parses a formatted string and returns the time unix nano value it represents.
// The layout defines the format by showing how the reference time.
func ParseUnixNano(layout, value string) (int64, error) {
	if strings.IsBlank(value) {
		return Default, nil
	}
	t, err := time.ParseInLocation(layout, value, time.Local)
	if err != nil {
		return 0, err
	}
	return t.UnixNano(), nil
}

// ParseGMT parses a formatted string and returns the time unix value it represents.
// The layout is "Mon, 02 Jan 2006 15:04:05 GMT"
func ParseGMT(value string) (int64, error) {
	return ParseUnix(GMT, value)
}
