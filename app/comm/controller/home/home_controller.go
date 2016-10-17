package home

import (
	"gopkg.in/goyy/goyy.v0/web/controller"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

func init() {
	xhttp.GET("/", ctl.home)
	xhttp.GET("/favicon.ico", ctl.favicon)
}

func (me *Controller) home(c xhttp.Context) {
	if !c.Session().IsLogin() {
		c.Redirect(xhttp.Conf.Secure.LoginUrl)
		return
	}
	c.Redirect(xhttp.Conf.Secure.SuccessUrl)
}

func (me *Controller) favicon(c xhttp.Context) {
	c.Redirect(xhttp.Conf.Asset.URL + "/favicon.ico")
}

type Controller struct {
	controller.Controller
}

var ctl = &Controller{}
