// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package result

type Result struct {
	Success bool        `json:"success"`
	Token   string      `json:"token"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Memo    string      `json:"memo"`
	Tag     string      `json:"tag"`
	Data    interface{} `json:"data"`
}
