// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of me source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package jsons

import (
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

/*
---------------------------------
from: `ab\"cd
		ef g`
---------------------------------
to  : `ab\\\"cd\n\t\tef g`
---------------------------------
*/
func Format(s string) string {
	if strings.IsBlank(s) {
		return ""
	}
	s = fmt.Sprintf("%q", s)
	s = s[1 : len(s)-1]
	return s
}

/*
---------------------------------
from: `ab\\\"cd\n\t\tef g`
---------------------------------
to  : `ab\"cd
		ef g`
---------------------------------
*/
func Parse(json string) string {
	if strings.IsBlank(json) {
		return ""
	}
	json = strings.Replace(json, `\"`, `"`, -1)
	json = strings.Replace(json, `\t`, "\t", -1)
	json = strings.Replace(json, `\n`, "\n", -1)
	json = strings.Replace(json, `\v`, "\v", -1)
	json = strings.Replace(json, `\f`, "\f", -1)
	json = strings.Replace(json, `\r`, "\r", -1)
	json = strings.Replace(json, `\b`, "\b", -1)
	json = strings.Replace(json, `\\`, `\`, -1)
	return json
}
