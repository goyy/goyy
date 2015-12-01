package area

import (
	"gopkg.in/goyy/goyy.v0/comm/log"
)

var logger = log.New("[sys-area]")

func SetPriority(value int) {
	logger.SetPriority(value)
}
