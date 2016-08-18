// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cookies

import (
	"net/http"
	"time"

	"gopkg.in/goyy/goyy.v0/util/errors"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

// Cookie returns the named cookie provided in the request or
// http.ErrNoCookie if not found.
func Cookie(r *http.Request, name string) (*http.Cookie, error) {
	return r.Cookie(name)
}

// Value returns the value of cookie by name or http.ErrNoCookie if not found.
func Value(r *http.Request, name string) (string, error) {
	c, err := r.Cookie(name)
	if err != nil {
		return "", err
	}
	return c.Value, nil
}

// SetCookie adds a Set-Cookie header to the provided ResponseWriter's headers.
func SetCookie(w http.ResponseWriter, cookie *http.Cookie) {
	http.SetCookie(w, cookie)
}

// SetValue adds a cookie to a user's browser with a name, value.
func SetValue(w http.ResponseWriter, name, value string) {
	if strings.IsBlank(name) {
		return
	}
	c := &http.Cookie{
		Name:  name,
		Value: value,
		Path:  "/",
	}
	http.SetCookie(w, c)
}

// SetValueMaxAge adds a cookie to a user's browser with a name, value
// and maxAge.
// MaxAge=0 means no 'Max-Age' attribute specified.
// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
// MaxAge>0 means Max-Age attribute present and given in seconds
func SetValueMaxAge(w http.ResponseWriter, name, value string, maxAge int) {
	if strings.IsBlank(name) {
		return
	}
	c := &http.Cookie{
		Name:   name,
		Value:  value,
		Path:   "/",
		MaxAge: maxAge,
	}
	http.SetCookie(w, c)
}

// SetValueExpires adds a cookie to a user's browser with a name, value
// and expires.
func SetValueExpires(w http.ResponseWriter, name, value string, expires time.Time) {
	if strings.IsBlank(name) {
		return
	}
	c := &http.Cookie{
		Name:    name,
		Value:   value,
		Path:    "/",
		Expires: expires,
	}
	http.SetCookie(w, c)
}

// Set adds a cookie to a user's browser with a name, value, path and expiry.
func Set(w http.ResponseWriter, name, value, path string, expires time.Time) {
	if strings.IsBlank(name) {
		return
	}
	c := &http.Cookie{
		Name:    name,
		Value:   value,
		Path:    path,
		Expires: expires,
	}
	http.SetCookie(w, c)
}

// Reset reset a cookie specified by name,
// only modify the value property, other attributes remain unchanged.
func Reset(w http.ResponseWriter, r *http.Request, name, value string) error {
	if strings.IsBlank(name) {
		return errors.NewNotBlank("name")
	}
	c, err := r.Cookie(name)
	if err != nil {
		return err
	}
	c.Value = value
	http.SetCookie(w, c)
	return nil
}

// Remove removes a cookie specified by name
func Remove(w http.ResponseWriter, r *http.Request, name string) error {
	if strings.IsBlank(name) {
		return errors.NewNotBlank("name")
	}
	c, err := r.Cookie(name)
	if err != nil {
		return err
	}
	c.MaxAge = 0
	c.Path = "/"
	c.Value = ""
	c.HttpOnly = true
	http.SetCookie(w, c)
	return nil
}
