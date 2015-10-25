// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package controller

import (
	"gopkg.in/goyy/goyy.v0/util/strings"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

func (me *HTMLController) Error(c xhttp.Context, err error) {
	//go errorSave(c.Request(), err)
	logger.Error(err.Error())
	if strings.IsNotBlank(xhttp.Conf.Err.Err500) {
		c.Redirect(xhttp.Conf.Err.Err500, xhttp.StatusFound)
		return
	} else {
		c.ResponseWriter().WriteHeader(500)
		c.ResponseWriter().Write([]byte(default500Body))
		return
	}
}
