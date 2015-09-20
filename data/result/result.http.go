// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package result

type Http struct {
	Success bool              `json:"success"`
	Token   string            `json:"token"`
	Code    string            `json:"code"`    // view message code
	Message string            `json:"message"` // view message content
	State   string            `json:"state"`   // view display state
	Params  map[string]string `json:"params"`
	Data    interface{}       `json:"data"`
}
