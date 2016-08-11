// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xhttp

import (
	"math"
)

const (
	defaultStatus = 200

	ifModifiedSince = "If-Modified-Since"
	lastModified    = "Last-Modified"

	drtBegin    = "<!--#"
	drtEnd      = `-->`
	drtEndKey   = "end"
	drtIf       = "if"
	drtProfile  = "profile"
	drtInclude  = "include"
	drtSettings = "settings"

	tplBegin    = "{%"
	tplEnd      = "%}"
	tplEqParams = "eq .params"

	attrExpr    = "expr"
	attrAccepts = "accepts"
	attrFile    = "file"
	attrParams  = "params"
	attrProject = "project"
	attrModule  = "module"
	attrTitle   = "title"

	tagBeginPre         = `<`
	tagEndPre           = `>`
	tagTextEndPre       = `</`
	tagTextTitle        = `title`
	tagTextType         = `type`
	tagAttrClass        = `class`
	tagAttrHref         = `href`
	tagAttrSrc          = `src`
	tagAttrAction       = `action`
	tagAttrOnerror      = `onerror`
	tagAttrOnclick      = `onclick`
	tagAttrData         = "data-"
	tagAttrPre          = ` go:`
	tagAttrPost         = `="`
	tagAttrEnd          = `"`
	tagParams           = `{%params%}`
	tagParams2          = `{%.params%}`
	tagProfile          = `{%profile%}`
	tagApis             = `{%apis%}`
	tagAssets           = `{%assets%}`
	tagAssetsStatics    = `{%statics%}`
	tagAssetsDevelopers = `{%developers%}`
	tagAssetsOperations = `{%operations%}`
	tagAssetsUploads    = `{%uploads%}`
	tagAssetsVer        = `{%ver%}`

	noWritten = -1

	AbortIndex int8 = math.MaxInt8 / 2

	MIMEJSON              = "application/json"
	MIMEHTML              = "text/html"
	MIMEXML               = "application/xml"
	MIMEXML2              = "text/xml"
	MIMEPlain             = "text/plain"
	MIMEPOSTForm          = "application/x-www-form-urlencoded"
	MIMEMultipartPOSTForm = "multipart/form-data"
)
