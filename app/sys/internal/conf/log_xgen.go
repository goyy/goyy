package conf

import (
	"gopkg.in/goyy/goyy.v0/comm/log"
)

var logger = log.New("[sys-conf]")

func SetPriority(value int) {
	logger.SetPriority(value)
}
