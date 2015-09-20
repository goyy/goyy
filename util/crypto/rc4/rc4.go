// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package rc4 implements RC4 encryption, as defined in Bruce Schneier's
// Applied Cryptography.
package rc4

import (
	"crypto/rc4"
	"encoding/hex"
)

// Encrypt encrypts the first block in src into dst.
// Dst and src may point at the same memory.
func Encrypt(src, key []byte) ([]byte, error) {
	cipher, err := rc4.NewCipher(key)
	if err != nil {
		return nil, err
	}
	dst := make([]byte, len(src))
	cipher.XORKeyStream(dst, src)
	return dst, nil
}

// Decrypt decrypts the first block in src into dst.
// Dst and src may point at the same memory.
func Decrypt(src, key []byte) ([]byte, error) {
	cipher, err := rc4.NewCipher(key)
	if err != nil {
		return nil, err
	}
	dst := make([]byte, len(src))
	cipher.XORKeyStream(dst, src)
	return dst, nil
}

// EncryptHex encrypts the first block in src into dst.
func EncryptHex(src, key string) (string, error) {
	ciphertext, err := Encrypt([]byte(src), []byte(key))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(ciphertext), nil
}

// DecryptHex decrypts the first block in src into dst.
func DecryptHex(src, key string) (string, error) {
	plaintext, err := hex.DecodeString(src)
	if err != nil {
		return "", err
	}
	dst, err := Decrypt(plaintext, []byte(key))
	return string(dst), nil
}
