// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package aes implements AES encryption (formerly Rijndael), as defined in
// U.S. Federal Information Processing Standards Publication 197.
package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"gopkg.in/goyy/goyy.v0/util/crypto/md5"
)

const DefaultKey = "aes-key"

// Encrypt encrypts the first block in src into dst.
// Dst and src may point at the same memory.
func Encrypt(src, key []byte) (dst []byte, err error) {
	keytext := key
	keylen := len(key)
	if keylen != 16 || keylen != 24 || keylen != 32 {
		keytext, err = md5.Digest(key)
		if err != nil {
			return
		}
	}
	var block cipher.Block
	block, err = aes.NewCipher(keytext)
	if err != nil {
		return
	}
	dst = make([]byte, len(src))
	iv := keytext[:aes.BlockSize] // aes.BlockSize = 16
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(dst, src)
	return
}

// Decrypt decrypts the first block in src into dst.
// Dst and src may point at the same memory.
func Decrypt(src, key []byte) (dst []byte, err error) {
	keytext := key
	keylen := len(key)
	if keylen != 16 || keylen != 24 || keylen != 32 {
		keytext, err = md5.Digest(key)
		if err != nil {
			return
		}
	}
	var block cipher.Block
	block, err = aes.NewCipher(keytext)
	if err != nil {
		return
	}
	dst = make([]byte, len(src))
	iv := keytext[:aes.BlockSize] // aes.BlockSize = 16
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
