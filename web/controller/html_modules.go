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

func (me *HTMLController) TmplBy(name string) string {
	return TmplBy(me.Project, me.Module, name)
}

func (me *HTMLController) TmplDefault() string {
	return TmplDefault(me.Project, me.Module)
}

func (me *HTMLController) TmplIndex() string {
	return TmplIndex(me.Project, me.Module)
}

func (me *HTMLController) TmplList() string {
	return TmplList(me.Project, me.Module)
}

func (me *HTMLController) TmplForm() string {
	return TmplForm(me.Project, me.Module)
}

// ----------------------------------------------------------
// path
// ----------------------------------------------------------

func (me *HTMLController) PathBy(name string) string {
	return PathBy(me.Project, me.Module, name+".html")
}

func (me *HTMLController) PathIndex() string {
	return me.PathBy(pIndex)
}

func (me *HTMLController) PathShow() string {
	return me.PathBy(pShow)
}

func (me *HTMLController) PathAdd() string {
	return me.PathBy(pAdd)
}

func (me *HTMLController) PathEdit() string {
	return me.PathBy(pEdit)
}

// ----------------------------------------------------------
// permission
// ----------------------------------------------------------

func (me *HTMLController) PermitBy(name string) *xtype.Permission {
	return PermitBy(me.Project, me.Module, name)
}

func (me *HTMLController) PermitView() *xtype.Permission {
	return PermitView(me.Project, me.Module)
}

func (me *HTMLController) PermitAdd() *xtype.Permission {
	return PermitAdd(me.Project, me.Module)
}

func (me *HTMLController) PermitEdit() *xtype.Permission {
	return PermitEdit(me.Project, me.Module)
}

func (me *HTMLController) PermitDisable() *xtype.Permission {
	return PermitDisable(me.Project, me.Module)
}
