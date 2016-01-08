// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package result_test

import (
	"testing"

	"gopkg.in/goyy/goyy.v0/data/domain"
	"gopkg.in/goyy/goyy.v0/data/result"
)

func TestPageJSON(t *testing.T) {
	expected := `{"success":true,"id":"","token":"","code":"1","message":"ok","memo":"","tag":"","data":{"pageNo":1,"pageSize":10,"pageFn":"page","totalPages":1,"totalElements":2,"length":8,"slider":1,"slice":[{"id":"1","name":"admin","passwd":"1ap93md","age":18,"email":"admin@gmail.com","version":0},{"id":"2","name":"sa","passwd":"3df69ku7h","age":20,"email":"sa@gmail.com","version":0}]}}`
	u1 := NewUser()
	u1.SetId("1")
	u1.SetName("admin")
	u1.SetPasswd("1ap93md")
	u1.SetAge(18)
	u1.SetEmail("admin@gmail.com")
	u2 := NewUser()
	u2.SetId("2")
	u2.SetName("sa")
	u2.SetPasswd("3df69ku7h")
	u2.SetAge(20)
	u2.SetEmail("sa@gmail.com")
	users := NewUserEntities(2)
	users.Append(u1)
	users.Append(u2)
	pageable := domain.NewPageable(1, 10)
	page := domain.NewPage(pageable, users, 2)
	r := result.Page{
		Success: true,
		Code:    "1",
		Message: "ok",
		Data:    page,
	}
	if out := r.JSON(); out != expected {
		t.Errorf(`Entities.JSON() = "%v", want "%v"`, out, expected)
	}
}

func TestPageParseJSON(t *testing.T) {
	json := `{"success":true,"code":"1","message":"ok","memo":"","tag":"","data":{"pageNo":1,"pageSize":10,"totalElements":2,"function":"page","length":8,"slider":1,"slice":[{"id":"1","name":"admin","passwd":"1ap93md","age":18,"email":"admin@gmail.com","version":0},{"id":"2","name":"sa","passwd":"3df69ku7h","age":20,"email":"sa@gmail.com","version":0}]}}`
	users := NewUserEntities(2)
	pageable := domain.NewPageable(1, 10)
	page := domain.NewPage(pageable, users, 2)
	r := result.Page{Data: page}
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
	u := users.Index(0)
	expected = "1"
	if out := u.Get("id").(string); out != expected {
		t.Errorf(`ParseJSON->Id = "%v", want "%v"`, out, expected)
	}
	expected = "admin"
	if out := u.Get("name").(string); out != expected {
		t.Errorf(`ParseJSON->Name = "%v", want "%v"`, out, expected)
	}
	expected = "1ap93md"
	if out := u.Get("passwd").(string); out != expected {
		t.Errorf(`ParseJSON->Passwd = "%v", want "%v"`, out, expected)
	}
	got := 18
	if out := u.Get("age").(int); out != got {
		t.Errorf(`ParseJSON->Age = "%v", want "%v"`, out, got)
	}
	expected = "admin@gmail.com"
	if out := u.Get("email").(string); out != expected {
		t.Errorf(`ParseJSON->Email = "%v", want "%v"`, out, expected)
	}
	got = 0
	if out := u.Get("version").(int); out != got {
		t.Errorf(`ParseJSON->Version = "%v", want "%v"`, out, got)
	}
}
