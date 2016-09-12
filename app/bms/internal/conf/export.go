package conf

import (
	"gopkg.in/goyy/goyy.v0/comm/env"
	"gopkg.in/goyy/goyy.v0/comm/log"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

func init() {
	if v, err := env.Developer("goyy"); err == nil {
		xhttp.Conf.Export.Dir = v.Dir + "/export/excel"
	} else {
		log.Println(err.Error())
	}
}
