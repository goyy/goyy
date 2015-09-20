// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package log

// Prefix returns the output prefix for the logger.
func (me *Logger) Prefix() string {
	return me.prefix
}

// SetPrefix sets the output prefix for the logger.
func (me *Logger) SetPrefix(prefix string) {
	me.prefix = prefix
}

// Priority returns the output priority for the logger.
func (me *Logger) Priority() int {
	return me.priority
}

// SetPriority sets the output priority for the logger.
func (me *Logger) SetPriority(priority int) {
	me.priority = priority
}

// Layouts returns the output layouts for the logger.
func (me *Logger) Layouts() int {
	return me.logger.Flags()
}

// SetLayouts sets the output layouts for the logger.
func (me *Logger) SetLayouts(layouts int) {
	me.logger.SetFlags(layouts)
}
