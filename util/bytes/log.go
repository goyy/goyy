// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package bytes

import (
	"gopkg.in/goyy/goyy.v0/comm/log"
)

var logger = log.New("[bytes]")

// SetPriority sets the output priority for the logging.
func SetPriority(value int) {
	logger.SetPriority(value)
}
