// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package controller

import (
	"gopkg.in/goyy/goyy.v0/data/service"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

type JSONController struct {
	baseController
	pre
	post
	Settings
	Mgr service.Service
}

func (me *JSONController) Index(c xhttp.Context) {
	r, err := me.baseController.Index(c, me.Mgr, me.PreIndex, me.PostIndex)
	if err != nil {
		me.Error(c, err)
		return
	}
	err = c.Text(xhttp.StatusOK, r.JSON())
	if err != nil {
		me.Error(c, err)
		return
	}
}

func (me *JSONController) Show(c xhttp.Context) {
	r, err := me.baseController.Show(c, me.Mgr, me.PreShow, me.PostShow)
	if err != nil {
		me.Error(c, err)
		return
	}
	err = c.Text(xhttp.StatusOK, r.JSON())
	if err != nil {
		me.Error(c, err)
		return
	}
}

func (me *JSONController) Add(c xhttp.Context) {
	r, err := me.baseController.Add(c, me.Mgr, me.PreAdd, me.PostAdd)
	if err != nil {
		me.Error(c, err)
		return
	}
	err = c.Text(xhttp.StatusOK, r.JSON())
	if err != nil {
		me.Error(c, err)
		return
	}
}

func (me *JSONController) Edit(c xhttp.Context) {
	r, err := me.baseController.Edit(c, me.Mgr, me.PreEdit, me.PostEdit)
	if err != nil {
		me.Error(c, err)
		return
	}
	err = c.Text(xhttp.StatusOK, r.JSON())
	if err != nil {
		me.Error(c, err)
		return
	}
}

func (me *JSONController) Save(c xhttp.Context) {
	r, err := me.baseController.Save(c, me.Mgr, me.PreSave, me.PostSave)
	if err != nil {
		me.Error(c, err)
		return
	}
	err = c.Text(xhttp.StatusOK, r.JSON())
	if err != nil {
		me.Error(c, err)
		return
	}
}

func (me *JSONController) Saved(c xhttp.Context) {
	r, err := me.baseController.Saved(c, me.Mgr, me.PreSaved, me.PostSaved)
	if err != nil {
		me.Error(c, err)
		return
	}
	err = c.Text(xhttp.StatusOK, r.JSON())
	if err != nil {
		me.Error(c, err)
		return
	}
}

func (me *JSONController) Disable(c xhttp.Context) {
	r, err := me.baseController.Disable(c, me.Mgr, me.PreDisable, me.PostDisable)
	if err != nil {
		me.Error(c, err)
		return
	}
	err = c.Text(xhttp.StatusOK, r.JSON())
	if err != nil {
		me.Error(c, err)
		return
	}
}

func (me *JSONController) Disabled(c xhttp.Context) {
	r, err := me.baseController.Disabled(c, me.Mgr, me.PreDisabled, me.PostDisabled)
	if err != nil {
		me.Error(c, err)
		return
	}
	err = c.Text(xhttp.StatusOK, r.JSON())
	if err != nil {
		me.Error(c, err)
		return
	}
}

func (me *JSONController) Box(c xhttp.Context) {
	out, err := me.baseController.Box(c, me.Mgr)
	if err != nil {
		me.Error(c, err)
		return
	}
	err = c.JSON(xhttp.StatusOK, me.Success(c, out))
	if err != nil {
		me.Error(c, err)
		return
	}
}
