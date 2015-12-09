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
func Format(layout string, t time.Time) string {
	if t.IsZero() {
		return ""
	} else {
		return t.Format(layout)
	}
}

// Time formatted as Mon, 02 Jan 2006 15:04:05 GMT.
func FormatGMT(t time.Time) string {
	return Format(GMT, t)
}

// Time formatted as 2006-01-02.
func FormatYYMD(t time.Time) string {
	return Format(YYMD, t)
}

// Time formatted as 2006-01-02 15:04:05.
func FormatYYMDHMS(t time.Time) string {
	return Format(YYMDHMS, t)
}

// Time formatted as 2006-01-02 15:04.
func FormatYYMDHM(t time.Time) string {
	return Format(YYMDHM, t)
}
