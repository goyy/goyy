package order

import (
	"gopkg.in/goyy/goyy.v0/comm/log"
)

var logger = log.New("[example-order]")

func SetPriority(value int) {
	logger.SetPriority(value)
}
