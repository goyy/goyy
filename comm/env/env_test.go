// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package env_test

import (
	"testing"

	"gopkg.in/goyy/goyy.v0/comm/env"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

func TestSettings(t *testing.T) {
	out1 := "env"
	out2 := "development"
	out3 := "comm"
	v, _ := env.Settings()
	if v.Name != out1 || v.Profile.Default != out2 || v.Profile.Actives != out3 {
		format := "env.Settings() = %s, %s, %s; want %s, %s, %s"
		t.Errorf(format, v.Name, v.Profile.Default, v.Profile.Actives, out1, out2, out3)
	}
}

func TestAPI(t *testing.T) {
	in := "env"
	out1 := "env"
	out2 := "/apis"
	v, _ := env.API(in)
	if v.Name != out1 || v.URL != out2 {
		format := "env.API(%s) = %s, %s; want %s, %s"
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
	out2 := true
	out3 := "static"
	out4 := "/statics"
	v, _ := env.Static(in)
	if v.Name != out1 || v.Enable != out2 || v.Dir != out3 || v.URL != out4 {
		format := "env.Static(%s) = %s, %t, %s, %s; want %s, %t, %s, %s"
		t.Errorf(format, in, v.Name, v.Enable, v.Dir, v.URL, out1, out2, out3, out4)
	}
}

func TestUpload(t *testing.T) {
	in := "env"
	out1 := "env"
	out2 := true
	out3 := "/assets/upls"
	out4 := "/upls"
	out5 := "5242880"
	v, _ := env.Upload(in)
	if v.Name != out1 || v.Enable != out2 || v.Dir != out3 || v.URL != out4 || v.MaxSize != out5 {
		format := "env.Upload(%s) = %s, %t, %s, %s, %s; want %s, %t, %s, %s, %s"
		t.Errorf(format, in, v.Name, v.Enable, v.Dir, v.URL, v.MaxSize, out1, out2, out3, out4, out5)
	}
}

func TestTemplate(t *testing.T) {
	in := "env"
	out1 := "env"
	out2 := true
	out3 := true
	v, _ := env.Template(in)
	if v.Name != out1 || v.Enable != out2 || v.Reloaded != out3 {
		format := "env.Template(%s) = %s, %t, %t; want %s, %t, %t"
		t.Errorf(format, in, v.Name, v.Enable, v.Reloaded, out1, out2, out3)
	}
}

func TestIllegal(t *testing.T) {
	in := "env"
	out1 := "env"
	out2 := true

	notExpected := -1
	exclude := "/login"
	value := "FileOutputStream,"

	v, _ := env.Illegal(in)
	out3 := strings.Index(v.Excludes, exclude)
	out4 := strings.Index(v.Values, value)
	if v.Name != out1 || v.Enable != out2 || out3 == notExpected || out4 == notExpected {
		format := "env.Illegal(%s) = %s, %t, %v, %v; want %s, %t, %s, %s"
		t.Errorf(format, in, v.Name, v.Enable, out3, out4, out1, out2, ">0", ">0")
	}
}

func TestLog(t *testing.T) {
	in := "env"
	out1 := "env"
	out2 := 2
	out3 := 75
	out4 := 1
	out5 := "logs"

	v, _ := env.Log(in)
	if v.Name != out1 || v.Priority != out2 || v.Layout != out3 || v.Output != out4 || v.Dir != out5 {
		format := "env.Log(%s) = %s, %v, %v, %v, %s; want %s, %v, %v, %v, %s"
		t.Errorf(format, in, v.Name, v.Priority, v.Layout, v.Output, v.Dir, out1, out2, out3, out4, out5)
	}
}

func TestSecure(t *testing.T) {
	in := "env"
	out1 := "env"
	out2 := true
	out3 := "/login.html"
	out4 := "/err/403.html"
	out5 := "/home.html"
	out6 := "/login"
	out7 := "anon"

	v, _ := env.Secure(in)
	if v.Name != out1 || v.Enable != out2 || v.LoginURL != out3 || v.ForbidURL != out4 || v.SuccessURL != out5 || v.Filters.InterceptURL[3].Pattern != out6 || v.Filters.InterceptURL[3].Access != out7 {
		format := "env.Secure(%s) = %s, %t, %s, %s, %s, %s, %s; want %s, %t, %s, %s, %s, %s, %s"
		t.Errorf(format, in, v.Name, v.Enable, v.LoginURL, v.ForbidURL, v.SuccessURL, v.Filters.InterceptURL[3].Pattern, v.Filters.InterceptURL[3].Access, out3, out4, out1, out2, out3, out5, out6, out7)
	}
}
