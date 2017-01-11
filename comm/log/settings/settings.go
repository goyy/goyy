// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package settings

import (
	"gopkg.in/goyy/goyy.v0/comm/env"
	"gopkg.in/goyy/goyy.v0/comm/log"
)

func init() {
	if v, err := env.Settings(); err == nil {
		if l, err := env.Log(v.Name); err == nil {
			log.SetDefaultPriority(l.Priority)
			log.SetDefaultLayout(l.Layout)
			log.SetDefaultOutput(l.Output)
			log.SetDefaultDir(l.Dir)
		} else {
			log.Println(err.Error())
		}
	}
}
