// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package i18n_test

import (
	"gopkg.in/goyy/goyy.v0/comm/i18n"
	"testing"
)

var en_US = map[string]string{
	"say":  "Hello, world!",
	"sayf": "Hello, %s!",
}
var locales = map[string]map[string]string{
	i18n.Locale_en_US: en_US,
}
var i18N = i18n.New(locales, i18n.Locale_en_US)

func TestMessage(t *testing.T) {
	key := "say"
	expected := "Hello, world!"
	if out := i18N.Message(key); out != expected {
		t.Errorf(`i18N.Message("%s") = "%s", want "%s"`, key, out, expected)
	}
}

func TestMessagef(t *testing.T) {
	key := "sayf"
	arg := "goyy"
	expected := "Hello, goyy!"
	if out := i18N.Messagef("sayf", "goyy"); out != expected {
		t.Errorf(`i18N.Messagef("%s", "%s") = "%s", want "%s"`, key, arg, out, expected)
	}
}

func TestError(t *testing.T) {
	key := "say"
	expected := "Hello, world!"
	if out := i18N.Error(key); out.Error() != expected {
		t.Errorf(`i18N.Error("%s") = error:"%s", want "%s"`, key, out, expected)
	}
}

func TestErrorf(t *testing.T) {
	key := "sayf"
	arg := "goyy"
	expected := "Hello, goyy!"
	if out := i18N.Errorf("sayf", "goyy"); out.Error() != expected {
		t.Errorf(`i18N.Errorf("%s", "%s") = error:"%s", want "%s"`, key, arg, out, expected)
	}
}
