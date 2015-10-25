// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package controller

import (
	"gopkg.in/goyy/goyy.v0/data/result"
	"gopkg.in/goyy/goyy.v0/util/templates"
	"gopkg.in/goyy/goyy.v0/web/client"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

type ClientController struct {
	prefn
	postfn
	Settings
	DTO  func() interface{}
	DTOs func() interface{}
}

func (me *ClientController) NewDTO() interface{} {
	if me.DTO != nil {
		return me.DTO()
	}
	return nil
}

func (me *ClientController) NewDTOs() interface{} {
	if me.DTOs != nil {
		return me.DTOs()
	}
	return nil
}

func (me *ClientController) NewDTOResult() *result.Result {
	return &result.Result{Data: me.NewDTO()}
}

func (me *ClientController) NewDTOsResult() *result.Result {
	return &result.Result{Data: me.NewDTOs()}
}

func (me *ClientController) NewPageResult() *result.Result {
	out := &result.Pagination{Slice: me.NewDTOs()}
	return &result.Result{Data: out}
}

func (me *ClientController) Index(c xhttp.Context) {
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
			if me.PostIndex != nil {
				if err = me.PostIndex(c, r); err != nil {
					me.Error(c, err)
					return
				}
			}
			err = c.HTML(xhttp.StatusOK, me.TmplDefault(), me.SuccessResult(c, templates.EnIndex, r))
			if err != nil {
				me.Error(c, err)
				return
			}
		},
	}
	cli.DoPost()
}

func (me *ClientController) Show(c xhttp.Context) {
	if me.PreShow != nil {
		if err := me.PreShow(c); err != nil {
			me.Error(c, err)
			return
		}
	}
	cli := &client.Client{
		URL:     me.URL(c, me.ApiShow()),
		Params:  c.Params(),
		Header:  c.Request().Header,
		Cookies: c.Request().Cookies(),
		OnError: func(err error) {
			me.Error(c, err)
			return
		},
		OnCompleted: func(rc *result.Client) {
			r := me.NewDTOResult()
			err := rc.ParseResult(r)
			if err != nil {
				me.Error(c, err)
				return
			}
			if me.PostShow != nil {
				if err = me.PostShow(c, r); err != nil {
					me.Error(c, err)
					return
				}
			}
			err = c.HTML(xhttp.StatusOK, me.TmplDefault(), me.SuccessResult(c, templates.EnShow, r))
			if err != nil {
				me.Error(c, err)
				return
			}
		},
	}
	cli.DoGet()
}

func (me *ClientController) Add(c xhttp.Context) {
	if me.PreAdd != nil {
		if err := me.PreAdd(c); err != nil {
			me.Error(c, err)
			return
		}
	}
	cli := &client.Client{
		URL:     me.URL(c, me.ApiAdd()),
		Params:  c.Params(),
		Header:  c.Request().Header,
		Cookies: c.Request().Cookies(),
		OnError: func(err error) {
			me.Error(c, err)
			return
		},
		OnCompleted: func(rc *result.Client) {
			r := me.NewDTOResult()
			err := rc.ParseResult(r)
			if err != nil {
				me.Error(c, err)
				return
			}
			if me.PostAdd != nil {
				if err = me.PostAdd(c, r); err != nil {
					me.Error(c, err)
					return
				}
			}
			err = c.HTML(xhttp.StatusOK, me.TmplDefault(), me.SuccessResult(c, templates.EnAdd, r))
			if err != nil {
				me.Error(c, err)
				return
			}
		},
	}
	cli.DoPost()
}

func (me *ClientController) Edit(c xhttp.Context) {
	if me.PreEdit != nil {
		if err := me.PreEdit(c); err != nil {
			me.Error(c, err)
			return
		}
	}
	cli := &client.Client{
		URL:     me.URL(c, me.ApiEdit()),
		Params:  c.Params(),
		Header:  c.Request().Header,
		Cookies: c.Request().Cookies(),
		OnError: func(err error) {
			me.Error(c, err)
			return
		},
		OnCompleted: func(rc *result.Client) {
			r := me.NewDTOResult()
			err := rc.ParseResult(r)
			if err != nil {
				me.Error(c, err)
				return
			}
			if me.PostEdit != nil {
				if err = me.PostEdit(c, r); err != nil {
					me.Error(c, err)
					return
				}
			}
			err = c.HTML(xhttp.StatusOK, me.TmplDefault(), me.SuccessResult(c, templates.EnEdit, r))
			if err != nil {
				me.Error(c, err)
				return
			}
		},
	}
	cli.DoPost()
}

func (me *ClientController) Save(c xhttp.Context) {
	if me.PreSave != nil {
		if err := me.PreSave(c); err != nil {
			me.Error(c, err)
			return
		}
	}
	cli := &client.Client{
		URL:     me.URL(c, me.ApiSave()),
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
			if me.PostSave != nil {
				if err = me.PostSave(c, r); err != nil {
					me.Error(c, err)
					return
				}
			}
			r.Message = i18N.Message("msg.save")
			err = c.HTML(xhttp.StatusOK, me.TmplDefault(), me.SuccessResult(c, templates.EnIndex, r))
			if err != nil {
				me.Error(c, err)
				return
			}
		},
	}
	cli.DoPost()
}

func (me *ClientController) Disable(c xhttp.Context) {
	if me.PreDisable != nil {
		if err := me.PreDisable(c); err != nil {
			me.Error(c, err)
			return
		}
	}
	cli := &client.Client{
		URL:     me.URL(c, me.ApiDisable()),
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
			if me.PostDisable != nil {
				if err = me.PostDisable(c, r); err != nil {
					me.Error(c, err)
					return
				}
			}
			r.Message = i18N.Message("msg.disable")
			err = c.HTML(xhttp.StatusOK, me.TmplDefault(), me.SuccessResult(c, templates.EnIndex, r))
			if err != nil {
				me.Error(c, err)
				return
			}
		},
	}
	cli.DoPost()
}

func (me *ClientController) Box(c xhttp.Context) {
	cli := &client.Client{
		URL:     me.URL(c, me.ApiBox()),
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

func (me *ClientController) URL(c xhttp.Context, path string) string {
	return "http://127.0.0.1:9097" + path
}
