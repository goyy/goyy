package menu

import (
	"gopkg.in/goyy/goyy.v0/comm/log"
)

var logger = log.New("[sys-menu]")

func SetPriority(value int) {
	logger.SetPriority(value)
}
