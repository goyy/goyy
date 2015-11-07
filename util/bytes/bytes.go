// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package bytes implements functions for the manipulation of byte slices.
// It is analogous to the facilities of the strings package.
package bytes

import (
	"bytes"
)

func TrimRightNul(src []byte) []byte {
	return bytes.TrimRight(src, "\x00")
}
