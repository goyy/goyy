// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cache

import (
	"github.com/garyburd/redigo/redis"
)

// Set set the value of the cache based on the key.
func Set(key, value string) error {
	err := send("SET", key, value)
	if err != nil {
		logger.Error(err.Error())
	}
	return err
}

// Get get the value of the cache based on the key.
func Get(key string) (string, error) {
	v, err := redis.String(do("GET", key))
	if err != nil {
		logger.Error(err.Error())
	}
	return v, err
}
