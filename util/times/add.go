// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package times

import (
	"time"
)

// Duration returns the time.Duration
type Duration time.Duration

// Add returns the time t+d.
func Add(t time.Time, d Duration) time.Time {
	return t.Add(time.Duration(d))
}

// AddUnix returns the time t+d.
func AddUnix(t int64, d Duration) int64 {
	n := time.Unix(t, 0)
	return n.Add(time.Duration(d)).Unix()
}

// AddStr returns the time t+d of unix string.
func AddStr(layout, t string, d Duration) (string, error) {
	n, err := Parse(layout, t)
	if err == nil {
		u := n.Add(time.Duration(d))
		return Format(layout, u), nil
	}
	return "", err
}

// AddGMT returns the time t+d of unix string.
// The layout is "Mon, 02 Jan 2006 15:04:05 GMT"
func AddGMT(t string, d Duration) (string, error) {
	return AddStr(GMT, t, d)
}

// AddYYMD returns the time t+d of unix string.
// The layout is "2006-01-02"
func AddYYMD(t string, d Duration) (string, error) {
	return AddStr(YYMD, t, d)
}

// AddYYMDHMS returns the time t+d of unix string.
// The layout is "2006-01-02 15:04:05"
func AddYYMDHMS(t string, d Duration) (string, error) {
	return AddStr(YYMDHMS, t, d)
}

// AddYYMDHM returns the time t+d of unix string.
// The layout is "2006-01-02 15:04"
func AddYYMDHM(t string, d Duration) (string, error) {
	return AddStr(YYMDHM, t, d)
}
