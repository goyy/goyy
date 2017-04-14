// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

type configuration struct {
	Settings *settings
	projects []*project
	modules  []*module
	buttons  []*button
	domains  []*domain
	columns  []*column
	parent   []*table
	tables   []*table
}

type xConfiguration struct {
	Settings *xSettings `xml:"settings"`
	Projects *xProjects `xml:"projects"`
	Modules  *xModules  `xml:"modules"`
	Buttons  *xButtons  `xml:"buttons"`
	Domains  *xDomains  `xml:"domains"`
	Columns  *xColumns  `xml:"columns"`
	Tables   *xTables   `xml:"tables"`
}
