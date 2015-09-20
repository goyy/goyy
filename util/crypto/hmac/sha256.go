// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package hmac

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

// Sha256Hex calculates the SHA-256 digest and returns the resulting hex string.
func Sha256Hex(src, key string) (string, error) {
	return Sha256SumHex(src, key, "")
}

// Sha256 calculates the SHA-256 digest and returns the value as a byte[].
func Sha256(src, key []byte) ([]byte, error) {
	return Sha256Sum(src, key, nil)
}

// Sha256SumHex calculates the SHA-256 digest to sum and returns the resulting hex string.
func Sha256SumHex(src, key, sum string) (string, error) {
	var sumtext []byte
	if strings.IsNotBlank(sum) {
		sumtext = []byte(sum)
	}
	out, err := Sha256Sum([]byte(src), []byte(key), sumtext)
	if err != nil {
		return "", err
	}
	dst := fmt.Sprintf("%x", out)
	return dst, nil
}

// Sha256Sum calculates the SHA-256 digest to sum and returns the resulting slice.
func Sha256Sum(src, key, sum []byte) ([]byte, error) {
	h := hmac.New(sha256.New, key)
	if _, err := h.Write(src); err != nil {
		return nil, err
	}
	return h.Sum(sum), nil
}
