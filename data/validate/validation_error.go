// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package validate

import (
	"fmt"
)

type Error struct {
	field   string
	typ     string
	message string
}

func (me *Error) Error() string {
	return fmt.Sprintf("field:%s,typ:%s,message:%s", me.field, me.typ, me.message)
}
