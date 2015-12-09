// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package db

import (
	"database/sql"
	"errors"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"log"
	"reflect"
	"strconv"
)

// add space string:{before is add:yes, after is add:no}
type query struct {
	Db       *sql.DB
	Dialect  Dialect
	qSelect  string
	qFrom    string
	qWhere   string
	qGroupBy string
	qHaving  string
	qOrderBy string
	qLimit   string
	qOffset  string
	args     []interface{}
}

// Select Query
func (me *query) Select(es ...Expression) Query {
	size := len(es)
	for i, e := range es {
		if e == "" {
			me.qSelect += " *"
		} else {
			me.qSelect += " " + e.String()
		}
		if i < size-1 {
			me.qSelect += ","
		}
	}
	return me
}

// From Query
func (me *query) From(t Table) Query {
	if t == "" {
		log.Fatal(errors.New(i18N.Message("err.empty.table.name")))
	}
	me.qFrom = " " + t.String()
	return me
}

// Where Query
func (me *query) Where(c Condition) Query {
	me.qWhere += " " + c.String()
	me.args = c.Args()
	return me
}

// GroupBy Query
func (me *query) GroupBy(es ...Expression) Query {
	size := len(es)
	for i, e := range es {
		if e == "" {
			continue
		} else {
			me.qGroupBy += " " + string(e)
		}
		if i < size-1 {
			me.qGroupBy += ","
		}
	}
	return me
}

// Having Query
func (me *query) Having(c Condition) Query {
	me.qHaving += " " + c.String()
	return me
}

// OrderBy Query
func (me *query) OrderBy(es ...Expression) Query {
	size := len(es)
	for i, e := range es {
		if e == "" {
			continue
		} else {
			me.qOrderBy += " " + e.String()
		}
		if i < size-1 {
			me.qOrderBy += ","
		}
	}
	return me
}

// Limit Query
func (me *query) Limit(l uint8) Query {
	me.qLimit = " " + strconv.Itoa(int(l))
	return me
}

// Offset Query
func (me *query) Offset(o uint8) Query {
	me.qOffset = " " + strconv.Itoa(int(o))
	return me
}

// Query List Data
func (me *query) List(out interface{}) error {
	if x := reflect.TypeOf(out).Kind(); x != reflect.Ptr {
		i18N.Panic("err.mapping.pointer")
	}
	sliceValue := reflect.Indirect(reflect.ValueOf(out))
	if x := sliceValue.Kind(); x != reflect.Slice {
		i18N.Panic("err.mapping.slice")
	}
	sliceType := sliceValue.Type().Elem()
	if x := sliceType.Kind(); x != reflect.Struct {
		i18N.Panic("err.mapping.struct")
	}
	stmt, err := me.Db.Prepare(me.String())
	if err != nil {
		return err
	}
	defer stmt.Close()
	rows, err := stmt.Query(me.args...)
	if err != nil {
		return err
	}
	defer rows.Close()
	cols, err := rows.Columns()
	if err != nil {
		return err
	}
	for rows.Next() {
		containers := make([]interface{}, 0, len(cols))
		for i := 0; i < cap(containers); i++ {
			var v interface{}
			containers = append(containers, &v)
		}
		err := rows.Scan(containers...)
		if err != nil {
			return err
		}
		// create a new row and fill
		rowValue := reflect.New(sliceType)
		for i, v := range containers {
			key := cols[i]
			value := reflect.Indirect(reflect.ValueOf(v))
			name := strings.Camel(key)
			field := rowValue.Elem().FieldByName(name)
			if field.IsValid() {
				vfield := field.FieldByName("Value")
				if vfield.IsValid() {
					err = me.Dialect.SetModelValue(value, vfield)
					if err != nil {
						return err
					}
				} else {
					err = me.Dialect.SetModelValue(value, field)
					if err != nil {
						return err
					}
				}
			}
		}
		// append to output
		sliceValue.Set(reflect.Append(sliceValue, rowValue.Elem()))
	}
	return nil
}

// Query One Data
func (me *query) One(out interface{}) error {
	if x := reflect.TypeOf(out).Kind(); x != reflect.Ptr {
		i18N.Panic("err.mapping.pointer")
	}
	sliceValue := reflect.Indirect(reflect.ValueOf(out))
	sliceType := sliceValue.Type()
	if x := sliceType.Kind(); x != reflect.Struct {
		i18N.Panic("err.mapping.struct")
	}
	stmt, err := me.Db.Prepare(me.String())
	if err != nil {
		return err
	}
	defer stmt.Close()
	rows, err := stmt.Query(me.args...)
	if err != nil {
		return err
	}
	defer rows.Close()
	cols, err := rows.Columns()
	if err != nil {
		return err
	}
	for rows.Next() {
		containers := make([]interface{}, 0, len(cols))
		for i := 0; i < cap(containers); i++ {
			var v interface{}
			containers = append(containers, &v)
		}
		err := rows.Scan(containers...)
		if err != nil {
			return err
		}
		// create a new row and fill
		rowValue := reflect.New(sliceType)
		for i, v := range containers {
			key := cols[i]
			value := reflect.Indirect(reflect.ValueOf(v))
			name := strings.Camel(key)
			field := rowValue.Elem().FieldByName(name)
			if field.IsValid() {
				vfield := field.FieldByName("Value")
				if vfield.IsValid() {
					err = me.Dialect.SetModelValue(value, vfield)
					if err != nil {
						return err
					}
				} else {
					err = me.Dialect.SetModelValue(value, field)
					if err != nil {
						return err
					}
				}
			}
		}
		sliceValue.Set(rowValue.Elem())
		return nil
	}
	return nil
}

// Query Reset
func (me *query) Reset() {
	me.qSelect = ""
	me.qFrom = ""
	me.qWhere = ""
	me.qGroupBy = ""
	me.qHaving = ""
	me.qOrderBy = ""
	me.qLimit = ""
	me.qOffset = ""
}

// Query To String
func (me *query) String() string {
	var result string
	if strings.IsBlank(me.qSelect) {
		result += "select *"
	} else {
		result += "select" + me.qSelect
	}
	if strings.IsBlank(me.qFrom) {
		log.Fatal(errors.New(i18N.Message("err.empty.table.name")))
	} else {
		result += " from" + me.qFrom
	}
	if strings.IsNotBlank(me.qWhere) {
		result += " where" + me.qWhere
	}
	if strings.IsNotBlank(me.qGroupBy) {
		result += " group by" + me.qGroupBy
		if strings.IsNotBlank(me.qHaving) {
			result += " having" + me.qHaving
		}
	}
	if strings.IsNotBlank(me.qOrderBy) {
		result += " order by" + me.qOrderBy
	}
	if strings.IsNotBlank(me.qLimit) {
		result += " limit" + me.qLimit
		if strings.IsNotBlank(me.qOffset) {
			result += " offset" + me.qOffset
		}
	}
	return result
}
