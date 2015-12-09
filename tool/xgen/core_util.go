// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/axgle/mahonia"
	"gopkg.in/goyy/goyy.v0/util/errors"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

func tagItemValue(tag, name string) string {
	return strings.Between(tag, name+`:"`, `"`)
}

func tagAttrValue(tag, name string) (string, error) {
	attributes := strings.Split(tag, "&")
	for _, attr := range attributes {
		pair := strings.Split(attr, "=")
		if len(pair) != 2 {
			return "", errors.Newf("Malformed tag: '%s'", attr)
		}
		switch strings.ToLower(pair[0]) {
		case name:
			return pair[1], nil
		}
	}
	return "", errors.Newf("%s was not found in tag", name)
}

func convertUTF8(in string) string {
	enc := mahonia.NewEncoder("gbk")
	return enc.ConvertString(in)
}
