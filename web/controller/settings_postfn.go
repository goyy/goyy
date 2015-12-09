// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package controller

import (
	"gopkg.in/goyy/goyy.v0/data/result"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

type postfn struct {
	PostIndex   func(c xhttp.Context, r *result.Result) error
	PostShow    func(c xhttp.Context, r *result.Result) error
	PostAdd     func(c xhttp.Context, r *result.Result) error
	PostEdit    func(c xhttp.Context, r *result.Result) error
	PostSave    func(c xhttp.Context, r *result.Result) error
	PostDisable func(c xhttp.Context, r *result.Result) error
	PostExp     func(c xhttp.Context, r *result.Result) error
}
