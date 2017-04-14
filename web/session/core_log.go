// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package session

import (
	"gopkg.in/goyy/goyy.v0/comm/log"
)

var logger = log.New("[session]")

// SetPriority set the priority of the logger.
func SetPriority(value int) {
	logger.SetPriority(value)
}
