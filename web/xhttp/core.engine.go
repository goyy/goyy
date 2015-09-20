// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xhttp

import (
	"net/http"
)

type engine struct {
	Router Router
}

// ServeHTTP makes the router implement the http.Handler interface.
func (me *engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if sec.ServeHTTP(w, r) { // secureServeMux
		return
	}
	if ssm.ServeHTTP(w, r) { // staticServeMux
		return
	}
	if hsm.ServeHTTP(w, r) { // htmlServeMux
		return
	}
	me.Router.ServeHTTP(w, r)
}
