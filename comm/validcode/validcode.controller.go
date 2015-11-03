// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package validcode

import (
	"gopkg.in/goyy/goyy.v0/web/controller"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

type Controller struct {
	controller.JSONController
}

// Get verification code
func (me *Controller) Build(c xhttp.Context) {
	createImage(c)
}

// Ajax verification code
func (me *Controller) Judge(c xhttp.Context) {
	r := Judge(c)
	if r {
		c.JSON(xhttp.StatusOK, me.SuccessMessage(c, "Verification code matching correctly"))
		return
	}
	c.JSON(xhttp.StatusOK, me.FaultMessage(c, "Verification code matching error"))
	return
}
