// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package times

import (
	"strconv"
	"time"
)

type Duration time.Duration

// Add returns the time t+d of unix string.
func Add(layout, t string, d Duration) (string, error) {
	if n, err := time.ParseInLocation(layout, t, time.Local); err == nil {
		u := n.Add(time.Duration(d)).Unix()
		return strconv.FormatInt(u, 10), nil
	} else {
		return "", err
	}
}

// Add returns the time t+d.
// The layout is "Mon, 02 Jan 2006 15:04:05 GMT"
func AddGmt(t string, d Duration) (string, error) {
	return Add(GMT, t, d)
}

// Add returns the time t+d.
func UAdd(t int64, d Duration) int64 {
	n := time.Unix(t, 0)
	return n.Add(time.Duration(d)).Unix()
}
