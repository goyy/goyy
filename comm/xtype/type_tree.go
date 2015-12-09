// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xtype

type Tree struct {
	Id       string `json:"id"`
	ParentId string `json:"pId"`
	Name     string `json:"name"`
	Open     bool   `json:"open"`
	IsParent bool   `json:"isParent"`
	Checked  bool   `json:"checked"`
}
