// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	e "gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

// jsonField information for the field of the entity struct.
type jsonField struct {
	Tag       string // Tag
	Name      string // Field appears in JSON as key name
	Omitempty bool   // The field is omitted from the object if its value is empty
	Ignored   bool   // Field is ignored by this package
}

// newJsonField analysis of tag to create a new jsonField struct.
func newJsonField(f *field, tag string) (*jsonField, bool) {
	if strings.IsBlank(tag) {
		return nil, false
	}
	v := e.NewJsonBy(tag)
	jf := &jsonField{
		Tag:       v.Tag(),
		Name:      v.Name(),
		Omitempty: v.Omitempty(),
		Ignored:   v.Ignored(),
	}
	return jf, true
}
