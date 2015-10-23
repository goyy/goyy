// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package controller

import (
	"gopkg.in/goyy/goyy.v0/data/result"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

func (me *Controller) Error(c xhttp.Context, err error) {
	//go errorSave(c.Request(), err)
	logger.Error(err.Error())
	c.HTML(xhttp.StatusBadRequest, tmplErr, result.Http{Message: err.Error()})
}

func (me *Controller) ErrorJson(c xhttp.Context, err error) {
	//go errorSave(c.Request(), err)
	logger.Error(err.Error())
	c.JSON(xhttp.StatusBadRequest, result.Http{Message: err.Error()})
}
