// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package internal

import (
	"gopkg.in/goyy/goyy.v0/util/errors"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

func (me *UserManager) CheckPasswd(loginName, passwd string) bool {
	if strings.IsNotBlank(loginName) && strings.IsNotBlank(passwd) {
		count, err := me.SelectInt(countByLogin, loginName, passwd)
		if err != nil {
			logger.Error(err.Error())
			return false
		}
		if count == 1 {
			return true
		}
	}
	return false
}

func (me *UserManager) SelectUser(loginName string) (*User, error) {
	if strings.IsNotBlank(loginName) {
		out := NewUser()
		err := me.SelectOne(out, selectUserByLoginName, loginName)
		if err != nil {
			logger.Error(err.Error())
			return nil, err
		}
		return out, nil
	}
	return nil, errors.NewNotBlank("loginName")
}
