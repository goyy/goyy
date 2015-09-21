// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dto

type Sys struct {
	Pk
	Memo      string `json:"memo"`
	Creates   string `json:"creates"`
	Creater   string `json:"creater"`
	Created   int64  `json:"created"`
	Modifier  string `json:"modifier"`
	Modified  int64  `json:"modified"`
	Version   int    `json:"version"`
	Deletion  int    `json:"deletion"`
	Artifical int    `json:"artifical"`
	History   int    `json:"history"`
}
