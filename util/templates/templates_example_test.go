// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package templates_test

import (
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/templates"
)

func ExampleProcess() {
	data := map[string]interface{}{"name": "goyy", "updatedAt": "2014-03-19"}
	out, _ := templates.Process(sql, data)
	fmt.Println(out)

	// Output: select * from demo where 1=1 and name like :name and updated_at < :updatedAt order by id
}
