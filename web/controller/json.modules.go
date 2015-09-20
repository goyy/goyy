// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of me source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package controller

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

func (me *JSONController) ApiDisable() string {
	return ApiDisable(me.Project, me.Module)
}

func (me *JSONController) ApiTree() string {
	return ApiTree(me.Project, me.Module)
}

func (me *JSONController) ApiBox() string {
	return ApiBox(me.Project, me.Module)
}

// ----------------------------------------------------------
// permission
// ----------------------------------------------------------

func (me *JSONController) PermitBy(name string) string {
	return PermitBy(me.Project, me.Module, name)
}

func (me *JSONController) PermitView() string {
	return PermitView(me.Project, me.Module)
}

func (me *JSONController) PermitAdd() string {
	return PermitAdd(me.Project, me.Module)
}

func (me *JSONController) PermitEdit() string {
	return PermitEdit(me.Project, me.Module)
}

func (me *JSONController) PermitDisable() string {
	return PermitDisable(me.Project, me.Module)
}
