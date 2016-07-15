// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
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
	return encode(s)
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

func encode(src string) string {
	var encodeString = []struct {
		in  string
		out string
	}{
		{`\x00`, `\u0000`},
		{`\x01`, `\u0001`},
		{`\x02`, `\u0002`},
		{`\x03`, `\u0003`},
		{`\x04`, `\u0004`},
		{`\x05`, `\u0005`},
		{`\x06`, `\u0006`},
		{`\x07`, `\u0007`},
		{`\x08`, `\u0008`},
		{`\x09`, `\t`},
		{`\x0a`, `\n`},
		{`\x0b`, `\u000b`},
		{`\x0c`, `\u000c`},
		{`\x0d`, `\r`},
		{`\x0e`, `\u000e`},
		{`\x0f`, `\u000f`},
		{`\x10`, `\u0010`},
		{`\x11`, `\u0011`},
		{`\x12`, `\u0012`},
		{`\x13`, `\u0013`},
		{`\x14`, `\u0014`},
		{`\x15`, `\u0015`},
		{`\x16`, `\u0016`},
		{`\x17`, `\u0017`},
		{`\x18`, `\u0018`},
		{`\x19`, `\u0019`},
		{`\x1a`, `\u001a`},
		{`\x1b`, `\u001b`},
		{`\x1c`, `\u001c`},
		{`\x1d`, `\u001d`},
		{`\x1e`, `\u001e`},
		{`\x1f`, `\u001f`},
	}
	for _, v := range encodeString {
		src = strings.Replace(src, v.in, v.out, -1)
	}
	return src
}
