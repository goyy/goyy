// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dialect_test

import (
	"gopkg.in/goyy/goyy.v0/data/dialect"
	"testing"
)

func TestType(t *testing.T) {
	expected := dialect.MYSQL
	mysql := &dialect.MySQL{}
	if out := mysql.Type(); out != expected {
		t.Errorf(`mysql.Type() = "%v", want "%v"`, out, expected)
	}
}
