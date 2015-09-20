// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package domain_test

import (
	"gopkg.in/goyy/goyy.v0/data/domain"
	"net/http"
	"testing"
)

func TestNewSift(t *testing.T) {
	s := []struct {
		name, value, key, operator string
		expected                   bool
	}{
		{"", "1", "", "", false},
		{"s", "2", "", "", false},
		{"sName", "3", "Name", "EQ", true},
		{"sNameEQ", "4", "Name", "EQ", true},
		{"sNameEQST", "5", "Name", "EQ", true},
		{"sNameSTEQ", "6", "Name", "EQ", true},
		{"sAgeGT", "7", "Age", "GT", true},
		{"sMemoLI", "8", "Memo", "LI", true},
	}
	for _, v := range s {
		if out, ok := domain.NewSift(v.name, v.value); ok {
			if ok != v.expected {
				t.Errorf(`domain.NewSift():"%v", want:"%v"`, ok, v.expected)
			}
			if out.Key() != v.key {
				t.Errorf(`domain.NewSift().Key():"%v", want:"%v"`, out.Key(), v.key)
			}
			if out.Value() != v.value {
				t.Errorf(`domain.NewSift().Value():"%v", want:"%v"`, out.Value(), v.value)
			}
			if out.Value() != v.value {
				t.Errorf(`domain.NewSift().Operator():"%v", want:"%v"`, out.Operator(), v.operator)
			}
		} else {
			if ok != v.expected {
				t.Errorf(`domain.NewSift():"%v", want:"%v"`, ok, v.expected)
			}
		}
	}
}

func TestNewSifts(t *testing.T) {
	url := "http://www.goyy.org/t?sNameEQ=goyy&sAgeGT=18"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		println(err.Error())
	}

	expected := 0
	if out, _ := domain.NewSifts(req); len(out) != expected {
		t.Errorf(`domain.NewSifts(req) = "%v", want "%v"`, len(out), expected)
	}
}
