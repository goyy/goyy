package conf

import (
	"gopkg.in/goyy/goyy.v0/comm/profile"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

func init() {
	xhttp.Conf.Html.Enable = true
	xhttp.Conf.Template.Enable = true
	if profile.Accepts(profile.DEV) {
		xhttp.Conf.Html.Reloaded = true
		xhttp.Conf.Template.Reloaded = true
	} else {
		xhttp.Conf.Html.Reloaded = false
		xhttp.Conf.Template.Reloaded = false
	}
}
