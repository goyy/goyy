package user

import (
	"gopkg.in/goyy/goyy.v0/comm/log"
)

var logger = log.New("[sys-user]")

func SetPriority(value int) {
	logger.SetPriority(value)
}
