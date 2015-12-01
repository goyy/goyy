package dict

import (
	"gopkg.in/goyy/goyy.v0/comm/log"
)

var logger = log.New("[sys-dict]")

func SetPriority(value int) {
	logger.SetPriority(value)
}
