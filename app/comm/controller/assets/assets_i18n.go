package assets

import (
	"gopkg.in/goyy/goyy.v0/comm/i18n"
)

func SetLocale(locale string) {
	i18N.SetLocale(locale)
}

var i18N = i18n.New(locales)

var locales = map[string]map[string]string{
	i18n.Locale_zh_CN: map[string]string{
		"upload.img.err": "上传图片失败:",
	},
	i18n.Locale_en_US: map[string]string{
		"upload.img.err": "Upload image failed:",
	},
}
