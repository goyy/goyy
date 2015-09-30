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
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(500)
			w.Write([]byte(default500Body))
		}
	}()
	if sec.ServeHTTP(w, r) { // secureServeMux
		return
	}
	if Conf.Static.Enable {
		if ssm == nil {
			ssm = &staticServeMux{
				static: http.StripPrefix(Conf.Static.Assets, http.FileServer(http.Dir(Conf.Static.Dir))),
			}
		}
		if ssm.ServeHTTP(w, r) { // staticServeMux
			return
		}
		if usm == nil {
			usm = &uploadServeMux{
				static: http.StripPrefix(Conf.Static.Consumers, http.FileServer(http.Dir(Conf.Upload.Dir))),
			}
		}
		if usm.ServeHTTP(w, r) { // uploadServeMux
			return
		}
	}
	if Conf.Html.Enable {
		if hsm.ServeHTTP(w, r) { // htmlServeMux
			return
		}
	}
	me.Router.ServeHTTP(w, r)
}
