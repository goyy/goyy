// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dml

import (
	"bytes"
	"fmt"
	"gopkg.in/goyy/goyy.v0/data/entity"
	"strconv"
)

type oracle struct {
}

func (me *oracle) Update(e entity.Interface) (dml string, args []interface{}) {
	dmlf := "update %s set %s where %s = :1"
	var b bytes.Buffer
	args = make([]interface{}, 0)
	i := 2
	for _, c := range e.Columns() {
		if t, ok := e.Type(c.Name()); ok {
			if t.HasUpdate() == false {
				continue
			}
		}
		if !c.IsPrimary() {
			if i > 2 {
				b.WriteString(",")
			}
			b.WriteString(c.Name() + " = :" + strconv.Itoa(i))
			args = append(args, e.Get(c.Name()))
			i++
		}
	}
	pk := e.Table().Primary().Name()
	dml = fmt.Sprintf(dmlf, e.Table().Name(), b.String(), pk)
	args = append(args, e.Get(pk))
	return
}

func (me *oracle) Insert(e entity.Interface) (dml string, args []interface{}) {
	dmlf := "insert into %s (%s) values (%s)"
	var bc bytes.Buffer
	var bv bytes.Buffer
	args = make([]interface{}, 0)
	i := 0
	for _, c := range e.Columns() {
		if t, ok := e.Type(c.Name()); ok {
			if t.HasInsert() == false {
				continue
			}
		}
		if i > 0 {
			bc.WriteString(",")
			bv.WriteString(",")
		}
		bc.WriteString(c.Name())
		bv.WriteString(":" + strconv.Itoa(i))
		args = append(args, e.Get(c.Name()))
		i++
	}
	dml = fmt.Sprintf(dmlf, e.Table().Name(), bc.String(), bv.String())
	return
}

func (me *oracle) Delete(e entity.Interface) (dml string, arg interface{}) {
	dmlf := "delete from %s where %s = :1"
	pk := e.Table().Primary().Name()
	dml = fmt.Sprintf(dmlf, e.Table().Name(), pk)
	arg = e.Get(pk)
	return
}

func (me *oracle) Disable(e entity.Interface) (dml string, arg interface{}) {
	dmlf := "update %s set %s = %v where %s = :1"
	pk := e.Table().Primary().Name()
	deletion := e.Table().Deletion().Name()
	dml = fmt.Sprintf(dmlf, e.Table().Name(), deletion, entity.DeletionDisable, pk)
	arg = e.Get(pk)
	return
}
