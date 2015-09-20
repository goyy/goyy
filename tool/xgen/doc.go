// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

/*
Package xgen implements generate utility functions.

Usage

	//go:generate xgen.v0 $GOFILE

	// User stores user account information.
	// @entity
	type User struct {
		table   schema.Table
		id      entity.String `xgen:"column=id&primary=true"`
		email   entity.String `xgen:"column=email"`
		created entity.Time   `xgen:"column=created"`
	}
*/
package main
