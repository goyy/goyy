// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package md5 implements the MD5 hash algorithm as defined in RFC 1321.
package md5

import (
	"crypto/md5"
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

// DigestHex calculates the MD5 digest and returns the resulting hex string.
func DigestHex(src string) (string, error) {
	return DigestSumHex(src, "")
}

// Digest calculates the MD5 digest and returns the resulting slice.
func Digest(src []byte) ([]byte, error) {
	return DigestSum(src, nil)
}

// DigestSumHex calculates the MD5 digest to sum and returns the resulting hex string.
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

// DigestSum calculates the MD5 digest to sum and returns the resulting slice.
func DigestSum(src, sum []byte) (dst []byte, err error) {
	h := md5.New()
	if _, err = h.Write(src); err != nil {
		return
	}
	dst = h.Sum(sum)
	return
}
