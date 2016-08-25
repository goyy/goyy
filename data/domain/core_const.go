// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package domain

const (
	defaultPageNo       = 1
	defaultPageSize     = 10
	defaultPageNoName   = "sPageNoTR"
	defaultPageSizeName = "sPageSizeTR"
	defaultPageSizeMax  = 500
	defaultRePage       = "sRePageTR"
	defaultRePageNo     = "sRePageNoTR"
	defaultPageLength   = 8
	defaultPageSlider   = 1
	defaultPageFn       = "page"
	defaultSiftPrefix   = "s"

	op_eq = "EQ" // equal( = )
	op_ne = "NE" // not equal( <> )
	op_gt = "GT" // greater than( > )
	op_lt = "LT" // less than( < )
	op_ge = "GE" // greater than or equal( >= )
	op_le = "LE" // less than or equal( <= )
	op_li = "LI" // like '?'
	op_lk = "LK" // like '%?%'
	op_ll = "LL" // like '%?'
	op_lr = "LR" // like '?%'
	op_be = "BE" // between 'a,b'
	op_in = "IN" // in 'a,b,c,d'
	op_nu = "NU" // is null
	op_nn = "NN" // is not null
	op_tr = "TR" // transient
	op_oa = "OA" // order by asc
	op_od = "OD" // order by desc

	ot_st = "ST" // string
	ot_bl = "BL" // bool
	ot_i0 = "I0" // int
	ot_i8 = "I8" // int8
	ot_i1 = "I1" // int16
	ot_i3 = "I3" // int32
	ot_i6 = "I6" // int64
	ot_u0 = "U0" // uint
	ot_u8 = "U8" // uint8
	ot_u1 = "U1" // uint16
	ot_u3 = "U3" // uint32
	ot_u6 = "U6" // uint64
	ot_f3 = "F3" // float32
	ot_f6 = "F6" // float64
	ot_td = "TD" // time.Time:2006-01-02
	ot_tt = "TT" // time.Time:2006-01-02 15:04:05
	ot_tm = "TM" // time.Time:2006-01-02 15:04
	ot_t2 = "T2" // time.Time.Unix():2006-01-02
	ot_t5 = "T5" // time.Time.Unix():2006-01-02 15:04:05
	ot_t4 = "T4" // time.Time.Unix():2006-01-02 15:04
)
