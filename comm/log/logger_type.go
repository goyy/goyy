// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package log

import (
	"io"
	"log"
)

// Logger logger.
type Logger struct {
	isConsole bool
	prefix    string
	priority  int
	logger    *log.Logger
}

// NewLogger creates a new Logger.
func NewLogger(out io.Writer) *Logger {
	return &Logger{logger: log.New(out, "", log.LstdFlags)}
}
