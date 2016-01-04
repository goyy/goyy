package home

import (
	"gopkg.in/goyy/goyy.v0/web/controller"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

func (me *Controller) home(c xhttp.Context) {
	c.Redirect("home.html")
}

func (me *Controller) favicon(c xhttp.Context) {
	c.Redirect(xhttp.Conf.Asset.URL + "/favicon.ico")
}

func init() {
	xhttp.GET("/", ctl.home)
	xhttp.GET("/favicon.ico", ctl.favicon)
}

type Controller struct {
	controller.ClientTreeController
}

var ctl = &Controller{}
