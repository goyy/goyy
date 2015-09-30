// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package uuids implements uuid utility functions.
package uuids

import (
	"github.com/satori/go.uuid"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

// New returns the id string.
func New() string {
	id := uuid.NewV1()
	return strings.Replace(id.String(), "-", "", -1)
}

// New returns the id string.
func NewV1() string {
	return uuid.NewV1().String()
}
