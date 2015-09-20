// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package hmac

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

// Sha1Hex calculates the SHA-1 digest and returns the resulting hex string.
func Sha1Hex(src, key string) (string, error) {
	return Sha1SumHex(src, key, "")
}

// Sha1 calculates the SHA-1 digest and returns the value as a byte[].
func Sha1(src, key []byte) ([]byte, error) {
	return Sha1Sum(src, key, nil)
}

// Sha1SumHex calculates the SHA-1 digest to sum and returns the resulting hex string.
func Sha1SumHex(src, key, sum string) (string, error) {
	var sumtext []byte
	if strings.IsNotBlank(sum) {
		sumtext = []byte(sum)
	}
	out, err := Sha1Sum([]byte(src), []byte(key), sumtext)
	if err != nil {
		return "", err
	}
	dst := fmt.Sprintf("%x", out)
	return dst, nil
}

// Sha1Sum calculates the SHA-1 digest to sum and returns the resulting slice.
func Sha1Sum(src, key, sum []byte) ([]byte, error) {
	h := hmac.New(sha1.New, key)
	if _, err := h.Write(src); err != nil {
		return nil, err
	}
	return h.Sum(sum), nil
}
