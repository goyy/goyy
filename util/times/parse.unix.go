// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package times

import (
	"gopkg.in/goyy/goyy.v0/util/strings"
	"strconv"
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

// ParseUnixStr parses a formatted string and returns the time unix value it represents.
// The layout defines the format by showing how the reference time.
func ParseUnixStr(layout, value string) (string, error) {
	if v, err := ParseUnix(layout, value); err == nil {
		return strconv.FormatInt(v, 10), nil
	} else {
		return "", err
	}
}

// ParseGMT parses a formatted string and returns the time unix value it represents.
// The layout is "Mon, 02 Jan 2006 15:04:05 GMT"
func ParseUnixGMT(value string) (int64, error) {
	return ParseUnix(GMT, value)
}

// ParseYYMD parses a formatted string and returns the time unix value it represents.
// The layout is "2006-01-02"
func ParseUnixYYMD(value string) (int64, error) {
	return ParseUnix(YYMD, value)
}

// ParseGMT parses a formatted string and returns the time unix value it represents.
// The layout is "2006-01-02 15:04:05"
func ParseUnixYYMDHMS(value string) (int64, error) {
	return ParseUnix(YYMDHMS, value)
}

// ParseGMT parses a formatted string and returns the time unix value it represents.
// The layout is "2006-01-02 15:04"
func ParseUnixYYMDHM(value string) (int64, error) {
	return ParseUnix(YYMDHM, value)
}

// ParseGmt parses a formatted string and returns the time unix value it represents.
// The layout is "Mon, 02 Jan 2006 15:04:05 GMT"
func ParseUnixGmt(value string) (string, error) {
	return ParseUnixStr(GMT, value)
}

// ParseYymd parses a formatted string and returns the time unix value it represents.
// The layout is "2006-01-02"
func ParseUnixYymd(value string) (string, error) {
	return ParseUnixStr(YYMD, value)
}

// ParseYymdhms parses a formatted string and returns the time unix value it represents.
// The layout is "2006-01-02 15:04:05"
func ParseUnixYymdhms(value string) (string, error) {
	return ParseUnixStr(YYMDHMS, value)
}

// ParseYymdhm parses a formatted string and returns the time unix value it represents.
// The layout is "2006-01-02 15:04"
func ParseUnixYymdhm(value string) (string, error) {
	return ParseUnixStr(YYMDHM, value)
}
