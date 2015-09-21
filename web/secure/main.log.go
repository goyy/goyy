// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package secure

import (
	"gopkg.in/goyy/goyy.v0/comm/log"
	"gopkg.in/goyy/goyy.v0/web/secure/internal"
)

var logger = log.New("[secure]")

func SetPriority(value int) {
	logger.SetPriority(value)
	internal.SetPriority(value)
}
