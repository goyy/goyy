package conf

import (
	"gopkg.in/goyy/goyy.v0/comm/profile"
	"gopkg.in/goyy/goyy.v0/util/templates"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

func init() {
	xhttp.Conf.Html.Enable = true
	xhttp.Conf.Template.Enable = true
	xhttp.Conf.Template.Funcs = append(xhttp.Conf.Template.Funcs, templates.Html.FuncMap)
	if profile.Accepts(profile.DEV) {
		xhttp.Conf.Html.Reloaded = true
		xhttp.Conf.Template.Reloaded = true
	} else {
		xhttp.Conf.Html.Reloaded = false
		xhttp.Conf.Template.Reloaded = false
	}
}
