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
			me.staticMappings(&Conf.Static.Mappings)
		}
		for _, v := range stas.ServeMux {
			if v.ServeHTTP(w, r) {
				return
			}
		}
	}
	if Conf.Developer.Enable { // developerServeMux
		if devs.ServeMux == nil {
			me.staticMappings(&Conf.Developer.Mappings)
		}
		for _, v := range devs.ServeMux {
			if v.ServeHTTP(w, r) {
				return
			}
		}
	}
	if Conf.Operation.Enable { // operationServeMux
		if oprs.ServeMux == nil {
			me.staticMappings(&Conf.Operation.Mappings)
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

func (me *engine) staticMappings(mappings *xtype.Mappings) {
	u := mappings.URL
	d := mappings.Dir
	m := mappings.Mapping
	if m != nil && len(m) > 0 {
		for _, v := range m {
			p := u + v.Path
			ssm := &staticServeMux{
				urlPrefix: p,
				static:    http.StripPrefix(p, http.FileServer(http.Dir(v.Dir))),
			}
			stas.ServeMux = append(stas.ServeMux, ssm)
		}
	}

	ssm := &staticServeMux{
		urlPrefix: u,
		static:    http.StripPrefix(u, http.FileServer(http.Dir(d))),
	}
	stas.ServeMux = append(stas.ServeMux, ssm)
}
