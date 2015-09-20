// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xhttp

type httpMethod string

const (
	httpMethodGet     = httpMethod("GET")
	httpMethodPost    = httpMethod("POST")
	httpMethodPut     = httpMethod("PUT")
	httpMethodDelete  = httpMethod("DELETE")
	httpMethodPatch   = httpMethod("PATCH")
	httpMethodHead    = httpMethod("HEAD")
	httpMethodOptions = httpMethod("OPTIONS")
)

var httpMethods = []httpMethod{
	httpMethodGet,
	httpMethodPost,
	httpMethodPut,
	httpMethodDelete,
	httpMethodPatch,
	httpMethodHead,
	httpMethodOptions,
}
