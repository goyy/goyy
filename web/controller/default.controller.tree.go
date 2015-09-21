// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package controller

import (
	"gopkg.in/goyy/goyy.v0/util/templates"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

type TreeController struct {
	baseTreeController
	Controller
}

func (me *TreeController) Index(c xhttp.Context) {
	out, err := me.baseTreeController.Index(c, me.Mgr, me.PreIndex, me.PostIndex)
	if err != nil {
		me.Error(c, err)
		return
	}
	err = c.HTML(xhttp.StatusOK, me.TmplDefault(), me.Success(c, templates.EnIndex, out))
	if err != nil {
		me.Error(c, err)
		return
	}
}

func (me *TreeController) Save(c xhttp.Context) {
	out, err := me.baseTreeController.Save(c, me.Mgr, me.PreSave, me.PostSave)
	if err != nil {
		me.Error(c, err)
		return
	}
	err = c.HTML(xhttp.StatusOK, me.TmplDefault(), me.SuccessMsg(c, "Save success", templates.EnIndex, out))
	if err != nil {
		me.Error(c, err)
		return
	}
}

func (me *TreeController) Disable(c xhttp.Context) {
	out, err := me.baseTreeController.Disable(c, me.Mgr, me.PreDisable, me.PostDisable)
	if err != nil {
		me.Error(c, err)
		return
	}
	err = c.HTML(xhttp.StatusOK, me.TmplDefault(), me.SuccessMsg(c, "Delete success", templates.EnIndex, out))
	if err != nil {
		me.Error(c, err)
		return
	}
}

func (me *TreeController) Tree(c xhttp.Context) {
	out, err := me.baseTreeController.Tree(c, me.Mgr)
	if err != nil {
		me.Error(c, err)
		return
	}
	err = c.JSON(xhttp.StatusOK, out)
	if err != nil {
		me.Error(c, err)
		return
	}
}
