// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xtype

import (
	"encoding/json"
	"errors"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"io/ioutil"
	"net/http"
)

// Boxes returns the []xtype.Box by the http.
func Boxes(url string) (boxes []Box, err error) {
	resp, gerr := http.Get(url)
	if gerr != nil {
		err = gerr
		return
	}
	defer resp.Body.Close()
	body, rerr := ioutil.ReadAll(resp.Body)
	if rerr != nil {
		err = rerr
		return
	}
	if strings.IsBlank(string(body)) {
		err = errors.New("Invalid url")
		return
	}
	err = json.Unmarshal(body, &boxes)
	if err != nil {
		return
	}
	return
}
