package home

import (
	"gopkg.in/goyy/goyy.v0/util/cookies"
	"gopkg.in/goyy/goyy.v0/util/files"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"gopkg.in/goyy/goyy.v0/web/controller"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

func init() {
	xhttp.GET("/", ctl.home)
	xhttp.GET("/favicon.ico", ctl.favicon)
}

func ver() string {
	var ver string = "ver=1"
	fver := xhttp.Conf.Html.Dir + "/version.html"
	if files.IsExist(fver) {
		if c, err := files.Read(fver); err == nil {
			ver = c
		}
	}
	if strings.Index(xhttp.Conf.Secure.SuccessUrl, "?") == -1 {
		return "?" + ver
	} else {
		return "&" + ver
	}
}

func (me *Controller) home(c xhttp.Context) {
	if !c.Session().IsLogin() {
		c.Redirect(xhttp.Conf.Secure.LoginUrl)
		return
	}
	// Set cookie,
	// for the computer version and touch screen version of the switch
	m := c.Param("mode")
	switch m {
	case "pc", "webapp":
		cookies.SetValue(c.ResponseWriter(), "mode", m)
	}
	c.Redirect(xhttp.Conf.Secure.SuccessUrl + ver())
}

func (me *Controller) favicon(c xhttp.Context) {
	c.Redirect(xhttp.Conf.Asset.URL + "/favicon.ico")
}

type Controller struct {
	controller.Controller
}

var ctl = &Controller{}
