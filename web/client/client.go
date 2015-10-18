// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package client

import (
	"gopkg.in/goyy/goyy.v0/data/result"
	"gopkg.in/goyy/goyy.v0/util/errors"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client struct {
	URL         string
	Params      url.Values
	Header      http.Header
	Cookies     []*http.Cookie
	Timeout     int
	OnTimeout   func()
	OnError     func(error)
	OnCompleted func(*result.Client)
}

func (me *Client) DoGet() {
	me.do("GET")
}

func (me *Client) DoPost() {
	me.do("POST")
}

func (me *Client) GoGet() {
	go me.do("GET")
}

func (me *Client) GoPost() {
	go me.do("POST")
}

func (me *Client) QueueGet() {
	go me.do("GET")
}

func (me *Client) QueuePost() {
	go me.do("POST")
}

func (me *Client) onError(err error) {
	logger.Error(err.Error())
	if me.OnError != nil {
		me.OnError(err)
	}
}

func (me *Client) do(method string) {
	if strings.IsBlank(me.URL) {
		me.onError(errors.NewNotBlank("URL"))
		return
	}
	client := &http.Client{}
	req, err := me.getRequest(method)
	if err != nil {
		me.onError(err)
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		logger.Debug(err.Error())
		me.onError(err)
		return
	}
	if resp.StatusCode >= 400 {
		me.onError(errors.New(resp.Status))
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Debug(err.Error())
		me.onError(err)
		return
	}
	if me.OnCompleted != nil {
		r := &result.Client{
			Body:       body,
			Status:     resp.Status,
			StatusCode: resp.StatusCode,
			Header:     resp.Header,
			Cookies:    resp.Cookies(),
		}
		me.OnCompleted(r)
	}
}

func (me *Client) getRequest(method string) (*http.Request, error) {
	if method == "GET" {
		url := me.URL
		if me.Params != nil {
			if strings.Contains(me.URL, "?") {
				url = me.URL + "&" + me.Params.Encode()
			} else {
				url = me.URL + "?" + me.Params.Encode()
			}
		}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			logger.Debug(err.Error())
			return req, err
		}
		req.Header = me.Header
		for _, c := range me.Cookies {
			req.AddCookie(c)
		}
		return req, err
	} else {
		req, err := http.NewRequest("POST", me.URL, strings.NewReader(me.Params.Encode()))
		if err != nil {
			logger.Debug(err.Error())
			return req, err
		}
		req.Header = me.Header
		if strings.IsBlank(req.Header.Get("Content-Type")) {
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		}
		for _, c := range me.Cookies {
			req.AddCookie(c)
		}
		return req, err
	}
}
