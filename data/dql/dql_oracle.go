// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dql

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"

	"gopkg.in/goyy/goyy.v0/data/domain"
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

type oracle struct {
}

func (me *oracle) SelectOne(e entity.Interface) (dql string, arg interface{}) {
	dqlf := "select * from %s where %s = :1"
	pk := e.Table().Primary().Name()
	dql = fmt.Sprintf(dqlf, e.Table().Name(), pk)
	arg = e.Get(pk)
	return
}

func (me *oracle) SelectPage(dql string, pageable domain.Pageable) string {
	dqlf := "select * from ( select row_.*, rownum rownum_ from (%s) row_ where rownum <= %d) where rownum_ > %d"
	return fmt.Sprintf(dqlf, dql, pageable.Offset()+pageable.PageSize(), pageable.Offset())
}

func (me *oracle) SelectListBySift(e entity.Interface, sifts ...domain.Sift) (dql string, args []interface{}, err error) {
	return me.selectBySift(e, "select * from ", sifts...)
}

func (me *oracle) SelectCountBySift(e entity.Interface, sifts ...domain.Sift) (dql string, args []interface{}, err error) {
	return me.selectBySift(e, "select count(1) from ", sifts...)
}

func (me *oracle) selectBySift(e entity.Interface, begin string, sifts ...domain.Sift) (dql string, args []interface{}, err error) {
	args = make([]interface{}, 0)
	w := bytes.NewBuffer([]byte(begin))
	w.WriteString(e.Table().Name() + " ")
	var o bytes.Buffer
	i := 0
	omap := make(map[string]domain.Sift)
	okeys := make([]string, 0)
	for _, v := range sifts {
		if v == nil || strings.IsBlank(v.Value()) {
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
			okeys = append(okeys, v.Value())
			omap[v.Value()] = v
		} else {
			if i == 0 {
				w.WriteString("where ")
			} else {
				w.WriteString("and ")
			}
			if v.Operator() == "NU" || v.Operator() == "NN" {
				w.WriteString(key + op[v.Operator()])
			} else {
				switch v.Operator() {
				case "IN", "NI":
					values := strings.Split(v.Value(), ",")
					var b bytes.Buffer
					for k := range values {
						if k > 0 {
							b.WriteString(",")
						}
						b.WriteString(":" + strconv.Itoa(i))
						i++
					}
					w.WriteString(key + op[v.Operator()] + "(" + b.String() + ") ")
				case "BE", "NB":
					values := strings.Split(v.Value(), ",")
					var b bytes.Buffer
					for k := range values {
						if k > 1 {
							break
						}
						if k > 0 {
							b.WriteString(" and ")
						}
						b.WriteString(":" + strconv.Itoa(i))
						i++
					}
					w.WriteString(key + op[v.Operator()] + b.String() + " ")
				default:
					w.WriteString(key + op[v.Operator()] + ":" + strconv.Itoa(i) + " ")
				}
				if typ, ok := e.Type(key); ok {
					val, err := toValue(v.Value(), typ.Name())
					if err != nil {
						return "", nil, err
					}
					switch v.Operator() {
					case "IN", "NI":
						values := strings.Split(v.Value(), ",")
						for _, inVal := range values {
							args = append(args, inVal)
						}
					case "BE", "NB":
						values := strings.Split(v.Value(), ",")
						for i, inVal := range values {
							if i > 1 {
								break
							}
							args = append(args, inVal)
						}
					default:
						args = append(args, val)
					}
				}
			}
			i++
		}
	}
	// Order By : The first order of the field is determined by the sift.value size
	y := 0
	if len(okeys) > 0 {
		sort.Strings(okeys)
		for _, k := range okeys {
			v := omap[k]
			field := strings.ToLowerFirst(v.Key())
			var key string
			if column, ok := e.Column(field); ok {
				key = column.Name()
			} else {
				continue
			}
			if y == 0 {
				o.WriteString("order by " + key + op[v.Operator()])
			} else {
				o.WriteString(", " + key + op[v.Operator()])
			}
			y++
		}
	}
	if y == 0 {
		return w.String(), args, nil
	}
	return w.String() + o.String(), args, nil
}
