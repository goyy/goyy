// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package times

import (
	"time"
)

// Format returns a textual representation of the time value formatted
// according to layout, which defines the format by showing how the reference
// time, defined to be
//	Mon Jan 2 15:04:05 -0700 MST 2006
// would be displayed if it were the value; it serves as an example of the
// desired output. The same display rules will then be applied to the time
// value.
//
// A fractional second is represented by adding a period and zeros
// to the end of the seconds section of layout string, as in "15:04:05.000"
// to format a time stamp with millisecond precision.
//
// Predefined layouts ANSIC, UnixDate, RFC3339 and others describe standard
// and convenient representations of the reference time. For more information
// about the formats and the definition of the reference time, see the
// documentation for ANSIC and the other constants defined by this package.
func Format(t time.Time, layout string) string {
	if t.IsZero() {
		return ""
	} else {
		return t.Format(layout)
	}
}

// Time formatted as Mon, 02 Jan 2006 15:04:05 GMT.
func Gmt(t time.Time) string {
	return Format(t, GMT)
}

// Time formatted as 2006-01-02.
func Yymd(t time.Time) string {
	return Format(t, YYMD)
}

// Time formatted as 2006-01-02 15:04:05.
func Yymdhms(t time.Time) string {
	return Format(t, YYMDHMS)
}

// Time formatted as 2006-01-02 15:04.
func Yymdhm(t time.Time) string {
	return Format(t, YYMDHM)
}

// Uformat returns a textual representation of the time value formatted
// according to layout, which defines the format by showing how the reference
// time, defined to be
//	Mon Jan 2 15:04:05 -0700 MST 2006
// would be displayed if it were the value; it serves as an example of the
// desired output. The same display rules will then be applied to the time
// value.
//
// A fractional second is represented by adding a period and zeros
// to the end of the seconds section of layout string, as in "15:04:05.000"
// to format a time stamp with millisecond precision.
//
// Predefined layouts ANSIC, UnixDate, RFC3339 and others describe standard
// and convenient representations of the reference time. For more information
// about the formats and the definition of the reference time, see the
// documentation for ANSIC and the other constants defined by this package.
func Uformat(i int64, layout string) string {
	if i == 0 || i == Default {
		return ""
	}
	t := time.Unix(i, 0)
	return Format(t, layout)
}

// Time formatted as Mon, 02 Jan 2006 15:04:05 GMT.
func Ugmt(i int64) string {
	return Uformat(i, GMT)
}

// Time formatted as 2006-01-02.
func Uyymd(i int64) string {
	return Uformat(i, YYMD)
}

// Time formatted as 2006-01-02 15:04:05.
func Uyymdhms(i int64) string {
	return Uformat(i, YYMDHMS)
}

// Time formatted as 2006-01-02 15:04.
func Uyymdhm(i int64) string {
	return Uformat(i, YYMDHM)
}
