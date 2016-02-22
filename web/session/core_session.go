// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package session

import (
	"net/http"

	"github.com/satori/go.uuid"
	"gopkg.in/goyy/goyy.v0/data/cache"
	"gopkg.in/goyy/goyy.v0/util/errors"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

func New(w http.ResponseWriter, r *http.Request, o *Options) Interface {
	// ignore error -> http: named cookie not present
	cookie, _ := r.Cookie(cookieKey)
	if cookie == nil {
		sid := strings.Replace(uuid.NewV4().String(), "-", "", -1)
		cookie = &http.Cookie{
			Name:     cookieKey,
			Value:    sid,
			Path:     o.Path,
			Domain:   o.Domain,
			Secure:   o.Secure,
			HttpOnly: o.HttpOnly,
		}
		http.SetCookie(w, cookie)
	}
	s := &session{
		id:      cookie.Value,
		key:     sessionPrefix + ":" + cookie.Value,
		options: o,
		request: r,
	}
	return s
}

type session struct {
	id      string
	key     string
	options *Options
	request *http.Request
}

func (me *session) Get(key string) (string, error) {
	v, err := cache.HGet(me.key, key)
	logger.Debugf("Get %v %v %v\r\n", me.key, key, err)
	if err != nil {
		return "", err
	}
	me.expire(me.options.MaxAge)
	return v, nil
}

func (me *session) Set(key string, val string) error {
	err := cache.HSet(me.key, key, val)
	logger.Debugf("Set %v %v %v %v\r\n", me.key, key, val, err)
	if err != nil {
		return err
	}
	me.expire(me.options.MaxAge)
	return nil
}

func (me *session) Delete(key string) error {
	err := cache.HDelete(me.key, key)
	logger.Debugf("Delete %v %v %v\r\n", me.key, key, err)
	if err != nil {
		return err
	}
	return nil
}

func (me *session) Clear() error {
	err := cache.Delete(me.key)
	logger.Debugf("Clear %v %v\r\n", me.key, err)
	if err != nil {
		return err
	}
	return nil
}

func (me *session) AddFlash(value string, vars ...string) {
}

func (me *session) Flashes(vars ...string) []string {
	return []string{}
}

func (me *session) Options(options *Options) {
	me.options = options
}

func (me *session) IsLogin() bool {
	if !me.exists(principalId) {
		return false
	}
	if id, err := me.Get(principalId); err == nil {
		if strings.IsNotBlank(id) {
			return true
		}
	}
	return false
}

func (me *session) Principal() (Principal, error) {
	p := Principal{}
	if !me.exists(principalId) {
		return p, errors.New("Not logged in")
	}
	id, err := me.Get(principalId)
	if err != nil {
		return p, err
	}
	name, err := me.Get(principalName)
	if err != nil {
		return p, err
	}
	loginName, err := me.Get(principalLoginName)
	if err != nil {
		return p, err
	}
	loginTime, err := me.Get(principalLoginTime)
	if err != nil {
		return p, err
	}
	permissions, err := me.Get(principalPermissions)
	if err != nil {
		return p, err
	}
	p.Id = id
	p.Name = name
	p.LoginName = loginName
	p.LoginTime = loginTime
	p.Permissions = permissions
	return p, nil
}

func (me *session) SetPrincipal(value Principal) error {
	if err := me.Set(principalId, value.Id); err != nil {
		return err
	}
	if err := me.Set(principalName, value.Name); err != nil {
		return err
	}
	if err := me.Set(principalLoginName, value.LoginName); err != nil {
		return err
	}
	if err := me.Set(principalLoginTime, value.LoginTime); err != nil {
		return err
	}
	if err := me.Set(principalPermissions, value.Permissions); err != nil {
		return err
	}
	return nil
}

func (me *session) ResetPrincipal() error {
	if err := me.Delete(principalId); err != nil {
		return err
	}
	if err := me.Delete(principalName); err != nil {
		return err
	}
	if err := me.Delete(principalLoginName); err != nil {
		return err
	}
	if err := me.Delete(principalLoginTime); err != nil {
		return err
	}
	if err := me.Delete(principalPermissions); err != nil {
		return err
	}
	return nil
}

func (me *session) exists(key string) bool {
	v := cache.HExists(me.key, key)
	logger.Debugf("exists %v %v %v\r\n", me.key, key, v)
	return v
}

func (me *session) expire(second int) error {
	err := cache.Expire(me.key, second)
	logger.Debugf("expire %v %v\r\n", me.key, err)
	if err != nil {
		return err
	}
	return nil
}
