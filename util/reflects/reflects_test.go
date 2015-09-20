// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package reflects_test

import (
	"gopkg.in/goyy/goyy.v0/util/reflects"
	"reflect"
	"testing"
)

type demo struct {
	attr int `tag:"attr"`
	Name int `tag:"Name"`
}

func (me *demo) Attr() int {
	return me.attr
}

func (me *demo) SetAttr(value int) {
	me.attr = value
}

func TestCallAttr(t *testing.T) {
	expected := 2
	d := &demo{attr: expected}
	if outs, err := reflects.Call(d, "Attr"); err == nil {
		if outs[0] != expected {
			t.Errorf(`reflects.Call(...) = "%v", want "%v"`, outs[0], expected)
		}
	} else {
		t.Error(err.Error())
	}
}

func TestCallSetAttr(t *testing.T) {
	expected := 2
	d := &demo{}
	if _, err := reflects.Call(d, "SetAttr", expected); err == nil {
		if d.Attr() != expected {
			t.Errorf(`reflects.Call(...) = "%v", want "%v"`, d.Attr(), expected)
		}
	} else {
		t.Error(err.Error())
	}
}

/*
func TestSetattr(t *testing.T) {
	expected := 2
	d := &demo{}
	if err := reflects.Set(d, "attr", expected); err == nil {
		if d.Name != expected {
			t.Errorf(`reflects.Set(...) = "%v", want "%v"`, d.attr, expected)
		}
	} else {
		t.Error(err.Error())
	}
}
*/

func TestSet(t *testing.T) {
	expected := 2
	d := &demo{}
	if err := reflects.Set(d, "Name", expected); err == nil {
		if d.Name != expected {
			t.Errorf(`reflects.Set(...) = "%v", want "%v"`, d.Name, expected)
		}
	} else {
		t.Error(err.Error())
	}
}

func TestInterface(t *testing.T) {
	expected := 2
	d := &demo{Name: expected}
	if out, err := reflects.Interface(d, "Name"); err == nil {
		if out != expected {
			t.Errorf(`reflects.Interface(...) = "%v", want "%v"`, out, expected)
		}
	} else {
		t.Error(err.Error())
	}
}

func TestTagAttr(t *testing.T) {
	expected := "attr"
	d := &demo{}
	if out, err := reflects.Tag(d, expected, "tag"); err == nil {
		if out != expected {
			t.Errorf(`reflects.Tag(...) = "%v", want "%v"`, out, expected)
		}
	} else {
		t.Error(err.Error())
	}
}

func TestTagName(t *testing.T) {
	expected := "Name"
	d := &demo{}
	if out, err := reflects.Tag(d, expected, "tag"); err == nil {
		if out != expected {
			t.Errorf(`reflects.Tag(...) = "%v", want "%v"`, out, expected)
		}
	} else {
		t.Error(err.Error())
	}
}

func TestField(t *testing.T) {
	expected := "Name"
	d := &demo{}
	println(reflect.ValueOf(demo{}).FieldByName("Name").CanSet())
	if out, err := reflects.Tag(d, expected, "tag"); err == nil {
		if out != expected {
			t.Errorf(`reflects.Tag(...) = "%v", want "%v"`, out, expected)
		}
	} else {
		t.Error(err.Error())
	}
}
