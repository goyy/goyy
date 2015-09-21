// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xhttp

func serveError(c Context, code int, defaultMessage []byte) {
	c.ResponseWriter().WriteHeader(code)
	c.Next()
	c.ResponseWriter().Header()["Content-Type"] = []string{MIMEPlain}
	c.ResponseWriter().Write(defaultMessage)
}
