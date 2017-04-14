// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package conf

import (
	"html/template"

	"gopkg.in/goyy/goyy.v0/comm/xtype"
	"gopkg.in/goyy/goyy.v0/util/envs"
	"gopkg.in/goyy/goyy.v0/util/templates"
	"gopkg.in/goyy/goyy.v0/web/session"
)

// Conf conf.Conf.
var Conf = &conf{
	Addr: ":9090",
	Err: &errOptions{
		Err401: "/err/401.html",
		Err403: "/err/403.html",
		Err404: "/err/404.html",
		Err500: "/err/500.html",
	},
	Api: &apiOptions{
		URL: "/apis",
	},
	Asset: &staticOptions{
		Enable: false,
		Ver:    "ver=1",
		URL:    "/assets",
		Dir:    "/app/assets",
	},
	Static: &staticOptions{
		Enable: false,
		Ver:    "ver=1",
		URL:    "/static",
		Dir:    "static",
	},
	Developer: &staticOptions{
		Enable: false,
		Ver:    "ver=1",
		URL:    "/gydev",
		Dir:    "/app/assets/gydev",
	},
	Operation: &staticOptions{
		Enable: false,
		Ver:    "ver=1",
		URL:    "/gyopr",
		Dir:    "/app/assets/gyopr",
	},
	Upload: &uploadOptions{
		Enable:  false,
		Dir:     "/app/assets/gyupl",
		URL:     "/gyupl",
		MaxSize: 5242880,
	},
	Export: &exportOptions{
		Dir: "/app/assets/gydev/export",
		Del: 600,
	},
	Sensitive: &sensitiveOptions{
		Enable:   false,
		Values:   []string{},
		Excludes: []string{},
	},
	Html: &htmlOptions{
		Enable:     false,
		Dir:        "templates",
		Extensions: []string{"html"},
		Reloaded:   true,
	},
	Session: &sessionOptions{
		Enable: true,
		Addr:   "127.0.0.1:6379",
		Options: &session.Options{
			Path:     "/",
			Domain:   "",
			MaxAge:   30 * 60,
			Secure:   false,
			HttpOnly: true,
		},
	},
	Secure: &secureOptions{
		Enable:     true,
		LoginUrl:   "/login.html",
		SuccessUrl: "/home.html",
		ForbidUrl:  "/err/403.html",
		Filters: []xtype.Map{
			{"/**", "anon"},
		},
	},
	Template: &templateOptions{
		Enable:     true,
		Dir:        "templates",
		Extensions: []string{"html"},
		Funcs:      []template.FuncMap{templates.Html.FuncMap},
		Delims: templateDelims{
			Left:  "<%",
			Right: "%>",
		},
		Reloaded: true,
	},
}

type conf struct {
	Addr      string            // the TCP network address
	Err       *errOptions       // Error options
	Api       *apiOptions       // Apis options
	Asset     *staticOptions    // Asset options
	Static    *staticOptions    // Static resource options
	Developer *staticOptions    // Developer static resource options
	Operation *staticOptions    // Operation static resource options
	Upload    *uploadOptions    // Upload options
	Export    *exportOptions    // Export options
	Sensitive *sensitiveOptions // Sensitive word options
	Html      *htmlOptions      // Html resource options
	Session   *sessionOptions   // the session TCP network address
	Secure    *secureOptions    // url secure options
	Template  *templateOptions  // template options
}

type errOptions struct {
	Err401 string // 401 URL prefix
	Err403 string // 403 URL prefix
	Err404 string // 404 URL prefix
	Err500 string // 500 URL prefix
}

type apiOptions struct {
	URL string // APIs URL prefix
}

type mappingOptions struct {
	path string
	dir  string
}

type mappingsOptions struct {
	mappings []*mappingOptions
}

func (me *mappingsOptions) Len() int {
	if me.mappings == nil {
		return 0
	}
	return len(me.mappings)
}

func (me *mappingsOptions) Add(path, dir string) {
	m := &mappingOptions{
		path: path,
		dir:  envs.ParseGOPATH(dir),
	}
	me.mappings = append(me.mappings, m)
}

func (me *mappingsOptions) Each(fn func(path, dir string) (isbreak bool, err error)) error {
	if me.mappings != nil && len(me.mappings) > 0 && fn != nil {
		for _, v := range me.mappings {
			isbreak, err := fn(v.path, v.dir)
			if isbreak {
				break
			}
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type staticOptions struct {
	Enable   bool            // Whether service is enabled
	Ver      string          // Static resource version
	URL      string          // Static resource URL prefix
	Dir      string          // Static resource directory
	Mappings mappingsOptions // Path and URL mapping, support for using %GOPATH% environment variables
}

type uploadOptions struct {
	Enable  bool   // Whether service is enabled
	URL     string // Upload URL prefix
	Dir     string // Upload directory
	MaxSize int    // Max upload size
}

type exportOptions struct {
	Dir string // Export directory
	Del int    // Interval delete time(second)
}

type sensitiveOptions struct {
	// Whether service is enabled
	Enable bool
	// Need to scan the sensitive word(Multiple characters using "," separated)
	Values []string
	// Excluded URL
	Excludes []string
}

type htmlOptions struct {
	Enable     bool            // Whether service is enabled
	Dir        string          // Directory to load templates. Default is "templates"
	Extensions []string        // Extensions to parse template files from. Defaults to ["html"]
	Reloaded   bool            // Reloaded sets up the template for each reload
	Mappings   mappingsOptions // Path and URL mapping, support for using %GOPATH% environment variables
}

type sessionOptions struct {
	Enable bool // Whether service is enabled
	*session.Options
	Addr     string
	Password string
}

type secureOptions struct {
	Enable     bool // Whether service is enabled
	LoginUrl   string
	SuccessUrl string
	ForbidUrl  string
	Filters    []xtype.Map
}

// Options is a struct for specifying configuration options for the html render
type templateOptions struct {
	// Whether service is enabled
	Enable bool
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
