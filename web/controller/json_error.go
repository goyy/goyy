// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package controller

import (
	"gopkg.in/goyy/goyy.v0/data/result"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

func (me *JSONController) Error(c xhttp.Context, err error) {
	//go errorSave(c.Request(), err)
	logger.Error(err.Error())
	switch t := err.(type) {
	case *PreError:
		c.JSON(xhttp.StatusOK, result.Http{Success: false, Code: t.Code, Message: t.Message})
	default:
		c.JSON(xhttp.StatusInternalServerError, result.Http{Message: err.Error()})
	}
}
