// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplNewProjSchemaBinDbSh = `#!/bin/sh

echo [INFO] Initialize database.

cd ..
export I18N_LOCALE=zh_CN
xtab -proj=<%.NewProjName%> -pkg=<%.NewProjPkg%> -db
`
