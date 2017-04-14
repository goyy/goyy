// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

type xButtons struct {
	Button []*xButton `xml:"button"`
}

type xButton struct {
	ID      string `xml:"id,attr"`
	Name    string `xml:"name,attr"`
	Comment string `xml:"comment,attr"`
}

type button struct {
	id      string
	name    string
	comment string
}

func (me *button) ID() string {
	return me.id
}

func (me *button) SetID(value string) {
	me.id = value
}

func (me *button) Name() string {
	return me.name
}

func (me *button) SetName(value string) {
	me.name = value
}

func (me *button) Comment() string {
	return me.comment
}

func (me *button) SetComment(value string) {
	me.comment = value
}
