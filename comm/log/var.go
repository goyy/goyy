// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package log

// Default settings.
var (
	DefaultPriority = Perror
	DefaultLayout   = LstdFlags | Lpriority | Llongfile
	DefaultOutput   = Ostd
)
