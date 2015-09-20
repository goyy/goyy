// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xhttp

import (
	"gopkg.in/goyy/goyy.v0/util/strings"
	"net/http"
)

type staticServeMux struct {
	static http.Handler
}

var ssm = &staticServeMux{
	static: http.StripPrefix(Conf.Static.Assets, http.FileServer(http.Dir(Conf.Static.Dir))),
}

func (me *staticServeMux) ServeHTTP(w http.ResponseWriter, req *http.Request) bool {
	if me.isStatic(req.URL.Path) {
		me.static.ServeHTTP(w, req)
		return true
	}
	return false
}

func (me *staticServeMux) isStatic(path string) bool {
	if strings.HasPrefix(path, Conf.Static.Assets+"/") {
		return true
	}
	return false
}
