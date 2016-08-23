// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xhttp

import (
	"net/http"
	"time"

	"gopkg.in/goyy/goyy.v0/comm/xtype"
	"gopkg.in/goyy/goyy.v0/data/cache"
)

var preRuns []func() = make([]func(), 10)

// GET adds a route for a HTTP GET request to the specified matching pattern.
func GET(path string, handle Handle, permissions ...*xtype.Permission) {
	defaultEngine.Router.GET(path, handle, permissions...)
}

// POST adds a route for a HTTP POST request to the specified matching pattern.
func POST(path string, handle Handle, permissions ...*xtype.Permission) {
	defaultEngine.Router.POST(path, handle, permissions...)
}

// PUT adds a route for a HTTP PUT request to the specified matching pattern.
func PUT(path string, handle Handle, permissions ...*xtype.Permission) {
	defaultEngine.Router.PUT(path, handle, permissions...)
}

// DELETE adds a route for a HTTP DELETE request to the specified matching pattern.
func DELETE(path string, handle Handle, permissions ...*xtype.Permission) {
	defaultEngine.Router.DELETE(path, handle, permissions...)
}

// PATCH adds a route for a HTTP PATCH request to the specified matching pattern.
func PATCH(path string, handle Handle, permissions ...*xtype.Permission) {
	defaultEngine.Router.PATCH(path, handle, permissions...)
}

// HEAD adds a route for a HTTP HEAD request to the specified matching pattern.
func HEAD(path string, handle Handle, permissions ...*xtype.Permission) {
	defaultEngine.Router.HEAD(path, handle, permissions...)
}

// OPTIONS adds a route for a HTTP OPTIONS request to the specified matching pattern.
func OPTIONS(path string, handle Handle, permissions ...*xtype.Permission) {
	defaultEngine.Router.OPTIONS(path, handle, permissions...)
}

// Use adds a middleware Handler to the stack.
func Use(middlewares ...Handler) Router {
	defaultEngine.Router.Use(middlewares...)
	return defaultEngine.Router
}

func RegisterPreRun(preRun func()) {
	if preRun != nil {
		preRuns = append(preRuns, preRun)
	}
}

// Run the http server. Listening on Conf.Addr or 9090 by default.
func Run() error {
	cache.Init(cache.Conf{
		Address:     Conf.Session.Addr,
		Password:    Conf.Session.Password,
		MaxIdle:     80,
		MaxActive:   12000,
		IdleTimeout: 240 * time.Second,
	})
	for _, preRun := range preRuns {
		if preRun != nil {
			preRun()
		}
	}
	logger.Printf("Listening and serving HTTP on %s\n", Conf.Addr)
	err := http.ListenAndServe(Conf.Addr, defaultEngine)
	return err
}
