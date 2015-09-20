// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package env

import (
	"gopkg.in/goyy/goyy.v0/comm/i18n"
)

var i18N = i18n.New(locales, i18n.Locale_en_US)

var locales = map[string]map[string]string{
	i18n.Locale_en_US: en_US,
}

var en_US = map[string]string{
	"empty.environments.default":              "environments.default: can not be empty!",
	"empty.environments.environment.database": "environments.environment.database[@name='%s']: can not be empty!",
	"empty.environments.environment.mail":     "environments.environment.mail[@name='%s']: can not be empty!",
}
