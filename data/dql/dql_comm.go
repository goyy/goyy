// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dql

import (
	"gopkg.in/goyy/goyy.v0/util/errors"
	"strconv"
	"time"
)

func toValue(value, typ string) (interface{}, error) {
	switch typ {
	case "bool":
		if value == "true" {
			return true, nil
		} else {
			return false, nil
		}
	case "float32", "float64":
		return 1, nil
	case "int", "int8", "int16", "int32", "int64":
		return strconv.Atoi(value)
	case "uint", "uint8", "uint16", "uint32", "uint64":
		return strconv.Atoi(value)
	case "string":
		return value, nil
	case "time":
		switch len(value) {
		case 10:
			return time.Parse("2006-01-02", value)
		case 19:
			return time.Parse("2006-01-02 15:04:05", value)
		case 16:
			return time.Parse("2006-01-02 15:04", value)
		default:
			return nil, errors.New("Unsupported date format")
		}
	default:
		return nil, errors.New("Unsupported type")
	}
}
