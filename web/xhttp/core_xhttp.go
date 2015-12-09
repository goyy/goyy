// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xhttp

import (
	"fmt"
	"gopkg.in/goyy/goyy.v0/data/cache"
	"net/http"
	"time"
)

func GET(path string, handle Handle, permissions ...string) {
	defaultEngine.Router.GET(path, handle, permissions...)
}

func POST(path string, handle Handle, permissions ...string) {
	defaultEngine.Router.POST(path, handle, permissions...)
}

func PUT(path string, handle Handle, permissions ...string) {
	defaultEngine.Router.PUT(path, handle, permissions...)
}

func DELETE(path string, handle Handle, permissions ...string) {
	defaultEngine.Router.DELETE(path, handle, permissions...)
}

func PATCH(path string, handle Handle, permissions ...string) {
	defaultEngine.Router.PATCH(path, handle, permissions...)
}

func HEAD(path string, handle Handle, permissions ...string) {
	defaultEngine.Router.HEAD(path, handle, permissions...)
}

func OPTIONS(path string, handle Handle, permissions ...string) {
	defaultEngine.Router.OPTIONS(path, handle, permissions...)
}

// Attachs a global middleware to the router. ie. the middlewares attached though Use() will be
// included in the handlers chain for every single request. Even 404, 405, static files...
func Use(middlewares ...Handler) Router {
	defaultEngine.Router.Use(middlewares...)
	return defaultEngine.Router
}

func Run() error {
	cache.Init(cache.Conf{
		Address:     Conf.Session.Addr,
		MaxIdle:     80,
		MaxActive:   12000,
		IdleTimeout: 240 * time.Second,
	})
	fmt.Printf("Listening and serving HTTP on %s\n", Conf.Addr)
	return http.ListenAndServe(Conf.Addr, defaultEngine)
}
