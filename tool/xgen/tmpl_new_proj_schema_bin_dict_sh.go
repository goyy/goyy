// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplNewProjSchemaBinDictSh = `#!/bin/sh

echo [INFO] Create Dict Datas.

cd ..
export I18N_LOCALE=zh_CN
xtab -proj=<%.NewProjName%> -pkg=<%.NewProjPkg%> -dict
`
