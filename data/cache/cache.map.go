// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cache

import (
	"github.com/garyburd/redigo/redis"
)

func HSet(key, field, value string) error {
	err := send("HSET", key, field, value)
	if err != nil {
		logger.Error(err.Error())
	}
	return err
}

func HGet(key, field string) (string, error) {
	v, err := redis.String(do("HGET", key, field))
	if err != nil {
		logger.Error(err.Error())
	}
	return v, err
}

func HDelete(key, field string) error {
	err := send("HDEL", key, field)
	if err != nil {
		logger.Error(err.Error())
	}
	return err
}

func HExists(key, field string) bool {
	v, err := redis.Bool(do("HEXISTS", key, field))
	if err != nil {
		logger.Error(err.Error())
		return false
	}
	return v
}

func HLen(key string) (int, error) {
	v, err := redis.Int(do("HLEN", key))
	if err != nil {
		logger.Error(err.Error())
	}
	return v, err
}
