// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package secure

import (
	"gopkg.in/goyy/goyy.v0/util/errors"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"gopkg.in/goyy/goyy.v0/util/times"
	"gopkg.in/goyy/goyy.v0/web/secure/internal"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
	"strconv"
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
		c.Session().Set(principalId, u.Id())
		c.Session().Set(principalName, u.Name())
		c.Session().Set(principalLoginName, u.LoginName())
		c.Session().Set(principalLoginTime, strconv.FormatInt(times.Now(), 10))
		c.Session().Set(principalPermissions, ps)
	} else {
		return errors.New(i18N.Message("err.login"))
	}
	return nil
}

func Logout(c xhttp.Context) error {
	err := c.Session().Delete(principalId)
	if err != nil {
		return err
	}
	err = c.Session().Delete(principalName)
	if err != nil {
		return err
	}
	err = c.Session().Delete(principalLoginName)
	if err != nil {
		return err
	}
	err = c.Session().Delete(principalPermissions)
	if err != nil {
		return err
	}
	return nil
}

func IsLogin(c xhttp.Context) bool {
	if c == nil {
		return false
	}
	if id, err := c.Session().Get(principalId); err == nil {
		if strings.IsNotBlank(id) {
			return true
		}
	}
	return false
}

func GetPrincipal(c xhttp.Context) (Principal, error) {
	p := Principal{}
	id, err := c.Session().Get(principalId)
	if err != nil {
		return p, err
	}
	name, err := c.Session().Get(principalName)
	if err != nil {
		return p, err
	}
	loginName, err := c.Session().Get(principalLoginName)
	if err != nil {
		return p, err
	}
	permissions, err := c.Session().Get(principalPermissions)
	if err != nil {
		return p, err
	}
	p.Id = id
	p.Name = name
	p.LoginName = loginName
	p.Permissions = permissions
	return p, nil
}
