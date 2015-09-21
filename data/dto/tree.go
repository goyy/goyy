// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dto

type Tree struct {
	Sys
	Code        string `json:"code"`
	Name        string `json:"name"`
	Fullname    string `json:"fullname"`
	Genre       string `json:"genre"`
	Leaf        int    `json:"leaf"`
	Grade       int    `json:"grade"`
	Ordinal     string `json:"ordinal"`
	ParentId    string `json:"parentId"`
	ParentIds   string `json:"parentIds"`
	ParentCodes string `json:"parentCodes"`
	ParentNames string `json:"parentNames"`
}
