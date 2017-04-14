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

	opEQ = "EQ" // equal( = )
	opNE = "NE" // not equal( <> )
	opGT = "GT" // greater than( > )
	opLT = "LT" // less than( < )
	opGE = "GE" // greater than or equal( >= )
	opLE = "LE" // less than or equal( <= )
	opLI = "LI" // like '?'
	opLK = "LK" // like '%?%'
	opLL = "LL" // like '%?'
	opLR = "LR" // like '?%'
	opBE = "BE" // between 'a,b'
	opNB = "NB" // not between 'a,b'
	opIN = "IN" // in 'a,b,c,d'
	opNI = "NI" // not in 'a,b,c,d'
	opNU = "NU" // is null
	opNN = "NN" // is not null
	opTR = "TR" // transient
	opOA = "OA" // order by asc
	opOD = "OD" // order by desc

	otST = "ST" // string
	otBL = "BL" // bool
	otI0 = "I0" // int
	otI8 = "I8" // int8
	otI1 = "I1" // int16
	otI3 = "I3" // int32
	otI6 = "I6" // int64
	otU0 = "U0" // uint
	otU8 = "U8" // uint8
	otU1 = "U1" // uint16
	otU3 = "U3" // uint32
	otU6 = "U6" // uint64
	otF3 = "F3" // float32
	otF6 = "F6" // float64
	otTD = "TD" // time.Time:2006-01-02
	otTT = "TT" // time.Time:2006-01-02 15:04:05
	otTM = "TM" // time.Time:2006-01-02 15:04
	otT2 = "T2" // time.Time.Unix():2006-01-02
	otT5 = "T5" // time.Time.Unix():2006-01-02 15:04:05
	otT4 = "T4" // time.Time.Unix():2006-01-02 15:04
)
