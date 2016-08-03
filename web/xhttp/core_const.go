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

	directiveArgEnd            = `"-->`
	directiveIncludeBegin      = `<!--#include file="`
	directiveIncludeParamBegin = `" param="`
	directiveIncludeEnd        = `<!--#endinclude-->`
	directiveIfBegin           = `<!--#if expr="`
	directiveIfEnd             = `<!--#endif-->`
	directiveProfileBegin      = `<!--#profile accepts="`
	directiveProfileEnd        = `<!--#endprofile-->`

	tplBegin  = "{{if param `"
	tplArgEnd = "`}}"
	tplEnd    = "{{end}}"

	tagBeginPre         = `<`
	tagEndPre           = `>`
	tagTextEndPre       = `</`
	tagTextTitle        = `title`
	tagAttrHref         = `href`
	tagAttrSrc          = `src`
	tagAttrAction       = `action`
	tagAttrPre          = ` go:`
	tagAttrPost         = `="`
	tagAttrEnd          = `"`
	tagProfile          = `{{profile}}`
	tagApis             = `{{apis}}`
	tagAssets           = `{{assets}}`
	tagAssetsStatics    = `{{statics}}`
	tagAssetsDevelopers = `{{developers}}`
	tagAssetsOperations = `{{operations}}`
	tagAssetsUploads    = `{{uploads}}`

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
