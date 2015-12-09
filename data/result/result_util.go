// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package result

import (
	"gopkg.in/goyy/goyy.v0/util/jsons"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

func jreplace(json string) string {
	replace := "&-q-u-o-t-;"
	return strings.Replace(json, `\"`, replace, -1)
}

func jparse(json string) string {
	replace := "&-q-u-o-t-;"
	content := strings.Replace(json, replace, `\"`, -1)
	return jsons.Parse(content)
}
