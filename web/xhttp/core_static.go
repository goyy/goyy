// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xhttp

import (
	"net/http"

	"gopkg.in/goyy/goyy.v0/util/strings"
)

var asts staticServeMuxs
var stas staticServeMuxs
var devs staticServeMuxs
var oprs staticServeMuxs

type staticServeMuxs struct {
	ServeMux []*staticServeMux
}

type staticServeMux struct {
	urlPrefix string
	static    http.Handler
}

func (me *staticServeMux) ServeHTTP(w http.ResponseWriter, req *http.Request) bool {
	if me.isStatic(req.URL.Path) {
		me.static.ServeHTTP(w, req)
		return true
	}
	return false
}

func (me *staticServeMux) isStatic(path string) bool {
	if strings.HasPrefix(path, me.urlPrefix+"/") {
		return true
	}
	return false
}
