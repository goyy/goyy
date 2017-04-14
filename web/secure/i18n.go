// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package secure

import (
	"gopkg.in/goyy/goyy.v0/comm/i18n"
)

// SetLocale set the i18n locale.
func SetLocale(locale string) {
	i18N.SetLocale(locale)
}

var i18N = i18n.New(locales)

var locales = map[string]map[string]string{
	i18n.LocaleZhCN: map[string]string{
		"err.login": "无效的登录名或密码",
	},
	i18n.LocaleEnUS: map[string]string{
		"err.login": "Invalid user name or password",
	},
}
