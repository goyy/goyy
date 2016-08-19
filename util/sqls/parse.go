// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package sqls

import (
	"bytes"
	"fmt"
	"text/template"

	"gopkg.in/goyy/goyy.v0/comm/xtype"
	"gopkg.in/goyy/goyy.v0/data/dialect"
	"gopkg.in/goyy/goyy.v0/util/errors"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"gopkg.in/goyy/goyy.v0/util/templates"
)

// select ... from ... -> select count(1) from ...
func ParseCountSql(sql string) string {
	stack := &xtype.Stack{}
	ss := strings.Split(sql, " ")
	p := 0
	for _, v := range ss {
		if strings.Contains(strings.ToLower(v), "select") {
			p++
			stack.Push(p)
			continue
		}
		if strings.Contains(strings.ToLower(v), "from") {
			if stack.Len() == 1 {
				break
			}
			stack.Pop()
			continue
		}
	}
	pfrom := strings.IndexOrdinal(strings.ToLower(sql), " from ", p)
	return "select count(1)" + sql[pfrom:]
}

// ParseNamedSql takes a query using named parameters and an argument and
// returns a new query with a list of args that can be executed by a database.
func ParseNamedSql(dia dialect.Interface, sql string, args map[string]interface{}) (sqlout string, argsout []interface{}, err error) {
	if dia == nil || strings.IsBlank(sql) || args == nil {
		err = errors.NewNotBlank("dia/sql/args")
		return
	}
	if !strings.Contains(sql, "#{") {
		sqlout = sql
		argsout = make([]interface{}, 0)
		return
	}
	sqls := strings.Betweens(sql, "#{", "}")
	if sqls != nil && len(sqls) > 0 {
		i := 0
		for _, v := range sqls {
			if strings.IsNotBlank(v) {
				if dia.Type() == dialect.ORACLE {
					sql = strings.Replace(sql, "#{"+v+"}", fmt.Sprintf(":%d", i), -1)
					i++
				} else {
					sql = strings.Replace(sql, "#{"+v+"}", "?", -1)
				}
				if _, ok := args[v]; ok {
					argsout = append(argsout, args[v])
				} else {
					err = errors.NewNotBlank("map[" + v + "]")
					sqlout = sql
					return
				}
			}
		}
		sqlout = sql
	}
	return
}

// ParseTemplateSql takes a query using named parameters and an argument and
// returns a new query with a list of args that can be executed by a database.
func ParseTemplateSql(sql string, args map[string]interface{}) (out string, err error) {
	t, err := template.New("sqls-tmpl").Funcs(templates.Text.FuncMap).Parse(sql)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	var v bytes.Buffer
	err = t.Execute(&v, args)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	out = v.String()
	return
}
