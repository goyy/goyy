// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xhttp

import (
	"net/http"
)

type Router interface {
	Use(middlewares ...Handler) Router

	GET(pattern string, handle Handle, permissions ...string)
	POST(pattern string, handle Handle, permissions ...string)
	PUT(pattern string, handle Handle, permissions ...string)
	DELETE(pattern string, handle Handle, permissions ...string)
	PATCH(pattern string, handle Handle, permissions ...string)
	HEAD(pattern string, handle Handle, permissions ...string)
	OPTIONS(pattern string, handle Handle, permissions ...string)

	ServeHTTP(w http.ResponseWriter, req *http.Request)
}
