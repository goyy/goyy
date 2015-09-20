// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xhttp

import (
	"gopkg.in/goyy/goyy.v0/util/strings"
	"gopkg.in/goyy/goyy.v0/util/webs"
	"net/http"
)

type router struct {
	methods  map[httpMethod]map[string]Handle
	handlers Handlers
}

func (me *router) GET(pattern string, handle Handle) {
	me.addRouter(httpMethodGet, pattern, handle)
}

func (me *router) POST(pattern string, handle Handle) {
	me.addRouter(httpMethodPost, pattern, handle)
}

func (me *router) PUT(pattern string, handle Handle) {
	me.addRouter(httpMethodPut, pattern, handle)
}

func (me *router) DELETE(pattern string, handle Handle) {
	me.addRouter(httpMethodDelete, pattern, handle)
}

func (me *router) PATCH(pattern string, handle Handle) {
	me.addRouter(httpMethodPatch, pattern, handle)
}

func (me *router) HEAD(pattern string, handle Handle) {
	me.addRouter(httpMethodHead, pattern, handle)
}

func (me *router) OPTIONS(pattern string, handle Handle) {
	me.addRouter(httpMethodOptions, pattern, handle)
}

// Attachs a global middleware to the router. ie. the middlewares attached though Use() will be
// included in the handlers chain for every single request. Even 404, 405, static files...
func (me *router) Use(middlewares ...Handler) Router {
	me.handlers = append(me.handlers, middlewares...)
	return me
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
			c.Next()
			a(c)
		} else {
			logger.Errorf("No match for router:%s:%s", r.Method, r.RequestURI)
			c := me.newContext(w, r)
			c.Next()
			serveError(c, 404, []byte(default404Body))
		}
	} else {
		logger.Error("No match for router:", r.Method)
		c := me.newContext(w, r)
		c.Next()
		serveError(c, 404, []byte(default404Body))
	}
}

func (me *router) addRouter(method httpMethod, pattern string, handle Handle) {
	if me.methods == nil {
		me.methods = make(map[httpMethod]map[string]Handle)
	}
	if _, ok := me.methods[method]; !ok {
		me.methods[method] = make(map[string]Handle)
	}
	me.addAction(method, pattern, handle)
}

func (me *router) addAction(method httpMethod, pattern string, handle Handle) {
	actions := me.methods[method]
	if _, ok := actions[pattern]; !ok {
		actions[pattern] = handle
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
