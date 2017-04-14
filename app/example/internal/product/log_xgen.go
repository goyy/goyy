package product

import (
	"gopkg.in/goyy/goyy.v0/comm/log"
)

var logger = log.New("[example-product]")

func SetPriority(value int) {
	logger.SetPriority(value)
}
