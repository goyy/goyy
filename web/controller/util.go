// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package controller

import (
	"fmt"

	"gopkg.in/goyy/goyy.v0/comm/xtype"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

const (
	pIndex   = "index"
	pShow    = "show"
	pView    = "view"
	pAdd     = "add"
	pEdit    = "edit"
	pSave    = "save"
	pDisable = "disable"
	pTree    = "tree"
	pBox     = "box"
	pExport  = "export"

	vList = "list"
	vForm = "form"
)

// ----------------------------------------------------------
// template
// ----------------------------------------------------------

// TmplBy gets template by name.
func TmplBy(project, module, name string) string {
	if strings.Contains(module, ".") {
		module = strings.Replace(module, ".", "/", -1)
		return fmt.Sprintf("%s/%s.%s", project, module, name)
	}
	return fmt.Sprintf("%s/%s/%s", project, module, name)
}

// TmplDefault gets the template with the default name.
func TmplDefault(project, module string) string {
	if strings.Contains(module, ".") {
		modules := strings.Split(module, ".")
		return TmplBy(project, modules[0], module)
	}
	return TmplBy(project, module, module)
}

// TmplIndex gets the template with the name index.
func TmplIndex(project, module string) string {
	return TmplBy(project, module, pIndex)
}

// TmplList gets the template with the name list.
func TmplList(project, module string) string {
	return TmplBy(project, module, vList)
}

// TmplForm gets the template with the name form.
func TmplForm(project, module string) string {
	return TmplBy(project, module, vForm)
}

// ----------------------------------------------------------
// path
// ----------------------------------------------------------

// PathBy gets path by name.
func PathBy(project, module, name string) string {
	if strings.Contains(module, ".") {
		module = strings.Replace(module, ".", "/", -1)
	}
	if strings.IsBlank(name) {
		return fmt.Sprintf("/%s/%s", project, module)
	}
	return fmt.Sprintf("/%s/%s/%s", project, module, name)
}

// PathModule gets the path from project and module.
func PathModule(project, module string) string {
	return PathBy(project, module, "")
}

// PathIndex gets the path with the name index.
func PathIndex(project, module string) string {
	return PathBy(project, module, pIndex)
}

// PathShow gets the path with the name show.
func PathShow(project, module string) string {
	return PathBy(project, module, pShow)
}

// PathAdd gets the path with the name add.
func PathAdd(project, module string) string {
	return PathBy(project, module, pAdd)
}

// PathEdit gets the path with the name edit.
func PathEdit(project, module string) string {
	return PathBy(project, module, pEdit)
}

// PathSave gets the path with the name save.
func PathSave(project, module string) string {
	return PathBy(project, module, pSave)
}

// PathDisable gets the path with the name disable.
func PathDisable(project, module string) string {
	return PathBy(project, module, pDisable)
}

// PathTree gets the path with the name tree.
func PathTree(project, module string) string {
	return PathBy(project, module, pTree)
}

// PathBox gets the path with the name box.
func PathBox(project, module string) string {
	return PathBy(project, module, pBox)
}

// PathExport gets the path with the name export.
func PathExport(project, module string) string {
	return PathBy(project, module, pExport)
}

// ----------------------------------------------------------
// api
// ----------------------------------------------------------

// ApiBy gets api path by name.
func ApiBy(project, module, name string) string {
	if strings.Contains(module, ".") {
		module = strings.Replace(module, ".", "/", -1)
	}
	if strings.IsBlank(name) {
		return fmt.Sprintf("/apis/%s/%s", project, module)
	}
	return fmt.Sprintf("/apis/%s/%s/%s", project, module, name)
}

// ApiModule gets the api path from project and module.
func ApiModule(project, module string) string {
	return ApiBy(project, module, "")
}

// ApiIndex gets the api path with the name index.
func ApiIndex(project, module string) string {
	return ApiBy(project, module, pIndex)
}

// ApiShow gets the api path with the name show.
func ApiShow(project, module string) string {
	return ApiBy(project, module, pShow)
}

// ApiAdd gets the api path with the name add.
func ApiAdd(project, module string) string {
	return ApiBy(project, module, pAdd)
}

// ApiEdit gets the api path with the name edit.
func ApiEdit(project, module string) string {
	return ApiBy(project, module, pEdit)
}

// ApiSave gets the api path with the name save.
func ApiSave(project, module string) string {
	return ApiBy(project, module, pSave)
}

// ApiDisable gets the api path with the name disable.
func ApiDisable(project, module string) string {
	return ApiBy(project, module, pDisable)
}

// ApiTree gets the api path with the name tree.
func ApiTree(project, module string) string {
	return ApiBy(project, module, pTree)
}

// ApiBox gets the api path with the name box.
func ApiBox(project, module string) string {
	return ApiBy(project, module, pBox)
}

// ApiExport gets the api path with the name export.
func ApiExport(project, module string) string {
	return ApiBy(project, module, pExport)
}

// ----------------------------------------------------------
// permission
// ----------------------------------------------------------

// Permit gets permission by id and profiles.
func Permit(id string, profiles ...string) *xtype.Permission {
	return &xtype.Permission{Id: id, Profiles: profiles}
}

// PermitBy gets permission by name and profiles.
func PermitBy(project, module, name string, profiles ...string) *xtype.Permission {
	if strings.Contains(module, ".") {
		module = strings.Replace(module, ".", ":", -1)
	}
	id := fmt.Sprintf("%s:%s:%s", project, module, name)
	return Permit(id, profiles...)
}

// PermitView obtain the view type permission from the profiles.
func PermitView(project, module string, profiles ...string) *xtype.Permission {
	return PermitBy(project, module, pView, profiles...)
}

// PermitAdd obtain the add type permission from the profiles.
func PermitAdd(project, module string, profiles ...string) *xtype.Permission {
	return PermitBy(project, module, pAdd, profiles...)
}

// PermitEdit obtain the edit type permission from the profiles.
func PermitEdit(project, module string, profiles ...string) *xtype.Permission {
	return PermitBy(project, module, pEdit, profiles...)
}

// PermitDisable obtain the disable type permission from the profiles.
func PermitDisable(project, module string, profiles ...string) *xtype.Permission {
	return PermitBy(project, module, pDisable, profiles...)
}

// PermitExport obtain the export type permission from the profiles.
func PermitExport(project, module string, profiles ...string) *xtype.Permission {
	return PermitBy(project, module, pExport, profiles...)
}
