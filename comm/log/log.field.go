// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package log

// Prefix returns the output prefix for the logger.
func Prefix() string {
	return console.Prefix()
}

// SetPrefix sets the output prefix for the logger.
func SetPrefix(value string) {
	console.SetPrefix(value)
}

// Priority returns the output priority for the logger.
func Priority() int {
	return console.Priority()
}

// SetPriority sets the output priority for the logger.
func SetPriority(value int) {
	console.SetPriority(value)
}

// Layouts returns the output layouts for the logger.
func Layouts() int {
	return console.Layouts()
}

// SetLayouts sets the output layouts for the logger.
func SetLayouts(value int) {
	console.SetLayouts(value)
}
