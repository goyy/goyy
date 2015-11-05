// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package controller

import (
	"gopkg.in/goyy/goyy.v0/data/result"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

func (me *Controller) Error(c xhttp.Context, err error) {
	//go errorSave(c.Request(), err)
	logger.Error(err.Error())
	if strings.IsNotBlank(xhttp.Conf.Err.Err500) {
		c.Redirect(xhttp.Conf.Err.Err500, xhttp.StatusFound)
		return
	} else {
		status := xhttp.StatusInternalServerError
		switch err.(type) {
		case *PreError:
			status = xhttp.StatusPreconditionFailed
		default:
			status = xhttp.StatusInternalServerError
		}
		c.ResponseWriter().WriteHeader(status)
		c.ResponseWriter().Write([]byte(default500Body))
		return
	}
}

func (me *Controller) ErrorJson(c xhttp.Context, err error) {
	//go errorSave(c.Request(), err)
	logger.Error(err.Error())
	switch t := err.(type) {
	case *PreError:
		c.JSON(xhttp.StatusOK, result.Http{Success: false, Code: t.Code, Message: t.Message})
	default:
		c.JSON(xhttp.StatusInternalServerError, result.Http{Message: err.Error()})
	}
}
