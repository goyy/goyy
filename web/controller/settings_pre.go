// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package controller

import (
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

type pre struct {
	PreIndex   func(c xhttp.Context) error
	PreShow    func(c xhttp.Context) error
	PreAdd     func(c xhttp.Context) error
	PreEdit    func(c xhttp.Context) error
	PreSave    func(c xhttp.Context) error
	PreDisable func(c xhttp.Context) error
	PreExport  func(c xhttp.Context) error
}
