// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xhttp

import (
	"gopkg.in/goyy/goyy.v0/comm/profile"
	"gopkg.in/goyy/goyy.v0/comm/xtype"
	"gopkg.in/goyy/goyy.v0/web/session"
	"html/template"
)

var Conf = &conf{
	Addr:    ":9090",
	Actives: []string{profile.DEV},
	Static: &staticOptions{
		Enable:     false,
		Dir:        "static",
		Apis:       "/api",
		Assets:     "/static",
		Consumers:  "/assets",
		Operations: "/assets",
	},
	Upload: &uploadOptions{
		Dir:     "/assets/upload",
		MaxSize: 5242880,
	},
	Html: &htmlOptions{
		Enable: false,
	},
	Templates: &templateOptions{
		Dir:        "templates",
		Extensions: []string{"html"},
		Funcs:      []template.FuncMap{},
		Delims: templateDelims{
			Left:  "{{",
			Right: "}}",
		},
		Reloaded: true,
	},
	Session: &sessionOptions{
		Addr: ":6379",
		Options: &session.Options{
			Path:     "",
			Domain:   "",
			MaxAge:   30 * 60,
			Secure:   false,
			HttpOnly: true,
		},
	},
	Secures: &secureOptions{
		LoginUrl:   "/login",
		SuccessUrl: "/",
		Filters: []xtype.Map{
			{"/**", "anon"},
		},
	},
}

type conf struct {
	Addr      string           // the TCP network address
	Actives   []string         // Active profile
	Session   *sessionOptions  // the session TCP network address
	Static    *staticOptions   // Static resource options
	Html      *htmlOptions     // Html resource options
	Upload    *uploadOptions   // Upload options
	Templates *templateOptions // template options
	Secures   *secureOptions
}

type sessionOptions struct {
	*session.Options
	Addr string
}

type staticOptions struct {
	Enable     bool   // Whether service is enabled
	Dir        string // Static resource directory
	Apis       string // APIs URL
	Assets     string // Static resource URL
	Consumers  string // Consumer uploaded static resource URL
	Operations string // Operations uploaded static resource URL
}

type htmlOptions struct {
	Enable bool // Whether service is enabled
}

type uploadOptions struct {
	Dir     string // Upload directory
	MaxSize int    // Max upload size
}

type secureOptions struct {
	LoginUrl   string
	SuccessUrl string
	Filters    []xtype.Map
}

// Options is a struct for specifying configuration options for the html render
type templateOptions struct {
	// Directory to load templates. Default is "templates"
	Dir string
	// Extensions to parse template files from. Defaults to ["tmpl"]
	Extensions []string
	// Funcs is a slice of FuncMaps to apply to the template upon compilation.
	// This is useful for helper functions. Defaults to [].
	Funcs []template.FuncMap
	// Delims sets the action delimiters to the specified strings in the templateDelims struct.
	Delims templateDelims
	// Reloaded sets up the template for each reload
	Reloaded bool
}

// templateDelims represents a set of Left and Right delimiters for HTML template rendering
type templateDelims struct {
	// Left delimiter, defaults to {{
	Left string
	// Right delimiter, defaults to }}
	Right string
}
