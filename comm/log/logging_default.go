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
)

func SetDefaultPriority(value int) {
	isResetDefault = true
	defaultPriority = value
}

func SetDefaultLayout(value int) {
	isResetDefault = true
	defaultLayout = value
}

func SetDefaultOutput(value int) {
	isResetDefault = true
	defaultOutput = value
}
