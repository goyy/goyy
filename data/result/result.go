// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package result

import (
	"encoding/json"

	"gopkg.in/goyy/goyy.v0/util/strings"
)

// Result result.Result.
type Result struct {
	Success bool        `json:"success"`
	Token   string      `json:"token"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Memo    string      `json:"memo"`
	Tag     string      `json:"tag"`
	Data    interface{} `json:"data"`
}

// JSON result.Result to JSON.
func (me *Result) JSON() (string, error) {
	b, err := json.Marshal(me)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// ParseJSON JSON to result.Result.
func (me *Result) ParseJSON(jsons string) error {
	if strings.IsBlank(jsons) {
		return nil
	}
	return json.Unmarshal([]byte(jsons), &me)
}
