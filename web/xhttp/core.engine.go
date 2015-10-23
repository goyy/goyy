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
	if Conf.Asset.Enable { // assetServeMux
		if asts == nil {
			asts = &staticServeMux{
				urlPrefix: Conf.Asset.URL,
				static:    http.StripPrefix(Conf.Asset.URL, http.FileServer(http.Dir(Conf.Asset.Dir))),
			}
		}
		if asts.ServeHTTP(w, r) {
			return
		}
	}
	if Conf.Developer.Enable { // developerServeMux
		if devs == nil {
			devs = &staticServeMux{
				urlPrefix: Conf.Developer.URL,
				static:    http.StripPrefix(Conf.Developer.URL, http.FileServer(http.Dir(Conf.Developer.Dir))),
			}
		}
		if devs.ServeHTTP(w, r) {
			return
		}
	}
	if Conf.Operation.Enable { // operationServeMux
		if oprs == nil {
			oprs = &staticServeMux{
				urlPrefix: Conf.Operation.URL,
				static:    http.StripPrefix(Conf.Operation.URL, http.FileServer(http.Dir(Conf.Operation.Dir))),
			}
		}
		if oprs.ServeHTTP(w, r) {
			return
		}
	}
	if Conf.Upload.Enable { // uploadServeMux
		if upls == nil {
			upls = &uploadServeMux{
				static: http.StripPrefix(Conf.Upload.URL, http.FileServer(http.Dir(Conf.Upload.Dir))),
			}
		}
		if upls.ServeHTTP(w, r) {
			return
		}
	}
	if Conf.Html.Enable {
		if hsm.ServeHTTP(w, r) { // htmlServeMux
			return
		}
	}
	if Conf.Secure.Enable { // secureServeMux
		if sec.ServeHTTP(w, r) {
			return
		}
	}
	if Conf.Template.Enable {
		me.Router.ServeHTTP(w, r)
	}
}
