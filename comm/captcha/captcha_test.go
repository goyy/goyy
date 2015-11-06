// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package captcha

import (
	"bytes"
	"testing"
)

func TestRandomDigits(t *testing.T) {
	d1 := RandomDigits(10)
	for _, v := range d1 {
		if v > 9 {
			t.Errorf("digits not in range 0-9: %v", d1)
		}
	}
	d2 := RandomDigits(10)
	if bytes.Equal(d1, d2) {
		t.Errorf("digits seem to be not random")
	}
}
