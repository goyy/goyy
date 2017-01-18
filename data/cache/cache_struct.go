// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cache

import (
	"encoding/json"

	"github.com/garyburd/redigo/redis"
)

// SSet set the value of the struct type according to key.
func SSet(key string, value interface{}) error {
	b, err := json.Marshal(value)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	err = send("SET", key, b)
	if err != nil {
		logger.Error(err.Error())
	}
	return err
}

// SGet get the value of the struct type according to key.
func SGet(key string, out interface{}) error {
	v, err := redis.Bytes(do("GET", key))
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	err = json.Unmarshal(v, out)
	if err != nil {
		logger.Error(err.Error())
	}
	return err
}
