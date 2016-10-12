// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package env_test

import (
	"testing"

	"gopkg.in/goyy/goyy.v0/comm/env"
)

func TestApi(t *testing.T) {
	in := "env"
	out1 := "env"
	out2 := "/apis"
	v, _ := env.Api(in)
	if v.Name != out1 || v.URL != out2 {
		format := "env.Api(%s) = %s, %s; want %s, %s"
		t.Errorf(format, in, v.Name, v.URL, out1, out2)
	}
}

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

func TestExport(t *testing.T) {
	in := "env"
	out1 := "env"
	out2 := "/assets/dev/export"
	v, _ := env.Export(in)
	if v.Name != out1 || v.Dir != out2 {
		format := "env.Export(%s) = %s, %s; want %s, %s"
		t.Errorf(format, in, v.Name, v.Dir, out1, out2)
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

func TestStatic(t *testing.T) {
	in := "env"
	out1 := "env"
	out2 := "static"
	out3 := "/statics"
	v, _ := env.Static(in)
	if v.Name != out1 || v.Dir != out2 || v.URL != out3 {
		format := "env.Static(%s) = %s, %s, %s; want %s, %s, %s"
		t.Errorf(format, in, v.Name, v.Dir, v.URL, out1, out2, out3)
	}
}

func TestUpload(t *testing.T) {
	in := "env"
	out1 := "env"
	out2 := "/assets/upls"
	out3 := "/upls"
	out4 := "5242880"
	v, _ := env.Upload(in)
	if v.Name != out1 || v.Dir != out2 || v.URL != out3 || v.MaxSize != out4 {
		format := "env.Upload(%s) = %s, %s, %s, %s; want %s, %s, %s, %s"
		t.Errorf(format, in, v.Name, v.Dir, v.URL, v.MaxSize, out1, out2, out3, out4)
	}
}
