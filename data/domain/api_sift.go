// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package domain

import (
	"net/http"
	"net/url"

	"gopkg.in/goyy/goyy.v0/util/strings"
	"gopkg.in/goyy/goyy.v0/util/times"
	"gopkg.in/goyy/goyy.v0/util/webs"
)

// Query conditional filtering.
// eg: sNameEQST=goyy
type Sift interface {
	// Returns the name.
	// eg: sNameEQST
	// @return
	Name() string

	// Returns the name of query conditional filtering.
	// eg: Name
	// @return
	Key() string

	// Returns the value of query conditional filtering.
	// eg: goyy
	// @return
	Value() string

	// Returns the operator of query conditional filtering.
	// eg: EQ
	// @return
	Operator() string

	// Returns the type of query conditional filtering.
	// eg: ST
	// @return
	Type() string
}

// NewSift returns the Sift from name, value.
// Prefix[s] + Name + Operator[EQ|NE|...|OD] + Type[ST|BL|...|TS]
// example:
// sName
// sNameEQ
// sNameEQST
// sNameSTEQ
// sAgeGT
// sMemoLI
func NewSift(name, value string, prefix ...string) (Sift, bool) {
	s := &sift{name: name}
	var prefixOk string
	if len(prefix) == 0 {
		prefixOk = defaultSiftPrefix
	}

	if strings.IsBlank(name) {
		return nil, false
	}
	if len(prefix) > 0 {
		if p, ok := strings.HasAnyPrefix(name, prefix...); ok {
			prefixOk = p
		} else {
			return nil, false
		}
	}
	if !strings.HasPrefix(name, prefixOk) {
		return nil, false
	}
	k := strings.After(name, prefixOk)
	if strings.IsBlank(k) {
		return nil, false
	}
	if strings.IsBlank(value) {
		return nil, false
	}
	size := 0
	op12 := strings.Right(k, 2)
	op34 := strings.Slice(k, -4, -2)
	if op_tr == op12 || op_tr == op34 {
		return nil, false
	}
	if strings.ContainsSliceAny(op12, ops) {
		s.operator = op12
		size = size + 2
	}
	if strings.ContainsSliceAny(op12, ots) {
		s.typ = op12
		size = size + 2
	}
	if strings.ContainsSliceAny(op34, ops) {
		s.operator = op34
		size = size + 2
	}
	if strings.ContainsSliceAny(op34, ots) {
		s.typ = op34
		size = size + 2
	}
	if strings.IsBlank(s.operator) {
		s.operator = op_eq
	}
	if strings.IsBlank(s.typ) {
		s.typ = ot_st
	}
	s.key = strings.Left(k, len(k)-size)
	s.value = convertValue(s.operator, s.typ, value)
	return s, true
}

func NewSifts(values url.Values, prefix ...string) ([]Sift, error) {
	ss := make([]Sift, 0)
	for k, v := range values {
		value := strings.JoinIgnoreBlank(v, ",")
		if s, ok := NewSift(k, value, prefix...); ok {
			ss = append(ss, s)
		}
	}
	return ss, nil
}

func NewSiftsReq(req *http.Request, prefix ...string) ([]Sift, error) {
	values, err := webs.Values(req)
	if err != nil {
		return nil, err
	}
	return NewSifts(values, prefix...)
}

func SiftsToParams(sifts ...Sift) map[string]string {
	result := make(map[string]string, 0)
	for _, sift := range sifts {
		result[sift.Name()] = sift.Value()
	}
	return result
}

func SiftsToMap(sifts ...Sift) map[string]interface{} {
	result := make(map[string]interface{}, 0)
	for _, sift := range sifts {
		result[sift.Name()] = sift.Value()
	}
	return result
}

func convertValue(operator, typ, value string) string {
	switch operator {
	case op_lk:
		return "%" + value + "%"
	case op_ll:
		return "%" + value
	case op_lr:
		return value + "%"
	}
	switch typ {
	case ot_t2:
		if operator == "LT" {
			if val, err := times.AddYYMD(value, times.Day); err == nil {
				if v, err := times.ParseUnixYymd(val); err == nil {
					return v
				} else {
					logger.Errorln(err.Error())
				}
			} else {
				logger.Errorln(err.Error())
			}
		} else if operator == "LE" {
			val := strings.TrimSpace(value) + " 23:59:59"
			if v, err := times.ParseUnixYymdhms(val); err == nil {
				return v
			} else {
				logger.Errorln(err.Error())
			}
		} else {
			if v, err := times.ParseUnixYymd(value); err == nil {
				return v
			} else {
				logger.Errorln(err.Error())
			}
		}
	case ot_t5:
		if v, err := times.ParseUnixYymdhms(value); err == nil {
			return v
		} else {
			logger.Errorln(err.Error())
		}
	case ot_t4:
		if v, err := times.ParseUnixYymdhm(value); err == nil {
			return v
		} else {
			logger.Errorln(err.Error())
		}
	}
	return value
}
