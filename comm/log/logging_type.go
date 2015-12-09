// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package log

type Logging struct {
	prefix        string
	priority      int
	layouts       int
	outputs       int
	dailyfilename string
	console       *Logger
	dailyfile     *Logger
}

// New creates a new Logging.
func New(prefix string) *Logging {
	return &Logging{
		prefix:   prefix,
		priority: DefaultPriority,
		layouts:  DefaultLayout,
		outputs:  DefaultOutput,
	}
}

// New creates a new Logging.
func NewLogging(prefix string, priority, layouts, outputs int) *Logging {
	return &Logging{
		prefix:   prefix,
		priority: priority,
		layouts:  layouts,
		outputs:  outputs,
	}
}
