// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xhttp

import (
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"gopkg.in/goyy/goyy.v0/web/session"
	"net/http"
	"net/url"
)

func newContext(w http.ResponseWriter, s session.Interface, r *http.Request, vs url.Values, hs Handlers) Context {
	c := &context{}
	c.request = r
	c.session = s
	c.response = w
	c.params = vs
	c.attributes = make(map[string]interface{}, 0)
	c.handlers = hs
	return c
}

type context struct {
	request    *http.Request
	session    session.Interface
	response   http.ResponseWriter
	params     url.Values
	attributes map[string]interface{}
	handlers   Handlers
	index      int8
}

func (me *context) Request() *http.Request {
	return me.request
}

func (me *context) Session() session.Interface {
	return me.session
}

func (me *context) ResponseWriter() http.ResponseWriter {
	return me.response
}

func (me *context) Params() url.Values {
	return me.params
}

func (me *context) Param(key string) string {
	return strings.JoinIgnoreBlank(me.params[key], ",")
}

func (me *context) Bind(out entity.Interface) error {
	for k, v := range me.params {
		value := strings.JoinIgnoreBlank(v, ",")
		// value isNotBlank -> exec SetString(k, value)
		// value isBlank -> exec SetString(k, value)
		if err := out.SetString(k, value); err != nil {
			logger.Error(err.Error())
			return err
		}
	}
	return nil
}

func (me *context) Attribute(key string) interface{} {
	return me.attributes[key]
}

func (me *context) Attributes() map[string]interface{} {
	return me.attributes
}

func (me *context) SetAttribute(key string, value interface{}) {
	if strings.IsNotBlank(key) && value != nil {
		me.attributes[key] = value
	}
}

// Next should be used only in the middlewares.
// It executes the pending handlers in the chain inside the calling handler.
func (me *context) Next() {
	s := int8(len(me.handlers))
	if s > 0 && me.index < s {
		i := me.index
		me.index++
		me.handlers[i](me)
	}
}

// Returns if the currect context was aborted.
func (me *context) IsAborted() bool {
	return me.index == AbortIndex
}

// Stops the system to continue calling the pending handlers in the chain.
// Let's say you have an authorization middleware that validates if the request is authorized
// if the authorization fails (the password does not match). This method (Abort()) should be called
// in order to stop the execution of the actual handler.
func (me *context) Abort() {
	me.index = AbortIndex
}

// It calls Abort() and writes the headers with the specified status code.
// For example, a failed attempt to authentificate a request could use: context.AbortWithStatus(401).
func (me *context) AbortWithStatus(code int) {
	me.response.WriteHeader(code)
	me.Abort()
}

func (me *context) HTML(status int, name string, v interface{}) error {
	return r.HTML(me.ResponseWriter(), status, name, v)
}

func (me *context) JSON(status int, v interface{}) error {
	return r.JSON(me.ResponseWriter(), status, v)
}

func (me *context) JSONP(status int, callback string, v interface{}) error {
	return r.JSONP(me.ResponseWriter(), status, callback, v)
}

func (me *context) XML(status int, v interface{}) error {
	return r.XML(me.ResponseWriter(), status, v)
}

func (me *context) Text(status int, format string, values ...interface{}) error {
	return r.Text(me.ResponseWriter(), status, format, values...)
}

func (me *context) Error(status int) error {
	return r.Error(me.ResponseWriter(), status)
}

func (me *context) Redirect(location string, status ...int) error {
	return r.Redirect(me.ResponseWriter(), me.Request(), location, status...)
}
