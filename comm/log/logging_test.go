// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package log_test

import (
	"testing"

	"gopkg.in/goyy/goyy.v0/comm/log"
)

func TestLogging(t *testing.T) {
	logging := log.NewLogging("[logging]", log.Perror, log.LstdFlags, log.Ostd)

	logging.Trace("=====trace=====")
	logging.Debug("=====debug=====")
	logging.Info("=====info=====")
	logging.Warn("=====warn=====")
	logging.Error("=====error=====")
	logging.Critical("=====critical=====")
	logging.Print("=====print=====")
}
