package area

import (
	"gopkg.in/goyy/goyy.v0/comm/i18n"
)

func SetLocale(locale string) {
	i18N.SetLocale(locale)
}

var i18N = i18n.New(locales)

var locales = map[string]map[string]string{
	i18n.Locale_zh_CN: map[string]string{
		"get.name.err":        "获取名称失败！",
		"get.fullname.err":    "获取全称失败！",
		"get.parentid.err":    "获取上一级标识失败！",
		"get.parentpid.err":   "获取上上一级标识失败！",
		"get.parentpname.err": "获取上上一级名称失败！",
	},
	i18n.Locale_en_US: map[string]string{
		"get.name.err":        "Failed to get the name!",
		"get.fullname.err":    "Failed to get the full name!",
		"get.parentid.err":    "Failed to get the superior ID!",
		"get.parentpid.err":   "Failed to get the superior ID of the superior!",
		"get.parentpname.err": "Failed to get the superior name of the superior!",
	},
}
