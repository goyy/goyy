// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package log

// Prefix returns the output prefix for the logging.
func (me *Logging) Prefix() string {
	return me.prefix
}

// SetPrefix sets the output prefix for the logging.
func (me *Logging) SetPrefix(value string) {
	me.prefix = value
	if me.console != nil {
		me.console.SetPrefix(me.prefix)
	}
	if me.dailyfile != nil {
		me.dailyfile.SetPrefix(me.prefix)
	}
}

// Priority returns the output priority for the logging.
func (me *Logging) Priority() int {
	return me.priority
}

// SetPriority sets the output priority for the logging.
func (me *Logging) SetPriority(value int) {
	me.priority = value
	if me.console != nil {
		me.console.SetPriority(me.priority)
	}
	if me.dailyfile != nil {
		me.dailyfile.SetPriority(me.priority)
	}
}

// Layouts returns the output layouts for the logging.
func (me *Logging) Layouts() int {
	return me.layouts
}

// SetLayouts sets the output layouts for the logging.
func (me *Logging) SetLayouts(value int) {
	me.layouts = value
	if me.console != nil {
		me.console.SetLayouts(me.layouts)
	}
	if me.dailyfile != nil {
		me.dailyfile.SetLayouts(me.layouts)
	}
}

// Outputs returns the output outputs for the logging.
func (me *Logging) Outputs() int {
	return me.outputs
}

// SetOutputs sets the output outputs for the logging.
func (me *Logging) SetOutputs(value int) {
	me.outputs = value
}
