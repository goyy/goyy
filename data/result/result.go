// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package result

import (
	"encoding/json"

	"gopkg.in/goyy/goyy.v0/util/strings"
)

type Result struct {
	Success bool        `json:"success"`
	Token   string      `json:"token"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Memo    string      `json:"memo"`
	Tag     string      `json:"tag"`
	Data    interface{} `json:"data"`
}

func (me *Result) JSON() (string, error) {
	b, err := json.Marshal(me)
	if err == nil {
		return string(b), nil
	} else {
		return "", err
	}
}

func (me *Result) ParseJSON(jsons string) error {
	if strings.IsBlank(jsons) {
		return nil
	}
	return json.Unmarshal([]byte(jsons), &me)
}
