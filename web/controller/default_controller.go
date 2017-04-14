// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package controller

import (
	"gopkg.in/goyy/goyy.v0/data/service"
	"gopkg.in/goyy/goyy.v0/util/templates"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

// Controller controller.Controller.
type Controller struct {
	baseController
	pre
	post
	Settings
	Mgr service.Service
}

// Index displays the list of pages.
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

// Show displays a non-editable form page.
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

// Add the form page for adding is displayed.
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

// Edit the form page for editing is displayed.
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

// Save used to save the form, but do not automatically commit the transaction.
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

// SaveAndTx used to save the form, but automatically commit the transaction.
func (me *Controller) SaveAndTx(c xhttp.Context) {
	out, err := me.baseController.SaveAndTx(c, me.Mgr, me.PreSave, me.PostSave)
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

// Disable tsed to delete the records corresponding to the form, but do not automatically commit the transaction.
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

// DisableAndTx used to delete the record corresponding to the form, but automatically commit the transaction.
func (me *Controller) DisableAndTx(c xhttp.Context) {
	out, err := me.baseController.DisableAndTx(c, me.Mgr, me.PreDisable, me.PostDisable)
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

// Box gets a list of box types and converts them to JSON.
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
