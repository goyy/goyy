// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package validate

import (
	"fmt"
)

// Error validate.Error.
type Error struct {
	field   string
	typ     string
	message string
}

// Field gets the Error.field.
func (me *Error) Field() string {
	return me.field
}

// Type gets the Error.type.
func (me *Error) Type() string {
	return me.typ
}

// Message gets the Error.message.
func (me *Error) Message() string {
	return me.message
}

func (me *Error) Error() string {
	return fmt.Sprintf("field:%s,typ:%s,message:%s", me.field, me.typ, me.message)
}
