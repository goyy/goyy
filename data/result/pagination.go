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

// Pagination result.Pagination.
type Pagination struct {
	PageNo        int         `json:"pageNo"`
	PageSize      int         `json:"pageSize"`
	PageFn        string      `json:"pageFn"`
	TotalElements int         `json:"totalElements"`
	Length        int         `json:"length"`
	Slider        int         `json:"slider"`
	Slice         interface{} `json:"slice"`
}

// First get the first page.
func (me *Pagination) First() int {
	return defaultPageNo
}

// Previous get the previous page.
func (me *Pagination) Previous() int {
	if me.PageNo == me.First() {
		return me.PageNo
	}
	return me.PageNo - 1
}

// Last get the last page.
func (me *Pagination) Last() int {
	if me.PageSize == defaultPageNot || me.TotalElements == defaultPageNot || me.TotalElements == 0 {
		return defaultPageNo
	}
	last := me.TotalElements / me.PageSize
	if me.TotalElements%me.PageSize == 0 {
		return last
	}
	return last + 1
}

// Next get the next page.
func (me *Pagination) Next() int {
	if me.PageNo == me.Last() {
		return me.PageNo
	}
	return me.PageNo + 1
}

// JSON result.Pagination to JSON.
func (me *Pagination) JSON() (string, error) {
	b, err := json.Marshal(me)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// ParseJSON JSON to result.Pagination.
func (me *Pagination) ParseJSON(jsons string) error {
	if strings.IsBlank(jsons) {
		return nil
	}
	return json.Unmarshal([]byte(jsons), me)
}
