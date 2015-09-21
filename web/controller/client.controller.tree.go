// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package controller

import (
	"encoding/json"
	"gopkg.in/goyy/goyy.v0/comm/xtype"
	"gopkg.in/goyy/goyy.v0/data/result"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"gopkg.in/goyy/goyy.v0/util/templates"
	"gopkg.in/goyy/goyy.v0/web/client"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

type ClientTreeController struct {
	ClientController
}

func (me *ClientTreeController) Index(c xhttp.Context) {
	if me.PreIndex != nil {
		if err := me.PreIndex(c); err != nil {
			me.Error(c, err)
			return
		}
	}
	cli := &client.Client{
		URL:     me.URL(c, me.ApiIndex()),
		Params:  c.Params(),
		Header:  c.Request().Header,
		Cookies: c.Request().Cookies(),
		OnError: func(err error) {
			me.Error(c, err)
			return
		},
		OnCompleted: func(rc *result.Client) {
			r := me.NewPageResult()
			err := rc.ParseResult(r)
			if err != nil {
				me.Error(c, err)
				return
			}
			if strings.IsNotBlank(r.Tag) {
				boxes := make([]xtype.Box, 0)
				err := json.Unmarshal([]byte(r.Tag), &boxes)
				if err != nil {
					me.Error(c, err)
					return
				}
				c.SetAttribute(defaultParents, boxes)
			}
			if me.PostIndex != nil {
				if err = me.PostIndex(c, r); err != nil {
					me.Error(c, err)
					return
				}
			}
			err = c.HTML(xhttp.StatusOK, me.TmplDefault(), me.Success(c, templates.EnIndex, r.Data))
			if err != nil {
				me.Error(c, err)
				return
			}
		},
	}
	cli.DoPost()
}

func (me *ClientTreeController) Tree(c xhttp.Context) {
	cli := &client.Client{
		URL:     me.URL(c, me.ApiTree()),
		Params:  c.Params(),
		Header:  c.Request().Header,
		Cookies: c.Request().Cookies(),
		OnError: func(err error) {
			me.Error(c, err)
			return
		},
		OnCompleted: func(rc *result.Client) {
			err := c.Text(xhttp.StatusOK, string(rc.Body))
			if err != nil {
				me.Error(c, err)
				return
			}
		},
	}
	cli.DoGet()
}
