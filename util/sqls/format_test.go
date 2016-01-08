// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package sqls_test

import (
	"testing"

	"gopkg.in/goyy/goyy.v0/util/sqls"
)

func TestFormatSpace(t *testing.T) {
	s := []struct {
		in, expected string
	}{
		{
			"select id,  name from   table  where   id = 1",
			"select id, name from table where id = 1",
		},
		{
			"select id,  name FROM   table  where   id  like '%a b c%'",
			"select id, name FROM table where id like '%a b c%'",
		},
		{
			"select id,  name FROM   table  where   id  like '''%a ''b c%'''",
			"select id, name FROM table where id like '''%a ''b c%'''",
		},
		{
			`select id,  name 
			 FROM   table  
			 where   id  like '%a b c%'`,
			"select id, name FROM table where id like '%a b c%'",
		},
	}
	for _, v := range s {
		if out := sqls.FormatSpace(v.in); out != v.expected {
			t.Errorf(`sqls.FormatSpace("%s") = "%s", want "%s"`, v.in, out, v.expected)
		}
	}
}
