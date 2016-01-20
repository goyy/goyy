// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package schema

import (
	"gopkg.in/goyy/goyy.v0/util/strings"
)

// new Table
func TABLE(name, comment string) Table {
	if strings.IsBlank(name) {
		panic("Table name are not allowed to empty.")
	}
	t := new(table)
	t.name = strings.ToLower(name)
	t.comment = comment
	t.columns = make(map[string]*column)
	return t
}
