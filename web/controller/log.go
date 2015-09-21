// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package controller

import (
	"gopkg.in/goyy/goyy.v0/comm/log"
)

var logger = log.New("[controller]")

func SetPriority(value int) {
	logger.SetPriority(value)
}
