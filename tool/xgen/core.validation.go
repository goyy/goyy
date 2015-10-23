// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

// validation is a field validate in an entity struct.
type validation struct {
	Name  string
	Value string
}

type validations []*validation
