// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package i18n

import (
	"errors"
	"fmt"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"log"
)

// I18n is the interface that wraps the operation i18n method.
type I18n interface {
	Message(key string) string
	Messagef(key string, args ...interface{}) string
	Error(key string) error
	Errorf(key string, args ...interface{}) error
	Panic(key string)
	Panicf(key string, args ...interface{})
	Locale() string
	SetLocale(locale string)
}

// New creates a new I18n by map.
func New(locales map[string]map[string]string, locale string) I18n {
	if strings.IsBlank(locale) {
		locale = Locale_en_US
	}
	if locales == nil {
		log.Fatalln("i18n.New:the locales not be nil!")
	}
	return &i18N{locales: locales, locale: locale}
}

type i18N struct {
	locale  string
	locales map[string]map[string]string
}

// Try to resolve the message. Return "" if no message was found.
func (me *i18N) Message(key string) string {
	if v, ok := me.locales[me.Locale()]; ok {
		if s, ok := v[key]; ok {
			return s
		}
	}
	return ""
}

// Try to resolve and format the message. Return "" if no message was found.
func (me *i18N) Messagef(key string, args ...interface{}) string {
	s := me.Message(key)
	if strings.IsBlank(s) {
		return ""
	}
	return fmt.Sprintf(s, args...)
}

// Try to resolve the error.
func (me *i18N) Error(key string) error {
	return errors.New(me.Message(key))
}

// Try to resolve and format the error.
func (me *i18N) Errorf(key string, args ...interface{}) error {
	return errors.New(me.Messagef(key, args...))
}

// Try to resolve the panic.
func (me *i18N) Panic(key string) {
	panic(errors.New(me.Message(key)))
}

// Try to resolve and format the panic.
func (me *i18N) Panicf(key string, args ...interface{}) {
	panic(errors.New(me.Messagef(key, args...)))
}

// Gets the current locale for the i18n.
func (me *i18N) Locale() string {
	if strings.IsBlank(me.locale) {
		me.locale = Locale_en_US
	}
	return me.locale
}

// Sets the current locale for the i18n.
func (me *i18N) SetLocale(locale string) {
	if strings.IsNotBlank(locale) {
		me.locale = locale
	}
}
