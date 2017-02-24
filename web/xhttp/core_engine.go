// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xhttp

import (
	"net/http"
	"runtime/debug"

	"gopkg.in/goyy/goyy.v0/comm/xtype"
)

type engine struct {
	Router Router
}

// ServeHTTP makes the router implement the http.Handler interface.
func (me *engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error(err)
			debug.PrintStack()
			err500(w, r)
		}
	}()
	if Conf.Static.Enable { // staticServeMux
		if stas.ServeMux == nil {
			me.staticMappings(Conf.Static.URL, Conf.Static.Dir, Conf.Static.Mappings)
		}
		for _, v := range stas.ServeMux {
			if v.ServeHTTP(w, r) {
				return
			}
		}
	}
	if Conf.Developer.Enable { // developerServeMux
		if devs.ServeMux == nil {
			me.staticMappings(Conf.Developer.URL, Conf.Developer.Dir, Conf.Developer.Mappings)
		}
		for _, v := range devs.ServeMux {
			if v.ServeHTTP(w, r) {
				return
			}
		}
	}
	if Conf.Operation.Enable { // operationServeMux
		if oprs.ServeMux == nil {
			me.staticMappings(Conf.Operation.URL, Conf.Operation.Dir, Conf.Operation.Mappings)
		}
		for _, v := range oprs.ServeMux {
			if v.ServeHTTP(w, r) {
				return
			}
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
	if Conf.Sensitive.Enable {
		if ssm.ServeHTTP(w, r) { // sensitiveServeMux
			return
		}
	}
	if Conf.Secure.Enable { // secureServeMux
		if sec.ServeHTTP(w, r) {
			return
		}
	}
	if Conf.Html.Enable {
		if hsm.ServeHTTP(w, r) { // htmlServeMux
			return
		}
	}
	if Conf.Template.Enable {
		me.Router.ServeHTTP(w, r)
	}
}

func (me *engine) staticMappings(url, dir string, mappings []xtype.Mapping) {
	if mappings != nil && len(mappings) > 0 {
		for _, v := range mappings {
			path := url + v.Path
			ssm := &staticServeMux{
				urlPrefix: path,
				static:    http.StripPrefix(path, http.FileServer(http.Dir(v.Dir))),
			}
			stas.ServeMux = append(stas.ServeMux, ssm)
		}
	}

	ssm := &staticServeMux{
		urlPrefix: url,
		static:    http.StripPrefix(url, http.FileServer(http.Dir(dir))),
	}
	stas.ServeMux = append(stas.ServeMux, ssm)
}
