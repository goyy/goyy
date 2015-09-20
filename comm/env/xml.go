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
	Id        string      `xml:"id,attr"`
	Databases []xDatabase `xml:"database"`
	Mails     []xMail     `xml:"mail"`
}

type xDatabase struct {
	Name           string `xml:"name,attr"`
	DriverName     string `xml:"driverName"`
	DataSourceName string `xml:"dataSourceName"`
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
