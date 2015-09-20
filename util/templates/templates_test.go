// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package templates_test

import (
	"gopkg.in/goyy/goyy.v0/util/templates"
	"testing"
)

func TestProcess(t *testing.T) {
	data := map[string]interface{}{"name": "goyy", "updatedAt": "2014-03-19"}
	expected := "select * from demo where 1=1 and name like :name and updated_at < :updatedAt order by id"
	if out, _ := templates.Process(sql, data); out != expected {
		t.Errorf(`templates.Parse:out = "%s", want "%s"`, out, expected)
	}
}
