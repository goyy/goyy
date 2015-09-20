# util-crypto [![GoDoc](http://godoc.org/gopkg.in/goyy/goyy.v0?status.png)](http://godoc.org/gopkg.in/goyy/goyy.v0/util/crypto)
General encoding/decoding algorithms library for Go

# Installation
`go get gopkg.in/goyy/goyy.v0/util/crypto`

# Usage
*MD5*

	dst, _ := md5.DigestHex("goyy")
	fmt.Println(dst)
	// Output:db9e2a3e99dbace8332b3042a6beb793

*SHA1*

	dst, _ := sha1.DigestHex("goyy")
	fmt.Println(dst)
	// Output:9a5de4d2e62e2c0f3018eeff35e09ab3d41781fb

*AES*

	dst, _ := aes.EncryptHex("goyy", "key")
	fmt.Println(dst)
	// Output:efe9f867
	
	dst, _ := aes.DecryptHex("efe9f867", "key")
	fmt.Println(dst)
	// Output:goyy

*RSA*

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
	
	in := "goyy"
	out, _ := rsa.EncryptHex(in, publicKey)
	dst, _ := rsa.DecryptHex(out, privateKey)
	fmt.Println(in == dst)
	// Output:true
