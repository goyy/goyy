// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package entity

import (
	"gopkg.in/goyy/goyy.v0/util/errors"
	"gopkg.in/goyy/goyy.v0/util/jsons"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

func FormatJSON(e Interface) string {
	return e.JSON()
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
