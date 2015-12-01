package conf

import (
	"gopkg.in/goyy/goyy.v0/comm/profile"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

func init() {
	xhttp.Conf.Html.Enable = false
	xhttp.Conf.Template.Enable = true
	xhttp.Conf.Template.Funcs = xhttp.Conf.Template.Funcs
	if profile.Accepts(profile.DEV) {
		xhttp.Conf.Html.Reloaded = true
		xhttp.Conf.Template.Reloaded = true
	} else {
		xhttp.Conf.Html.Reloaded = false
		xhttp.Conf.Template.Reloaded = false
	}
}
