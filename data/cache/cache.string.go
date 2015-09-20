// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cache

import (
	"github.com/garyburd/redigo/redis"
)

func Set(key, value string) error {
	err := send("SET", key, value)
	logging.Debugf("Set %v %v %v", key, value, err)
	return err
}

func Get(key string) (string, error) {
	v, err := redis.String(do("GET", key))
	logging.Debugf("Get %v %v", key, err)
	return v, err
}
