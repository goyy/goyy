// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xhttp

import (
	"net/http"

	"gopkg.in/goyy/goyy.v0/util/files"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

func serveError(c Context, code int, defaultMessage []byte) {
	c.ResponseWriter().WriteHeader(code)
	c.Next()
	c.ResponseWriter().Header()["Content-Type"] = []string{MIMEPlain}
	c.ResponseWriter().Write(defaultMessage)
}

func err401(w http.ResponseWriter, r *http.Request, c Context) {
	if strings.IsNotBlank(Conf.Err.Err401) {
		if files.IsExist(Conf.Template.Dir + Conf.Err.Err401) {
			http.Redirect(w, r, Conf.Err.Err401, http.StatusFound)
			return
		}
	}
	msg := i18N.Message("err.401")
	serveError(c, StatusNotFound, []byte(msg))
}

func err404(w http.ResponseWriter, r *http.Request, c Context) {
	if strings.IsNotBlank(Conf.Err.Err404) {
		if files.IsExist(Conf.Template.Dir + Conf.Err.Err404) {
			http.Redirect(w, r, Conf.Err.Err404, http.StatusFound)
			return
		}
	}
	msg := i18N.Message("err.404")
	serveError(c, StatusNotFound, []byte(msg))
}

func err500(w http.ResponseWriter, r *http.Request) {
	if strings.IsNotBlank(Conf.Err.Err500) {
		if files.IsExist(Conf.Template.Dir + Conf.Err.Err500) {
			http.Redirect(w, r, Conf.Err.Err500, http.StatusFound)
			return
		}
	}
	msg := i18N.Message("err.500")
	w.WriteHeader(500)
	w.Write([]byte(msg))
}

func err407(w http.ResponseWriter, r *http.Request) {
	msg := i18N.Message("err.407")
	c := `{"success":false,"message":"` + msg + `"}`
	w.WriteHeader(StatusProxyAuthRequired)
	w.Write([]byte(c))
}
