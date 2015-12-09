package role

import (
	"gopkg.in/goyy/goyy.v0/comm/log"
)

var logger = log.New("[sys-role]")

func SetPriority(value int) {
	logger.SetPriority(value)
}
