// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package env

import (
	"gopkg.in/goyy/goyy.v0/comm/i18n"
)

func SetLocale(locale string) {
	i18N.SetLocale(locale)
}

var i18N = i18n.New(locales)

var locales = map[string]map[string]string{
	i18n.Locale_zh_CN: map[string]string{
		"empty.environments.default":               "environments.default: 不允许为空！",
		"empty.environments.environment.database":  "environments.environment.database[@name='%s']: 不允许为空！",
		"empty.environments.environment.mail":      "environments.environment.mail[@name='%s']: 不允许为空！",
		"empty.environments.environment.session":   "environments.environment.session[@name='%s']: 不允许为空！",
		"empty.environments.environment.api":       "environments.environment.api[@name='%s']: 不允许为空！",
		"empty.environments.environment.asset":     "environments.environment.asset[@name='%s']: 不允许为空！",
		"empty.environments.environment.static":    "environments.environment.static[@name='%s']: 不允许为空！",
		"empty.environments.environment.developer": "environments.environment.developer[@name='%s']: 不允许为空！",
		"empty.environments.environment.operation": "environments.environment.operation[@name='%s']: 不允许为空！",
		"empty.environments.environment.upload":    "environments.environment.upload[@name='%s']: 不允许为空！",
		"empty.environments.environment.export":    "environments.environment.export[@name='%s']: 不允许为空！",
		"empty.environments.environment.html":      "environments.environment.html[@name='%s']: 不允许为空！",
		"empty.environments.environment.template":  "environments.environment.template[@name='%s']: 不允许为空！",
	},
	i18n.Locale_en_US: map[string]string{
		"empty.environments.default":               "environments.default: can not be empty!",
		"empty.environments.environment.database":  "environments.environment.database[@name='%s']: can not be empty!",
		"empty.environments.environment.mail":      "environments.environment.mail[@name='%s']: can not be empty!",
		"empty.environments.environment.session":   "environments.environment.session[@name='%s']: can not be empty!",
		"empty.environments.environment.api":       "environments.environment.api[@name='%s']: can not be empty!",
		"empty.environments.environment.asset":     "environments.environment.asset[@name='%s']: can not be empty!",
		"empty.environments.environment.static":    "environments.environment.static[@name='%s']: can not be empty!",
		"empty.environments.environment.developer": "environments.environment.developer[@name='%s']: can not be empty!",
		"empty.environments.environment.operation": "environments.environment.operation[@name='%s']: can not be empty!",
		"empty.environments.environment.upload":    "environments.environment.upload[@name='%s']: can not be empty!",
		"empty.environments.environment.export":    "environments.environment.export[@name='%s']: can not be empty!",
		"empty.environments.environment.html":      "environments.environment.html[@name='%s']: can not be empty!",
		"empty.environments.environment.template":  "environments.environment.template[@name='%s']: can not be empty!",
	},
}
