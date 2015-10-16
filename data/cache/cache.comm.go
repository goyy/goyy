// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cache

import (
	"github.com/garyburd/redigo/redis"
)

func send(cmd string, args ...interface{}) error {
	conn := pool.Get()
	defer conn.Close()
	if err := conn.Err(); err != nil {
		logger.Error(err.Error())
		return err
	}
	err := conn.Send(cmd, args...)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return conn.Flush()
}

func do(cmd string, args ...interface{}) (interface{}, error) {
	conn := pool.Get()
	defer conn.Close()
	if err := conn.Err(); err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	v, err := conn.Do(cmd, args...)
	if err != nil {
		logger.Error(err.Error())
	}
	return v, err
}

func Delete(key string) error {
	err := send("DEL", key)
	if err != nil {
		logger.Error(err.Error())
	}
	return err
}

func Exists(key string) bool {
	v, err := redis.Bool(do("EXISTS", key))
	if err != nil {
		logger.Error(err.Error())
		return false
	}
	return v
}

func Expire(key string, second int) error {
	err := send("EXPIRE", key, second)
	if err != nil {
		logger.Error(err.Error())
	}
	return err
}
