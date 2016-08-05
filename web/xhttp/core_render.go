// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xhttp

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"gopkg.in/goyy/goyy.v0/util/files"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"gopkg.in/goyy/goyy.v0/util/templates"
)

var (
	r     = &renderer{}
	mutex sync.Mutex
)

type renderer struct {
	t *template.Template
}

func (me *renderer) HTML(w http.ResponseWriter, status int, name string, v interface{}) error {
	mutex.Lock()
	if Conf.Template.Reloaded || me.t == nil {
		err := me.compile()
		if err != nil {
			mutex.Unlock()
			return err
		}
	}
	mutex.Unlock()
	me.writeHeader(w, status, "text/html")
	return me.t.ExecuteTemplate(w, name, v)
}

func (me *renderer) JSON(w http.ResponseWriter, status int, v interface{}) error {
	me.writeHeader(w, status, "application/json")
	return json.NewEncoder(w).Encode(v)
}

func (me *renderer) JSONP(w http.ResponseWriter, status int, callback string, v interface{}) error {
	result, err := json.Marshal(v)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	// JSON marshaled fine, write out the result.
	me.writeHeader(w, status, "application/json")
	w.Write([]byte(callback + "("))
	w.Write(result)
	w.Write([]byte(");"))

	return nil
}

func (me *renderer) XML(w http.ResponseWriter, status int, v interface{}) error {
	me.writeHeader(w, status, "application/xml")
	return xml.NewEncoder(w).Encode(v)
}

func (me *renderer) Text(w http.ResponseWriter, status int, format string, values ...interface{}) (err error) {
	me.writeHeader(w, status, "text/plain")
	if len(values) > 0 {
		_, err = w.Write([]byte(fmt.Sprintf(format, values...)))
	} else {
		_, err = w.Write([]byte(format))
	}
	return
}

// Error writes the given HTTP status to the current ResponseWriter
func (me *renderer) Error(w http.ResponseWriter, status int) error {
	me.writeHeader(w, status, "text/html")
	return nil
}

func (me *renderer) Redirect(w http.ResponseWriter, req *http.Request, location string, status ...int) error {
	code := http.StatusFound
	if len(status) == 1 {
		code = status[0]
	}
	http.Redirect(w, req, location, code)
	return nil
}

func (me *renderer) writeHeader(w http.ResponseWriter, status int, contentType string) {
	contentType = contentType + "; charset=utf-8"
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(status)
}

func (me *renderer) compile() error {
	options := Conf.Template
	dir := options.Dir
	me.t = template.New(dir)
	me.t.Delims(options.Delims.Left, options.Delims.Right)
	// parse an initial template in case we don't have any
	me.t = template.Must(me.t.Parse("xhttp"))

	// add our funcmaps
	me.t.Funcs(templates.Html.FuncMap)
	for _, funcs := range options.Funcs {
		me.t.Funcs(funcs)
	}

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		r, err := filepath.Rel(dir, path)
		if err != nil {
			logger.Error(err.Error())
			return err
		} else {
			r = strings.Replace(r, "\\", "/", -1)
		}

		ext := strings.ToLower(files.Extension(r))

		for _, extension := range options.Extensions {
			if ext == extension {
				buf, err := ioutil.ReadFile(path)
				if err != nil {
					logger.Error(err.Error())
					return err
				}

				name := (r[0 : len(r)-len(ext)-1])
				me.t = me.t.New(filepath.ToSlash(name))

				// Bomb out if parse fails. We don't want any silent server starts.
				me.t = template.Must(me.t.Parse(string(buf)))
				break
			}
		}
		return nil
	})
	return nil
}

func (me *renderer) execute(name string, binding interface{}) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	return buf, me.t.ExecuteTemplate(buf, name, binding)
}
