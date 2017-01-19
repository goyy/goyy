// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package settings

import (
	"strings"

	"gopkg.in/goyy/goyy.v0/comm/env"
	"gopkg.in/goyy/goyy.v0/comm/profile"
)

func init() {
	if v, err := env.ParseSettings(); err == nil {
		profile.SetDefault(v.Profile.Default)
		actives := strings.Split(v.Profile.Actives, ",")
		profile.SetActives(actives...)
	}
}
