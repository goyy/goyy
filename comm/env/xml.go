// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package env

// XMLConfiguration configuration.
type XMLConfiguration struct {
	Settings     XMLSettings     `xml:"settings"`
	Environments XMLEnvironments `xml:"environments"`
}

// XMLEnvironments environments.
type XMLEnvironments struct {
	Default     string           `xml:"default,attr"`
	Environment []XMLEnvironment `xml:"environment"`
}

// XMLSettings settings.
type XMLSettings struct {
	Name    string     `xml:"name,attr"`
	Profile XMLProfile `xml:"profile"`
}

// XMLProfile profile.
type XMLProfile struct {
	Default string `xml:"default,attr"`
	Actives string `xml:"actives,attr"`
}

// XMLEnvironment environment.
type XMLEnvironment struct {
	ID         string        `xml:"id,attr"`
	Databases  []XMLDatabase `xml:"database"`
	Mails      []XMLMail     `xml:"mail"`
	Sessions   []XMLSession  `xml:"session"`
	APIs       []XMLAPI      `xml:"api"`
	Assets     []XMLStatic   `xml:"asset"`
	Statics    []XMLStatic   `xml:"static"`
	Developers []XMLStatic   `xml:"developer"`
	Operations []XMLStatic   `xml:"operation"`
	Uploads    []XMLUpload   `xml:"upload"`
	Exports    []XMLExport   `xml:"export"`
	Htmls      []XMLTemplate `xml:"html"`
	Templates  []XMLTemplate `xml:"template"`
	Illegals   []XMLIllegal  `xml:"illegal"`
	Logs       []XMLLog      `xml:"log"`
	Secures    []XMLSecure   `xml:"secure"`
}

// XMLDatabase database.
type XMLDatabase struct {
	Name           string `xml:"name,attr"`
	DriverName     string `xml:"driverName"`
	DataSourceName string `xml:"dataSourceName"`
	MaxIdleConns   int    `xml:"maxIdleConns"`
	MaxOpenConns   int    `xml:"maxOpenConns"`
}

// XMLMail mail.
type XMLMail struct {
	Name     string `xml:"name,attr"`
	Secret   string `xml:"secret"`
	Protocol string `xml:"protocol"`
	Username string `xml:"username"`
	Password string `xml:"password"`
	Host     string `xml:"host"`
	Port     string `xml:"port"`
}

// XMLSession session.
type XMLSession struct {
	Name     string `xml:"name,attr"`
	Addr     string `xml:"addr"`
	Password string `xml:"password"`
}

// XMLAPI api.
type XMLAPI struct {
	Name string `xml:"name,attr"`
	URL  string `xml:"url"`
}

// XMLStatic static.
type XMLStatic struct {
	Name   string `xml:"name,attr"`
	Enable bool   `xml:"enable"`
	Dir    string `xml:"dir"`
	URL    string `xml:"url"`
}

// XMLUpload upload.
type XMLUpload struct {
	Name    string `xml:"name,attr"`
	Enable  bool   `xml:"enable"`
	Dir     string `xml:"dir"`
	URL     string `xml:"url"`
	MaxSize string `xml:"maxSize"`
}

// XMLExport export.
type XMLExport struct {
	Name string `xml:"name,attr"`
	Dir  string `xml:"dir"`
}

// XMLTemplate template.
type XMLTemplate struct {
	Name     string `xml:"name,attr"`
	Enable   bool   `xml:"enable"`
	Reloaded bool   `xml:"reloaded"`
}

// XMLIllegal illegal.
type XMLIllegal struct {
	Name     string `xml:"name,attr"`
	Enable   bool   `xml:"enable"`
	Excludes string `xml:"excludes"`
	Values   string `xml:"values"`
}

// XMLLog log.
type XMLLog struct {
	Name     string `xml:"name,attr"`
	Priority int    `xml:"priority"`
	Layout   int    `xml:"layout"`
	Output   int    `xml:"output"`
	Dir      string `xml:"dir"`
}

// XMLSecure secure.
type XMLSecure struct {
	Name       string     `xml:"name,attr"`
	Enable     bool       `xml:"enable"`
	LoginURL   string     `xml:"login-url"`
	ForbidURL  string     `xml:"forbid-url"`
	SuccessURL string     `xml:"success-url"`
	Filters    XMLFilters `xml:"filters"`
}

// XMLFilters filters.
type XMLFilters struct {
	InterceptURL []XMLInterceptURL `xml:"intercept-url"`
}

// XMLInterceptURL interceptURL.
type XMLInterceptURL struct {
	Pattern string `xml:"pattern,attr"`
	Access  string `xml:"access,attr"`
}
