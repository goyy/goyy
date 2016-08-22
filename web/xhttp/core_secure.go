// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xhttp

import (
	"net/http"
	"regexp"

	"gopkg.in/goyy/goyy.v0/util/crypto/aes"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"gopkg.in/goyy/goyy.v0/util/webs"
)

type secureServeMux struct {
	loginUrl   string
	successUrl string
}

type box struct {
	Key   string
	Value string
}

var sec = &secureServeMux{}

func (me *secureServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) bool {
	path := r.URL.Path
	s := newSession4Redis(w, r)
	if me.isForbidden(path) {
		if webs.IsXMLHttpRequest(r) { // AJAX
			msg := i18N.Message("err.403")
			c := `{"success":false,"message":"` + msg + `"}`
			w.WriteHeader(StatusForbidden)
			w.Write([]byte(c))
			return true
		} else {
			http.Redirect(w, r, Conf.Secure.ForbidUrl, http.StatusFound)
			return true
		}
	}
	if !s.IsLogin() {
		if me.isRedirectLogin(path) {
			if webs.IsXMLHttpRequest(r) { // AJAX
				msg := i18N.Message("err.401")
				c := `{"success":false,"message":"` + msg + `"}`
				w.WriteHeader(StatusUnauthorized)
				w.Write([]byte(c))
				return true
			} else {
				loginUrl := me.loginUrl
				if strings.IsBlank(loginUrl) {
					loginUrl = Conf.Secure.LoginUrl
				}
				// After login support to redirect to the URL before
				if url, err := aes.EncryptHex(r.URL.String(), aes.DefaultKey); err == nil {
					if strings.Index(loginUrl, "?") > 0 {
						loginUrl = loginUrl + "&redirect=" + url
					} else {
						loginUrl = loginUrl + "?redirect=" + url
					}
				}
				http.Redirect(w, r, loginUrl, http.StatusFound)
				return true
			}
		}
	}
	return false
}

func (me *secureServeMux) isForbidden(path string) bool {
	for _, v := range Conf.Secure.Filters {
		if me.isMatch(path, v.Key) {
			if v.Value == "forbid" {
				return true
			}
		}
	}
	return false
}

func (me *secureServeMux) isRedirectLogin(path string) bool {
	for _, v := range Conf.Secure.Filters {
		if me.isMatch(path, v.Key) {
			switch v.Value {
			case "authc":
				return true
			case "anon":
				return false
			}
		}
	}
	return false
}

func (me *secureServeMux) isMatch(path, conf string) bool {
	str := strings.Replace(conf, ".", `\.`, -1)
	str = strings.Replace(str, "**", ".*", -1)
	reg := regexp.MustCompile(str)
	if reg.MatchString(path) {
		return true
	}
	return false
}
