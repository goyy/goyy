// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package db

// equal( = )
func (me Column) Eq(value interface{}) Condition {
	return &condition{Column: me, Operator: o_eq, Operand: value}
}

// not equal( <> )
func (me Column) Ne(value interface{}) Condition {
	return &condition{Column: me, Operator: o_ne, Operand: value}
}

// greater than( > )
func (me Column) Gt(value interface{}) Condition {
	return &condition{Column: me, Operator: o_gt, Operand: value}
}

// less than( < )
func (me Column) Lt(value interface{}) Condition {
	return &condition{Column: me, Operator: o_lt, Operand: value}
}

// greater than or equal( >= )
func (me Column) Ge(value interface{}) Condition {
	return &condition{Column: me, Operator: o_ge, Operand: value}
}

// less than or equal( <= )
func (me Column) Le(value interface{}) Condition {
	return &condition{Column: me, Operator: o_le, Operand: value}
}

// like ?
func (me Column) Li(value interface{}) Condition {
	return &condition{Column: me, Operator: o_li, Operand: value}
}

// like %?%
func (me Column) Lk(value interface{}) Condition {
	if value.(string) != "" {
		value = "%" + value.(string) + "%"
	}
	return &condition{Column: me, Operator: o_lk, Operand: value}
}

// like %?
func (me Column) Ll(value interface{}) Condition {
	if value.(string) != "" {
		value = "%" + value.(string)
	}
	return &condition{Column: me, Operator: o_ll, Operand: value}
}

// like ?%
func (me Column) Lr(value interface{}) Condition {
	if value.(string) != "" {
		value = value.(string) + "%"
	}
	return &condition{Column: me, Operator: o_lr, Operand: value}
}

// between
func (me Column) Be(lo interface{}, hi interface{}) Condition {
	return &condition{Column: me, Operator: o_be, Operand: lo, Operand2: hi}
}

// in
func (me Column) In(values ...interface{}) Condition {
	return &condition{Column: me, Operator: o_in, Operand: values}
}

// is null
func (me Column) Nu() Condition {
	return &condition{Column: me, Operator: o_nu}
}

// is not null
func (me Column) Nn() Condition {
	return &condition{Column: me, Operator: o_nn}
}

// order by asc
func (me Column) Asc() Expression {
	return Expression(me + " asc")
}

// order by desc
func (me Column) Desc() Expression {
	return Expression(me + " desc")
}

// count
func (me Column) Count() Expression {
	return Expression("count(" + me + ")")
}

// max
func (me Column) Max() Expression {
	return Expression("max(" + me + ")")
}

// min
func (me Column) Min() Expression {
	return Expression("min(" + me + ")")
}

// sum
func (me Column) Sum() Expression {
	return Expression("sum(" + me + ")")
}

// avg
func (me Column) Avg() Expression {
	return Expression("avg(" + me + ")")
}

// To String
func (me Column) String() string {
	return string(me)
}
