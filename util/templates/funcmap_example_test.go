// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package templates_test

import (
	"bytes"
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/templates"
)

func ExampleExist() {
	buf := new(bytes.Buffer)
	data := map[string]interface{}{"name": "goyy"}
	tmpl, _ := templates.New("tmpl").Parse(text)
	tmpl.Execute(buf, data)
	fmt.Println(buf.String())

	// Output: true
}
