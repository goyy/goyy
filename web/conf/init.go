// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package conf

import _ "gopkg.in/goyy/goyy.v0/comm/profile/settings"
import _ "gopkg.in/goyy/goyy.v0/comm/log/settings"

import (
	"gopkg.in/goyy/goyy.v0/comm/env"
	"gopkg.in/goyy/goyy.v0/comm/log"
	"gopkg.in/goyy/goyy.v0/comm/xtype"
	"gopkg.in/goyy/goyy.v0/data/service"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"gopkg.in/goyy/goyy.v0/util/templates"
)

func init() {
	if v, err := env.Settings(); err == nil {
		initDB(v.Name)
		initApi(v.Name)
		initAsset(v.Name)
		initExport(v.Name)
		initSession(v.Name)
		initTemplate(v.Name)
		initIllegal(v.Name)
		initSecure(v.Name)
	} else {
		log.Println(err.Error())
	}
}

func initDB(envName string) {
	service.DB = service.NewDB(envName)
}

func initApi(envName string) {
	if v, err := env.Api(envName); err == nil {
		Conf.Api.URL = v.URL
	} else {
		log.Println(err.Error())
	}
}

func initAsset(envName string) {
	if v, err := env.Asset(envName); err == nil {
		Conf.Asset.Enable = v.Enable
		Conf.Asset.Dir = v.Dir
		Conf.Asset.URL = v.URL
	} else {
		log.Println(err.Error())
	}

	if v, err := env.Static(envName); err == nil {
		Conf.Static.Enable = v.Enable
		Conf.Static.Dir = v.Dir
		Conf.Static.URL = v.URL
	} else {
		log.Println(err.Error())
	}

	if v, err := env.Developer(envName); err == nil {
		Conf.Developer.Enable = v.Enable
		Conf.Developer.Dir = v.Dir
		Conf.Developer.URL = v.URL
	} else {
		log.Println(err.Error())
	}

	if v, err := env.Operation(envName); err == nil {
		Conf.Operation.Enable = v.Enable
		Conf.Operation.Dir = v.Dir
		Conf.Operation.URL = v.URL
	} else {
		log.Println(err.Error())
	}

	if v, err := env.Upload(envName); err == nil {
		Conf.Upload.Enable = v.Enable
		Conf.Upload.Dir = v.Dir
		Conf.Upload.URL = v.URL
	} else {
		log.Println(err.Error())
	}
}

func initExport(envName string) {
	if v, err := env.Export(envName); err == nil {
		Conf.Export.Dir = v.Dir
	} else {
		log.Println(err.Error())
	}
}

func initSession(envName string) {
	if v, err := env.Session(envName); err == nil {
		Conf.Session.Addr = v.Addr
		Conf.Session.Password = v.Password
	} else {
		log.Println(err.Error())
	}
}

func initTemplate(envName string) {
	if v, err := env.Html(envName); err == nil {
		Conf.Html.Enable = v.Enable
		Conf.Html.Reloaded = v.Reloaded
	} else {
		log.Println(err.Error())
	}

	if v, err := env.Template(envName); err == nil {
		Conf.Template.Enable = v.Enable
		Conf.Template.Reloaded = v.Reloaded
		if v.Enable {
			templates.GetApis = func() string { return Conf.Api.URL }
			templates.GetAssets = func() string { return Conf.Asset.URL }
			templates.GetDevelopers = func() string { return Conf.Developer.URL }
			templates.GetOperations = func() string { return Conf.Operation.URL }
			templates.GetStatics = func() string { return Conf.Static.URL }
			templates.GetUploads = func() string { return Conf.Upload.URL }
		}
	} else {
		log.Println(err.Error())
	}
}

func initIllegal(envName string) {
	if v, err := env.Illegal(envName); err == nil {
		Conf.Illegal.Enable = v.Enable
		if v.Enable {
			Conf.Illegal.Excludes = []string{v.Excludes}
			Conf.Illegal.Values = []string{v.Values}
		}
	} else {
		log.Println(err.Error())
	}
}

func initSecure(envName string) {
	if v, err := env.Secure(envName); err == nil {
		Conf.Secure.Enable = v.Enable
		if v.Enable {
			if strings.IsNotBlank(v.LoginUrl) {
				Conf.Secure.LoginUrl = v.LoginUrl
			}
			if strings.IsNotBlank(v.ForbidUrl) {
				Conf.Secure.ForbidUrl = v.ForbidUrl
			}
			if strings.IsNotBlank(v.SuccessUrl) {
				Conf.Secure.SuccessUrl = v.SuccessUrl
			}
			l := len(v.Filters.InterceptUrl)
			if v.Filters.InterceptUrl != nil && l > 0 {
				filters := make([]xtype.Map, l)
				for i, v := range v.Filters.InterceptUrl {
					m := xtype.Map{
						Key:   v.Pattern,
						Value: v.Access,
					}
					filters[i] = m
				}
			}
		}
	} else {
		log.Println(err.Error())
	}
}
