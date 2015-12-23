// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xhttp

import (
	"net/http"
	"time"

	"gopkg.in/goyy/goyy.v0/data/cache"
)

// GET adds a route for a HTTP GET request to the specified matching pattern.
func GET(path string, handle Handle, permissions ...string) {
	defaultEngine.Router.GET(path, handle, permissions...)
}

// POST adds a route for a HTTP POST request to the specified matching pattern.
func POST(path string, handle Handle, permissions ...string) {
	defaultEngine.Router.POST(path, handle, permissions...)
}

// PUT adds a route for a HTTP PUT request to the specified matching pattern.
func PUT(path string, handle Handle, permissions ...string) {
	defaultEngine.Router.PUT(path, handle, permissions...)
}

// DELETE adds a route for a HTTP DELETE request to the specified matching pattern.
func DELETE(path string, handle Handle, permissions ...string) {
	defaultEngine.Router.DELETE(path, handle, permissions...)
}

// PATCH adds a route for a HTTP PATCH request to the specified matching pattern.
func PATCH(path string, handle Handle, permissions ...string) {
	defaultEngine.Router.PATCH(path, handle, permissions...)
}

// HEAD adds a route for a HTTP HEAD request to the specified matching pattern.
func HEAD(path string, handle Handle, permissions ...string) {
	defaultEngine.Router.HEAD(path, handle, permissions...)
}

// OPTIONS adds a route for a HTTP OPTIONS request to the specified matching pattern.
func OPTIONS(path string, handle Handle, permissions ...string) {
	defaultEngine.Router.OPTIONS(path, handle, permissions...)
}

// Use adds a middleware Handler to the stack.
func Use(middlewares ...Handler) Router {
	defaultEngine.Router.Use(middlewares...)
	return defaultEngine.Router
}

// Run the http server. Listening on Conf.Addr or 9090 by default.
func Run() error {
	cache.Init(cache.Conf{
		Address:     Conf.Session.Addr,
		MaxIdle:     80,
		MaxActive:   12000,
		IdleTimeout: 240 * time.Second,
	})
	logger.Printf("Listening and serving HTTP on %s\n", Conf.Addr)
	return http.ListenAndServe(Conf.Addr, defaultEngine)
}
