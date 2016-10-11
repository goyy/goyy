package blacklist

import (
	"gopkg.in/goyy/goyy.v0/comm/log"
)

var logger = log.New("[sys-blacklist]")

func SetPriority(value int) {
	logger.SetPriority(value)
}
