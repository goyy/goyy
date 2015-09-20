// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xhttp

import (
	"github.com/satori/go.uuid"
	"gopkg.in/goyy/goyy.v0/data/cache"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"net/http"
)

const (
	sessionPrefix        = "session"
	cookieKey     string = "GSESSIONID"
)

func newSession4Redis(w http.ResponseWriter, r *http.Request) Session {
	cookie, _ := r.Cookie(cookieKey)
	if cookie == nil {
		sid := strings.Replace(uuid.NewV4().String(), "-", "", -1)
		cookie = &http.Cookie{
			Name:     cookieKey,
			Value:    sid,
			Path:     Conf.Session.Path,
			Domain:   Conf.Session.Domain,
			Secure:   Conf.Session.Secure,
			HttpOnly: Conf.Session.HttpOnly,
		}
		http.SetCookie(w, cookie)
	}
	s := &session{
		id:      cookie.Value,
		key:     sessionPrefix + ":" + cookie.Value,
		options: Conf.Session.Options,
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
	logger.Debugf("Get %v %v %v", me.key, key, err)
	if err != nil {
		return "", err
	}
	me.expire(Conf.Session.MaxAge)
	return v, nil
}

func (me *session) Set(key string, val string) error {
	err := cache.HSet(me.key, key, val)
	logger.Debugf("Set %v %v %v %v", me.key, key, val, err)
	if err != nil {
		return err
	}
	me.expire(Conf.Session.MaxAge)
	return nil
}

func (me *session) Delete(key string) error {
	err := cache.HDelete(me.key, key)
	logger.Debugf("Delete %v %v %v", me.key, key, err)
	if err != nil {
		return err
	}
	return nil
}

func (me *session) Clear() error {
	err := cache.Delete(me.key)
	logger.Debugf("Clear %v %v", me.key, err)
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

func (me *session) exists(key string) bool {
	v := cache.HExists(me.key, key)
	logger.Debugf("exists %v %v %v", me.key, key, v)
	return v
}

func (me *session) expire(second int) error {
	err := cache.Expire(me.key, second)
	logger.Debugf("expire %v %v", me.key, err)
	if err != nil {
		return err
	}
	return nil
}
