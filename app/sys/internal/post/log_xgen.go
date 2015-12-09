package post

import (
	"gopkg.in/goyy/goyy.v0/comm/log"
)

var logger = log.New("[sys-post]")

func SetPriority(value int) {
	logger.SetPriority(value)
}
