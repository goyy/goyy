// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package controller

import (
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

var DemoCtl = &DemoController{}

type DemoController struct {
}

func (me *DemoController) JSON(ctx xhttp.Context) {
	ctx.JSON(xhttp.StatusOK, map[string]string{"key1": "value1", "key2": "value2"})
}

func (me *DemoController) JSONP(ctx xhttp.Context) {
	ctx.JSONP(xhttp.StatusOK, "callback", map[string]string{"key1": "value1", "key2": "value2"})
}

func (me *DemoController) XML(ctx xhttp.Context) {
	type users struct {
		Name  string
		Email string
	}
	data := &users{Name: "admin", Email: "admin@gmail.com"}
	ctx.XML(xhttp.StatusOK, data)
}

func (me *DemoController) List(ctx xhttp.Context) {
	v, _ := ctx.Session().Get("user")
	ctx.HTML(xhttp.StatusOK, "list", v)
}

func (me *DemoController) Form(ctx xhttp.Context) {
	ctx.HTML(xhttp.StatusOK, "form", "!!!")
}
