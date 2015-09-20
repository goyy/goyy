// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rsa_test

import (
	"gopkg.in/goyy/goyy.v0/util/crypto/rsa"
	"strings"
	"testing"
)

func TestGeneratingKey(t *testing.T) {
	prikey, pubkey, _ := rsa.GeneratingKey(512)
	if !strings.Contains(prikey, "-----BEGIN RSA PRIVATE KEY-----") {
		t.Errorf(`privateKey = "%s"`, prikey)
	}
	if !strings.Contains(pubkey, "-----BEGIN RSA PUBLIC KEY-----") {
		t.Errorf(`publicKey = "%s"`, pubkey)
	}
}

var privateKey = `
-----BEGIN RSA PRIVATE KEY-----
MIIBOwIBAAJBAM7ews9ZrWeDybHsThj03sUa3AqpQ+aR0wcDfzNf1QgorJ8u0u+5
WebCh8Fl2OZkuJvkqNDO+QeUQGvenLHHCbsCAwEAAQJBAJ2LywNNAaG/HUFSfNvG
yU2FOiUoaZzUW8mQoTQH/N67dHm6kQqTDTp0ppRkB8DyXjdpgyqquhGfFujfnrmj
51ECIQDwU3oSC+AmQy/YmKjfu6eL8SZaKsAf6CWcOjhqkMDHMwIhANxcsttcWgvl
VtztF/FcS8YR8rAkaiczpTEsY53cIxNZAiEAxgXyssYIR17bIN0BYYEtmFj3IhrR
vji6LNWoQN7PihMCIFfnRpXIxkbeioMAtT9bwQJXIIdxT0MqD+iIu4g6S2epAiA3
aI9Exlhsv47JCGs3ZCN5JpD7XVAdo2t629s7pwPfnA==
-----END RSA PRIVATE KEY-----
`

var publicKey = `
-----BEGIN RSA PUBLIC KEY-----
MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBAM7ews9ZrWeDybHsThj03sUa3AqpQ+aR
0wcDfzNf1QgorJ8u0u+5WebCh8Fl2OZkuJvkqNDO+QeUQGvenLHHCbsCAwEAAQ==
-----END RSA PUBLIC KEY-----
`
