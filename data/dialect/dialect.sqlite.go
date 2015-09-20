// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dialect

type Sqlite struct {
	typ string
}

func (me *Sqlite) Type() string {
	return SQLITE
}
