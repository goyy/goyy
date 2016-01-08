// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package result

import (
	"encoding/json"

	"gopkg.in/goyy/goyy.v0/util/strings"
)

type Http struct {
	Success bool              `json:"success"`
	Token   string            `json:"token"`
	Code    string            `json:"code"`    // view message code
	Message string            `json:"message"` // view message content
	State   string            `json:"state"`   // view display state
	Memo    string            `json:"memo"`
	Tag     string            `json:"tag"`
	Params  map[string]string `json:"params"`
	Data    interface{}       `json:"data"`
}

func (me *Http) JSON() (string, error) {
	b, err := json.Marshal(me)
	if err == nil {
		return string(b), nil
	} else {
		return "", err
	}
}

func (me *Http) ParseJSON(jsons string) error {
	if strings.IsBlank(jsons) {
		return nil
	}
	return json.Unmarshal([]byte(jsons), me)
}
