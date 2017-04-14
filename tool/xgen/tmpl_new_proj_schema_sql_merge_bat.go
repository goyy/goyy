// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplNewProjSchemaSQLMergeBat = `@echo off
echo [INFO] Merged SQL Files.

cd %~dp0
set I18N_LOCALE=zh_CN
call xtab -regexp=^insert.[\S]+.sql$ -newfile=merge-file.sql -merge

pause
`
