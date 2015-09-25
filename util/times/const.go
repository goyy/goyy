// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package times

const (
	Default int64 = -62135596800

	YYMD    = "2006-01-02"
	YYMDHMS = "2006-01-02 15:04:05"
	YYMDHM  = "2006-01-02 15:04"
	GMT     = "Mon, 02 Jan 2006 15:04:05 GMT"

	Nanosecond  Duration = 1
	Microsecond          = 1000 * Nanosecond
	Millisecond          = 1000 * Microsecond
	Second               = 1000 * Millisecond
	Minute               = 60 * Second
	Hour                 = 60 * Minute
	Day                  = 24 * Hour
)
