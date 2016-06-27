// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package env

type xConfiguration struct {
	Environments xEnvironments `xml:"environments"`
}

type xEnvironments struct {
	Default     string         `xml:"default,attr"`
	Environment []xEnvironment `xml:"environment"`
}

type xEnvironment struct {
	Id         string      `xml:"id,attr"`
	Databases  []xDatabase `xml:"database"`
	Mails      []xMail     `xml:"mail"`
	Sessions   []xSession  `xml:"session"`
	Apis       []xApi      `xml:"api"`
	Assets     []xStatic   `xml:"asset"`
	Statics    []xStatic   `xml:"static"`
	Developers []xStatic   `xml:"developer"`
	Operations []xStatic   `xml:"operation"`
	Uploads    []xUpload   `xml:"upload"`
}

type xDatabase struct {
	Name           string `xml:"name,attr"`
	DriverName     string `xml:"driverName"`
	DataSourceName string `xml:"dataSourceName"`
	MaxIdleConns   int    `xml:"maxIdleConns"`
	MaxOpenConns   int    `xml:"maxOpenConns"`
}

type xMail struct {
	Name     string `xml:"name,attr"`
	Secret   string `xml:"secret"`
	Protocol string `xml:"protocol"`
	Username string `xml:"username"`
	Password string `xml:"password"`
	Host     string `xml:"host"`
	Port     string `xml:"port"`
}

type xSession struct {
	Name     string `xml:"name,attr"`
	Addr     string `xml:"addr"`
	Password string `xml:"password"`
}

type xApi struct {
	Name string `xml:"name,attr"`
	URL  string `xml:"url"`
}

type xStatic struct {
	Name string `xml:"name,attr"`
	Dir  string `xml:"dir"`
	URL  string `xml:"url"`
}

type xUpload struct {
	Name    string `xml:"name,attr"`
	Dir     string `xml:"dir"`
	URL     string `xml:"url"`
	MaxSize string `xml:"maxSize"`
}
