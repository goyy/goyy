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

// ApiBy gets api path by name.
func (me *JSONController) ApiBy(name string) string {
	return ApiBy(me.Project, me.Module, name)
}

// ApiIndex gets the api path with the name index.
func (me *JSONController) ApiIndex() string {
	return ApiIndex(me.Project, me.Module)
}

// ApiShow gets the api path with the name show.
func (me *JSONController) ApiShow() string {
	return ApiShow(me.Project, me.Module)
}

// ApiAdd gets the api path with the name add.
func (me *JSONController) ApiAdd() string {
	return ApiAdd(me.Project, me.Module)
}

// ApiEdit gets the api path with the name edit.
func (me *JSONController) ApiEdit() string {
	return ApiEdit(me.Project, me.Module)
}

// ApiSave gets the api path with the name save.
func (me *JSONController) ApiSave() string {
	return ApiSave(me.Project, me.Module)
}

// ApiDisable gets the api path with the name disable.
func (me *JSONController) ApiDisable() string {
	return ApiDisable(me.Project, me.Module)
}

// ApiTree gets the api path with the name tree.
func (me *JSONController) ApiTree() string {
	return ApiTree(me.Project, me.Module)
}

// ApiBox gets the api path with the name box.
func (me *JSONController) ApiBox() string {
	return ApiBox(me.Project, me.Module)
}

// ApiExport gets the api path with the name export.
func (me *JSONController) ApiExport() string {
	return ApiExport(me.Project, me.Module)
}

// ----------------------------------------------------------
// permission
// ----------------------------------------------------------

// Permit gets permission by id and profiles.
func (me *JSONController) Permit(id string, profiles ...string) *xtype.Permission {
	return Permit(id, profiles...)
}

// PermitBy gets permission by name and profiles.
func (me *JSONController) PermitBy(name string, profiles ...string) *xtype.Permission {
	return PermitBy(me.Project, me.Module, name, profiles...)
}

// PermitView obtain the view type permission from the profiles.
func (me *JSONController) PermitView(profiles ...string) *xtype.Permission {
	return PermitView(me.Project, me.Module, profiles...)
}

// PermitAdd obtain the add type permission from the profiles.
func (me *JSONController) PermitAdd(profiles ...string) *xtype.Permission {
	return PermitAdd(me.Project, me.Module, profiles...)
}

// PermitEdit obtain the edit type permission from the profiles.
func (me *JSONController) PermitEdit(profiles ...string) *xtype.Permission {
	return PermitEdit(me.Project, me.Module, profiles...)
}

// PermitDisable obtain the disable type permission from the profiles.
func (me *JSONController) PermitDisable(profiles ...string) *xtype.Permission {
	return PermitDisable(me.Project, me.Module, profiles...)
}

// PermitExport obtain the export type permission from the profiles.
func (me *JSONController) PermitExport(profiles ...string) *xtype.Permission {
	return PermitExport(me.Project, me.Module, profiles...)
}
