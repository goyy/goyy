package conf

import (
	"gopkg.in/goyy/goyy.v0/comm/env"
	"gopkg.in/goyy/goyy.v0/comm/log"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

func init() {
	if v, err := env.Session("goyy"); err == nil {
		xhttp.Conf.Session.Addr = v.Addr
	} else {
		log.Println(err.Error())
	}
}
