package conf

import (
	"gopkg.in/goyy/goyy.v0/comm/env"
	"gopkg.in/goyy/goyy.v0/comm/log"
	"gopkg.in/goyy/goyy.v0/comm/profile"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

func init() {
	if profile.Accepts(profile.DEV) {
		xhttp.Conf.Asset.Enable = true
	} else {
		xhttp.Conf.Asset.Enable = false
	}
	if v, err := env.Asset("goyy"); err == nil {
		xhttp.Conf.Asset.Dir = v.Dir
		xhttp.Conf.Asset.URL = v.URL
	} else {
		log.Println(err.Error())
	}

	if profile.Accepts(profile.DEV) {
		xhttp.Conf.Static.Enable = true
	} else {
		xhttp.Conf.Static.Enable = false
	}
	if v, err := env.Static("goyy"); err == nil {
		xhttp.Conf.Static.Dir = v.Dir
		xhttp.Conf.Static.URL = v.URL
	} else {
		log.Println(err.Error())
	}

	if profile.Accepts(profile.DEV) {
		xhttp.Conf.Developer.Enable = true
	} else {
		xhttp.Conf.Developer.Enable = false
	}
	if v, err := env.Developer("goyy"); err == nil {
		xhttp.Conf.Developer.Dir = v.Dir
		xhttp.Conf.Developer.URL = v.URL
	} else {
		log.Println(err.Error())
	}

	if profile.Accepts(profile.DEV) {
		xhttp.Conf.Operation.Enable = true
	} else {
		xhttp.Conf.Operation.Enable = false
	}
	if v, err := env.Operation("goyy"); err == nil {
		xhttp.Conf.Operation.Dir = v.Dir
		xhttp.Conf.Operation.URL = v.URL
	} else {
		log.Println(err.Error())
	}

	if profile.Accepts(profile.DEV) {
		xhttp.Conf.Upload.Enable = true
	} else {
		xhttp.Conf.Upload.Enable = false
	}
	if v, err := env.Upload("goyy"); err == nil {
		xhttp.Conf.Upload.Dir = v.Dir
		xhttp.Conf.Upload.URL = v.URL
	} else {
		log.Println(err.Error())
	}
}
