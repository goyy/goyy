package conf

import (
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

func init() {
	xhttp.Conf.Err.Err403 = "/err/403.html"
	xhttp.Conf.Err.Err404 = "/err/404.html"
	xhttp.Conf.Err.Err500 = "/err/500.html"
}
