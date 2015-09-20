// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of me source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package controller

import (
	"gopkg.in/goyy/goyy.v0/data/result"
	"gopkg.in/goyy/goyy.v0/util/webs"
	"gopkg.in/goyy/goyy.v0/web/secure"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

// ----------------------------------------------------------
// result
// ----------------------------------------------------------

func (me *ClientController) Result(c xhttp.Context, r result.Http) map[string]interface{} {
	params := webs.ToParams(c.Params())
	p := secure.Principal{}
	if secure.IsLogin(c) {
		v, err := secure.GetPrincipal(c)
		if err != nil {
			logger.Error(err.Error())
		} else {
			p = v
		}
	}
	return map[string]interface{}{
		"Ctx":         "",
		"Project":     me.Project,
		"Module":      me.Module,
		"Title":       me.Title,
		"Sifts":       me.Sifts,
		"Tag":         me.Tag,
		"Success":     r.Success,
		"Token":       r.Token,
		"Code":        r.Code,
		"Message":     r.Message,
		"State":       r.State,
		"Params":      params,
		"Attributes":  c.Attributes(),
		"Data":        r.Data,
		"LoginName":   p.LoginName,
		"Permissions": p.Permissions,
	}
}

func (me *ClientController) result(c xhttp.Context, success bool, code, msg, token string, state string, data interface{}) map[string]interface{} {
	r := result.Http{}
	r.Success = success
	r.Code = code
	r.Message = msg
	r.Token = token
	r.State = state
	r.Data = data
	return me.Result(c, r)
}

// ----------------------------------------------------------
// success
// ----------------------------------------------------------

func (me *ClientController) Success(c xhttp.Context, state string, data interface{}) map[string]interface{} {
	return me.result(c, true, "0", "", "", state, data)
}

func (me *ClientController) SuccessMsg(c xhttp.Context, msg, state string, data interface{}) map[string]interface{} {
	return me.result(c, true, "0", msg, "", state, data)
}

func (me *ClientController) SuccessStatus(c xhttp.Context, code, state string, data interface{}) map[string]interface{} {
	return me.result(c, true, code, "", "", state, data)
}

func (me *ClientController) SuccessToken(c xhttp.Context, token, state string, data interface{}) map[string]interface{} {
	return me.result(c, true, "0", "", token, state, data)
}

func (me *ClientController) SuccessStatusMsg(c xhttp.Context, code, msg, state string, data interface{}) map[string]interface{} {
	return me.result(c, true, code, msg, "", state, data)
}

func (me *ClientController) SuccessStatusMsgToken(c xhttp.Context, code, msg, token, state string, data interface{}) map[string]interface{} {
	return me.result(c, true, code, msg, token, state, data)
}

// ----------------------------------------------------------
// fault
// ----------------------------------------------------------

func (me *ClientController) Fault(c xhttp.Context, state string, data interface{}) map[string]interface{} {
	return me.result(c, false, "0", "", "", state, data)
}

func (me *ClientController) FaultMsg(c xhttp.Context, msg, state string, data interface{}) map[string]interface{} {
	return me.result(c, false, "0", msg, "", state, data)
}

func (me *ClientController) FaultStatus(c xhttp.Context, code, state string, data interface{}) map[string]interface{} {
	return me.result(c, false, code, "", "", state, data)
}

func (me *ClientController) FaultToken(c xhttp.Context, token, state string, data interface{}) map[string]interface{} {
	return me.result(c, false, "0", "", token, state, data)
}

func (me *ClientController) FaultStatusMsg(c xhttp.Context, code, msg, state string, data interface{}) map[string]interface{} {
	return me.result(c, false, code, msg, "", state, data)
}

func (me *ClientController) FaultStatusMsgToken(c xhttp.Context, code, msg, token, state string, data interface{}) map[string]interface{} {
	return me.result(c, false, code, msg, token, state, data)
}
