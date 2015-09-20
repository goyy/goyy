// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package db

import (
	"time"
)

type (
	String struct {
		Column
		Value string
	}
	Int struct {
		Column
		Value int
	}
	Uint struct {
		Column
		Value uint
	}
	Float struct {
		Column
		Value float32
	}
	Bool struct {
		Column
		Value bool
	}
	Time struct {
		Column
		Value time.Time
	}

	Id struct {
		Column
		Value string
	}
	Creater struct {
		Column
		Value string
	}
	// Created denotes a timestamp field that is automatically set on insert.
	Created struct {
		Column
		Value time.Time
	}
	Modifier struct {
		Column
		Value string
	}
	// Modified denotes a timestamp field that is automatically set on update.
	Modified struct {
		Column
		Value time.Time
	}
	Version struct {
		Column
		Value int
	}
	Disabled struct {
		Column
		Value int
	}

	Parent struct {
		Column
		Value string
	}
	User struct {
		Column
		Value string
	}
	Role struct {
		Column
		Value string
	}
	Post struct {
		Column
		Value string
	}
	Org struct {
		Column
		Value string
	}
	Area struct {
		Column
		Value string
	}
)
