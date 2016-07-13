// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package entity

import (
	"bytes"

	"gopkg.in/goyy/goyy.v0/util/errors"
	"gopkg.in/goyy/goyy.v0/util/jsons"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

func FormatJSON(e Interface) string {
	var b bytes.Buffer
	b.WriteString("{")
	i := 0
	for _, f := range e.Names() {
		if col, ok := e.Column(f); ok {
			if typ, ok := e.Type(col.Name()); ok {
				if i > 0 {
					b.WriteString(",")
				}
				v := jsons.Format(typ.String())
				switch typ.Name() {
				case "string", "time":
					b.WriteString(`"` + f + `":"` + v + `"`)
				default:
					b.WriteString(`"` + f + `":` + v)
				}
				i++
			}
		}
	}
	b.WriteString(`}`)
	return jencode(b.String())
}

func ParseJSON(out Interface, json string) error {
	// JSON format validation
	if strings.IsBlank(json) {
		return nil
	}
	if !strings.HasPrefix(json, `{"`) {
		return errors.New("JSON format is not legal : The beginning must be {")
	}
	if !strings.HasSuffix(json, "}") {
		return errors.New("JSON format is not legal : The end must be }")
	}
	content := jreplace(json[:len(json)-1] + ",}")
	for _, f := range out.Names() {
		if col, ok := out.Column(f); ok {
			if typ, ok := out.Type(col.Name()); ok {
				switch typ.Name() {
				case "string", "time":
					v := strings.Between(content, `"`+f+`":"`, `",`)
					err := out.SetString(f, jparse(v))
					if err != nil {
						return err
					}
				default:
					v := strings.Between(content, `"`+f+`":`, `,`)
					err := out.SetString(f, jparse(v))
					if err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}

func jreplace(json string) string {
	replace := "&-q-u-o-t-;"
	return strings.Replace(json, `\"`, replace, -1)
}

func jparse(json string) string {
	replace := "&-q-u-o-t-;"
	content := strings.Replace(json, replace, `\"`, -1)
	return jsons.Parse(content)
}

func jencode(src string) string {
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
