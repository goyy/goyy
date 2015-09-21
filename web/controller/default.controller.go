// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package controller

import (
	"gopkg.in/goyy/goyy.v0/data/service"
	"gopkg.in/goyy/goyy.v0/util/templates"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

type Controller struct {
	baseController
	pre
	post
	Settings
	Mgr service.Service
}

func (me *Controller) Index(c xhttp.Context) {
	out, err := me.baseController.Index(c, me.Mgr, me.PreIndex, me.PostIndex)
	if err != nil {
		me.Error(c, err)
		return
	}
	err = c.HTML(xhttp.StatusOK, me.TmplDefault(), me.Success(c, templates.EnIndex, out.Data))
	if err != nil {
		me.Error(c, err)
		return
	}
}

func (me *Controller) Show(c xhttp.Context) {
	out, err := me.baseController.Show(c, me.Mgr, me.PreShow, me.PostShow)
	if err != nil {
		me.Error(c, err)
		return
	}
	err = c.HTML(xhttp.StatusOK, me.TmplDefault(), me.Success(c, templates.EnShow, out.Data))
	if err != nil {
		me.Error(c, err)
		return
	}
}

func (me *Controller) Add(c xhttp.Context) {
	out, err := me.baseController.Add(c, me.Mgr, me.PreAdd, me.PostAdd)
	if err != nil {
		me.Error(c, err)
		return
	}
	err = c.HTML(xhttp.StatusOK, me.TmplDefault(), me.Success(c, templates.EnAdd, out.Data))
	if err != nil {
		me.Error(c, err)
		return
	}
}

func (me *Controller) Edit(c xhttp.Context) {
	out, err := me.baseController.Edit(c, me.Mgr, me.PreEdit, me.PostEdit)
	if err != nil {
		me.Error(c, err)
		return
	}
	err = c.HTML(xhttp.StatusOK, me.TmplDefault(), me.Success(c, templates.EnEdit, out.Data))
	if err != nil {
		me.Error(c, err)
		return
	}
}

func (me *Controller) Save(c xhttp.Context) {
	out, err := me.baseController.Save(c, me.Mgr, me.PreSave, me.PostSave)
	if err != nil {
		me.Error(c, err)
		return
	}
	err = c.HTML(xhttp.StatusOK, me.TmplDefault(), me.SuccessMsg(c, i18N.Message("msg.save"), templates.EnIndex, out.Data))
	if err != nil {
		me.Error(c, err)
		return
	}
}

func (me *Controller) Disable(c xhttp.Context) {
	out, err := me.baseController.Disable(c, me.Mgr, me.PreDisable, me.PostDisable)
	if err != nil {
		me.Error(c, err)
		return
	}
	err = c.HTML(xhttp.StatusOK, me.TmplDefault(), me.SuccessMsg(c, i18N.Message("msg.disable"), templates.EnIndex, out.Data))
	if err != nil {
		me.Error(c, err)
		return
	}
}

func (me *Controller) Box(c xhttp.Context) {
	out, err := me.baseController.Box(c, me.Mgr)
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
