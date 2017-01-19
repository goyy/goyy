// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package controller

import (
	"gopkg.in/goyy/goyy.v0/util/templates"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

// TreeController controller.TreeController.
type TreeController struct {
	baseTreeController
	Controller
}

// Index displays the list of pages.
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

// Save used to save the form, but do not automatically commit the transaction.
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

// SaveAndTx used to save the form, but automatically commit the transaction.
func (me *TreeController) SaveAndTx(c xhttp.Context) {
	out, err := me.baseTreeController.SaveAndTx(c, me.Mgr, me.PreSave, me.PostSave)
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

// Disable tsed to delete the records corresponding to the form, but do not automatically commit the transaction.
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

// DisableAndTx used to delete the record corresponding to the form, but automatically commit the transaction.
func (me *TreeController) DisableAndTx(c xhttp.Context) {
	out, err := me.baseTreeController.DisableAndTx(c, me.Mgr, me.PreDisable, me.PostDisable)
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

// Tree gets a list of tree types and converts them to JSON.
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
