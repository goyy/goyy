// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package controller

import (
	"gopkg.in/goyy/goyy.v0/util/templates"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

type HTMLController struct {
	Settings
}

func (me *HTMLController) Index(c xhttp.Context) {
	err := c.HTML(xhttp.StatusOK, me.TmplIndex(), me.Success(c, templates.EnIndex, ""))
	if err != nil {
		me.Error(c, err)
		return
	}
}

func (me *HTMLController) Show(c xhttp.Context) {
	err := c.HTML(xhttp.StatusOK, me.TmplForm(), me.Success(c, templates.EnShow, ""))
	if err != nil {
		me.Error(c, err)
		return
	}
}

func (me *HTMLController) Add(c xhttp.Context) {
	err := c.HTML(xhttp.StatusOK, me.TmplForm(), me.Success(c, templates.EnAdd, ""))
	if err != nil {
		me.Error(c, err)
		return
	}
}

func (me *HTMLController) Edit(c xhttp.Context) {
	err := c.HTML(xhttp.StatusOK, me.TmplForm(), me.Success(c, templates.EnEdit, ""))
	if err != nil {
		me.Error(c, err)
		return
	}
}
