// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package db

import (
	"gopkg.in/goyy/goyy.v0/comm/i18n"
)

var i18N = i18n.New(locales, i18n.Locale_en_US)

var locales = map[string]map[string]string{
	i18n.Locale_en_US: en_US,
}

var en_US = map[string]string{
	"err.empty.table.name":  "table name can not be empty!",
	"err.empty.column.name": "column name can not be empty!",
	"err.assign.exp":        "cannot use %s as type *expression in assignment",
	"err.file.exist":        "file(%s) not exists",
	"err.mapping.pointer":   "not pointer type",
	"err.mapping.slice":     "not slice type",
	"err.mapping.struct":    "not struct type",
}
