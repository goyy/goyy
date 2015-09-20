// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xhttp

import (
	"gopkg.in/goyy/goyy.v0/comm/profile"
	"gopkg.in/goyy/goyy.v0/data/cache"
	"net/http"
	"time"
)

func GET(path string, handle Handle) {
	defaultEngine.Router.GET(path, handle)
}

func POST(path string, handle Handle) {
	defaultEngine.Router.POST(path, handle)
}

func PUT(path string, handle Handle) {
	defaultEngine.Router.PUT(path, handle)
}

func DELETE(path string, handle Handle) {
	defaultEngine.Router.DELETE(path, handle)
}

func PATCH(path string, handle Handle) {
	defaultEngine.Router.PATCH(path, handle)
}

func HEAD(path string, handle Handle) {
	defaultEngine.Router.HEAD(path, handle)
}

func OPTIONS(path string, handle Handle) {
	defaultEngine.Router.OPTIONS(path, handle)
}

// Attachs a global middleware to the router. ie. the middlewares attached though Use() will be
// included in the handlers chain for every single request. Even 404, 405, static files...
func Use(middlewares ...Handler) Router {
	defaultEngine.Router.Use(middlewares...)
	return defaultEngine.Router
}

func Run() error {
	profile.SetActives(Conf.Actives...)
	cache.Init(cache.Conf{
		Address:     Conf.Session.Addr,
		MaxIdle:     80,
		MaxActive:   12000,
		IdleTimeout: 240 * time.Second,
	})
	logger.Printf("Listening and serving HTTP on %s\n", Conf.Addr)
	return http.ListenAndServe(Conf.Addr, defaultEngine)
}
