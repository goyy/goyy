// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xhttp

import (
	"net/http"
	"net/url"
	"sync"

	"gopkg.in/goyy/goyy.v0/util/strings"
	"gopkg.in/goyy/goyy.v0/util/webs"
)

// illegal character
type illegalServeMux struct {
	values []string
}

var ism = &illegalServeMux{}

var illegalMutex sync.Mutex

func (me *illegalServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) bool {
	// Excluded URL
	if Conf.Illegal.Excludes != nil && len(Conf.Illegal.Excludes) > 0 {
		for _, e := range Conf.Illegal.Excludes {
			if e == r.URL.Path {
				return false
			}
		}
	}
	me.setValues()
	if me.values != nil && len(me.values) > 0 {
		params, err := webs.Values(r)
		if err != nil {
			logger.Error(err.Error())
			return false
		}
		for _, vs := range params {
			for _, v := range vs {
				if strings.IsNotBlank(v) {
					v = webs.ParseUrlSpecialChars(v)
					unescape, err := url.QueryUnescape(v)
					if err != nil {
						logger.Error(err.Error())
						me.write(w, r)
						return true
					}
					unescape = webs.ParseUrlSpecialChars(unescape)
					unescape2, err := url.QueryUnescape(unescape)
					if err != nil {
						logger.Error(err.Error())
						me.write(w, r)
						return true
					}
					value := strings.TrimSpace(strings.ToLower(unescape2))
					for _, val := range me.values {
						if strings.Contains(value, val) {
							logger.Printf("The content of the input contains illegal characters:%s -> %s \r\n", val, r.URL.Path)
							me.write(w, r)
							return true
						}
					}
				}
			}
		}
	}
	return false
}

func (me *illegalServeMux) write(w http.ResponseWriter, r *http.Request) {
	msg := i18N.Message("err.illegal")
	if webs.IsXMLHttpRequest(r) { // AJAX
		c := `{"success":false,"message":"` + msg + `"}`
		w.WriteHeader(451)
		w.Write([]byte(c))
	} else {
		c := `<script language="javascript">alert("` + msg + `");window.history.go(-1);</script>`
		w.WriteHeader(451)
		w.Write([]byte(c))
	}
}

func (me *illegalServeMux) setValues() {
	if Conf.Illegal.Enable && me.values == nil {
		illegalMutex.Lock()
		if me.values == nil {
			if Conf.Illegal.Values != nil && len(Conf.Illegal.Values) > 0 {
				for _, val := range Conf.Illegal.Values {
					if strings.IsNotBlank(val) {
						vs := strings.Split(val, ",")
						for _, v := range vs {
							if strings.IsNotBlank(v) {
								v = strings.ToLower(v)
								me.values = append(me.values, v)
							}
						}
					}
				}
			}
		}
		illegalMutex.Unlock()
	}
}
