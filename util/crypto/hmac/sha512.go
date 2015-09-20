// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package hmac

import (
	"crypto/hmac"
	"crypto/sha512"
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

// Sha512Hex calculates the SHA-512 digest and returns the resulting hex string.
func Sha512Hex(src, key string) (string, error) {
	return Sha512SumHex(src, key, "")
}

// Sha512 calculates the SHA-512 digest and returns the value as a byte[].
func Sha512(src, key []byte) ([]byte, error) {
	return Sha512Sum(src, key, nil)
}

// Sha512SumHex calculates the SHA-512 digest to sum and returns the resulting hex string.
func Sha512SumHex(src, key, sum string) (string, error) {
	var sumtext []byte
	if strings.IsNotBlank(sum) {
		sumtext = []byte(sum)
	}
	out, err := Sha512Sum([]byte(src), []byte(key), sumtext)
	if err != nil {
		return "", err
	}
	dst := fmt.Sprintf("%x", out)
	return dst, nil
}

// Sha512Sum calculates the SHA-512 digest to sum and returns the resulting slice.
func Sha512Sum(src, key, sum []byte) ([]byte, error) {
	h := hmac.New(sha512.New, key)
	if _, err := h.Write(src); err != nil {
		return nil, err
	}
	return h.Sum(sum), nil
}
