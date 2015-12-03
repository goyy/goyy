// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xhttp

import (
	"gopkg.in/goyy/goyy.v0/util/strings"
	"net/http"
)

var asts *staticServeMux
var stas *staticServeMux
var devs *staticServeMux
var oprs *staticServeMux

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
