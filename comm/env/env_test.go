// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package env_test

import (
	"testing"

	"gopkg.in/goyy/goyy.v0/comm/env"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

func TestParseSettings(t *testing.T) {
	out1 := "env"
	out2 := "development"
	out3 := "comm"
	v, _ := env.ParseSettings()
	if v.Name != out1 || v.Profile.Default != out2 || v.Profile.Actives != out3 {
		format := "env.ParseSettings() = %s, %s, %s; want %s, %s, %s"
		t.Errorf(format, v.Name, v.Profile.Default, v.Profile.Actives, out1, out2, out3)
	}
}

func TestParseApi(t *testing.T) {
	in := "env"
	out1 := "env"
	out2 := "/apis"
	v, _ := env.ParseApi(in)
	if v.Name != out1 || v.URL != out2 {
		format := "env.ParseApi(%s) = %s, %s; want %s, %s"
		t.Errorf(format, in, v.Name, v.URL, out1, out2)
	}
}

func TestParseDatabase(t *testing.T) {
	in := "env"
	out1 := "mysql"
	out2 := "root:root@/env_development?charset=utf8"
	out3 := 10
	out4 := 100
	db, _ := env.ParseDatabase(in)
	if db.DriverName != out1 || db.DataSourceName != out2 || db.MaxIdleConns != out3 || db.MaxOpenConns != out4 {
		format := "env.ParseDatabase(%s) = %s, %s, %v, %v; want %s, %s, %v, %v"
		t.Errorf(format, in, db.DriverName, db.DataSourceName, db.MaxIdleConns, db.MaxOpenConns, out1, out2, out3, out4)
	}
}

func TestParseExport(t *testing.T) {
	in := "env"
	out1 := "env"
	out2 := "/assets/dev/export"
	v, _ := env.ParseExport(in)
	if v.Name != out1 || v.Dir != out2 {
		format := "env.ParseExport(%s) = %s, %s; want %s, %s"
		t.Errorf(format, in, v.Name, v.Dir, out1, out2)
	}
}

func TestParseMail(t *testing.T) {
	in := "env"
	out := "username@example.com"
	m, _ := env.ParseMail(in)
	if m.Username != out {
		format := "env.ParseMail(%s) = %s, _; want %s, _"
		t.Errorf(format, in, m.Username, out)
	}
}

func TestParseSession(t *testing.T) {
	in := "env"
	outAddr := ":6379"
	outPassword := "123456"
	s, _ := env.ParseSession(in)
	if s.Addr != outAddr {
		format := "env.ParseSession(%s) = %s, _; want %s, _"
		t.Errorf(format, in, s.Addr, outAddr)
	}
	if s.Password != outPassword {
		format := "env.ParseSession(%s) = %s, _; want %s, _"
		t.Errorf(format, in, s.Password, outPassword)
	}
}

func TestParseStatic(t *testing.T) {
	in := "env"
	out1 := "env"
	out2 := true
	out3 := "static"
	out4 := "/statics"
	v, _ := env.ParseStatic(in)
	if v.Name != out1 || v.Enable != out2 || v.Dir != out3 || v.URL != out4 {
		format := "env.ParseStatic(%s) = %s, %t, %s, %s; want %s, %t, %s, %s"
		t.Errorf(format, in, v.Name, v.Enable, v.Dir, v.URL, out1, out2, out3, out4)
	}
}

func TestParseUpload(t *testing.T) {
	in := "env"
	out1 := "env"
	out2 := true
	out3 := "/assets/upls"
	out4 := "/upls"
	out5 := "5242880"
	v, _ := env.ParseUpload(in)
	if v.Name != out1 || v.Enable != out2 || v.Dir != out3 || v.URL != out4 || v.MaxSize != out5 {
		format := "env.ParseUpload(%s) = %s, %t, %s, %s, %s; want %s, %t, %s, %s, %s"
		t.Errorf(format, in, v.Name, v.Enable, v.Dir, v.URL, v.MaxSize, out1, out2, out3, out4, out5)
	}
}

func TestParseTemplate(t *testing.T) {
	in := "env"
	out1 := "env"
	out2 := true
	out3 := true
	v, _ := env.ParseTemplate(in)
	if v.Name != out1 || v.Enable != out2 || v.Reloaded != out3 {
		format := "env.ParseTemplate(%s) = %s, %t, %t; want %s, %t, %t"
		t.Errorf(format, in, v.Name, v.Enable, v.Reloaded, out1, out2, out3)
	}
}

func TestParseIllegal(t *testing.T) {
	in := "env"
	out1 := "env"
	out2 := true

	notExpected := -1
	exclude := "/login"
	value := "FileOutputStream,"

	v, _ := env.ParseIllegal(in)
	out3 := strings.Index(v.Excludes, exclude)
	out4 := strings.Index(v.Values, value)
	if v.Name != out1 || v.Enable != out2 || out3 == notExpected || out4 == notExpected {
		format := "env.ParseIllegal(%s) = %s, %t, %v, %v; want %s, %t, %s, %s"
		t.Errorf(format, in, v.Name, v.Enable, out3, out4, out1, out2, ">0", ">0")
	}
}

func TestParseLog(t *testing.T) {
	in := "env"
	out1 := "env"
	out2 := 2
	out3 := 75
	out4 := 1
	out5 := "logs"

	v, _ := env.ParseLog(in)
	if v.Name != out1 || v.Priority != out2 || v.Layout != out3 || v.Output != out4 || v.Dir != out5 {
		format := "env.ParseLog(%s) = %s, %v, %v, %v, %s; want %s, %v, %v, %v, %s"
		t.Errorf(format, in, v.Name, v.Priority, v.Layout, v.Output, v.Dir, out1, out2, out3, out4, out5)
	}
}

func TestParseSecure(t *testing.T) {
	in := "env"
	out1 := "env"
	out2 := true
	out3 := "/login.html"
	out4 := "/err/403.html"
	out5 := "/home.html"
	out6 := "/login"
	out7 := "anon"

	v, _ := env.ParseSecure(in)
	if v.Name != out1 || v.Enable != out2 || v.LoginUrl != out3 || v.ForbidUrl != out4 || v.SuccessUrl != out5 || v.Filters.InterceptUrl[3].Pattern != out6 || v.Filters.InterceptUrl[3].Access != out7 {
		format := "env.ParseSecure(%s) = %s, %t, %s, %s, %s, %s, %s; want %s, %t, %s, %s, %s, %s, %s"
		t.Errorf(format, in, v.Name, v.Enable, v.LoginUrl, v.ForbidUrl, v.SuccessUrl, v.Filters.InterceptUrl[3].Pattern, v.Filters.InterceptUrl[3].Access, out3, out4, out1, out2, out3, out5, out6, out7)
	}
}
