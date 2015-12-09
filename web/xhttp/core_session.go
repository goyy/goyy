// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xhttp

import (
	"gopkg.in/goyy/goyy.v0/web/session"
	"net/http"
)

func newSession4Redis(w http.ResponseWriter, r *http.Request) session.Interface {
	o := &session.Options{
		Path:     Conf.Session.Path,
		Domain:   Conf.Session.Domain,
		MaxAge:   Conf.Session.MaxAge,
		Secure:   Conf.Session.Secure,
		HttpOnly: Conf.Session.HttpOnly,
	}
	return session.New(w, r, o)
}
