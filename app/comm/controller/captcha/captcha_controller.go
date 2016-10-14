package captcha

import (
	"gopkg.in/goyy/goyy.v0/comm/captcha"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

func init() {
	xhttp.GET("/captcha/build", ctl.Build)
	xhttp.GET("/captcha/verify", ctl.Verify)
}

var ctl = &captcha.Controller{}
