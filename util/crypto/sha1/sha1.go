// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package sha1 implements the SHA1 hash algorithm as defined in RFC 3174.
package sha1

import (
	"crypto/sha1"
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

// DigestHex calculates the SHA-1 digest and returns the resulting hex string.
func DigestHex(src string) (string, error) {
	return DigestSumHex(src, "")
}

// Digest calculates the SHA-1 digest and returns the resulting slice.
func Digest(src []byte) ([]byte, error) {
	return DigestSum(src, nil)
}

// DigestSumHex calculates the SHA-1 digest to sum and returns the resulting hex string.
func DigestSumHex(src, sum string) (dst string, err error) {
	var key []byte
	var out []byte
	if strings.IsNotBlank(sum) {
		key = []byte(sum)
	}
	out, err = DigestSum([]byte(src), key)
	if err != nil {
		return
	}
	dst = fmt.Sprintf("%x", out)
	return
}

// DigestSum calculates the SHA-1 digest to sum and returns the resulting slice.
func DigestSum(src, sum []byte) (dst []byte, err error) {
	h := sha1.New()
	if _, err = h.Write(src); err != nil {
		return
	}
	dst = h.Sum(sum)
	return
}
