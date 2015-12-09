// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package result_test

import (
	"gopkg.in/goyy/goyy.v0/data/result"
	"testing"
)

func TestEntityJSON(t *testing.T) {
	json := `{"success":true,"token":"","code":"1","message":"ok","memo":"","tag":"","data":{"id":"1","name":"admin","passwd":"1ap93md","age":18,"email":"admin@gmail.com","version":0}}`
	u := NewUser()
	u.SetId("1")
	u.SetName("admin")
	u.SetPasswd("1ap93md")
	u.SetAge(18)
	u.SetEmail("admin@gmail.com")
	r := result.Entity{
		Success: true,
		Code:    "1",
		Message: "ok",
		Data:    u,
	}
	if out := r.JSON(); out != json {
		t.Errorf(`Entity.JSON() = "%v", want "%v"`, out, json)
	}
}

func TestEntityParseJSON(t *testing.T) {
	json := `{"success":true,"code":"1","message":"ok","memo":"","tag":"","data":{"id":"1","name":"admin","passwd":"1ap93md","age":18,"email":"admin@gmail.com","version":0}}`
	u := NewUser()
	r := result.Entity{Data: u}
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
	if out := u.Id(); out != expected {
		t.Errorf(`ParseJSON->Id = "%v", want "%v"`, out, expected)
	}
	expected = "admin"
	if out := u.Name(); out != expected {
		t.Errorf(`ParseJSON->Name = "%v", want "%v"`, out, expected)
	}
	expected = "1ap93md"
	if out := u.Passwd(); out != expected {
		t.Errorf(`ParseJSON->Passwd = "%v", want "%v"`, out, expected)
	}
	got := 18
	if out := u.Age(); out != got {
		t.Errorf(`ParseJSON->Age = "%v", want "%v"`, out, got)
	}
	expected = "admin@gmail.com"
	if out := u.Email(); out != expected {
		t.Errorf(`ParseJSON->Email = "%v", want "%v"`, out, expected)
	}
	got = 0
	if out := u.Version(); out != got {
		t.Errorf(`ParseJSON->Version = "%v", want "%v"`, out, got)
	}
}
