// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dialect

// Oracle oracle dialect.
type Oracle struct {
	typ string
}

// Type return the type of dialect.
func (me *Oracle) Type() string {
	return ORACLE
}
