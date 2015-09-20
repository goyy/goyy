// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package errors

import (
	"errors"
	"fmt"
)

const (
	fNotUse            = "Cannot use %s on a non-struct pointer"
	fNotBlanck         = "The %s can't be blank"
	fReqNameNotPresent = "http.Request: named %s not present"
)

// New returns an error that formats as the given text.
func New(text string) error {
	return errors.New(text)
}

// New returns an error that formats as the given text.
func Newf(format string, v ...interface{}) error {
	return fmt.Errorf(format, v...)
}

// NewNonStructPtr returns non-struct pointer error.
func NewNonStructPtr(name string) error {
	return fmt.Errorf(fNotUse, name)
}

// NewNotBlank returns can't be blank error.
func NewNotBlank(name string) error {
	return fmt.Errorf(fNotBlanck, name)
}

// NewReqNotPresent returns http.Request: named not present error.
func NewReqNameNotPresent(name string) error {
	return fmt.Errorf(fReqNameNotPresent, name)
}
