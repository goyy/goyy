package org

import (
	"gopkg.in/goyy/goyy.v0/comm/log"
)

var logger = log.New("[sys-org]")

func SetPriority(value int) {
	logger.SetPriority(value)
}
