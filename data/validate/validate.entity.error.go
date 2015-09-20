// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package validate

type (
	Error struct {
		field, typ, message string
		isReMsg             bool
		reMsg               string
	}
)
