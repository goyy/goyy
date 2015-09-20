// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dialect

type PostgreSQL struct {
	typ string
}

func (me *PostgreSQL) Type() string {
	return POSTGRESQL
}
