// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cache_test

import (
	"gopkg.in/goyy/goyy.v0/data/cache"
	"testing"
)

func TestString(t *testing.T) {
	expected := "value1"
	key := "key1"
	cache.Set(key, expected)
	if out, _ := cache.Get(key); out != expected {
		t.Errorf(`cache.Get("%s") = "%s", want "%s"`, key, out, expected)
	}
}

func TestMap(t *testing.T) {
	expected := "hvalue1"
	key := "hkey1"
	filed := "hfiled1"
	cache.HSet(key, filed, expected)
	if out, _ := cache.HGet(key, filed); out != expected {
		t.Errorf(`cache.Get("%s", "%s") = "%s", want "%s"`, key, filed, out, expected)
	}
}

func TestStruct(t *testing.T) {
	key := "user"
	name := "admin"
	expected := "admin@gmail.com"
	type users struct {
		Name  string
		Email string
	}
	user := &users{Name: name, Email: expected}
	cache.SSet(key, user)
	out := &users{}
	if cache.SGet(key, out); out.Email != expected {
		t.Errorf(`cache.SGet("%s") = "%s", want "%s"`, key, out.Email, expected)
	}
}

func init() {
	cache.Init(cache.Conf{Address: "10.105.99.81:6379"})
}
