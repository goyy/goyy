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
	return b.String()
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
