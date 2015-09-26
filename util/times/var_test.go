// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package times_test

import (
	"strconv"
	"time"
)

var in = time.Date(2014, 4, 3, 13, 31, 45, 1234454, time.Local)
var i int64 = in.Unix()
var s string = strconv.FormatInt(i, 10)
