// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xtype

// Mappings mappings.
type Mappings struct {
	URL     string    `xml:"url,attr"`
	Dir     string    `xml:"dir,attr"`
	Mapping []Mapping `xml:"mapping"`
}

// Mapping mapping.
type Mapping struct {
	Path string `xml:"path,attr"`
	Dir  string `xml:"dir,attr"`
}
