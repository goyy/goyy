// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xhttp

import (
	"gopkg.in/goyy/goyy.v0/comm/profile"
	"gopkg.in/goyy/goyy.v0/web/conf"
)

func setConfProfile() {
	if profile.Accepts(profile.PROD) {
		Conf.Profile = profile.PROD
		return
	}
	if profile.Accepts(profile.TEST) {
		Conf.Profile = profile.TEST
		return
	}
}

var Conf = conf.Conf

func init() {
	RegisterPreRun(setConfProfile)
}
