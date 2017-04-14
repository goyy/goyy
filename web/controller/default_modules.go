// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package controller

import (
	"gopkg.in/goyy/goyy.v0/comm/xtype"
)

// ----------------------------------------------------------
// template
// ----------------------------------------------------------

// TmplBy gets template by name.
func (me *Controller) TmplBy(name string) string {
	return TmplBy(me.Project, me.Module, name)
}

// TmplDefault gets the template with the default name.
func (me *Controller) TmplDefault() string {
	return TmplDefault(me.Project, me.Module)
}

// TmplIndex gets the template with the name index.
func (me *Controller) TmplIndex() string {
	return TmplIndex(me.Project, me.Module)
}

// TmplList gets the template with the name list.
func (me *Controller) TmplList() string {
	return TmplList(me.Project, me.Module)
}

// TmplForm gets the template with the name form.
func (me *Controller) TmplForm() string {
	return TmplForm(me.Project, me.Module)
}

// ----------------------------------------------------------
// path
// ----------------------------------------------------------

// PathBy gets path by name.
func (me *Controller) PathBy(name string) string {
	return PathBy(me.Project, me.Module, name)
}

// PathIndex gets the path with the name index.
func (me *Controller) PathIndex() string {
	return PathIndex(me.Project, me.Module)
}

// PathShow gets the path with the name show.
func (me *Controller) PathShow() string {
	return PathShow(me.Project, me.Module)
}

// PathAdd gets the path with the name add.
func (me *Controller) PathAdd() string {
	return PathAdd(me.Project, me.Module)
}

// PathEdit gets the path with the name edit.
func (me *Controller) PathEdit() string {
	return PathEdit(me.Project, me.Module)
}

// PathSave gets the path with the name save.
func (me *Controller) PathSave() string {
	return PathSave(me.Project, me.Module)
}

// PathDisable gets the path with the name disable.
func (me *Controller) PathDisable() string {
	return PathDisable(me.Project, me.Module)
}

// PathTree gets the path with the name tree.
func (me *Controller) PathTree() string {
	return PathTree(me.Project, me.Module)
}

// PathBox gets the path with the name box.
func (me *Controller) PathBox() string {
	return PathBox(me.Project, me.Module)
}

// ----------------------------------------------------------
// permission
// ----------------------------------------------------------

// Permit gets permission by id and profiles.
func (me *Controller) Permit(id string, profiles ...string) *xtype.Permission {
	return Permit(id, profiles...)
}

// PermitBy gets permission by name and profiles.
func (me *Controller) PermitBy(name string, profiles ...string) *xtype.Permission {
	return PermitBy(me.Project, me.Module, name, profiles...)
}

// PermitView obtain the view type permission from the profiles.
func (me *Controller) PermitView(profiles ...string) *xtype.Permission {
	return PermitView(me.Project, me.Module, profiles...)
}

// PermitAdd obtain the add type permission from the profiles.
func (me *Controller) PermitAdd(profiles ...string) *xtype.Permission {
	return PermitAdd(me.Project, me.Module, profiles...)
}

// PermitEdit obtain the edit type permission from the profiles.
func (me *Controller) PermitEdit(profiles ...string) *xtype.Permission {
	return PermitEdit(me.Project, me.Module, profiles...)
}

// PermitDisable obtain the disable type permission from the profiles.
func (me *Controller) PermitDisable(profiles ...string) *xtype.Permission {
	return PermitDisable(me.Project, me.Module, profiles...)
}
