// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cache

import (
	"github.com/garyburd/redigo/redis"
)

func HSet(key, field, value string) error {
	err := send("HSET", key, field, value)
	logging.Debugf("HSet %v %v %v %v", key, field, value, err)
	return err
}

func HGet(key, field string) (string, error) {
	v, err := redis.String(do("HGET", key, field))
	logging.Debugf("HGet %v %v %v", key, field, err)
	return v, err
}

func HDelete(key, field string) error {
	err := send("HDEL", key, field)
	logging.Debugf("HDelete %v %v %v", key, field, err)
	return err
}

func HExists(key, field string) bool {
	v, err := redis.Bool(do("HEXISTS", key, field))
	logging.Debugf("HExists %v %v %v", key, field, err)
	if err != nil {
		return false
	}
	return v
}

func HLen(key string) (int, error) {
	v, err := redis.Int(do("HLEN", key))
	logging.Debugf("HLen %v %v", key, err)
	return v, err
}
