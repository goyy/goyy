// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package log

type Logging struct {
	hasSettings   bool
	Settings      func() // The function executed before the log is printed and executed only once.
	prefix        string
	isDefaultConf bool
	priority      int
	setPriority   bool
	layouts       int
	setLayouts    bool
	outputs       int
	setOutputs    bool
	dailyfilename string
	console       *Logger
	dailyfile     *Logger
}

// New creates a new Logging.
func New(prefix string) *Logging {
	return &Logging{
		prefix:        prefix,
		isDefaultConf: true,
		priority:      defaultPriority,
		layouts:       defaultLayout,
		outputs:       defaultOutput,
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
