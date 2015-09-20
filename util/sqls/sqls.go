// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package sqls

import (
	"gopkg.in/goyy/goyy.v0/comm/xtype"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

// select ... from ... -> select count(*) from ...
func ParseCountSql(sql string) string {
	intHead := &xtype.IntHeap{}
	ss := strings.Split(sql, " ")
	p := 0
	for _, v := range ss {
		if strings.Contains(strings.ToLower(v), "select") {
			p = p + 1
			intHead.Push(p)
			continue
		}
		if strings.Contains(strings.ToLower(v), "from") {
			if intHead.Len() == 1 {
				break
			}
			intHead.Pop()
			continue
		}
	}
	pfrom := strings.IndexOrdinal(strings.ToLower(sql), "from", p)
	return "select count(*) " + sql[pfrom:]
}
