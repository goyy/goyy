// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package controller

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

func (me *HTMLController) PermitBy(name string) string {
	return PermitBy(me.Project, me.Module, name)
}

func (me *HTMLController) PermitView() string {
	return PermitView(me.Project, me.Module)
}

func (me *HTMLController) PermitAdd() string {
	return PermitAdd(me.Project, me.Module)
}

func (me *HTMLController) PermitEdit() string {
	return PermitEdit(me.Project, me.Module)
}

func (me *HTMLController) PermitDisable() string {
	return PermitDisable(me.Project, me.Module)
}
