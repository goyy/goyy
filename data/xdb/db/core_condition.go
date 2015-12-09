// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package db

import (
	"errors"
	"log"
	"reflect"
)

type condition struct {
	Column   Column
	Operator string
	Operand  interface{}
	Operand2 interface{}

	Andor string
	Adds  []*condition
}

// And Condition
func (me *condition) And(e Condition) Condition {
	return me.andor("and", e)
}

// Or Condition
func (me *condition) Or(e Condition) Condition {
	return me.andor("or", e)
}

// Get Args
func (me *condition) Args() (args []interface{}) {
	args = append(args, me.Operand)
	if me.Operator == o_be {
		args = append(args, me.Operand2)
	}
	if len(me.Adds) > 0 {
		for _, add := range me.Adds {
			args = append(args, add.Args()...)
		}
	}
	return
}

// To String
func (me *condition) String() string {
	if me.Column == "" {
		log.Fatal(errors.New(i18N.Message("err.empty.column.name")))
	}
	var result, end, operator string
	switch me.Operator {
	case o_eq:
		operator = " = ?"
	case o_ne:
		operator = " <> ?"
	case o_gt:
		operator = " > ?"
	case o_lt:
		operator = " < ?"
	case o_ge:
		operator = " >= ?"
	case o_le:
		operator = " <= ?"
	case o_li, o_lk, o_ll, o_lr:
		operator = " like ?"
	case o_be:
		operator = " between ? and ?"
	case o_in:
		operator = " in ?"
	case o_nu:
		operator = " is null"
	case o_nn:
		operator = " is not null"
	default:
		operator = " = ?"
	}
	if len(me.Adds) == 0 {
		result = me.Column.String() + operator
	} else {
		if me.Andor == "" { // root
			result = me.Column.String() + operator
		} else {
			result = "(" + me.Column.String() + operator
			end = ")"
		}
		for _, add := range me.Adds {
			result += " " + add.Andor + " " + add.String()
		}
		result += end
	}
	return result
}

// And/Or Condition
func (me *condition) andor(andor string, e Condition) Condition {
	if t, ok := e.(*condition); ok {
		t.Andor = andor
		me.Adds = append(me.Adds, t)
	} else {
		log.Fatal(errors.New(i18N.Messagef("err.assign.exp", reflect.TypeOf(e))))
	}
	return me
}
