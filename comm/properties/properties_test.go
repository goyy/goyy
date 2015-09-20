// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package properties_test

import (
	"gopkg.in/goyy/goyy.v0/comm/properties"
	"testing"
)

func TestProperty(t *testing.T) {
	p, _ := properties.New("./example.properties")
	key := "say"
	expected := "Hello, world!"
	if out := p.Property(key); out != expected {
		t.Errorf(`Property("%s") = "%s", want "%s"`, key, out, expected)
	}
}

func TestPropertyf(t *testing.T) {
	p, _ := properties.New("./example.properties")
	key := "sayf"
	arg := "goyy"
	expected := "Hello, goyy!"
	if out := p.Propertyf(key, arg); out != expected {
		t.Errorf(`Propertyf("%s", "%s") = "%s", want "%s"`, key, arg, out, expected)
	}
}

func TestSetProperty(t *testing.T) {
	p, _ := properties.New("./example.properties")
	// read
	key := "set"
	expected := "Hello, world!"
	if out := p.Property(key); out != expected {
		t.Errorf(`Property("%s") = "%s", want "%s"`, key, out, expected)
	}
	// write
	expected = "Hello, goyy!"
	p.SetProperty(key, expected)
	if out := p.Property(key); out != expected {
		t.Errorf(`Property("%s") = "%s", want "%s"`, key, out, expected)
	}
	// revert
	p.SetProperty("set", "Hello, world!")
}
