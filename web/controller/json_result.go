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

// Success constructs a http.ResponseWriter parameter of the successful type.
func (me *JSONController) Success(c xhttp.Context, data interface{}) result.Http {
	return me.result(c, true, "0", "", "", "", "", "", data)
}

// SuccessMessage constructs a http.ResponseWriter parameter of the successful type.
func (me *JSONController) SuccessMessage(c xhttp.Context, msg string) result.Http {
	return me.result(c, true, "0", msg, "", "", "", "", nil)
}

// SuccessMsg constructs a http.ResponseWriter parameter of the successful type.
func (me *JSONController) SuccessMsg(c xhttp.Context, msg string, data interface{}) result.Http {
	return me.result(c, true, "0", msg, "", "", "", "", data)
}

// SuccessState constructs a http.ResponseWriter parameter of the successful type.
func (me *JSONController) SuccessState(c xhttp.Context, state string, data interface{}) result.Http {
	return me.result(c, true, "0", "", "", state, "", "", data)
}

// SuccessStatus constructs a http.ResponseWriter parameter of the successful type.
func (me *JSONController) SuccessStatus(c xhttp.Context, code string, data interface{}) result.Http {
	return me.result(c, true, code, "", "", "", "", "", data)
}

// SuccessToken constructs a http.ResponseWriter parameter of the successful type.
func (me *JSONController) SuccessToken(c xhttp.Context, token string, data interface{}) result.Http {
	return me.result(c, true, "0", "", token, "", "", "", data)
}

// SuccessStatusMsg constructs a http.ResponseWriter parameter of the successful type.
func (me *JSONController) SuccessStatusMsg(c xhttp.Context, code, msg string, data interface{}) result.Http {
	return me.result(c, true, code, msg, "", "", "", "", data)
}

// SuccessStatusMsgToken constructs a http.ResponseWriter parameter of the successful type.
func (me *JSONController) SuccessStatusMsgToken(c xhttp.Context, code, msg, token string, data interface{}) result.Http {
	return me.result(c, true, code, msg, token, "", "", "", data)
}

// SuccessResult constructs a http.ResponseWriter parameter of the successful type.
func (me *JSONController) SuccessResult(c xhttp.Context, r *result.Result) result.Http {
	return me.result(c, true, r.Code, r.Message, r.Token, "", r.Memo, r.Tag, r.Data)
}

// ----------------------------------------------------------
// fault
// ----------------------------------------------------------

// Fault constructs a http.ResponseWriter parameter of the failure type.
func (me *JSONController) Fault(c xhttp.Context, data interface{}) result.Http {
	return me.result(c, false, "0", "", "", "", "", "", data)
}

// FaultMessage constructs a http.ResponseWriter parameter of the failure type.
func (me *JSONController) FaultMessage(c xhttp.Context, msg string) result.Http {
	return me.result(c, false, "0", msg, "", "", "", "", nil)
}

// FaultMsg constructs a http.ResponseWriter parameter of the failure type.
func (me *JSONController) FaultMsg(c xhttp.Context, msg string, data interface{}) result.Http {
	return me.result(c, false, "0", msg, "", "", "", "", data)
}

// FaultState constructs a http.ResponseWriter parameter of the failure type.
func (me *JSONController) FaultState(c xhttp.Context, state string, data interface{}) result.Http {
	return me.result(c, false, "0", "", "", state, "", "", data)
}

// FaultStatus constructs a http.ResponseWriter parameter of the failure type.
func (me *JSONController) FaultStatus(c xhttp.Context, code string, data interface{}) result.Http {
	return me.result(c, false, code, "", "", "", "", "", data)
}

// FaultToken constructs a http.ResponseWriter parameter of the failure type.
func (me *JSONController) FaultToken(c xhttp.Context, token string, data interface{}) result.Http {
	return me.result(c, false, "0", "", token, "", "", "", data)
}

// FaultStatusMsg constructs a http.ResponseWriter parameter of the failure type.
func (me *JSONController) FaultStatusMsg(c xhttp.Context, code, msg string, data interface{}) result.Http {
	return me.result(c, false, code, msg, "", "", "", "", data)
}

// FaultStatusMsgToken constructs a http.ResponseWriter parameter of the failure type.
func (me *JSONController) FaultStatusMsgToken(c xhttp.Context, code, msg, token string, data interface{}) result.Http {
	return me.result(c, false, code, msg, token, "", "", "", data)
}

// FaultResult constructs a http.ResponseWriter parameter of the failure type.
func (me *JSONController) FaultResult(c xhttp.Context, r *result.Result) result.Http {
	return me.result(c, false, r.Code, r.Message, r.Token, "", r.Memo, r.Tag, r.Data)
}
