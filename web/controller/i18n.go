// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package controller

import (
	"gopkg.in/goyy/goyy.v0/comm/i18n"
)

func SetLocale(locale string) {
	i18N.SetLocale(locale)
}

var i18N = i18n.New(locales)

var locales = map[string]map[string]string{
	i18n.Locale_zh_CN: map[string]string{
		"msg.save":         "保存成功",
		"msg.disable":      "删除成功",
		"msg.disable.leaf": "非叶子节点不允许删除",
		"exp.limit":        "导出记录数不能大于10000",
		"exp.data.blank":   "导出数据参数不能为空",
	},
	i18n.Locale_en_US: map[string]string{
		"msg.save":         "Save success",
		"msg.disable":      "Delete success",
		"msg.disable.leaf": "Delete is not allowed because it is a non leaf node",
		"exp.limit":        "The number of records can not be more than 10000",
		"exp.data.blank":   "Export data parameters can not be empty",
	},
}
