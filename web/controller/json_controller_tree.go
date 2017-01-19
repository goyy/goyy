// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package controller

import (
	"encoding/json"

	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

// JSONTreeController controller.JSONTreeController.
type JSONTreeController struct {
	baseTreeController
	JSONController
}

// Index gets the list data for the JSON type.
func (me *JSONTreeController) Index(c xhttp.Context) {
	r, err := me.baseTreeController.Index(c, me.Mgr, me.PreIndex, me.PostIndex)
	if err != nil {
		me.Error(c, err)
		return
	}
	parents := c.Attribute(defaultParents)
	b, err := json.Marshal(parents)
	if err != nil {
		me.Error(c, err)
		return
	}
	r.Tag = string(b)
	err = c.Text(xhttp.StatusOK, r.JSON())
	if err != nil {
		me.Error(c, err)
		return
	}
}

// Save saves the data, but does not automatically commit the transaction.
func (me *JSONTreeController) Save(c xhttp.Context) {
	r, err := me.baseTreeController.Save(c, me.Mgr, me.PreSave, me.PostSave)
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

// SaveAndTx saves the data, but automatically commits the transaction.
func (me *JSONTreeController) SaveAndTx(c xhttp.Context) {
	r, err := me.baseTreeController.SaveAndTx(c, me.Mgr, me.PreSave, me.PostSave)
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

// Disable delete the data, but does not automatically commit the transaction.
func (me *JSONTreeController) Disable(c xhttp.Context) {
	r, err := me.baseTreeController.Disable(c, me.Mgr, me.PreDisable, me.PostDisable)
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

// DisableAndTx delete the data, but automatically commits the transaction.
func (me *JSONTreeController) DisableAndTx(c xhttp.Context) {
	r, err := me.baseTreeController.DisableAndTx(c, me.Mgr, me.PreDisable, me.PostDisable)
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

// Tree gets a list of tree types and converts them to JSON.
func (me *JSONTreeController) Tree(c xhttp.Context) {
	out, err := me.baseTreeController.Tree(c, me.Mgr)
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
