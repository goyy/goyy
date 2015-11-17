// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xhttp

import (
	"gopkg.in/goyy/goyy.v0/comm/i18n"
)

func SetLocale(locale string) {
	i18N.SetLocale(locale)
}

var i18N = i18n.New(locales, i18n.Locale_zh_CN)

var locales = map[string]map[string]string{
	i18n.Locale_zh_CN: map[string]string{
		"err.illegal": "存在非法字符，请检查输入内容!",
	},
	i18n.Locale_en_US: map[string]string{
		"err.illegal": "There are illegal characters, please check your input!",
	},
}
