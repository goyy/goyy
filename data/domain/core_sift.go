// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package domain

var ops = []string{op_eq, op_ne, op_gt, op_lt, op_ge,
	op_le, op_li, op_lk, op_ll, op_lr,
	op_be, op_in, op_nu, op_nn, op_tr, op_oa, op_od}

var ots = []string{ot_st, ot_bl,
	ot_i0, ot_i8, ot_i1, ot_i3, ot_i6,
	ot_u0, ot_u8, ot_u1, ot_u3, ot_u6,
	ot_f3, ot_f6,
	ot_td, ot_tt, ot_tm, ot_t2, ot_t5, ot_t4}

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
