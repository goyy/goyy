// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xhttp

import (
	"net/http"

	"gopkg.in/goyy/goyy.v0/comm/xtype"
)

type Router interface {
	Use(middlewares ...Handler) Router

	GET(pattern string, handle Handle, permissions ...*xtype.Permission)
	POST(pattern string, handle Handle, permissions ...*xtype.Permission)
	PUT(pattern string, handle Handle, permissions ...*xtype.Permission)
	DELETE(pattern string, handle Handle, permissions ...*xtype.Permission)
	PATCH(pattern string, handle Handle, permissions ...*xtype.Permission)
	HEAD(pattern string, handle Handle, permissions ...*xtype.Permission)
	OPTIONS(pattern string, handle Handle, permissions ...*xtype.Permission)

	ServeHTTP(w http.ResponseWriter, req *http.Request)
}
