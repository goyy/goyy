// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplNewProjSchemaBinSqlSh = `#!/bin/sh

echo [INFO] Create DDL SQL Files.

cd ..
export I18N_LOCALE=zh_CN
xtab -proj=<%.NewProjName%> -sql
`
