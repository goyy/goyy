// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cache

import (
	"github.com/garyburd/redigo/redis"
	"gopkg.in/goyy/goyy.v0/comm/log"
)

var pool *redis.Pool

var logger = log.New("[cache]")
