// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplEntityBat = `@echo off
echo [INFO] run go generate.

cd %~dp0
:set I18N_LOCALE=en_US
call go generate
pause`
