// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package times

const (
	// Default returns the default value of time.
	Default int64 = -62135596800

	// YYMD 2006-01-02
	YYMD = "2006-01-02"
	// YYMDHMS 2006-01-02 15:04:05
	YYMDHMS = "2006-01-02 15:04:05"
	// YYMDHM 2006-01-02 15:04
	YYMDHM = "2006-01-02 15:04"
	// GMT Mon, 02 Jan 2006 15:04:05 GMT
	GMT = "Mon, 02 Jan 2006 15:04:05 GMT"

	// Nanosecond 1
	Nanosecond Duration = 1
	// Microsecond 1000 * Nanosecond
	Microsecond = 1000 * Nanosecond
	// Millisecond 1000 * Microsecond
	Millisecond = 1000 * Microsecond
	// Second 1000 * Millisecond
	Second = 1000 * Millisecond
	// Minute 60 * Second
	Minute = 60 * Second
	// Hour 60 * Minute
	Hour = 60 * Minute
	// Day 24 * Hour
	Day = 24 * Hour
)
