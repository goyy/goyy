// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dql

var op map[string]string = map[string]string{
	"EQ": " = ",
	"NE": " <> ",
	"GT": " > ",
	"LT": " < ",
	"GE": " >= ",
	"LE": " <= ",
	"LI": " like ",
	"LK": " like ",
	"LL": " like ",
	"LR": " like ",
	"BE": " between ",
	"IN": " in ",
	"NU": " is null ",
	"NN": " is not null ",
	"OA": " asc",
	"OD": " desc",
}
