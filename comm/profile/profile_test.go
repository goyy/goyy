// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package profile_test

import (
	"gopkg.in/goyy/goyy.v0/comm/profile"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"testing"
)

func TestActives(t *testing.T) {
	expected := 0
	if out := len(profile.Actives()); out != expected {
		t.Errorf(`profile.Actives() = "%v", want "%v"`, out, expected)
	}
	want := true
	if out := profile.Accepts(profile.DEFAULT); out != want {
		t.Errorf(`profile.Accepts("%s") = "%v", want "%v"`, profile.PROD, out, want)
	}
}

func TestSetActives(t *testing.T) {
	expected := "production,test"
	profile.SetActives(profile.PROD, profile.TEST)
	if out := strings.Join(profile.Actives(), ","); out != expected {
		t.Errorf(`profile.Actives() = "%v", want "%v"`, out, expected)
	}
	profile.SetActives()
	if out := strings.Join(profile.Actives(), ","); out != expected {
		t.Errorf(`profile.Actives() = "%v", want "%v"`, out, expected)
	}
	profile.SetActives("")
	if out := strings.Join(profile.Actives(), ","); out != expected {
		t.Errorf(`profile.Actives() = "%v", want "%v"`, out, expected)
	}
	profile.SetActives("  ")
	if out := strings.Join(profile.Actives(), ","); out != expected {
		t.Errorf(`profile.Actives() = "%v", want "%v"`, out, expected)
	}
	profile.SetActives("", " ", "    ")
	if out := strings.Join(profile.Actives(), ","); out != expected {
		t.Errorf(`profile.Actives() = "%v", want "%v"`, out, expected)
	}
}

func TestAccepts(t *testing.T) {
	expected := true
	profile.SetActives(profile.PROD, profile.TEST)
	if out := profile.Accepts(profile.PROD); out != expected {
		t.Errorf(`profile.Accepts("%s") = "%v", want "%v"`, profile.PROD, out, expected)
	}
	if out := profile.Accepts(profile.TEST); out != expected {
		t.Errorf(`profile.Accepts("%s") = "%v", want "%v"`, profile.PROD, out, expected)
	}
	if out := profile.Accepts(profile.PROD, profile.TEST); out != expected {
		t.Errorf(`profile.Accepts("%s") = "%v", want "%v"`, profile.PROD, out, expected)
	}
	if out := profile.Accepts(profile.PROD, profile.DEV); out != expected {
		t.Errorf(`profile.Accepts("%s") = "%v", want "%v"`, profile.PROD, out, expected)
	}
	if out := profile.Accepts("!development"); out != expected {
		t.Errorf(`profile.Accepts("%s") = "%v", want "%v"`, profile.PROD, out, expected)
	}
	if out := profile.Accepts("tmp", "!development"); out != expected {
		t.Errorf(`profile.Accepts("%s") = "%v", want "%v"`, profile.PROD, out, expected)
	}
	want := false
	if out := profile.Accepts(profile.DEV); out != want {
		t.Errorf(`profile.Accepts("%s") = "%v", want "%v"`, profile.PROD, out, want)
	}
	if out := profile.Accepts(profile.DEV, profile.UNIT); out != want {
		t.Errorf(`profile.Accepts("%s") = "%v", want "%v"`, profile.PROD, out, want)
	}
}
