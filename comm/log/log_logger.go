// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package log

import (
	"os"
)

var console *Logger = NewLogger(os.Stderr)

func init() {
	console.isConsole = true
	console.SetLayouts(DefaultLayout)
}
