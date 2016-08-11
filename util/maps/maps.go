// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package maps implements map utility functions.
package maps

import (
	"gopkg.in/goyy/goyy.v0/util/strings"
)

// ParseURLQuery returns the map by the params(url query).
// eg. key1=value1&key2=value2  =>  map[key1:value1 key2:value2]
func ParseURLQuery(params string) map[string]string {
	maps := make(map[string]string, 0)
	if !strings.Contains(params, "=") {
		return maps
	}
	ps := strings.Split(params, "&")
	for _, p := range ps {
		param := strings.Split(p, "=")
		if len(param) == 2 && strings.IsNotBlank(param[0]) {
			maps[param[0]] = param[1]
		}
	}
	return maps
}
