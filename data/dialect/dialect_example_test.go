// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dialect_test

import (
	"fmt"
	"gopkg.in/goyy/goyy.v0/data/dialect"
)

func ExampleType() {
	mysql := &dialect.MySQL{}
	fmt.Println(mysql.Type())

	// Output:MySQL
}
