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

func (me *Controller) TmplBy(name string) string {
	return TmplBy(me.Project, me.Module, name)
}

func (me *Controller) TmplDefault() string {
	return TmplDefault(me.Project, me.Module)
}

func (me *Controller) TmplIndex() string {
	return TmplIndex(me.Project, me.Module)
}

func (me *Controller) TmplList() string {
	return TmplList(me.Project, me.Module)
}

func (me *Controller) TmplForm() string {
	return TmplForm(me.Project, me.Module)
}

// ----------------------------------------------------------
// path
// ----------------------------------------------------------

func (me *Controller) PathBy(name string) string {
	return PathBy(me.Project, me.Module, name)
}

func (me *Controller) PathIndex() string {
	return PathIndex(me.Project, me.Module)
}

func (me *Controller) PathShow() string {
	return PathShow(me.Project, me.Module)
}

func (me *Controller) PathAdd() string {
	return PathAdd(me.Project, me.Module)
}

func (me *Controller) PathEdit() string {
	return PathEdit(me.Project, me.Module)
}

func (me *Controller) PathSave() string {
	return PathSave(me.Project, me.Module)
}

func (me *Controller) PathDisable() string {
	return PathDisable(me.Project, me.Module)
}

func (me *Controller) PathTree() string {
	return PathTree(me.Project, me.Module)
}

func (me *Controller) PathBox() string {
	return PathBox(me.Project, me.Module)
}

// ----------------------------------------------------------
// permission
// ----------------------------------------------------------

func (me *Controller) Permit(id string, profiles ...string) *xtype.Permission {
	return Permit(id, profiles...)
}

func (me *Controller) PermitBy(name string, profiles ...string) *xtype.Permission {
	return PermitBy(me.Project, me.Module, name, profiles...)
}

func (me *Controller) PermitView(profiles ...string) *xtype.Permission {
	return PermitView(me.Project, me.Module, profiles...)
}

func (me *Controller) PermitAdd(profiles ...string) *xtype.Permission {
	return PermitAdd(me.Project, me.Module, profiles...)
}

func (me *Controller) PermitEdit(profiles ...string) *xtype.Permission {
	return PermitEdit(me.Project, me.Module, profiles...)
}

func (me *Controller) PermitDisable(profiles ...string) *xtype.Permission {
	return PermitDisable(me.Project, me.Module, profiles...)
}
