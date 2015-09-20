// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of me source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package controller

// ----------------------------------------------------------
// template
// ----------------------------------------------------------

func (me *ClientController) TmplBy(name string) string {
	return TmplBy(me.Project, me.Module, name)
}

func (me *ClientController) TmplDefault() string {
	return TmplDefault(me.Project, me.Module)
}

func (me *ClientController) TmplIndex() string {
	return TmplIndex(me.Project, me.Module)
}

func (me *ClientController) TmplList() string {
	return TmplList(me.Project, me.Module)
}

func (me *ClientController) TmplForm() string {
	return TmplForm(me.Project, me.Module)
}

// ----------------------------------------------------------
// api
// ----------------------------------------------------------

func (me *ClientController) ApiBy(name string) string {
	return ApiBy(me.Project, me.Module, name)
}

func (me *ClientController) ApiIndex() string {
	return ApiIndex(me.Project, me.Module)
}

func (me *ClientController) ApiShow() string {
	return ApiShow(me.Project, me.Module)
}

func (me *ClientController) ApiAdd() string {
	return ApiAdd(me.Project, me.Module)
}

func (me *ClientController) ApiEdit() string {
	return ApiEdit(me.Project, me.Module)
}

func (me *ClientController) ApiSave() string {
	return ApiSave(me.Project, me.Module)
}

func (me *ClientController) ApiDisable() string {
	return ApiDisable(me.Project, me.Module)
}

func (me *ClientController) ApiTree() string {
	return ApiTree(me.Project, me.Module)
}

func (me *ClientController) ApiBox() string {
	return ApiBox(me.Project, me.Module)
}

// ----------------------------------------------------------
// path
// ----------------------------------------------------------

func (me *ClientController) PathBy(name string) string {
	return PathBy(me.Project, me.Module, name)
}

func (me *ClientController) PathIndex() string {
	return PathIndex(me.Project, me.Module)
}

func (me *ClientController) PathShow() string {
	return PathShow(me.Project, me.Module)
}

func (me *ClientController) PathAdd() string {
	return PathAdd(me.Project, me.Module)
}

func (me *ClientController) PathEdit() string {
	return PathEdit(me.Project, me.Module)
}

func (me *ClientController) PathSave() string {
	return PathSave(me.Project, me.Module)
}

func (me *ClientController) PathDisable() string {
	return PathDisable(me.Project, me.Module)
}

func (me *ClientController) PathTree() string {
	return PathTree(me.Project, me.Module)
}

func (me *ClientController) PathBox() string {
	return PathBox(me.Project, me.Module)
}

// ----------------------------------------------------------
// permission
// ----------------------------------------------------------

func (me *ClientController) PermitBy(name string) string {
	return PermitBy(me.Project, me.Module, name)
}

func (me *ClientController) PermitView() string {
	return PermitView(me.Project, me.Module)
}

func (me *ClientController) PermitAdd() string {
	return PermitAdd(me.Project, me.Module)
}

func (me *ClientController) PermitEdit() string {
	return PermitEdit(me.Project, me.Module)
}

func (me *ClientController) PermitDisable() string {
	return PermitDisable(me.Project, me.Module)
}
