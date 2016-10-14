// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xhttp

import (
	"io"
	"net/http"
	"net/url"
	"time"

	"gopkg.in/goyy/goyy.v0/comm/log"
	"gopkg.in/goyy/goyy.v0/comm/xtype"
	"gopkg.in/goyy/goyy.v0/data/cache"
	"gopkg.in/goyy/goyy.v0/util/strings"
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
	if logger.Outputs()&log.Oconsole == 0 {
		log.Printf("Listening and serving HTTP on %s\n", Conf.Addr)
	}
	err := http.ListenAndServe(Conf.Addr, defaultEngine)
	return err
}

// NewRequest returns a new Request given a method, URL, and optional body.
//
// If the provided body is also an io.Closer, the returned
// Request.Body is set to body and will be closed by the Client
// methods Do, Post, and PostForm, and Transport.RoundTrip.
//
// NewRequest returns a Request suitable for use with Client.Do or
// Transport.RoundTrip.
// To create a request for use with testing a Server Handler use either
// ReadRequest or manually update the Request fields. See the Request
// type's documentation for the difference between inbound and outbound
// request fields.
func NewRequest(method, urlStr string, values url.Values) (*http.Request, error) {
	var body io.Reader
	switch method {
	case "GET":
		if values != nil {
			if strings.Index(urlStr, "?") != -1 {
				urlStr = urlStr + "&" + values.Encode()
			} else {
				urlStr = urlStr + "?" + values.Encode()
			}
		}
	case "POST":
		if values != nil {
			body = strings.NewReader(values.Encode())
		}
	}
	return http.NewRequest(method, urlStr, body)
}
