// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package secure

import (
	"gopkg.in/goyy/goyy.v0/util/errors"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"gopkg.in/goyy/goyy.v0/util/times"
	"gopkg.in/goyy/goyy.v0/web/secure/internal"
	"gopkg.in/goyy/goyy.v0/web/session"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

func Login(c xhttp.Context, loginName, passwd string) error {
	if strings.IsBlank(passwd) {
		return errors.NewNotBlank("passwd")
	}
	npasswd := EncryptPasswd(passwd)
	checked := internal.UserMgr.CheckPasswd(loginName, npasswd)
	if checked {
		u, err := internal.UserMgr.SelectUser(loginName)
		if err != nil {
			return err
		}
		ps, err := internal.PermissionMgr.SelectPermission(u.Id())
		if err != nil {
			return err
		}
		p := session.Principal{
			Id:          u.Id(),
			Name:        u.Name(),
			LoginName:   u.LoginName(),
			LoginTime:   times.NowStr(),
			Permissions: ps,
		}
		return c.Session().SetPrincipal(p)
	} else {
		return errors.New(i18N.Message("err.login"))
	}
	return nil
}

func Logout(c xhttp.Context) error {
	return c.Session().ResetPrincipal()
}
