// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cache

import (
	"github.com/garyburd/redigo/redis"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"time"
)

func Init(conf Conf) {
	if pool == nil {
		pool = &redis.Pool{
			MaxIdle:     conf.MaxIdle,
			MaxActive:   conf.MaxActive,
			IdleTimeout: conf.IdleTimeout,
			Dial: func() (redis.Conn, error) {
				conn, err := redis.Dial("tcp", conf.Address)
				if err != nil {
					logger.Error(err.Error())
					return nil, err
				}
				if strings.IsNotBlank(conf.Password) {
					if _, err := conn.Do("AUTH", conf.Password); err != nil {
						conn.Close()
						logger.Error(err.Error())
						return nil, err
					}
				}
				return conn, err
			},
			TestOnBorrow: func(conn redis.Conn, t time.Time) error {
				_, err := conn.Do("PING")
				if err != nil {
					logger.Error(err.Error())
				}
				return err
			},
		}
	}
}
