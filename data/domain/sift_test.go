// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package domain_test

import (
	"net/http"
	"testing"

	"gopkg.in/goyy/goyy.v0/data/domain"
)

func TestNewSift(t *testing.T) {
	s := []struct {
		name, value             string
		key, val, operator, typ string
		expected                bool
	}{
		{"", "1", "", "1", "", "", false},
		{"s", "2", "", "2", "", "", false},
		{"typesId", "2", "", "2", "", "", false},
		{"sName", "3", "Name", "3", "EQ", "ST", true},
		{"sNameEQ", "4", "Name", "4", "EQ", "ST", true},
		{"sNameEQST", "5", "Name", "5", "EQ", "ST", true},
		{"sNameSTEQ", "6", "Name", "6", "EQ", "ST", true},
		{"sAgeGTI0", "7", "Age", "7", "GT", "I0", true},
		{"sMemoLI", "8", "Memo", "8", "LI", "ST", true},
		{"sMemoLK", "8", "Memo", "%8%", "LK", "ST", true},
		{"sMemoLL", "8", "Memo", "%8", "LL", "ST", true},
		{"sMemoLR", "8", "Memo", "8%", "LR", "ST", true},
		{"sCreatedLIT2", "2014-04-03", "Created", "1396454400", "LI", "T2", true},
		{"sCreatedLTT2", "2014-04-03", "Created", "1396540800", "LT", "T2", true},
		{"sCreatedLIT5", "2014-04-03 13:31:45", "Created", "1396503105", "LI", "T5", true},
		{"sCreatedLIT4", "2014-04-03 13:31", "Created", "1396503060", "LI", "T4", true},
	}
	for _, v := range s {
		if out, ok := domain.NewSift(v.name, v.value); ok {
			if ok != v.expected {
				t.Errorf(`domain.NewSift(%s, %s):"%v", want:"%v"`, v.name, v.value, ok, v.expected)
			}
			if out.Key() != v.key {
				t.Errorf(`domain.NewSift(%s, %s).Key():"%v", want:"%v"`, v.name, v.value, out.Key(), v.key)
			}
			if out.Value() != v.val {
				t.Errorf(`domain.NewSift(%s, %s).Value():"%v", want:"%v"`, v.name, v.value, out.Value(), v.val)
			}
			if out.Operator() != v.operator {
				t.Errorf(`domain.NewSift(%s, %s).Operator():"%v", want:"%v"`, v.name, v.value, out.Operator(), v.operator)
			}
			if out.Type() != v.typ {
				t.Errorf(`domain.NewSift(%s, %s).Type():"%v", want:"%v"`, v.name, v.value, out.Type(), v.typ)
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

	expected := 2
	if out, _ := domain.NewSiftsReq(req); len(out) != expected {
		t.Errorf(`domain.NewSifts(req) = "%v", want "%v"`, len(out), expected)
	}
}
