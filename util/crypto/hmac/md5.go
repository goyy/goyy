// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package hmac

import (
	"crypto/hmac"
	"crypto/md5"
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

// Md5Hex calculates the MD5 digest and returns the resulting hex string.
func Md5Hex(src, key string) (string, error) {
	return Md5SumHex(src, key, "")
}

// Md5 calculates the MD5 digest and returns the value as a byte[].
func Md5(src, key []byte) ([]byte, error) {
	return Md5Sum(src, key, nil)
}

// Md5SumHex calculates the MD5 digest to sum and returns the resulting hex string.
func Md5SumHex(src, key, sum string) (string, error) {
	var sumtext []byte
	if strings.IsNotBlank(sum) {
		sumtext = []byte(sum)
	}
	out, err := Md5Sum([]byte(src), []byte(key), sumtext)
	if err != nil {
		return "", err
	}
	dst := fmt.Sprintf("%x", out)
	return dst, nil
}

// Md5Sum calculates the MD5 digest to sum and returns the resulting slice.
func Md5Sum(src, key, sum []byte) ([]byte, error) {
	h := hmac.New(md5.New, key)
	if _, err := h.Write(src); err != nil {
		return nil, err
	}
	return h.Sum(sum), nil
}
