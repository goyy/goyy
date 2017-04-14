// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplNewProjSchemaBinMenuBat = `@echo off
echo [INFO] Create Menu Files.

cd %~dp0
cd ..
set I18N_LOCALE=zh_CN
call xtab -proj=<%.NewProjName%> -pkg=<%.NewProjPkg%> -menu

pause
`
