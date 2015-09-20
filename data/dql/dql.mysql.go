// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dql

import (
	"bytes"
	"fmt"
	"gopkg.in/goyy/goyy.v0/data/domain"
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

type mysql struct {
}

func (me *mysql) SelectOne(e entity.Interface) (dql string, arg interface{}) {
	dqlf := "select * from %s where %s = ?"
	pk := e.Table().Primary().Name()
	dql = fmt.Sprintf(dqlf, e.Table().Name(), pk)
	arg = e.Get(pk)
	return
}

func (me *mysql) SelectPage(dql string, pageable domain.Pageable) string {
	dqlf := "%slimit %d,%d "
	return fmt.Sprintf(dqlf, dql, pageable.Offset(), pageable.PageSize())
}

func (me *mysql) SelectListBySift(e entity.Interface, sifts ...domain.Sift) (dql string, args []interface{}, err error) {
	op := map[string]string{
		"EQ": " = ",
		"NE": " <> ",
		"GT": " > ",
		"LT": " < ",
		"GE": " >= ",
		"LE": " <= ",
		"LI": " like ",
		"LK": " like ",
		"LL": " like ",
		"LR": " like ",
		"BE": " between ",
		"IN": " in ",
		"NU": " is null ",
		"NN": " is not null ",
		"OA": " asc ",
		"OD": " desc ",
	}
	args = make([]interface{}, 0)
	w := bytes.NewBuffer([]byte("select * from "))
	w.WriteString(e.Table().Name() + " ")
	var o bytes.Buffer
	i := 0
	y := 0
	for _, v := range sifts {
		if strings.IsBlank(v.Value()) {
			continue
		}
		field := strings.ToLowerFirst(v.Key())
		var key string
		if column, ok := e.Column(field); ok {
			key = column.Name()
		} else {
			continue
		}
		if v.Operator() == "OA" || v.Operator() == "OD" {
			if y == 0 {
				o.WriteString("order by " + key + op[v.Operator()])
			} else {
				o.WriteString("," + key + op[v.Operator()])
			}
			y++
		} else {
			if i == 0 {
				w.WriteString("where ")
			} else {
				w.WriteString("and ")
			}
			if v.Operator() == "NU" || v.Operator() == "NN" {
				w.WriteString(key + op[v.Operator()])
			} else {
				w.WriteString(key + op[v.Operator()] + "? ")
				if typ, ok := e.Type(key); ok {
					val, err := toValue(v.Value(), typ.Name())
					if err != nil {
						return "", nil, err
					}
					args = append(args, val)
				}
			}
			i++
		}
	}
	if y == 0 {
		return w.String(), args, nil
	} else {
		return w.String() + o.String(), args, nil
	}
}
