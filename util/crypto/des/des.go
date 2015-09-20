// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package des implements the Data Encryption Standard (DES) and the
// Triple Data Encryption Algorithm (TDEA) as defined
// in U.S. Federal Information Processing Standards Publication 46-3.
package des

import (
	"crypto/cipher"
	"crypto/des"
	"encoding/hex"
	"gopkg.in/goyy/goyy.v0/util/crypto/sha256"
)

// Encrypt encrypts the first block in src into dst.
// Dst and src may point at the same memory.
func Encrypt(src, key []byte) (dst []byte, err error) {
	keytext := key
	if len(key) != 24 {
		keytext, err = sha256.Digest(key)
		if err != nil {
			return
		}
	}
	var block cipher.Block
	block, err = des.NewTripleDESCipher(keytext[:24])
	if err != nil {
		return
	}
	dst = make([]byte, len(src))
	iv := keytext[:des.BlockSize] // aes.BlockSize = 8
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(dst, src)
	return
}

// Decrypt decrypts the first block in src into dst.
// Dst and src may point at the same memory.
func Decrypt(src, key []byte) (dst []byte, err error) {
	keytext := key
	if len(key) != 24 {
		keytext, err = sha256.Digest(key)
		if err != nil {
			return
		}
	}
	var block cipher.Block
	block, err = des.NewTripleDESCipher(keytext[:24])
	if err != nil {
		return
	}
	dst = make([]byte, len(src))
	iv := keytext[:des.BlockSize] // aes.BlockSize = 8
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(dst, src)
	return
}

// EncryptHex encrypts the first block in src into dst.
func EncryptHex(src, key string) (dst string, err error) {
	var ciphertext []byte
	ciphertext, err = Encrypt([]byte(src), []byte(key))
	if err != nil {
		return
	}
	dst = hex.EncodeToString(ciphertext)
	return
}

// DecryptHex decrypts the first block in src into dst.
func DecryptHex(src, key string) (dst string, err error) {
	var out, plaintext []byte
	plaintext, err = hex.DecodeString(src)
	if err != nil {
		return
	}
	out, err = Decrypt(plaintext, []byte(key))
	if err != nil {
		return
	}
	dst = string(out)
	return
}
