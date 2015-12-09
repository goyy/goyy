// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xhttp

import (
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/web/session"
	"net/http"
	"net/url"
)

type Context interface {
	Request() *http.Request
	Session() session.Interface
	ResponseWriter() http.ResponseWriter
	Params() url.Values

	Param(key string) string
	Bind(out entity.Interface) error

	Attribute(key string) interface{}
	Attributes() map[string]interface{}
	SetAttribute(key string, value interface{})

	Next()
	IsAborted() bool
	Abort()
	AbortWithStatus(code int)

	HTML(status int, name string, v interface{}) error
	JSON(status int, v interface{}) error
	JSONP(status int, callback string, v interface{}) error
	XML(status int, v interface{}) error
	Text(status int, format string, values ...interface{}) error
	Error(status int) error
	Redirect(location string, status ...int) error
}
