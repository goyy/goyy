// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package log_test

import (
	"testing"

	"gopkg.in/goyy/goyy.v0/comm/log"
)

func TestLog(t *testing.T) {
	log.SetPrefix("[log]")
	log.SetPriority(log.Perror)
	log.SetLayouts(log.LstdFlags | log.Llongfile)

	log.Trace("=====trace=====")
	log.Debug("=====debug=====")
	log.Info("=====info=====")
	log.Warn("=====warn=====")
	log.Error("=====error=====")
	log.Critical("=====critical=====")
	log.Print("=====print=====")
}
