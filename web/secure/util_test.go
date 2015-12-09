// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package secure_test

import (
	"gopkg.in/goyy/goyy.v0/web/secure"
	"testing"
)

func TestMethod(t *testing.T) {
	expected := "92d55a4a6b07"
	in := "111111"
	if out := secure.EncryptPasswd(in); out != expected {
		t.Errorf(`secure.EncryptPasswd(%s) = "%v", want "%v"`, in, out, expected)
	}
}
