// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xhttp

import (
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"gopkg.in/goyy/goyy.v0/util/webs"
	"net/http"
	"net/url"
	"sync"
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
					unescape, err := url.QueryUnescape(v)
					if err != nil {
						logger.Error(err.Error())
						return true
					}
					unescape2, err := url.QueryUnescape(unescape)
					if err != nil {
						logger.Error(err.Error())
						return true
					}
					value := strings.TrimSpace(strings.ToLower(unescape2))
					for _, val := range me.values {
						if strings.Contains(value, val) {
							logger.Printf("The content of the input contains illegal characters:%s -> %s \r\n", val, r.URL.Path)
							reqType := r.Header.Get("X-Requested-With")
							if "XMLHttpRequest" == reqType { // AJAX
								f := `{"success":false,"message":"%s"}`
								c := fmt.Sprintf(f, i18N.Message("err.illegal"))
								w.Write([]byte(c))
							} else {
								f := `<script language="javascript">alert("%s");window.history.go(-1);</script>`
								c := fmt.Sprintf(f, i18N.Message("err.illegal"))
								w.Write([]byte(c))
							}
							return true
						}
					}
				}
			}
		}
	}
	return false
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
