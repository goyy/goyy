# comm-env [![GoDoc](http://godoc.org/gopkg.in/goyy/goyy.v0?status.png)](http://godoc.org/gopkg.in/goyy/goyy.v0/comm/env)
environment library for Go

# Installation
`go get gopkg.in/goyy/goyy.v0/comm/env`

# Usage
	db, _ := env.Database("env")
	fmt.Println(db.DriverName)
	fmt.Println(db.DataSourceName)
	// Output:
	// mysql
	// root:root@/env_development?charset=utf8

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
