// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package captcha

const (
	key       = "captcha"
	stdWidth  = 100
	stdHeight = 40
	maxSkew   = 3
)

const (
	fontWidth  = 5
	fontHeight = 8
	blackChar  = 1
)

const (
	// Standard length of uniuri string to achive ~95 bits of entropy.
	StdLen = 16
	// Length of uniurl string to achive ~119 bits of entropy, closest
	// to what can be losslessly converted to UUIDv4 (122 bits).
	UUIDLen = 20
)
