// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package controller

import (
	"gopkg.in/goyy/goyy.v0/util/webs"
	"net/http"
	"net/url"
)

func errorSave(req *http.Request, err error) {
	if err == nil {
		return
	}
	message := err.Error()
	var params string
	values, err := webs.Values(req)
	if err != nil {
		logger.Println(err.Error())
		return
	}
	if values != nil {
		params = webs.ParseQuery(values)
	}
	genre := "20"
	remoteIp := webs.RemoteIP(req)
	userAgent := req.Header.Get("user-agent")
	requestMethod := req.Method
	requestUri := req.RequestURI
	requestParams := params
	location := ""
	urlvs := url.Values{
		"genre":         {genre},
		"remoteIp":      {remoteIp},
		"userAgent":     {userAgent},
		"requestMethod": {requestMethod},
		"requestUri":    {requestUri},
		"requestParams": {requestParams},
		"location":      {location},
		"message":       {message},
	}
	nurl := "http://" + req.Host + "/sys/log/save?" + urlvs.Encode()
	nreq, err := http.NewRequest("POST", nurl, nil)
	if err != nil {
		logger.Println(err.Error())
	}
	nreq.Header = req.Header
	client := &http.Client{}
	resp, err := client.Do(nreq)
	if err != nil {
		logger.Println(err.Error())
	}
	defer resp.Body.Close()
}
