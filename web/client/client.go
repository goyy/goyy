// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package client

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"gopkg.in/goyy/goyy.v0/data/result"
	"gopkg.in/goyy/goyy.v0/util/errors"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

// Client client.Client.
type Client struct {
	URL         string
	Params      url.Values
	PostBody    string
	Referer     string
	Header      http.Header
	Cookies     []*http.Cookie
	Transport   http.RoundTripper
	Timeout     int64
	OnTimeout   func()
	OnError     func(error)
	OnCompleted func(*result.Client)
}

// DoGet execute the get request.
func (me *Client) DoGet() {
	me.do("GET")
}

// DoPost execute the post request.
func (me *Client) DoPost() {
	me.do("POST")
}

// GoGet thread to execute the get request.
func (me *Client) GoGet() {
	go me.doRecover("GET")
}

// GoPost thread to execute the post request.
func (me *Client) GoPost() {
	go me.doRecover("POST")
}

// QueueGet Queue to execute the get request.
func (me *Client) QueueGet() {
	go me.doRecover("GET")
}

// QueuePost Queue to execute the post request.
func (me *Client) QueuePost() {
	go me.doRecover("POST")
}

func (me *Client) onError(err error) {
	logger.Error(err.Error())
	if me.OnError != nil {
		me.OnError(err)
	}
}

func (me *Client) doRecover(method string) {
	defer func() {
		if err := recover(); err != nil {
			logger.Printf("Panic recovery -> %s\n", err)
		}
	}()
	me.do(method)
}

func (me *Client) do(method string) {
	if strings.IsBlank(me.URL) {
		me.onError(errors.NewNotBlank("URL"))
		return
	}
	var timeout int64 = 30
	if me.Timeout > 0 {
		timeout = me.Timeout
	}
	client := &http.Client{
		Timeout: time.Duration(timeout * int64(time.Second)),
	}
	if me.Transport != nil {
		client.Transport = me.Transport
	}
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
	switch method {
	case "GET":
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
	default:
		var params io.Reader
		url := me.URL
		if strings.IsNotBlank(me.PostBody) {
			if me.Params != nil {
				if strings.Contains(me.URL, "?") {
					url = me.URL + "&" + me.Params.Encode()
				} else {
					url = me.URL + "?" + me.Params.Encode()
				}
			}
			params = strings.NewReader(me.PostBody)
		} else {
			params = strings.NewReader(me.Params.Encode())
		}
		req, err := http.NewRequest("POST", url, params)
		if err != nil {
			logger.Debug(err.Error())
			return req, err
		}
		req.Header = me.Header
		if strings.IsBlank(req.Header.Get("Content-Type")) {
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		}
		if strings.IsBlank(me.Referer) {
			req.Header.Set("Referer", me.URL)
		} else {
			req.Header.Set("Referer", me.Referer)
		}
		for _, c := range me.Cookies {
			req.AddCookie(c)
		}
		return req, err
	}
}
