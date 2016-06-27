// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package env_test

import (
	"testing"

	"gopkg.in/goyy/goyy.v0/comm/env"
)

func TestDatabase(t *testing.T) {
	in := "env"
	out1 := "mysql"
	out2 := "root:root@/env_development?charset=utf8"
	out3 := 10
	out4 := 100
	db, _ := env.Database(in)
	if db.DriverName != out1 || db.DataSourceName != out2 || db.MaxIdleConns != out3 || db.MaxOpenConns != out4 {
		format := "env.Database(%s) = %s, %s, %v, %v; want %s, %s, %v, %v"
		t.Errorf(format, in, db.DriverName, db.DataSourceName, db.MaxIdleConns, db.MaxOpenConns, out1, out2, out3, out4)
	}
}

func TestMail(t *testing.T) {
	in := "env"
	out := "username@example.com"
	m, _ := env.Mail(in)
	if m.Username != out {
		format := "env.Mail(%s) = %s, _; want %s, _"
		t.Errorf(format, in, m.Username, out)
	}
}

func TestSession(t *testing.T) {
	in := "env"
	outAddr := ":6379"
	outPassword := "123456"
	s, _ := env.Session(in)
	if s.Addr != outAddr {
		format := "env.Session(%s) = %s, _; want %s, _"
		t.Errorf(format, in, s.Addr, outAddr)
	}
	if s.Password != outPassword {
		format := "env.Session(%s) = %s, _; want %s, _"
		t.Errorf(format, in, s.Password, outPassword)
	}
}
