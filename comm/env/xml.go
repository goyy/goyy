// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package env

// Configuration configuration.
type Configuration struct {
	Settings     Settings     `xml:"settings"`
	Environments Environments `xml:"environments"`
}

// Environments environments.
type Environments struct {
	Default     string        `xml:"default,attr"`
	Environment []Environment `xml:"environment"`
}

// Settings settings.
type Settings struct {
	Name    string  `xml:"name,attr"`
	Profile Profile `xml:"profile"`
}

// Profile profile.
type Profile struct {
	Default string `xml:"default,attr"`
	Actives string `xml:"actives,attr"`
}

// Environment environment.
type Environment struct {
	Id         string      `xml:"id,attr"`
	Databases  []Database  `xml:"database"`
	Mails      []Mail      `xml:"mail"`
	Sessions   []Session   `xml:"session"`
	Apis       []Api       `xml:"api"`
	Assets     []Static    `xml:"asset"`
	Statics    []Static    `xml:"static"`
	Developers []Static    `xml:"developer"`
	Operations []Static    `xml:"operation"`
	Uploads    []Upload    `xml:"upload"`
	Exports    []Export    `xml:"export"`
	Htmls      []Template  `xml:"html"`
	Templates  []Template  `xml:"template"`
	Sensitives []Sensitive `xml:"sensitive"`
	Logs       []Log       `xml:"log"`
	Secures    []Secure    `xml:"secure"`
}

// Database database.
type Database struct {
	Name           string `xml:"name,attr"`
	DriverName     string `xml:"driverName"`
	DataSourceName string `xml:"dataSourceName"`
	MaxIdleConns   int    `xml:"maxIdleConns"`
	MaxOpenConns   int    `xml:"maxOpenConns"`
}

// Mail mail.
type Mail struct {
	Name     string `xml:"name,attr"`
	Secret   string `xml:"secret"`
	Protocol string `xml:"protocol"`
	Username string `xml:"username"`
	Password string `xml:"password"`
	Host     string `xml:"host"`
	Port     string `xml:"port"`
}

// Session session.
type Session struct {
	Name     string `xml:"name,attr"`
	Addr     string `xml:"addr"`
	Password string `xml:"password"`
}

// Api api.
type Api struct {
	Name string `xml:"name,attr"`
	URL  string `xml:"url"`
}

// Static static.
type Static struct {
	Name   string `xml:"name,attr"`
	Enable bool   `xml:"enable"`
	Dir    string `xml:"dir"`
	URL    string `xml:"url"`
}

// Upload upload.
type Upload struct {
	Name    string `xml:"name,attr"`
	Enable  bool   `xml:"enable"`
	Dir     string `xml:"dir"`
	URL     string `xml:"url"`
	MaxSize string `xml:"maxSize"`
}

// Export export.
type Export struct {
	Name string `xml:"name,attr"`
	Dir  string `xml:"dir"`
}

// Template template.
type Template struct {
	Name     string `xml:"name,attr"`
	Enable   bool   `xml:"enable"`
	Reloaded bool   `xml:"reloaded"`
}

// Sensitive sensitive word.
type Sensitive struct {
	Name     string `xml:"name,attr"`
	Enable   bool   `xml:"enable"`
	Excludes string `xml:"excludes"`
	Values   string `xml:"values"`
}

// Log log.
type Log struct {
	Name     string `xml:"name,attr"`
	Priority int    `xml:"priority"`
	Layout   int    `xml:"layout"`
	Output   int    `xml:"output"`
	Dir      string `xml:"dir"`
}

// Secure secure.
type Secure struct {
	Name       string  `xml:"name,attr"`
	Enable     bool    `xml:"enable"`
	LoginUrl   string  `xml:"login-url"`
	ForbidUrl  string  `xml:"forbid-url"`
	SuccessUrl string  `xml:"success-url"`
	Filters    Filters `xml:"filters"`
}

// Filters filters.
type Filters struct {
	InterceptUrl []InterceptUrl `xml:"intercept-url"`
}

// InterceptUrl interceptUrl.
type InterceptUrl struct {
	Pattern string `xml:"pattern,attr"`
	Access  string `xml:"access,attr"`
}
