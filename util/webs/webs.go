// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package webs

import (
	"net/http"
	"net/url"

	"gopkg.in/goyy/goyy.v0/util/bytes"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

// Form contains the parsed form data, including both the URL
// field's query parameters and the POST or PUT form data.
func ToParams(values url.Values, prefix ...string) map[string]string {
	result := make(map[string]string, 0)
	for k, v := range values {
		if _, ok := strings.HasAnyPrefix(k, prefix...); ok {
			result[k] = strings.JoinIgnoreBlank(v, ",")
		}
	}
	return result
}

// Form contains the parsed form data, including both the URL
// field's query parameters and the POST or PUT form data.
func Params(req *http.Request, prefix ...string) (map[string]string, error) {
	vs, err := Values(req)
	if err != nil {
		return nil, err
	}
	result := make(map[string]string, 0)
	for k, v := range vs {
		if _, ok := strings.HasAnyPrefix(k, prefix...); ok {
			result[k] = strings.JoinIgnoreBlank(v, ",")
		}
	}
	return result, nil
}

// Form contains the parsed form data, including both the URL
// field's query parameters and the POST or PUT form data.
func Values(req *http.Request) (values url.Values, err error) {
	if req.Method == "GET" {
		err := req.ParseForm()
		if err != nil {
			return nil, err
		}
		values = req.Form
	} else {
		if strings.Contains(req.Header.Get("Content-Type"), "multipart/form-data") {
			err := req.ParseMultipartForm(32 << 20)
			if err != nil {
				return nil, err
			}
			if values == nil {
				values = url.Values{}
			}
			m := req.MultipartForm
			if m != nil {
				for k, v := range m.Value {
					values[k] = v
				}
			}
		} else {
			err := req.ParseForm()
			if err != nil {
				return nil, err
			}
			values = req.PostForm
		}
		if values == nil {
			values = url.Values{}
		}
		u, err := url.Parse(req.RequestURI)
		if err != nil {
			return nil, err
		}
		if strings.IsNotBlank(u.RawQuery) {
			vs, err := url.ParseQuery(u.RawQuery)
			if err != nil {
				return nil, err
			}
			for k, v := range vs {
				for _, s := range v {
					if strings.IsNotBlank(s) {
						values.Add(k, s)
					}
				}
			}
		}
	}
	return
}

// RemoteIP returns the Remote IP
func RemoteIP(req *http.Request) string {
	unknown := "unknown"
	ip := req.Header.Get("x-forwarded-for")
	if strings.IsBlank(ip) || unknown == strings.ToLower(ip) {
		ip = req.Header.Get("X-Forwarded-For")
	}
	if strings.IsBlank(ip) || unknown == strings.ToLower(ip) {
		ip = req.Header.Get("Proxy-Client-IP")
	}
	if strings.IsBlank(ip) || unknown == strings.ToLower(ip) {
		ip = req.Header.Get("WL-Proxy-Client-IP")
	}
	if strings.IsBlank(ip) || unknown == strings.ToLower(ip) {
		ip = req.Header.Get("HTTP_CLIENT_IP")
	}
	if strings.IsBlank(ip) || unknown == strings.ToLower(ip) {
		ip = req.Header.Get("HTTP_X_FORWARDED_FOR")
	}
	if strings.IsBlank(ip) || unknown == strings.ToLower(ip) {
		ip = req.RemoteAddr
	}
	if strings.IsNotBlank(ip) && strings.Index(ip, ",") != -1 {
		ips := strings.Split(ip, ",")
		for i := 0; i < len(ips); i++ {
			if strings.IsNotBlank(ips[i]) && unknown != strings.ToLower(ip) {
				ip = ips[i]
				break
			}
		}
		if "0:0:0:0:0:0:1" == ip {
			ip = "127.0.0.1"
		}
	}
	if strings.IsNotBlank(ip) && strings.Index(ip, ":") != -1 {
		ip = strings.Before(ip, ":")
	}
	return ip
}

// ParseQuery returns the url.Query by the url.Values.
func ParseQuery(params url.Values) string {
	if params == nil {
		return ""
	}
	b := bytes.NewBuffer()
	i := 0
	for k, v := range params {
		param := strings.JoinIgnoreBlank(v, ",")
		if strings.IsNotBlank(param) {
			if i > 0 {
				b.WriteString("&")
			}
			b.WriteString(k + "=" + param)
			i++
		}
	}
	return b.String()
}

func ParseUrlSpecialChars(s string) string {
	b := bytes.NewBuffer()
	// Count %, check that they're well-formed.
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '%':
			if i+2 >= len(s) || !bytes.IsHex(s[i+1]) || !bytes.IsHex(s[i+2]) {
				b.WriteString("%25")
				continue
			}
		}
		b.WriteByte(s[i])
	}
	return b.String()
}

// IsXMLHttpRequest reports whether the request is a Ajax Request.
func IsXMLHttpRequest(r *http.Request) bool {
	reqType := r.Header.Get("X-Requested-With")
	if "XMLHttpRequest" == reqType { // AJAX
		return true
	}
	return false
}
