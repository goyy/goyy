package wxalogin

import (
	"gopkg.in/goyy/goyy.v0/comm/log"
)

var logger = log.New("[comm-controller-wxalogin]")

func SetPriority(value int) {
	logger.SetPriority(value)
}
