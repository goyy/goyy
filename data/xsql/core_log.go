// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xsql

import (
	"bytes"
	"fmt"
	"time"

	"gopkg.in/goyy/goyy.v0/comm/log"
	"gopkg.in/goyy/goyy.v0/comm/profile"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

var isResetPriority bool

func init() {
	logger.Settings = func() {
		if !isResetPriority {
			isResetPriority = true
			if profile.Accepts(profile.PROD) {
				logger.SetPriority(log.Perror)
			} else {
				logger.SetPriority(log.Pdebug)
			}
		}
	}
}

// SetPriority set the priority of the logger.
func SetPriority(value int) {
	isResetPriority = true
	logger.SetPriority(value)
}

var logger = log.New("[xsql]")

func debugLog(started time.Time, sql string, args ...interface{}) {
	t := time.Now().Sub(started)
	logger.Debugf("%s [%s] (%s)", clearSpace(sql), argsString(args...), t)
}

func errorLog(started time.Time, sql string, args ...interface{}) {
	t := time.Now().Sub(started)
	logger.Errorf("%s [%s] (%s)", clearSpace(sql), argsString(args...), t)
}

func isDebug() bool {
	return log.Pdebug <= logger.Priority()
}

func isError() bool {
	return log.Perror <= logger.Priority()
}

func argsString(args ...interface{}) string {
	var b bytes.Buffer
	for i, arg := range args {
		if i > 0 {
			b.WriteString(" ")
		}
		v := fmt.Sprintf("%d:%v", i, arg)
		b.WriteString(v)
	}
	return b.String()
}

func clearSpace(sql string) string {
	sql = strings.Replace(sql, "\n", "", -1)
	return sql
}
