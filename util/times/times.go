// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package times

import (
	"strconv"
	"time"
)

// Now returns the current local time of unix.
func Now() time.Time {
	return time.Now()
}

// Now returns the current local time of unix.
func NowUnix() int64 {
	return time.Now().Unix()
}

// Now returns the current local time of unix.
func NowUnixStr() string {
	return strconv.FormatInt(NowUnix(), 10)
}

// Unix returns the local Time corresponding to the given Unix time,
// sec seconds and nsec nanoseconds since January 1, 1970 UTC.
// It is valid to pass nsec outside the range [0, 999999999].
// Not all sec values have a corresponding time value. One such
// value is 1<<63-1 (the largest int64 value).
func Unix(sec int64, nsec int64) time.Time {
	return time.Unix(sec, nsec)
}
