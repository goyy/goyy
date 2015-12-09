// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xhttp

import (
	"net/http"
)

// Render provides functions for easily writing JSON and HTML templates out to
// a http Response.
type render interface {
	// HTML renders a html template specified by the name and writes the result
	// and given status to the http.ResponseWriter.
	HTML(w http.ResponseWriter, status int, name string, v interface{}) error
	// JSON writes the given status and JSON serialized version of the given
	// value to the http.ResponseWriter.
	JSON(w http.ResponseWriter, status int, v interface{}) error
	// JSONP writes the given status and JSONP serialized version of the given
	// value to the http.ResponseWriter.
	JSONP(w http.ResponseWriter, status int, callback string, v interface{}) error
	// XML writes the given status and XML serialized version of the given
	// value to the http.ResponseWriter.
	XML(w http.ResponseWriter, status int, v interface{}) error
	// Data writes the raw byte array to the http.ResponseWriter.
	Text(w http.ResponseWriter, status int, format string, values ...interface{}) error
	// Error is a convenience function that writes an http status to the
	// http.ResponseWriter.
	Error(w http.ResponseWriter, status int) error
	// Redirect is a convienience function that sends an HTTP redirect.
	// If status is omitted, uses 302 (Found)
	Redirect(w http.ResponseWriter, req *http.Request, location string, status ...int) error
}
