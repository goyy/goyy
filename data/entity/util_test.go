// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package entity_test

import (
	"gopkg.in/goyy/goyy.v0/data/entity"
	"testing"
)

func TestFormatJSON(t *testing.T) {
	expected := `{"id":"1","name":"admin","passwd":"1ap93md","age":18,"email":"admin@gmail.com","version":0}`
	u := NewUser()
	u.SetId("1")
	u.SetName("admin")
	u.SetPasswd("1ap93md")
	u.SetAge(18)
	u.SetEmail("admin@gmail.com")
	if out := entity.FormatJSON(u); out != expected {
		t.Errorf(`InterfaceResult.JSON() = "%v", want "%v"`, out, expected)
	}
}

func TestParseJSON(t *testing.T) {
	json := `{"id":"1","name":"admin","passwd":"1ap93md","age":18,"email":"admin@gmail.com","version":0}`
	u := NewUser()
	entity.ParseJSON(u, json)
	expected := "1"
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
