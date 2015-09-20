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
		logging.Debugf("1:send %v %v %v", cmd, args, err)
		return err
	} else {
		logging.Debugf("2:send %v %v", cmd, args)
	}
	err := conn.Send(cmd, args...)
	logging.Debugf("3:send %v %v %v", cmd, args, err)
	if err != nil {
		return err
	}
	return conn.Flush()
}

func do(cmd string, args ...interface{}) (interface{}, error) {
	conn := pool.Get()
	defer conn.Close()
	if err := conn.Err(); err != nil {
		logging.Debugf("1:do %v %v %v", cmd, args, err)
		return nil, err
	} else {
		logging.Debugf("2:do %v %v", cmd, args)
	}
	v, err := conn.Do(cmd, args...)
	logging.Debugf("3:do %v %v %v", cmd, args, err)
	return v, err
}

func Delete(key string) error {
	err := send("DEL", key)
	logging.Debugf("Delete %v %v", key, err)
	return err
}

func Exists(key string) bool {
	v, err := redis.Bool(do("EXISTS", key))
	logging.Debugf("Exists %v %v", key, err)
	if err != nil {
		return false
	}
	return v
}

func Expire(key string, second int) error {
	err := send("EXPIRE", key, second)
	logging.Debugf("Expire %v %v %v", key, second, err)
	return err
}
