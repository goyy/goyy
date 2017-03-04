package home

import (
	"gopkg.in/goyy/goyy.v0/util/cookies"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"gopkg.in/goyy/goyy.v0/web/controller"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

func init() {
	xhttp.GET("/", ctl.home)
	xhttp.GET("/favicon.ico", ctl.favicon)
}

func ver(url string) string {
	if strings.Index(url, "?") == -1 {
		return url + "?" + xhttp.Conf.Asset.Ver
	} else {
		return url + "&" + xhttp.Conf.Asset.Ver
	}
}

func (me *Controller) home(c xhttp.Context) {
	if !c.Session().IsLogin() {
		c.Redirect(ver(xhttp.Conf.Secure.LoginUrl))
		return
	}
	// Set cookie,
	// for the computer version and touch screen version of the switch
	m := c.Param("mode")
	switch m {
	case "pc", "webapp":
		cookies.SetValue(c.ResponseWriter(), "mode", m)
	}
	c.Redirect(ver(xhttp.Conf.Secure.SuccessUrl))
}

func (me *Controller) favicon(c xhttp.Context) {
	c.Redirect(xhttp.Conf.Static.URL + "/favicon.ico")
}

type Controller struct {
	controller.Controller
}

var ctl = &Controller{}
