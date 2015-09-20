// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// GeneratingKey generates an RSA keypair of the given bit size.
func GeneratingKey(bits int) (privateKey string, publicKey string, err error) {
	prikey, pubkey, err := GenerateKey(bits)
	return string(prikey), string(pubkey), err
}

// GenerateKey generates an RSA keypair of the given bit size.
func GenerateKey(bits int) (privateKey []byte, publicKey []byte, err error) {
	key, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	prikey := marshalPrivateKey(key)
	pubkey, err := marshalPublicKey(&key.PublicKey)
	if err != nil {
		return nil, nil, err
	}
	return prikey, pubkey, nil
}

// marshalPrivateKey converts a private key to ASN.1 DER encoded form.
func marshalPrivateKey(privateKey *rsa.PrivateKey) []byte {
	marshaled := x509.MarshalPKCS1PrivateKey(privateKey)
	encoded := pem.EncodeToMemory(&pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   marshaled,
	})
	return encoded
}

// marshalPublicKey serialises a public key to DER-encoded PKIX format.
func marshalPublicKey(publicKey *rsa.PublicKey) ([]byte, error) {
	marshaled, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, err
	}
	encoded := pem.EncodeToMemory(&pem.Block{
		Type:    "RSA PUBLIC KEY",
		Headers: nil,
		Bytes:   marshaled,
	})
	return encoded, nil
}

// parsePrivateKey returns an RSA private key from its ASN.1 PKCS#1 DER encoded form.
func parsePrivateKey(privateKey []byte) (*rsa.PrivateKey, error) {
	decoded, _ := pem.Decode(privateKey)
	if decoded == nil {
		return nil, errors.New("private key error")
	}
	prikey, err := x509.ParsePKCS1PrivateKey(decoded.Bytes)
	if err != nil {
		return nil, err
	}
	return prikey, nil
}

// parsePublicKey parses a DER encoded public key.
// These values are typically found in PEM blocks with "BEGIN PUBLIC KEY".
func parsePublicKey(publicKey []byte) (*rsa.PublicKey, error) {
	decoded, _ := pem.Decode(publicKey)
	if decoded == nil {
		return nil, errors.New("public key error")
	}
	pubkey, err := x509.ParsePKIXPublicKey(decoded.Bytes)
	if err != nil {
		return nil, err
	}
	return pubkey.(*rsa.PublicKey), nil
}
