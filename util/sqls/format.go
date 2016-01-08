// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package sqls

import (
	"bytes"
	"unicode"

	"gopkg.in/goyy/goyy.v0/util/strings"
)

// FormatSpace remove superfluous spaces in the SQL statement
func FormatSpace(sql string) string {
	if strings.IsBlank(sql) {
		return ""
	}
	var singleQuotes = '\''
	var b bytes.Buffer
	var moreSpace bool
	var singleQuotesStarting bool
	//var sqlEscapeChar bool
	var sqls []rune
	for _, r := range sql {
		sqls = append(sqls, r)
	}
	sqlsLen := len(sqls)
	for i, r := range sql {
		if r == singleQuotes { // Single quote in the content is not formatted
			if singleQuotesStarting {
				// To determine whether the SQL escape character
				if i+1 >= sqlsLen || sqls[i+1] != singleQuotes {
					singleQuotesStarting = false
				}
			} else {
				singleQuotesStarting = true
			}
		}
		if !singleQuotesStarting && unicode.IsSpace(r) {
			if moreSpace {
				continue
			}
			moreSpace = true
			b.WriteRune(' ')
			continue
		}
		if moreSpace {
			moreSpace = false
		}
		b.WriteRune(r)
	}
	return b.String()
}
