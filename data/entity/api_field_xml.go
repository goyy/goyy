// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package entity

// Xml information for the field of the entity struct.
type Xml interface {
	Tag() string     // Get Tag
	SetTag(v string) // Set Tag
	Name() string    // Field appears in XML as key name
	Omitempty() bool // The field is omitted from the object if its value is empty
	Ignored() bool   // Field is ignored by this package
}

func NewXml() Xml {
	return &xml{}
}

func NewXmlBy(name string) Xml {
	v := &xml{}
	v.SetTag(name)
	return v
}
