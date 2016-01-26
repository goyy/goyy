// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package controller

import (
	"fmt"

	"gopkg.in/goyy/goyy.v0/util/strings"
)

const (
	pIndex    = "index"
	pShow     = "show"
	pView     = "view"
	pAdd      = "add"
	pEdit     = "edit"
	pSave     = "save"
	pSaved    = "saved"
	pDisable  = "disable"
	pDisabled = "disabled"
	pTree     = "tree"
	pBox      = "box"
	pExport   = "export"

	vList = "list"
	vForm = "form"
)

// ----------------------------------------------------------
// template
// ----------------------------------------------------------

func TmplBy(project, module, name string) string {
	if strings.Contains(module, ".") {
		module = strings.Replace(module, ".", "/", -1)
		return fmt.Sprintf("%s/%s.%s", project, module, name)
	} else {
		return fmt.Sprintf("%s/%s/%s", project, module, name)
	}
}

func TmplDefault(project, module string) string {
	if strings.Contains(module, ".") {
		modules := strings.Split(module, ".")
		return TmplBy(project, modules[0], module)
	} else {
		return TmplBy(project, module, module)
	}
}

func TmplIndex(project, module string) string {
	return TmplBy(project, module, pIndex)
}

func TmplList(project, module string) string {
	return TmplBy(project, module, vList)
}

func TmplForm(project, module string) string {
	return TmplBy(project, module, vForm)
}

// ----------------------------------------------------------
// path
// ----------------------------------------------------------

func PathBy(project, module, name string) string {
	if strings.Contains(module, ".") {
		module = strings.Replace(module, ".", "/", -1)
	}
	if strings.IsBlank(name) {
		return fmt.Sprintf("/%s/%s", project, module)
	} else {
		return fmt.Sprintf("/%s/%s/%s", project, module, name)
	}
}

func PathModule(project, module string) string {
	return PathBy(project, module, "")
}

func PathIndex(project, module string) string {
	return PathBy(project, module, pIndex)
}

func PathShow(project, module string) string {
	return PathBy(project, module, pShow)
}

func PathAdd(project, module string) string {
	return PathBy(project, module, pAdd)
}

func PathEdit(project, module string) string {
	return PathBy(project, module, pEdit)
}

func PathSave(project, module string) string {
	return PathBy(project, module, pSave)
}

func PathSaved(project, module string) string {
	return PathBy(project, module, pSaved)
}

func PathDisable(project, module string) string {
	return PathBy(project, module, pDisable)
}

func PathDisabled(project, module string) string {
	return PathBy(project, module, pDisabled)
}

func PathTree(project, module string) string {
	return PathBy(project, module, pTree)
}

func PathBox(project, module string) string {
	return PathBy(project, module, pBox)
}

func PathExport(project, module string) string {
	return PathBy(project, module, pExport)
}

// ----------------------------------------------------------
// api
// ----------------------------------------------------------

func ApiBy(project, module, name string) string {
	if strings.Contains(module, ".") {
		module = strings.Replace(module, ".", "/", -1)
	}
	if strings.IsBlank(name) {
		return fmt.Sprintf("/apis/%s/%s", project, module)
	} else {
		return fmt.Sprintf("/apis/%s/%s/%s", project, module, name)
	}
}

func ApiModule(project, module string) string {
	return ApiBy(project, module, "")
}

func ApiIndex(project, module string) string {
	return ApiBy(project, module, pIndex)
}

func ApiShow(project, module string) string {
	return ApiBy(project, module, pShow)
}

func ApiAdd(project, module string) string {
	return ApiBy(project, module, pAdd)
}

func ApiEdit(project, module string) string {
	return ApiBy(project, module, pEdit)
}

func ApiSave(project, module string) string {
	return ApiBy(project, module, pSave)
}

func ApiSaved(project, module string) string {
	return ApiBy(project, module, pSaved)
}

func ApiDisable(project, module string) string {
	return ApiBy(project, module, pDisable)
}

func ApiDisabled(project, module string) string {
	return ApiBy(project, module, pDisabled)
}

func ApiTree(project, module string) string {
	return ApiBy(project, module, pTree)
}

func ApiBox(project, module string) string {
	return ApiBy(project, module, pBox)
}

func ApiExport(project, module string) string {
	return ApiBy(project, module, pExport)
}

// ----------------------------------------------------------
// permission
// ----------------------------------------------------------

func PermitBy(project, module, name string) string {
	if strings.Contains(module, ".") {
		module = strings.Replace(module, ".", ":", -1)
	}
	return fmt.Sprintf("%s:%s:%s", project, module, name)
}

func PermitView(project, module string) string {
	return PermitBy(project, module, pView)
}

func PermitAdd(project, module string) string {
	return PermitBy(project, module, pAdd)
}

func PermitEdit(project, module string) string {
	return PermitBy(project, module, pEdit)
}

func PermitDisable(project, module string) string {
	return PermitBy(project, module, pDisable)
}

func PermitExport(project, module string) string {
	return PermitBy(project, module, pExport)
}
