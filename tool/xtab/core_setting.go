// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

type xSettings struct {
	Statement *xStatement `xml:"statement"`
}

type xStatement struct {
	Seperator string `xml:"seperator,attr"`
	Case      string `xml:"case,attr"`
	Comment   bool   `xml:"comment,attr"`
}

type settings struct {
	Statement *statement
}

type statement struct {
	Seperator string
	Case      string
	Comment   bool
}
