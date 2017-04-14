package role

import (
	"gopkg.in/goyy/goyy.v0/comm/i18n"
)

func SetLocale(locale string) {
	i18N.SetLocale(locale)
}

var i18N = i18n.New(locales)

var locales = map[string]map[string]string{
	i18n.LocaleZhCN: map[string]string{
		"not.logged.err": "未登录！",
	},
	i18n.LocaleEnUS: map[string]string{
		"not.logged.err": "Not logged in!",
	},
}
