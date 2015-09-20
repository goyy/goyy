// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package rsa implements RSA encryption as specified in PKCS#1.
package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/hex"
)

// Encrypt encrypts the first block in src into dst.
// Dst and src may point at the same memory.
func Encrypt(src, publicKey []byte) ([]byte, error) {
	pubkey, err := parsePublicKey(publicKey)
	if err != nil {
		return nil, err
	}
	return rsa.EncryptPKCS1v15(rand.Reader, pubkey, src)
}

// Decrypt decrypts the first block in src into dst.
// Dst and src may point at the same memory.
func Decrypt(src, privateKey []byte) ([]byte, error) {
	prikey, err := parsePrivateKey(privateKey)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, prikey, src)
}

// EncryptHex encrypts the first block in src into dst.
func EncryptHex(src, publicKey string) (string, error) {
	out, err := Encrypt([]byte(src), []byte(publicKey))
	if err != nil {
		return "", err
	}
	dst := hex.EncodeToString(out)
	return dst, nil
}

// DecryptHex decrypts the first block in src into dst.
func DecryptHex(src, privateKey string) (string, error) {
	plaintext, err := hex.DecodeString(src)
	if err != nil {
		return "", err
	}
	dst, err := Decrypt(plaintext, []byte(privateKey))
	if err != nil {
		return "", err
	}
	return string(dst), nil
}
