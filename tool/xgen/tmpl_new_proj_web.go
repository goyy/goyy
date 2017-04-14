// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplNewProjWeb = `package main

import _ "github.com/go-sql-driver/mysql"
import _ "gopkg.in/goyy/goyy.v0/web/conf"
import _ "gopkg.in/goyy/goyy.v0/app/sys/api/dict"

import (
	_ "<%.NewProjPkg%>/<%.NewProjName%>-web/internal"
	_ "gopkg.in/goyy/goyy.v0/app/comm/controller/captcha"
	_ "gopkg.in/goyy/goyy.v0/app/comm/controller/home"
	_ "gopkg.in/goyy/goyy.v0/app/comm/controller/login"
	_ "gopkg.in/goyy/goyy.v0/app/sys"
	"gopkg.in/goyy/goyy.v0/comm/log"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

func main() {
	xhttp.Use(xhttp.Recovery())
	err := xhttp.Run()
	if err != nil {
		log.Error(err.Error())
	}
}
`
