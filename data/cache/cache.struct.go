// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cache

import (
	"encoding/json"
	"github.com/garyburd/redigo/redis"
)

func SSet(key string, value interface{}) error {
	b, err := json.Marshal(value)
	if err != nil {
		logging.Debugf("1:SSet %v %v %v", key, value, err)
		return err
	} else {
		logging.Debugf("2:SSet %v %v", key, value)
	}
	err = send("SET", key, b)
	logging.Debugf("3:SSet %v %v %v", key, value, err)
	return err
}

func SGet(key string, out interface{}) error {
	v, err := redis.Bytes(do("GET", key))
	logging.Debugf("1:SGet %v %v", key, err)
	if err != nil {
		return err
	}
	err = json.Unmarshal(v, out)
	logging.Debugf("2:SGet %v %v %v", key, v, err)
	return err
}
