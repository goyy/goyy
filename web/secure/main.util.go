// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package secure

import (
	"gopkg.in/goyy/goyy.v0/util/crypto/aes"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

func EncryptPasswd(passwd string) string {
	if strings.IsNotBlank(passwd) {
		npasswd, _ := aes.EncryptHex(passwd, passwdKey)
		return npasswd
	}
	return passwd
}
