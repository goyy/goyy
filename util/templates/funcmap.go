// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package templates

import (
	htmpl "html/template"
	ttmpl "text/template"

	"gopkg.in/goyy/goyy.v0/util/strings"
	"gopkg.in/goyy/goyy.v0/util/times"
	"gopkg.in/goyy/goyy.v0/web/conf"
)

var Text = text{funcMapText}

var Html = html{funcMapHtml}

type text struct {
	FuncMap ttmpl.FuncMap
}

type html struct {
	FuncMap htmpl.FuncMap
}

////////////////////////////////////////////////////////////
// template funcmap
////////////////////////////////////////////////////////////

var funcMapText = ttmpl.FuncMap{
	"yymd":     times.FormatYYMD,
	"yymdhms":  times.FormatYYMDHMS,
	"yymdhm":   times.FormatYYMDHM,
	"uyymd":    times.FormatUnixYYMD,
	"uyymdhms": times.FormatUnixYYMDHMS,
	"uyymdhm":  times.FormatUnixYYMDHM,

	"blank":     strings.IsBlank,
	"notblank":  strings.IsNotBlank,
	"abbr":      strings.Abbr,
	"anon":      strings.Anon,
	"anonymity": strings.Anonymity,

	"divisible": divisible,
	"exist":     exist,
	"zhstate":   zhstate,
	"eqindex":   eqindex,
	"eqshow":    eqshow,
	"eqadd":     eqadd,
	"eqedit":    eqedit,
	"neindex":   neindex,
	"neshow":    neshow,
	"neadd":     neadd,
	"needit":    needit,

	"profile":    getProfile,    // profile[production|development|test]
	"apis":       getApis,       // The URL of the apis.
	"assets":     getAssets,     // The URL of the static file.
	"statics":    getStatics,    // static file[web|wap|bms]
	"developers": getDevelopers, // The URL of the development of relevant documents
	"operations": getOperations, // The URL of the operation related documents.
	"uploads":    getUploads,    // The URL of the uploaded file.
}

var funcMapHtml = htmpl.FuncMap{
	"yymd":     times.FormatYYMD,
	"yymdhms":  times.FormatYYMDHMS,
	"yymdhm":   times.FormatYYMDHM,
	"uyymd":    times.FormatUnixYYMD,
	"uyymdhms": times.FormatUnixYYMDHMS,
	"uyymdhm":  times.FormatUnixYYMDHM,

	"blank":     strings.IsBlank,
	"notblank":  strings.IsNotBlank,
	"abbr":      strings.Abbr,
	"anon":      strings.Anon,
	"anonymity": strings.Anonymity,

	"set":       set,
	"divisible": divisible,
	"exist":     exist,
	"zhstate":   zhstate,
	"eqindex":   eqindex,
	"eqshow":    eqshow,
	"eqadd":     eqadd,
	"eqedit":    eqedit,
	"neindex":   neindex,
	"neshow":    neshow,
	"neadd":     neadd,
	"needit":    needit,

	"profile":    getProfile,    // profile[production|development|test]
	"apis":       getApis,       // The URL of the apis.
	"assets":     getAssets,     // The URL of the static file.
	"statics":    getStatics,    // static file[web|wap|bms]
	"developers": getDevelopers, // The URL of the development of relevant documents
	"operations": getOperations, // The URL of the operation related documents.
	"uploads":    getUploads,    // The URL of the uploaded file.
}

////////////////////////////////////////////////////////////
// common
////////////////////////////////////////////////////////////

var set = func(args map[string]interface{}, key string, value interface{}) htmpl.HTML {
	args[key] = value
	return htmpl.HTML("")
}

var divisible = func(x, y int) bool {
	if x%y == 0 {
		return true
	}
	return false
}

var exist = func(m map[string]interface{}, key string) (ok bool) {
	_, ok = m[key]
	return
}

////////////////////////////////////////////////////////////
// template state
////////////////////////////////////////////////////////////

var zhstate = func(t string) string {
	switch t {
	case EnIndex:
		return zhAdd
	case EnShow:
		return zhShow
	case EnAdd:
		return zhAdd
	case EnEdit:
		return zhEdit
	}
	return ""
}

var eqindex = func(t string) bool {
	if EnIndex == t {
		return true
	}
	return false
}

var eqshow = func(t string) bool {
	if EnShow == t {
		return true
	}
	return false
}

var eqadd = func(t string) bool {
	if EnAdd == t {
		return true
	}
	return false
}

var eqedit = func(t string) bool {
	if EnEdit == t {
		return true
	}
	return false
}

var neindex = func(t string) bool {
	return !eqindex(t)
}

var neshow = func(t string) bool {
	return !eqshow(t)
}

var neadd = func(t string) bool {
	return !eqadd(t)
}

var needit = func(t string) bool {
	return !eqedit(t)
}

////////////////////////////////////////////////////////////
// static state
////////////////////////////////////////////////////////////

func getProfile() string {
	return conf.Conf.Profile
}

func getApis() string {
	return conf.Conf.Api.URL
}

func getAssets() string {
	return conf.Conf.Asset.URL
}

func getStatics() string {
	return conf.Conf.Static.URL
}

func getDevelopers() string {
	return conf.Conf.Developer.URL
}

func getOperations() string {
	return conf.Conf.Operation.URL
}

func getUploads() string {
	return conf.Conf.Upload.URL
}
