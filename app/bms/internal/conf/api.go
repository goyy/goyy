package conf

import (
	"gopkg.in/goyy/goyy.v0/comm/env"
	"gopkg.in/goyy/goyy.v0/comm/log"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

func init() {
	if v, err := env.Api("goyy"); err == nil {
		xhttp.Conf.Api.URL = v.URL
	} else {
		log.Println(err.Error())
	}
}
