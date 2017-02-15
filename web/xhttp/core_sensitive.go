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

// sensitive word
type sensitiveServeMux struct {
	values []string
}

var ssm = &sensitiveServeMux{}

var sensitiveMutex sync.Mutex

func (me *sensitiveServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) bool {
	// Excluded URL
	if Conf.Sensitive.Excludes != nil && len(Conf.Sensitive.Excludes) > 0 {
		for _, e := range Conf.Sensitive.Excludes {
			for _, v := range strings.Split(e, ",") {
				v = strings.TrimSpace(v)
				if strings.IsBlank(v) {
					continue
				}
				if v == r.URL.Path {
					return false
				}
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
							logger.Printf("The content of the input contains sensitive words:%s -> %s \r\n", val, r.URL.Path)
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

func (me *sensitiveServeMux) write(w http.ResponseWriter, r *http.Request) {
	msg := i18N.Message("err.sensitive")
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

func (me *sensitiveServeMux) setValues() {
	if Conf.Sensitive.Enable && me.values == nil {
		sensitiveMutex.Lock()
		if me.values == nil {
			if Conf.Sensitive.Values != nil && len(Conf.Sensitive.Values) > 0 {
				for _, val := range Conf.Sensitive.Values {
					if strings.IsNotBlank(val) {
						vs := strings.Split(val, ",")
						for _, v := range vs {
							if strings.IsNotBlank(v) {
								v = strings.ToLower(v)
								v = strings.TrimSpace(v)
								me.values = append(me.values, v)
							}
						}
					}
				}
			}
		}
		sensitiveMutex.Unlock()
	}
}
