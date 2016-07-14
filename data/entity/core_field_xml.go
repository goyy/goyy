// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package entity

import (
	"gopkg.in/goyy/goyy.v0/util/strings"
)

// xml information for the field of the entity struct.
type xml struct {
	name      string
	omitempty bool
	ignored   bool
}

func (me *xml) Name() string {
	return me.name
}

func (me *xml) SetName(v string) {
	vs := strings.Split(v, ",")
	if len(vs) == 2 && vs[1] == "omitempty" {
		me.omitempty = true
		me.name = vs[0]
	} else {
		me.name = v
	}
}

func (me *xml) Omitempty() bool {
	return me.omitempty
}

func (me *xml) Ignored() bool {
	return me.name == "-"
}
