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

var defaultLocale string = Locale_en_US

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
		if v, ok := me.locales[Locale_en_US]; ok {
			if s, ok := v[key]; ok {
				return s
			}
		} else {
			if v, ok := me.locales[Locale_zh_CN]; ok {
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

// Gets the default locale.
func DefaultLocale() string {
	return defaultLocale
}

// Sets the default locale.
func SetDefaultLocale(locale string) {
	if strings.IsNotBlank(locale) {
		defaultLocale = transformLocale(locale)
		if strings.IsBlank(defaultLocale) {
			defaultLocale = Locale_en_US
		}
	}
}

func transformLocale(locale string) string {
	switch locale {
	case "cs_CZ":
		return Locale_cs_CZ
	case "da_DK":
		return Locale_da_DK
	case "de_DE":
		return Locale_de_DE
	case "en_GB":
		return Locale_en_GB
	case "en_US":
		return Locale_en_US
	case "en_XM":
		return Locale_en_XM
	case "es_ES":
		return Locale_es_ES
	case "es_LA":
		return Locale_es_LA
	case "es_MX":
		return Locale_es_MX
	case "es_NA":
		return Locale_es_NA
	case "fi_FI":
		return Locale_fi_FI
	case "fr_CA":
		return Locale_fr_CA
	case "fr_FR":
		return Locale_fr_FR
	case "fr_XM":
		return Locale_fr_XM
	case "hu_HU":
		return Locale_hu_HU
	case "it_IT":
		return Locale_it_IT
	case "ja_JP":
		return Locale_ja_JP
	case "ko_KR":
		return Locale_ko_KR
	case "nb_NO":
		return Locale_nb_NO
	case "nl_NL":
		return Locale_nl_NL
	case "pl_PL":
		return Locale_pl_PL
	case "pt_BR":
		return Locale_pt_BR
	case "ro_RO":
		return Locale_ro_RO
	case "ru_RU":
		return Locale_ru_RU
	case "sv_SE":
		return Locale_sv_SE
	case "tr_TR":
		return Locale_tr_TR
	case "uk_UA":
		return Locale_uk_UA
	case "zh_CN":
		return Locale_zh_CN
	case "zh_TW":
		return Locale_zh_TW
	default:
		return "" // Get the latest default locale when the value is blank
		//return defaultLocale
	}
}
