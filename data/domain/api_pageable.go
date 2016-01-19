// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package domain

import (
	"net/http"
	"strconv"

	"gopkg.in/goyy/goyy.v0/util/cookies"
	"gopkg.in/goyy/goyy.v0/util/errors"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"gopkg.in/goyy/goyy.v0/util/webs"
)

// Abstract interface for pagination information.
type Pageable interface {

	// Returns the page to be returned.
	// @return the page to be returned.
	PageNo() int

	// Returns the number of items to be returned.
	// @return the number of items of that page
	PageSize() int

	// Returns the offset to be taken according to the underlying page and page size.
	// @return the offset to be taken
	Offset() int
}

// NewPageable returns the Pageable from pageNo, pageSize.
func NewPageable(pageNo, pageSize int) Pageable {
	p := &pageable{}
	p.SetPageNo(pageNo)
	p.SetPageSize(pageSize)
	return p
}

// NewPageablePageNo returns the Pageable from pageNo.
// PageSize defaults to 10
func NewPageablePageNo(pageNo int) Pageable {
	return NewPageable(pageNo, defaultPageSize)
}

// NewPageableHTTP returns the Pageable from HTTP.
func NewPageableHTTP(w http.ResponseWriter, r *http.Request) (Pageable, error) {
	values, err := webs.Values(r)
	if err != nil {
		return nil, err
	}
	p := &pageable{}
	if _, ok := values[defaultRePage]; ok { // Get pageable from http.Cookie
		if v, err := getByCookie(w, r, defaultPageNoName); err == nil {
			p.SetPageNo(v)
		} else {
			return nil, err
		}
		if v, err := getByCookie(w, r, defaultPageSizeName); err == nil {
			p.SetPageSize(v)
		} else {
			return nil, err
		}
		return p, err
	} else { // Get pageable from http.Request
		if v, err := getByRequest(w, r, defaultPageNoName); err == nil {
			p.SetPageNo(v)
		} else {
			return nil, err
		}
		if v, err := getByRequest(w, r, defaultPageSizeName); err == nil {
			p.SetPageSize(v)
		} else {
			return nil, err
		}
	}
	// When rePageNo=true, set pageNo=1
	if v, ok := values[defaultRePageNo]; ok && v[0] == "true" {
		p.SetPageNo(defaultPageNo)
	}
	return p, nil
}

func getByCookie(w http.ResponseWriter, r *http.Request, name string) (int, error) {
	if v, err := cookies.Value(r, name); err == nil {
		if strings.IsNotBlank(v) {
			if val, err := strconv.Atoi(name); err == nil {
				return val, nil
			} else {
				return 0, err
			}
		} else {
			return setDefaultPage(w, name)
		}
	} else {
		return 0, err
	}
}

func getByRequest(w http.ResponseWriter, r *http.Request, name string) (int, error) {
	values, err := webs.Values(r)
	if err != nil {
		return 0, err
	}
	if _, ok := values[name]; ok {
		v := values.Get(name)
		if strings.IsNotBlank(v) {
			if val, err := strconv.Atoi(v); err == nil {
				if name == defaultPageNoName || name == defaultPageSizeName {
					cookies.SetValue(w, name, v)
				}
				return val, nil
			} else {
				return 0, err
			}
		}
	}
	return setDefaultPage(w, name)
}

func setDefaultPage(w http.ResponseWriter, name string) (int, error) {
	if name == defaultPageNoName {
		cookies.SetValue(w, name, strconv.Itoa(defaultPageNo))
		return defaultPageNo, nil
	}
	if name == defaultPageSizeName {
		cookies.SetValue(w, name, strconv.Itoa(defaultPageSize))
		return defaultPageSize, nil
	}
	return 0, errors.NewReqNameNotPresent(name)
}
