// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xhttp_test

import (
	"testing"
)

func TestMethod(t *testing.T) {
	expected := true
	type httpMethod string
	httpMethodGet1 := httpMethod("GET")
	httpMethodGet2 := httpMethod("GET")
	if out := (httpMethodGet1 == httpMethodGet2); out != expected {
		t.Errorf(`(httpMethod("GET") == httpMethod("GET")) = "%v", want "%v"`, out, expected)
	}
}
