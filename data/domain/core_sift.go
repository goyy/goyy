// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package domain

var ops = []string{opEQ, opNE, opGT, opLT, opGE,
	opLE, opLI, opLK, opLL, opLR,
	opBE, opNB, opIN, opNI, opNU, opNN, opTR, opOA, opOD}

var ots = []string{otST, otBL,
	otI0, otI8, otI1, otI3, otI6,
	otU0, otU8, otU1, otU3, otU6,
	otF3, otF6,
	otTD, otTT, otTM, otT2, otT5, otT4}

// Query conditional filtering.
type sift struct {
	name     string
	key      string
	value    string
	typ      string
	operator string
}

// Returns the name.
// @return
func (me *sift) Name() string {
	return me.name
}

// Returns the name of query conditional filtering.
// @return
func (me *sift) Key() string {
	return me.key
}

// Returns the value of query conditional filtering.
// @return
func (me *sift) Value() string {
	return me.value
}

// Returns the operator of query conditional filtering.
// @return
func (me *sift) Operator() string {
	return me.operator
}

// Returns the type of query conditional filtering.
// @return
func (me *sift) Type() string {
	return me.typ
}
