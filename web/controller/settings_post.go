// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package controller

import (
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/result"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

type post struct {
	PostIndex   func(c xhttp.Context, r *result.Page) error
	PostShow    func(c xhttp.Context, r *result.Entity) error
	PostAdd     func(c xhttp.Context, r *result.Entity) error
	PostEdit    func(c xhttp.Context, r *result.Entity) error
	PostSave    func(c xhttp.Context, r *result.Entity) error
	PostDisable func(c xhttp.Context, r *result.Entity) error
	PostExport  func(c xhttp.Context, r entity.Interfaces) error
}
