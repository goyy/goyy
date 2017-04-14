// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package entity

import (
	"gopkg.in/goyy/goyy.v0/util/strings"
)

// json information for the field of the entity struct.
type json struct {
	tag       string
	name      string
	omitempty bool
	ignored   bool
}

func (me *json) Tag() string {
	return me.tag
}

func (me *json) SetTag(v string) {
	me.tag = v
	if v == "-" {
		me.ignored = true
		return
	}
	vs := strings.Split(v, ",")
	if len(vs) == 2 && vs[1] == "omitempty" {
		me.omitempty = true
		if vs[0] != "-" {
			me.name = vs[0]
		}
	} else {
		if v == "omitempty" {
			me.omitempty = true
		} else {
			me.name = v
		}
	}
}

func (me *json) Name() string {
	return me.name
}

func (me *json) Omitempty() bool {
	return me.omitempty
}

func (me *json) Ignored() bool {
	return me.ignored
}
