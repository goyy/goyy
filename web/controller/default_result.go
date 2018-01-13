// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package controller

import (
	"gopkg.in/goyy/goyy.v0/comm/xtype"
	"gopkg.in/goyy/goyy.v0/data/result"
	"gopkg.in/goyy/goyy.v0/util/webs"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

// ----------------------------------------------------------
// result
// ----------------------------------------------------------

// Result build the http.ResponseWriter parameters.
func (me *Controller) Result(c xhttp.Context, r result.Http) map[string]interface{} {
	params := webs.ToParams(c.Params())
	p := xtype.Principal{}
	if c.Session().IsLogin() {
		v, err := c.Session().Principal()
		if err != nil {
			logger.Error(err.Error())
		} else {
			p = v
		}
	}
	return map[string]interface{}{
		"Project":     me.Project,
		"Module":      me.Module,
		"Title":       me.Title,
		"Sifts":       me.Sifts,
		"Success":     r.Success,
		"Token":       r.Token,
		"Code":        r.Code,
		"Message":     r.Message,
		"State":       r.State,
		"Memo":        r.Memo,
		"Tag":         r.Tag,
		"Params":      params,
		"Attributes":  c.Attributes(),
		"Data":        r.Data,
		"LoginName":   p.LoginName,
		"Permissions": p.Permissions,
	}
}

func (me *Controller) result(c xhttp.Context, success bool, code int, msg, token, state, memo, tag string, data interface{}) map[string]interface{} {
	r := result.Http{}
	r.Success = success
	r.Code = code
	r.Message = msg
	r.Token = token
	r.State = state
	r.Memo = memo
	r.Tag = tag
	r.Data = data
	return me.Result(c, r)
}

// ----------------------------------------------------------
// success
// ----------------------------------------------------------

// Success constructs a http.ResponseWriter parameter of the successful type.
func (me *Controller) Success(c xhttp.Context, state string, data interface{}) map[string]interface{} {
	return me.result(c, true, 2000, "", "", state, "", "", data)
}

// SuccessMessage constructs a http.ResponseWriter parameter of the successful type.
func (me *Controller) SuccessMessage(c xhttp.Context, msg, state string) map[string]interface{} {
	return me.result(c, true, 2000, msg, "", state, "", "", nil)
}

// SuccessMsg constructs a http.ResponseWriter parameter of the successful type.
func (me *Controller) SuccessMsg(c xhttp.Context, msg, state string, data interface{}) map[string]interface{} {
	return me.result(c, true, 2000, msg, "", state, "", "", data)
}

// SuccessStatus constructs a http.ResponseWriter parameter of the successful type.
func (me *Controller) SuccessStatus(c xhttp.Context, code int, state string, data interface{}) map[string]interface{} {
	return me.result(c, true, code, "", "", state, "", "", data)
}

// SuccessToken constructs a http.ResponseWriter parameter of the successful type.
func (me *Controller) SuccessToken(c xhttp.Context, token, state string, data interface{}) map[string]interface{} {
	return me.result(c, true, 2000, "", token, state, "", "", data)
}

// SuccessStatusMsg constructs a http.ResponseWriter parameter of the successful type.
func (me *Controller) SuccessStatusMsg(c xhttp.Context, code int, msg, state string, data interface{}) map[string]interface{} {
	return me.result(c, true, code, msg, "", state, "", "", data)
}

// SuccessStatusMsgToken constructs a http.ResponseWriter parameter of the successful type.
func (me *Controller) SuccessStatusMsgToken(c xhttp.Context, code int, msg, token, state string, data interface{}) map[string]interface{} {
	return me.result(c, true, code, msg, token, state, "", "", data)
}

// SuccessResult constructs a http.ResponseWriter parameter of the successful type.
func (me *Controller) SuccessResult(c xhttp.Context, state string, r *result.Result) map[string]interface{} {
	return me.result(c, true, r.Code, r.Message, r.Token, state, r.Memo, r.Tag, r.Data)
}

// ----------------------------------------------------------
// fault
// ----------------------------------------------------------

// Fault constructs a http.ResponseWriter parameter of the failure type.
func (me *Controller) Fault(c xhttp.Context, state string, data interface{}) map[string]interface{} {
	return me.result(c, false, 4000, "", "", state, "", "", data)
}

// FaultMessage constructs a http.ResponseWriter parameter of the failure type.
func (me *Controller) FaultMessage(c xhttp.Context, msg, state string) map[string]interface{} {
	return me.result(c, false, 4000, msg, "", state, "", "", nil)
}

// FaultMsg constructs a http.ResponseWriter parameter of the failure type.
func (me *Controller) FaultMsg(c xhttp.Context, msg, state string, data interface{}) map[string]interface{} {
	return me.result(c, false, 4000, msg, "", state, "", "", data)
}

// FaultStatus constructs a http.ResponseWriter parameter of the failure type.
func (me *Controller) FaultStatus(c xhttp.Context, code int, state string, data interface{}) map[string]interface{} {
	return me.result(c, false, code, "", "", state, "", "", data)
}

// FaultToken constructs a http.ResponseWriter parameter of the failure type.
func (me *Controller) FaultToken(c xhttp.Context, token, state string, data interface{}) map[string]interface{} {
	return me.result(c, false, 4000, "", token, state, "", "", data)
}

// FaultStatusMsg constructs a http.ResponseWriter parameter of the failure type.
func (me *Controller) FaultStatusMsg(c xhttp.Context, code int, msg, state string, data interface{}) map[string]interface{} {
	return me.result(c, false, code, msg, "", state, "", "", data)
}

// FaultStatusMsgToken constructs a http.ResponseWriter parameter of the failure type.
func (me *Controller) FaultStatusMsgToken(c xhttp.Context, code int, msg, token, state string, data interface{}) map[string]interface{} {
	return me.result(c, false, code, msg, token, state, "", "", data)
}

// FaultResult constructs a http.ResponseWriter parameter of the failure type.
func (me *Controller) FaultResult(c xhttp.Context, state string, r *result.Result) map[string]interface{} {
	return me.result(c, false, r.Code, r.Message, r.Token, state, r.Memo, r.Tag, r.Data)
}
