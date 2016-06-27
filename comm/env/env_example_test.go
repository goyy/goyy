// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package env_test

import (
	"fmt"

	"gopkg.in/goyy/goyy.v0/comm/env"
)

func ExampleDatabase() {
	db, _ := env.Database("env")
	fmt.Println(db.DriverName)
	fmt.Println(db.DataSourceName)
	fmt.Println(db.MaxIdleConns)
	fmt.Println(db.MaxOpenConns)

	// Output:
	// mysql
	// root:root@/env_development?charset=utf8
	// 10
	// 100
}

func ExampleMail() {
	m, _ := env.Mail("env")
	fmt.Println(m.Secret)
	fmt.Println(m.Protocol)
	fmt.Println(m.Username)
	fmt.Println(m.Password)
	fmt.Println(m.Host)
	fmt.Println(m.Port)

	// Output:
	// 0ae36a2eha9p1e16
	// POP3
	// username@example.com
	// password
	// mail.example.com
	// 110
}

func ExampleSession() {
	s, _ := env.Session("env")
	fmt.Println(s.Addr)
	fmt.Println(s.Password)

	// Output:
	// :6379
	// 123456
}
