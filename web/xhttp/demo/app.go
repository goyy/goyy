// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"gopkg.in/goyy/goyy.v0/comm/log"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
	"gopkg.in/goyy/goyy.v0/web/xhttp/demo/internal/controller"
)

func main() {
	xhttp.GET("/", func(ctx xhttp.Context) {
		ctx.Session().Set("user", "goyy")
		ctx.Text(xhttp.StatusOK, "index")
	})
	xhttp.GET("/json", controller.DemoCtl.JSON)
	xhttp.GET("/jsonp", controller.DemoCtl.JSONP)
	xhttp.GET("/xml", controller.DemoCtl.XML)
	xhttp.GET("/list", controller.DemoCtl.List)
	xhttp.GET("/form", controller.DemoCtl.Form)
	xhttp.Conf.Session.Addr = "10.100.130.250:6379"
	err := xhttp.Run()
	if err != nil {
		log.Error(err.Error())
	}
}
