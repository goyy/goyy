// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package conf

import (
	"gopkg.in/goyy/goyy.v0/comm/env"
	"gopkg.in/goyy/goyy.v0/comm/log"
	"gopkg.in/goyy/goyy.v0/comm/profile"
)

func Init(envName, defaultProfile string, activesProfile ...string) {
	initProfile(defaultProfile, activesProfile...)
	initLog()
	initApi(envName)
}

func initProfile(defaults string, actives ...string) {
	profile.SetDefault(defaults)
	profile.SetActives(actives...)
}

func initLog() {
	if profile.Accepts(profile.DEV) {
		log.SetDefaultOutput(log.Oconsole)
	} else {
		log.SetDefaultOutput(log.Odailyfile)
	}
}

func initApi(envName string) {
	if v, err := env.Api(envName); err == nil {
		Conf.Api.URL = v.URL
	} else {
		log.Println(err.Error())
	}
}
