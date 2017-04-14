// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package log

// Default settings.
var (
	isResetDefault  bool
	defaultPriority = Perror
	defaultLayout   = LstdFlags | Lpriority | Llongfile
	defaultOutput   = Oconsole
	defaultDir      = "logs"
)

// SetDefaultPriority set the default priority.
func SetDefaultPriority(value int) {
	isResetDefault = true
	defaultPriority = value
}

// SetDefaultLayout set the default layout.
func SetDefaultLayout(value int) {
	isResetDefault = true
	defaultLayout = value
}

// SetDefaultOutput set the default output.
func SetDefaultOutput(value int) {
	isResetDefault = true
	defaultOutput = value
}

// SetDefaultDir set the default dir.
func SetDefaultDir(value string) {
	isResetDefault = true
	defaultDir = value
}
