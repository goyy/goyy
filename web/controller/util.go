// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package controller

import (
	"fmt"
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

func PathDisable(project, module string) string {
	return PathBy(project, module, pDisable)
}

func PathTree(project, module string) string {
	return PathBy(project, module, pTree)
}

func PathBox(project, module string) string {
	return PathBy(project, module, pBox)
}

// ----------------------------------------------------------
// api
// ----------------------------------------------------------

func ApiBy(project, module, name string) string {
	if strings.Contains(module, ".") {
		module = strings.Replace(module, ".", "/", -1)
	}
	if strings.IsBlank(name) {
		return fmt.Sprintf("/api/%s/%s", project, module)
	} else {
		return fmt.Sprintf("/api/%s/%s/%s", project, module, name)
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

func ApiDisable(project, module string) string {
	return ApiBy(project, module, pDisable)
}

func ApiTree(project, module string) string {
	return ApiBy(project, module, pTree)
}

func ApiBox(project, module string) string {
	return ApiBy(project, module, pBox)
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
