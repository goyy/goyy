package cache

import (
	"gopkg.in/goyy/goyy.v0/comm/log"
)

var logger = log.New("[sys-cache]")

func SetPriority(value int) {
	logger.SetPriority(value)
}
