// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package i18n

import (
	"errors"
	"fmt"
	"log"
	"os"

	"gopkg.in/goyy/goyy.v0/util/strings"
)

var defaultLocale = LocaleZhCN

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
// Get the latest default locale.
func New(locales map[string]map[string]string) I18n {
	if locales == nil {
		log.Fatalln("i18n.New:the locales not be nil!")
	}
	return &i18N{locales: locales}
}

// NewBy creates a new I18n by map and locale.
func NewBy(locales map[string]map[string]string, locale string) I18n {
	if locales == nil {
		log.Fatalln("i18n.New:the locales not be nil!")
	}
	return &i18N{locales: locales, locale: transformLocale(locale)}
}

// NewByEnv creates a new I18n by map.
// locale value from the I18N_LOCALE of the environment variable.
func NewByEnv(locales map[string]map[string]string) I18n {
	locale := os.Getenv("I18N_LOCALE")
	return NewBy(locales, transformLocale(locale))
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
	} else {
		if v, ok := me.locales[LocaleEnUS]; ok {
			if s, ok := v[key]; ok {
				return s
			}
		} else {
			if v, ok := me.locales[LocaleZhCN]; ok {
				if s, ok := v[key]; ok {
					return s
				}
			}
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
		me.locale = defaultLocale
	}
	return me.locale
}

// Sets the current locale for the i18n.
func (me *i18N) SetLocale(locale string) {
	if strings.IsNotBlank(locale) {
		me.locale = transformLocale(locale)
	}
}

// DefaultLocale gets the default locale.
func DefaultLocale() string {
	return defaultLocale
}

// SetDefaultLocale sets the default locale.
func SetDefaultLocale(locale string) {
	if strings.IsNotBlank(locale) {
		defaultLocale = transformLocale(locale)
		if strings.IsBlank(defaultLocale) {
			defaultLocale = LocaleZhCN
		}
	}
}

func transformLocale(locale string) string {
	switch locale {
	case "cs_CZ":
		return LocaleCsCZ
	case "da_DK":
		return LocaleDaDK
	case "de_DE":
		return LocaleDeDE
	case "en_GB":
		return LocaleEnGB
	case "en_US":
		return LocaleEnUS
	case "en_XM":
		return LocaleEnXM
	case "es_ES":
		return LocaleEsES
	case "es_LA":
		return LocaleEsLA
	case "es_MX":
		return LocaleEsMX
	case "es_NA":
		return LocaleEsNA
	case "fi_FI":
		return LocaleFiFI
	case "fr_CA":
		return LocaleFrCA
	case "fr_FR":
		return LocaleFrFR
	case "fr_XM":
		return LocaleFrXM
	case "hu_HU":
		return LocaleHuHU
	case "it_IT":
		return LocaleItIT
	case "ja_JP":
		return LocaleJaJP
	case "ko_KR":
		return LocaleKoKR
	case "nb_NO":
		return LocaleNbNO
	case "nl_NL":
		return LocaleNlNL
	case "pl_PL":
		return LocalePlPL
	case "pt_BR":
		return LocalePtBR
	case "ro_RO":
		return LocaleRoRO
	case "ru_RU":
		return LocaleRuRU
	case "sv_SE":
		return LocaleSvSE
	case "tr_TR":
		return LocaleTrTR
	case "uk_UA":
		return LocaleUkUA
	case "zh_CN":
		return LocaleZhCN
	case "zh_TW":
		return LocaleZhTW
	default:
		return "" // Get the latest default locale when the value is blank
		//return defaultLocale
	}
}
