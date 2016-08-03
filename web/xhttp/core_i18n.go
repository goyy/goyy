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

var i18N = i18n.New(locales)

var locales = map[string]map[string]string{
	i18n.Locale_zh_CN: map[string]string{
		"err.illegal": "存在非法字符，请检查输入内容！",
		"err.401":     "您没有权限访问该页面！",
		"err.403":     "访问被拒绝！",
		"err.404":     "抱歉，您访问的页面不存在！",
		"err.405":     "不允许此方法！",
		"err.500":     "抱歉，您访问的页面出现了问题！",
	},
	i18n.Locale_en_US: map[string]string{
		"err.illegal": "There are illegal characters, please check your input!",
		"err.401":     "401 Unauthorized",
		"err.403":     "403 Forbidden",
		"err.404":     "404 page not found",
		"err.405":     "405 method not allowed",
		"err.500":     "500 Internal Server Error",
	},
}
