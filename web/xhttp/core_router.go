// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xhttp

import (
	"net/http"

	"gopkg.in/goyy/goyy.v0/comm/profile"
	"gopkg.in/goyy/goyy.v0/comm/xtype"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"gopkg.in/goyy/goyy.v0/util/webs"
)

type router struct {
	methods  map[httpMethod]map[string]action
	handlers Handlers
}

func (me *router) GET(pattern string, handle Handle, permissions ...*xtype.Permission) {
	me.addRouter(httpMethodGet, pattern, handle, permissions...)
}

func (me *router) POST(pattern string, handle Handle, permissions ...*xtype.Permission) {
	me.addRouter(httpMethodPost, pattern, handle, permissions...)
}

func (me *router) PUT(pattern string, handle Handle, permissions ...*xtype.Permission) {
	me.addRouter(httpMethodPut, pattern, handle, permissions...)
}

func (me *router) DELETE(pattern string, handle Handle, permissions ...*xtype.Permission) {
	me.addRouter(httpMethodDelete, pattern, handle, permissions...)
}

func (me *router) PATCH(pattern string, handle Handle, permissions ...*xtype.Permission) {
	me.addRouter(httpMethodPatch, pattern, handle, permissions...)
}

func (me *router) HEAD(pattern string, handle Handle, permissions ...*xtype.Permission) {
	me.addRouter(httpMethodHead, pattern, handle, permissions...)
}

func (me *router) OPTIONS(pattern string, handle Handle, permissions ...*xtype.Permission) {
	me.addRouter(httpMethodOptions, pattern, handle, permissions...)
}

// Attachs a global middleware to the router. ie. the middlewares attached though Use() will be
// included in the handlers chain for every single request. Even 404, 405, static files...
func (me *router) Use(middlewares ...Handler) Router {
	me.handlers = append(me.handlers, middlewares...)
	return me
}

func (me *router) isPermission(c Context, permission *xtype.Permission) bool {
	if permission == nil || strings.IsBlank(permission.Id) {
		return false
	}
	if permission.Profiles != nil && len(permission.Profiles) > 0 && !profile.Accepts(permission.Profiles...) {
		return true
	}
	if c == nil || !c.Session().IsLogin() {
		return false
	}
	if p, err := c.Session().Principal(); err == nil {
		if permission.Profiles == nil || len(permission.Profiles) == 0 || profile.Accepts(permission.Profiles...) {
			ps := strings.Split(permission.Id, ",")
			if strings.ContainsSliceAny(p.Permissions, ps) {
				return true
			}
		}
	} else {
		logger.Error(err.Error())
		return false
	}
	return false
}

// ServeHTTP makes the router implement the http.Handler interface.
func (me *router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if v, ok := me.methods[httpMethod(r.Method)]; ok {
		requestURI := r.RequestURI
		if strings.Contains(requestURI, "?") {
			requestURI = strings.Before(r.RequestURI, "?")
		}
		if a, aok := v[requestURI]; aok {
			c := me.newContext(w, r)
			if a.Permissions != nil && len(a.Permissions) > 0 {
				isPermission := false
				for _, permission := range a.Permissions {
					if me.isPermission(c, permission) {
						isPermission = true
					}
				}
				if !isPermission {
					msg := i18N.Message("err.407")
					if webs.IsXMLHttpRequest(r) { // AJAX
						c := `{"success":false,"message":"` + msg + `"}`
						w.WriteHeader(StatusProxyAuthRequired)
						w.Write([]byte(c))
						return
					} else {
						if strings.IsNotBlank(Conf.Err.Err401) {
							http.Redirect(w, r, Conf.Err.Err401, http.StatusFound)
							return
						} else {
							w.WriteHeader(StatusProxyAuthRequired)
							w.Write([]byte(msg))
							return
						}
					}
				}
			}
			c.Next()
			a.Handle(c)
		} else {
			logger.Errorf("No match for router:%s:%s", r.Method, r.RequestURI)
			c := me.newContext(w, r)
			c.Next()
			if strings.IsNotBlank(Conf.Err.Err404) {
				http.Redirect(w, r, Conf.Err.Err404, http.StatusFound)
				return
			} else {
				msg := i18N.Message("err.404")
				serveError(c, StatusNotFound, []byte(msg))
			}
		}
	} else {
		logger.Error("No match for router:", r.Method)
		c := me.newContext(w, r)
		c.Next()
		if strings.IsNotBlank(Conf.Err.Err404) {
			http.Redirect(w, r, Conf.Err.Err404, http.StatusFound)
			return
		} else {
			msg := i18N.Message("err.404")
			serveError(c, StatusNotFound, []byte(msg))
		}
	}
}

func (me *router) addRouter(method httpMethod, pattern string, handle Handle, permissions ...*xtype.Permission) {
	if me.methods == nil {
		me.methods = make(map[httpMethod]map[string]action)
	}
	if _, ok := me.methods[method]; !ok {
		me.methods[method] = make(map[string]action)
	}
	me.addAction(method, pattern, handle, permissions...)
}

func (me *router) addAction(method httpMethod, pattern string, handle Handle, permissions ...*xtype.Permission) {
	actions := me.methods[method]
	if _, ok := actions[pattern]; !ok {
		a := action{
			Handle:      handle,
			Permissions: permissions,
		}
		actions[pattern] = a
	}
}

func (me *router) newContext(w http.ResponseWriter, r *http.Request) Context {
	values, err := webs.Values(r)
	if err != nil {
		logger.Error(err.Error())
	}
	s := newSession4Redis(w, r)
	return newContext(w, s, r, values, me.handlers)
}
