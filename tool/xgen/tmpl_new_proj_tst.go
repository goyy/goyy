// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplNewProjTst = `package tst

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "gopkg.in/goyy/goyy.v0/app/sys/api/dict"
	"gopkg.in/goyy/goyy.v0/comm/env"
	"gopkg.in/goyy/goyy.v0/comm/log"
	"gopkg.in/goyy/goyy.v0/comm/profile"
	"gopkg.in/goyy/goyy.v0/data/cache"
	"gopkg.in/goyy/goyy.v0/data/service"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

func init() {
	// profile
	profile.SetDefault(profile.TEST)

	// DataBase
	env.SetConf("<%.NewProjPath%>/<%.NewProjName%>-tst/conf")
	service.DB = service.NewDB("<%.NewProjName%>")

	// cache
	if v, err := env.Session("<%.NewProjName%>"); err == nil {
		xhttp.Conf.Session.Addr = v.Addr
		xhttp.Conf.Session.Password = v.Password
		cache.Init(cache.Conf{
			Address:     v.Addr,
			Password:    v.Password,
			MaxIdle:     80,
			MaxActive:   12000,
			IdleTimeout: 240 * time.Second,
		})
	} else {
		log.Println(err.Error())
	}
}
`
