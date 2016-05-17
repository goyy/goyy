// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"gopkg.in/goyy/goyy.v0/util/strings"
)

// validation is a field validate in an entity struct.
type validation struct {
	Name  string
	Value string
}

type validations []*validation

// Init sets the fields.
func newValidations(tag string) (validations, bool) {
	if strings.IsBlank(tag) {
		return nil, false
	}
	validates := strings.Split(tag, "&")
	if validates == nil || len(validates) == 0 {
		return nil, false
	}
	valids := make([]*validation, 0)
	ok := false
	for _, v := range validates {
		vs := strings.Split(v, "=")
		if len(vs) == 2 {
			name := strings.TrimSpace(vs[0])
			value := strings.TrimSpace(vs[1])
			switch strings.ToLower(name) {
			case "required":
				if strings.IsBlank(value) {
					continue
				}
				valid := &validation{
					Name: strings.ToLower(name),
				}
				if value == "true" {
					valid.Value = "true"
				} else {
					valid.Value = "false"
				}
				ok = true
				valids = append(valids, valid)
			case "min", "max", "range", "minlen", "maxlen", "rangelen",
				"email", "url", "ip", "mobile", "tel", "phone", "zipcode",
				"float", "integer", "alpha", "alrod", "alnum", "alnumrod",
				"alnumhan", "alnumhanrod", "alhan", "alhanrod", "han", "hanrod",
				"match", "nomatch", "minf", "maxf", "rangef":
				if strings.IsBlank(value) {
					continue
				}
				valid := &validation{
					Name:  strings.ToLower(name),
					Value: value,
				}
				ok = true
				valids = append(valids, valid)
			}
		}
	}
	return valids, ok
}
