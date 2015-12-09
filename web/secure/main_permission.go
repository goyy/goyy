// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package secure

import (
	"gopkg.in/goyy/goyy.v0/util/strings"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

func IsPermission(c xhttp.Context, permission string) bool {
	if strings.IsBlank(permission) {
		return false
	}
	if c == nil || !c.Session().IsLogin() {
		return false
	}
	if p, err := c.Session().Principal(); err == nil {
		if strings.Contains(p.Permissions, permission) {
			return true
		}
	} else {
		logger.Error(err.Error())
		return false
	}
	return false
}

func IsAnyPermission(c xhttp.Context, permissions string) bool {
	if strings.IsBlank(permissions) {
		return false
	}
	if c == nil || !c.Session().IsLogin() {
		return false
	}
	ps := strings.Split(permissions, ",")
	if p, err := c.Session().Principal(); err == nil {
		if strings.ContainsSliceAny(p.Permissions, ps) {
			return true
		}
	} else {
		logger.Error(err.Error())
		return false
	}
	return false
}
