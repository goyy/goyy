// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xhttp

import (
	"gopkg.in/goyy/goyy.v0/util/crypto/aes"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"net/http"
	"regexp"
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
	if !s.IsLogin() {
		if me.isRedirectLogin(path) {
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
	return false
}

func (me *secureServeMux) isRedirectLogin(path string) bool {
	for _, v := range Conf.Secure.Filters {
		str := strings.Replace(v.Key, ".", `\.`, -1)
		str = strings.Replace(str, "**", ".*", -1)
		reg := regexp.MustCompile(str)
		if reg.MatchString(path) {
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
