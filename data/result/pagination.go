// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package result

import (
	"encoding/json"

	"gopkg.in/goyy/goyy.v0/util/strings"
)

const (
	defaultPageNo  = 1
	defaultPageNot = -1
)

type Pagination struct {
	PageNo        int         `json:"pageNo"`
	PageSize      int         `json:"pageSize"`
	PageFn        string      `json:"pageFn"`
	TotalElements int         `json:"totalElements"`
	Length        int         `json:"length"`
	Slider        int         `json:"slider"`
	Slice         interface{} `json:"slice"`
}

func (me *Pagination) First() int {
	return defaultPageNo
}

func (me *Pagination) Previous() int {
	if me.PageNo == me.First() {
		return me.PageNo
	} else {
		return me.PageNo - 1
	}
}

func (me *Pagination) Last() int {
	if me.PageSize == defaultPageNot || me.TotalElements == defaultPageNot || me.TotalElements == 0 {
		return defaultPageNo
	} else {
		last := me.TotalElements / me.PageSize
		if me.TotalElements%me.PageSize == 0 {
			return last
		} else {
			return last + 1
		}
	}
}

func (me *Pagination) Next() int {
	if me.PageNo == me.Last() {
		return me.PageNo
	} else {
		return me.PageNo + 1
	}
}

func (me *Pagination) JSON() (string, error) {
	b, err := json.Marshal(me)
	if err == nil {
		return string(b), nil
	} else {
		return "", err
	}
}

func (me *Pagination) ParseJSON(jsons string) error {
	if strings.IsBlank(jsons) {
		return nil
	}
	return json.Unmarshal([]byte(jsons), me)
}
