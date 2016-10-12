// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package profile_test

import (
	"testing"

	"gopkg.in/goyy/goyy.v0/comm/profile"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

func TestActives(t *testing.T) {
	expected := 0
	if out := len(profile.Actives()); out != expected {
		t.Errorf(`profile.Actives() = "%v", want "%v"`, out, expected)
	}
	want := true
	if out := profile.Accepts(profile.DEFAULT); out != want {
		t.Errorf(`profile.Accepts("%s") = "%v", want "%v"`, profile.DEFAULT, out, want)
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
		t.Errorf(`profile.Accepts("%s") = "%v", want "%v"`, profile.TEST, out, expected)
	}
	if out := profile.Accepts(profile.PROD, profile.TEST); out != expected {
		t.Errorf(`profile.Accepts("%s") = "%v", want "%v"`, profile.PROD+","+profile.TEST, out, expected)
	}
	if out := profile.Accepts(profile.PROD, profile.DEV); out != expected {
		t.Errorf(`profile.Accepts("%s") = "%v", want "%v"`, profile.PROD+","+profile.DEV, out, expected)
	}
	notDev := "!development"
	if out := profile.Accepts(notDev); out != expected {
		t.Errorf(`profile.Accepts("%s") = "%v", want "%v"`, notDev, out, expected)
	}
	if out := profile.Accepts("tmp", "!development"); out != expected {
		t.Errorf(`profile.Accepts("%s") = "%v", want "%v"`, "tmp,"+notDev, out, expected)
	}
	want := false
	if out := profile.Accepts(profile.DEV); out != want {
		t.Errorf(`profile.Accepts("%s") = "%v", want "%v"`, profile.DEV, out, want)
	}
	if out := profile.Accepts(profile.DEV, profile.UNIT); out != want {
		t.Errorf(`profile.Accepts("%s") = "%v", want "%v"`, profile.DEV+","+profile.UNIT, out, want)
	}
}

func TestSetDefault(t *testing.T) {
	expected := true
	if out := profile.Accepts(profile.DEFAULT); out != expected {
		t.Errorf(`profile.Accepts("%s") = "%v", want "%v"`, profile.DEFAULT, out, expected)
	}
	profile.SetDefault(profile.PROD)
	if out := profile.Accepts(profile.PROD); out != expected {
		t.Errorf(`profile.Accepts("%s") = "%v", want "%v"`, profile.PROD, out, expected)
	}
	want := false
	if out := profile.Accepts(profile.DEV); out != want {
		t.Errorf(`profile.Accepts("%s") = "%v", want "%v"`, profile.DEV, out, want)
	}
	if out := profile.Accepts(profile.DEV, profile.UNIT); out != want {
		t.Errorf(`profile.Accepts("%s") = "%v", want "%v"`, profile.DEV+","+profile.UNIT, out, want)
	}
}
