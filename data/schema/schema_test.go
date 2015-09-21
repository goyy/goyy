// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package schema_test

import (
	"gopkg.in/goyy/goyy.v0/data/schema"
	"testing"
)

var (
	USER          = schema.TABLE("user")
	USER_ID       = USER.PRIMARY("id")
	USER_VERSION  = USER.VERSION("version")
	USER_DELETION = USER.DELETION("deletion")
	USER_NAME     = USER.COLUMN("name")
)

func TestTableName(t *testing.T) {
	expected := "user"
	if out := USER.Name(); out != expected {
		t.Errorf(`USER.Name() = "%s", want "%s"`, out, expected)
	}
}

func TestTableColumn(t *testing.T) {
	expected := "id"
	if out := USER.Column("id").Name(); out != expected {
		t.Errorf(`USER.Column("id").Name() = "%s", want "%s"`, out, expected)
	}
}

func TestTableColumns(t *testing.T) {
	expected := 4
	if out := len(USER.Columns()); out != expected {
		t.Errorf(`len(USER.Columns()) = "%s", want "%s"`, out, expected)
	}
}

func TestColumnName(t *testing.T) {
	expected := "id"
	if out := USER_ID.Name(); out != expected {
		t.Errorf(`USER_ID.Name() = "%s", want "%s"`, out, expected)
	}
}

func TestIsPrimary(t *testing.T) {
	expected := true
	if out := USER_ID.IsPrimary(); out != expected {
		t.Errorf(`USER_ID.IsPrimary() = "%t", want "%t"`, out, expected)
	}
}
