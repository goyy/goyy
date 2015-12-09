// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package controller

import (
	"gopkg.in/goyy/goyy.v0/data/result"
	"gopkg.in/goyy/goyy.v0/util/webs"
	"gopkg.in/goyy/goyy.v0/web/session"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

// ----------------------------------------------------------
// result
// ----------------------------------------------------------

func (me *HTMLController) Result(c xhttp.Context, r result.Http) map[string]interface{} {
	params := webs.ToParams(c.Params())
	p := session.Principal{}
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

func (me *HTMLController) result(c xhttp.Context, success bool, code, msg, token, state, memo, tag string, data interface{}) map[string]interface{} {
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

func (me *HTMLController) Success(c xhttp.Context, state string, data interface{}) map[string]interface{} {
	return me.result(c, true, "0", "", "", state, "", "", data)
}

func (me *HTMLController) SuccessMessage(c xhttp.Context, msg, state string) map[string]interface{} {
	return me.result(c, true, "0", msg, "", state, "", "", nil)
}

func (me *HTMLController) SuccessMsg(c xhttp.Context, msg, state string, data interface{}) map[string]interface{} {
	return me.result(c, true, "0", msg, "", state, "", "", data)
}

func (me *HTMLController) SuccessStatus(c xhttp.Context, code, state string, data interface{}) map[string]interface{} {
	return me.result(c, true, code, "", "", state, "", "", data)
}

func (me *HTMLController) SuccessToken(c xhttp.Context, token, state string, data interface{}) map[string]interface{} {
	return me.result(c, true, "0", "", token, state, "", "", data)
}

func (me *HTMLController) SuccessStatusMsg(c xhttp.Context, code, msg, state string, data interface{}) map[string]interface{} {
	return me.result(c, true, code, msg, "", state, "", "", data)
}

func (me *HTMLController) SuccessStatusMsgToken(c xhttp.Context, code, msg, token, state string, data interface{}) map[string]interface{} {
	return me.result(c, true, code, msg, token, state, "", "", data)
}

func (me *HTMLController) SuccessResult(c xhttp.Context, state string, r *result.Result) map[string]interface{} {
	return me.result(c, true, r.Code, r.Message, r.Token, state, r.Memo, r.Tag, r.Data)
}

// ----------------------------------------------------------
// fault
// ----------------------------------------------------------

func (me *HTMLController) Fault(c xhttp.Context, state string, data interface{}) map[string]interface{} {
	return me.result(c, false, "0", "", "", state, "", "", data)
}

func (me *HTMLController) FaultMessage(c xhttp.Context, msg, state string) map[string]interface{} {
	return me.result(c, false, "0", msg, "", state, "", "", nil)
}

func (me *HTMLController) FaultMsg(c xhttp.Context, msg, state string, data interface{}) map[string]interface{} {
	return me.result(c, false, "0", msg, "", state, "", "", data)
}

func (me *HTMLController) FaultStatus(c xhttp.Context, code, state string, data interface{}) map[string]interface{} {
	return me.result(c, false, code, "", "", state, "", "", data)
}

func (me *HTMLController) FaultToken(c xhttp.Context, token, state string, data interface{}) map[string]interface{} {
	return me.result(c, false, "0", "", token, state, "", "", data)
}

func (me *HTMLController) FaultStatusMsg(c xhttp.Context, code, msg, state string, data interface{}) map[string]interface{} {
	return me.result(c, false, code, msg, "", state, "", "", data)
}

func (me *HTMLController) FaultStatusMsgToken(c xhttp.Context, code, msg, token, state string, data interface{}) map[string]interface{} {
	return me.result(c, false, code, msg, token, state, "", "", data)
}

func (me *HTMLController) FaultResult(c xhttp.Context, state string, r *result.Result) map[string]interface{} {
	return me.result(c, false, r.Code, r.Message, "", state, r.Memo, r.Tag, r.Data)
}
