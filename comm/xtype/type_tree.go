// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xtype

// Tree tree.
type Tree struct {
	ID       string `json:"id"`
	ParentID string `json:"pId"`
	Name     string `json:"name"`
	Open     bool   `json:"open"`
	IsParent bool   `json:"isParent"`
	Checked  bool   `json:"checked"`
}
