// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xhttp

import (
	"net/http"
)

type Router interface {
	Use(middlewares ...Handler) Router

	GET(pattern string, handle Handle)
	POST(pattern string, handle Handle)
	PUT(pattern string, handle Handle)
	DELETE(pattern string, handle Handle)
	PATCH(pattern string, handle Handle)
	HEAD(pattern string, handle Handle)
	OPTIONS(pattern string, handle Handle)

	ServeHTTP(w http.ResponseWriter, req *http.Request)
}
