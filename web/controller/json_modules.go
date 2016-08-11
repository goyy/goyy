// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package controller

import (
	"gopkg.in/goyy/goyy.v0/comm/xtype"
)

// ----------------------------------------------------------
// api
// ----------------------------------------------------------

func (me *JSONController) ApiBy(name string) string {
	return ApiBy(me.Project, me.Module, name)
}

func (me *JSONController) ApiIndex() string {
	return ApiIndex(me.Project, me.Module)
}

func (me *JSONController) ApiShow() string {
	return ApiShow(me.Project, me.Module)
}

func (me *JSONController) ApiAdd() string {
	return ApiAdd(me.Project, me.Module)
}

func (me *JSONController) ApiEdit() string {
	return ApiEdit(me.Project, me.Module)
}

func (me *JSONController) ApiSave() string {
	return ApiSave(me.Project, me.Module)
}

func (me *JSONController) ApiSaved() string {
	return ApiSaved(me.Project, me.Module)
}

func (me *JSONController) ApiDisable() string {
	return ApiDisable(me.Project, me.Module)
}

func (me *JSONController) ApiDisabled() string {
	return ApiDisabled(me.Project, me.Module)
}

func (me *JSONController) ApiTree() string {
	return ApiTree(me.Project, me.Module)
}

func (me *JSONController) ApiBox() string {
	return ApiBox(me.Project, me.Module)
}

func (me *JSONController) ApiExport() string {
	return ApiExport(me.Project, me.Module)
}

// ----------------------------------------------------------
// permission
// ----------------------------------------------------------

func (me *JSONController) PermitBy(name string) *xtype.Permission {
	return PermitBy(me.Project, me.Module, name)
}

func (me *JSONController) PermitView() *xtype.Permission {
	return PermitView(me.Project, me.Module)
}

func (me *JSONController) PermitAdd() *xtype.Permission {
	return PermitAdd(me.Project, me.Module)
}

func (me *JSONController) PermitEdit() *xtype.Permission {
	return PermitEdit(me.Project, me.Module)
}

func (me *JSONController) PermitDisable() *xtype.Permission {
	return PermitDisable(me.Project, me.Module)
}

func (me *JSONController) PermitExport() *xtype.Permission {
	return PermitExport(me.Project, me.Module)
}
