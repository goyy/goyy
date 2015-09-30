// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package controller

import (
	"gopkg.in/goyy/goyy.v0/data/result"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

// ----------------------------------------------------------
// result
// ----------------------------------------------------------

func (me *JSONController) result(c xhttp.Context, success bool, code, msg, token, state, memo, tag string, data interface{}) result.Http {
	r := result.Http{}
	r.Success = success
	r.Code = code
	r.Message = msg
	r.Token = token
	r.State = state
	r.Memo = memo
	r.Tag = tag
	r.Data = data
	return r
}

// ----------------------------------------------------------
// success
// ----------------------------------------------------------

func (me *JSONController) Success(c xhttp.Context, data interface{}) result.Http {
	return me.result(c, true, "0", "", "", "", "", "", data)
}

func (me *JSONController) SuccessState(c xhttp.Context, state string, data interface{}) result.Http {
	return me.result(c, true, "0", "", "", state, "", "", data)
}

func (me *JSONController) SuccessMsg(c xhttp.Context, msg, state string, data interface{}) result.Http {
	return me.result(c, true, "0", msg, "", state, "", "", data)
}

func (me *JSONController) SuccessStatus(c xhttp.Context, code, state string, data interface{}) result.Http {
	return me.result(c, true, code, "", "", state, "", "", data)
}

func (me *JSONController) SuccessToken(c xhttp.Context, token, state string, data interface{}) result.Http {
	return me.result(c, true, "0", "", token, state, "", "", data)
}

func (me *JSONController) SuccessStatusMsg(c xhttp.Context, code, msg, state string, data interface{}) result.Http {
	return me.result(c, true, code, msg, "", state, "", "", data)
}

func (me *JSONController) SuccessStatusMsgToken(c xhttp.Context, code, msg, token, state string, data interface{}) result.Http {
	return me.result(c, true, code, msg, token, state, "", "", data)
}

func (me *JSONController) SuccessResult(c xhttp.Context, r *result.Result) result.Http {
	return me.result(c, true, r.Code, r.Message, r.Token, "", r.Memo, r.Tag, r.Data)
}

// ----------------------------------------------------------
// fault
// ----------------------------------------------------------

func (me *JSONController) Fault(c xhttp.Context, data interface{}) result.Http {
	return me.result(c, false, "0", "", "", "", "", "", data)
}

func (me *JSONController) FaultState(c xhttp.Context, state string, data interface{}) result.Http {
	return me.result(c, false, "0", "", "", state, "", "", data)
}

func (me *JSONController) FaultMsg(c xhttp.Context, msg, state string, data interface{}) result.Http {
	return me.result(c, false, "0", msg, "", state, "", "", data)
}

func (me *JSONController) FaultStatus(c xhttp.Context, code, state string, data interface{}) result.Http {
	return me.result(c, false, code, "", "", state, "", "", data)
}

func (me *JSONController) FaultToken(c xhttp.Context, token, state string, data interface{}) result.Http {
	return me.result(c, false, "0", "", token, state, "", "", data)
}

func (me *JSONController) FaultStatusMsg(c xhttp.Context, code, msg, state string, data interface{}) result.Http {
	return me.result(c, false, code, msg, "", state, "", "", data)
}

func (me *JSONController) FaultStatusMsgToken(c xhttp.Context, code, msg, token, state string, data interface{}) result.Http {
	return me.result(c, false, code, msg, token, state, "", "", data)
}

func (me *JSONController) FaultResult(c xhttp.Context, r *result.Result) result.Http {
	return me.result(c, false, r.Code, r.Message, r.Token, "", r.Memo, r.Tag, r.Data)
}
