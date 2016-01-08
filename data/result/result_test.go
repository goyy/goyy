// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package result_test

import (
	"testing"

	"gopkg.in/goyy/goyy.v0/data/result"
)

func TestResultJSON(t *testing.T) {
	json := `{"success":true,"token":"","code":"1","message":"ok","memo":"","tag":"","data":{"id":"1","name":"admin"}}`
	r := &result.Result{
		Success: true,
		Code:    "1",
		Message: "ok",
		Data: map[string]string{
			"id":   "1",
			"name": "admin",
		},
	}
	if out, _ := r.JSON(); out != json {
		t.Errorf(`Result.JSON() = "%v", want "%v"`, out, json)
	}
}

func TestResultParseJSON(t *testing.T) {
	json := `{"success":true,"token":"","code":"1","message":"ok","memo":"","tag":"","data":{"id":"1","name":"admin"}}`
	r := &result.Result{Data: map[string]string{}}
	if err := r.ParseJSON(json); err != nil {
		t.Errorf(`ParseJSON->error:"%v"`, err.Error())
		return
	}
	if out := r.Success; out != true {
		t.Errorf(`ParseJSON->Success = "%v", want "%v"`, out, true)
	}
	expected := "1"
	if out := r.Code; out != expected {
		t.Errorf(`ParseJSON->Code = "%v", want "%v"`, out, expected)
	}
	expected = "ok"
	if out := r.Message; out != expected {
		t.Errorf(`ParseJSON->Message = "%v", want "%v"`, out, expected)
	}
	expected = "1"
	if u, ok := r.Data.(map[string]string); ok {
		if out := u["id"]; out != expected {
			t.Errorf(`ParseJSON->Id = "%v", want "%v"`, out, expected)
		}
		expected = "admin"
		if out := u["name"]; out != expected {
			t.Errorf(`ParseJSON->Name = "%v", want "%v"`, out, expected)
		}
	}
}
